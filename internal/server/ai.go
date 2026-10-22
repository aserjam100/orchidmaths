package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/aserjam100/orchid/internal/ai"
)

// Bounds for the AI endpoints. They are small on purpose: an assistant message
// is a question, not an essay, and the goal is a sentence. A caller that
// exceeds them is rejected rather than truncated, so a visitor never gets a
// reply to something they did not ask.
const (
	maxAIBody           = 16 << 10
	maxAssistantMsgLen  = 2000
	maxAssistantHistory = 8
	maxAssistantLinks   = 4
	maxGoalLen          = 500
	maxGoalDestinations = 3
)

// aiGoalScopeRequest is the goal-scope endpoint's body.
type aiGoalScopeRequest struct {
	Goal string `json:"goal"`
}

// aiAssistantRequest is the assistant endpoint's body.
//
// History is the earlier turns of the same conversation, replayed so the app can
// hold a conversation at all: the fixed prompt tells the assistant to ask one
// clarifying question, and a question with no memory of it is unanswerable. The
// server keeps only user and assistant roles from it, so a client cannot add a
// second set of instructions.
type aiAssistantRequest struct {
	Message string       `json:"message"`
	History []aiChatTurn `json:"history"`
}

// aiChatTurn is one replayed turn as the client sends it.
type aiChatTurn struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// aiAssistantLink is one clickable destination in the reply.
//
// The href is never taken from the model. The model writes a marker naming an id,
// the server looks the id up in its own table, and an id it does not know is
// removed from the text rather than turned into a link. A generated URL would be
// one prompt injection away from sending a visitor anywhere.
type aiAssistantLink struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Href  string `json:"href"`
	// Description is shown to the model with the id, so it can tell the
	// diagnostic from the planner. It is never sent to the client.
	Description string `json:"-"`
}

// aiAssistantPart is one run of the reply: either plain text or a link.
//
// The reply is sent as ordered parts rather than as text plus a list of links
// because the list rendered the label twice — once where the model wrote the
// marker, and again in a row of boxes underneath — so a sentence read "you can
// start with our Calculus path here: Calculus" followed by a Calculus box. A
// part sequence puts the link where the model put it and nowhere else.
type aiAssistantPart struct {
	Text string           `json:"text,omitempty"`
	Link *aiAssistantLink `json:"link,omitempty"`
}

type aiAssistantResponse struct {
	// Reply is the same text as the parts, flattened. It is what a client that
	// does not render links shows, and what the conversation history replays.
	Reply    string            `json:"reply"`
	Segments []aiAssistantPart `json:"segments"`
}

// assistantAreas are the fixed places the assistant may point at, beside the
// destination catalogue. Hrefs live here and nowhere else.
var assistantAreas = []aiAssistantLink{
	{ID: "diagnose", Label: "Take the diagnostic", Href: "/onboard", Description: "Measures what you already know and starts you at the right concept."},
	{ID: "plan", Label: "Your study plan", Href: "/plan", Description: "Sets a destination and a daily pace."},
	{ID: "learn", Label: "Learn", Href: "/learn", Description: "Answer questions and get them graded."},
	{ID: "review", Label: "Review", Href: "/review", Description: "Scheduled review of concepts you have met."},
	{ID: "graph", Label: "Concept graph", Href: "/graph", Description: "Every concept and its prerequisites."},
	{ID: "glossary", Label: "Glossary", Href: "/glossary", Description: "Definitions of mathematical terms."},
	{ID: "leaderboard", Label: "Leaderboard", Href: "/leaderboard", Description: "Standings, when someone asks about them."},
}

// assistantContextAreas offers the linkable areas to the model. Without this
// the resolver knows the ids and the model does not, so a rule telling it to
// point at the diagnostic could only be obeyed by guessing an id or by writing
// a URL, and a URL in a reply is not a link: only ids from this list are.
func assistantContextAreas() []ai.Destination {
	out := make([]ai.Destination, 0, len(assistantAreas))
	for _, a := range assistantAreas {
		out = append(out, ai.Destination{ID: a.ID, Name: a.Label, Description: a.Description})
	}
	return out
}

// assistantMarker matches the link syntax the system prompt requests.
var assistantMarker = regexp.MustCompile(`\[\[([A-Za-z0-9_-]+)\]\]`)

// assistantLinks resolves the markers in a reply against the server's own table.
//
// It returns the reply as ordered parts, with at most maxAssistantLinks unique
// links. A marker this server does not know is removed and produces no link, and
// a marker past the cap keeps its label but is rendered as text: the assistant
// can name a destination it was given, and nothing else.
func assistantLinks(reply string, destinations []ai.Destination) (string, []aiAssistantPart) {
	byID := make(map[string]aiAssistantLink, len(assistantAreas)+len(destinations))
	for _, a := range assistantAreas {
		byID[a.ID] = a
	}
	for _, d := range destinations {
		if d.ID == "" {
			continue
		}
		// A destination whose id collides with an area keeps the area: those
		// ids are reserved, and silently redirecting them would be surprising.
		if _, taken := byID[d.ID]; taken {
			continue
		}
		label := d.Name
		if label == "" {
			label = d.ID
		}
		byID[d.ID] = aiAssistantLink{
			ID:    d.ID,
			Label: label,
			Href:  "/plan?destination=" + url.QueryEscape(d.ID),
		}
	}

	parts := make([]aiAssistantPart, 0, 3)
	seen := make(map[string]bool, maxAssistantLinks)
	appendText := func(s string) {
		if s == "" {
			return
		}
		// Adjacent text runs are merged: removing an unknown marker splits the
		// sentence, and three spans where one is enough is a rendering artefact.
		if n := len(parts); n > 0 && parts[n-1].Link == nil {
			parts[n-1].Text = tidySpacing(parts[n-1].Text + s)
			return
		}
		parts = append(parts, aiAssistantPart{Text: tidySpacing(s)})
	}
	last := 0
	for _, m := range assistantMarker.FindAllStringSubmatchIndex(reply, -1) {
		appendText(reply[last:m[0]])
		link, ok := byID[reply[m[2]:m[3]]]
		switch {
		case !ok:
			// Unknown id: removed, so the reply cannot name what it was not given.
		case seen[link.Href]:
			// Already linked once. The word still reads as a link, since the
			// same marker twice is a sentence, not a second destination.
			l := link
			parts = append(parts, aiAssistantPart{Link: &l})
		case len(seen) >= maxAssistantLinks:
			appendText(link.Label)
		default:
			seen[link.Href] = true
			l := link
			parts = append(parts, aiAssistantPart{Link: &l})
		}
		last = m[1]
	}
	appendText(reply[last:])

	// The flattened reply is built from the parts rather than recomputed, so the
	// text and the links cannot disagree about what the reply said.
	var flat strings.Builder
	for _, p := range parts {
		if p.Link != nil {
			flat.WriteString(p.Link.Label)
			continue
		}
		flat.WriteString(p.Text)
	}
	return tidyRemovedMarkers(flat.String()), parts
}

// tidySpacing repairs the double spaces and the space before punctuation that
// deleting a marker leaves behind, without touching the reply's own line breaks
// or the space that separates a run of text from the link that follows it.
func tidySpacing(s string) string {
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " .", ".")
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}

// tidyRemovedMarkers is tidySpacing for a whole reply, where the leading and
// trailing whitespace of the reply itself is noise.
func tidyRemovedMarkers(s string) string {
	return strings.TrimSpace(tidySpacing(s))
}

// sanitizeAssistantHistory bounds what a client may replay. It keeps the most
// recent turns, drops anything that is not a user or assistant turn, and drops
// over-long turns rather than truncating them, so a replayed message is answered
// as the visitor typed it or not at all.
func sanitizeAssistantHistory(in []aiChatTurn) []ai.Turn {
	if len(in) == 0 {
		return nil
	}
	if len(in) > maxAssistantHistory {
		in = in[len(in)-maxAssistantHistory:]
	}
	out := make([]ai.Turn, 0, len(in))
	for _, t := range in {
		content := strings.TrimSpace(t.Content)
		if content == "" || len([]rune(content)) > maxAssistantMsgLen {
			continue
		}
		switch t.Role {
		case string(ai.RoleUser), string(ai.RoleAssistant):
			out = append(out, ai.Turn{Role: t.Role, Content: content})
		}
	}
	return out
}

// POST /api/ai/assistant — an optional, best-effort product assistant. It
// answers "what is Orchid" and "where do I start", and it is not a tutor: it
// holds no concept context and never guides a graded question. The deterministic
// app does not depend on it: a disabled registry, an exhausted quota or a
// provider outage all surface as an error the client shows as "unavailable".
func (s *Server) handleAIAssistant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if s.aiChain == nil || !s.aiChain.Enabled() {
		writeError(w, "ai unavailable", http.StatusServiceUnavailable)
		return
	}

	var req aiAssistantRequest
	if err := decodeAIJSON(w, r, &req); err != nil {
		return
	}
	msg := strings.TrimSpace(req.Message)
	if msg == "" || len([]rune(msg)) > maxAssistantMsgLen {
		writeError(w, "message must be 1 to 2000 characters", http.StatusBadRequest)
		return
	}

	if !s.reserveAIUsage(w, r) {
		return
	}

	// The catalogue is the server's own destination list, so the assistant can
	// name a real starting point and cannot invent one. With no catalogue it
	// still answers, and asks where the visitor wants to go.
	actx := ai.AssistantContext{
		Message:      msg,
		Destinations: assistantDestinations(s.goalCandidates()),
		Areas:        assistantContextAreas(),
		History:      sanitizeAssistantHistory(req.History),
	}
	res, err := s.aiChain.Complete(r.Context(), ai.DataPseudonymous, ai.Request{
		Messages:  actx.Messages(),
		MaxTokens: s.aiAssistantMaxTokens,
	})
	if err != nil {
		s.writeAIError(w, "assistant", err)
		return
	}
	reply, segments := assistantLinks(res.Response.Text, actx.Destinations)
	writeJSON(w, aiAssistantResponse{Reply: reply, Segments: segments})
}

// Goal-scope reasons. They exist so a failure is not silently indistinguishable
// from a model that simply chose nothing: the client shows different copy, and
// the value is what a log or counter can be built on.
const (
	goalReasonDisabled     = "disabled"      // AI is off or has no eligible provider
	goalReasonNoCandidates = "no_candidates" // the server has no destinations to choose from
	goalReasonFailed       = "failed"        // every provider in the chain failed
	goalReasonNoMatch      = "no_match"      // the model answered, but chose no valid destination
)

// aiGoalScopeResponse is the goal-scope reply.
type aiGoalScopeResponse struct {
	Destinations []string `json:"destinations"`
	UsedAI       bool     `json:"used_ai"`
	// Reason is set only when UsedAI is false.
	Reason string `json:"reason,omitempty"`
}

// POST /api/ai/goal-scope — maps a learner's free-text goal to destination ids
// from the repository's own list. It is an enhancement, not the mechanism: when
// AI is unavailable it returns used_ai=false with a reason, and the client keeps
// its deterministic fuzzy match.
func (s *Server) handleAIGoalScope(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if s.aiChain == nil || !s.aiChain.Enabled() {
		writeJSON(w, aiGoalScopeResponse{Destinations: []string{}, Reason: goalReasonDisabled})
		return
	}

	var req aiGoalScopeRequest
	if err := decodeAIJSON(w, r, &req); err != nil {
		return
	}
	goal := strings.TrimSpace(req.Goal)
	if goal == "" || len([]rune(goal)) > maxGoalLen {
		writeError(w, "goal must be 1 to 500 characters", http.StatusBadRequest)
		return
	}

	// The candidate list comes from the server's own data, never from the
	// request, so a client cannot make the model pick an id that does not exist.
	candidates := s.goalCandidates()
	if len(candidates) == 0 {
		writeJSON(w, aiGoalScopeResponse{Destinations: []string{}, Reason: goalReasonNoCandidates})
		return
	}

	if !s.reserveAIUsage(w, r) {
		return
	}

	res, err := s.aiChain.Complete(r.Context(), ai.DataPseudonymous, ai.Request{
		Messages:  goalScopeMessages(goal, candidates),
		MaxTokens: s.aiGoalMaxTokens,
		JSON:      true,
	})
	if err != nil {
		// A provider failure is not worth an error page here: the client's
		// deterministic match is the fallback and is already on screen. It is
		// still reported, because "the model failed" and "the model chose
		// nothing" must not look the same in a response or in a log.
		log.Printf("ai goal-scope: failed: %v", err)
		writeJSON(w, aiGoalScopeResponse{Destinations: []string{}, Reason: goalReasonFailed})
		return
	}

	picked := validateDestinationIDs(parseDestinationIDs(res.Response.Text), candidates)
	if len(picked) == 0 {
		writeJSON(w, aiGoalScopeResponse{Destinations: []string{}, Reason: goalReasonNoMatch})
		return
	}
	writeJSON(w, aiGoalScopeResponse{Destinations: picked, UsedAI: true})
}

// goalCandidate is one destination offered to the model.
type goalCandidate struct {
	ID          string
	Name        string
	Description string
}

// goalCandidates reads the destination list from the engine, which loaded it
// from data/destinations.json at boot.
func (s *Server) goalCandidates() []goalCandidate {
	if s.eng == nil {
		return nil
	}
	dests, err := s.eng.Destinations("")
	if err != nil {
		log.Printf("ai goal-scope: load destinations: %v", err)
		return nil
	}
	out := make([]goalCandidate, 0, len(dests))
	for _, d := range dests {
		if d.Destination == nil {
			continue
		}
		out = append(out, goalCandidate{ID: d.ID, Name: d.Name, Description: d.Description})
	}
	return out
}

// goalScopeMessages builds the goal-scope prompt. The candidate list is authored
// course data, not learner data, so this request carries only the learner's own
// words plus public text.
func goalScopeMessages(goal string, candidates []goalCandidate) []ai.Message {
	var b strings.Builder
	b.WriteString("Destinations:\n")
	for _, c := range candidates {
		b.WriteString("- ")
		b.WriteString(c.ID)
		b.WriteString(": ")
		b.WriteString(c.Name)
		if c.Description != "" {
			b.WriteString(" — ")
			b.WriteString(c.Description)
		}
		b.WriteByte('\n')
	}
	return []ai.Message{
		{Role: ai.RoleSystem, Content: "You map a learner's goal to the destinations in the list that meet it.\n" +
			"Reply with JSON only, in the form {\"destinations\":[\"id\"]}.\n" +
			"Use only ids from the list. Choose at most 3. If none fit, reply with an empty list."},
		{Role: ai.RoleUser, Content: "Goal: " + goal + "\n\n" + b.String()},
	}
}

// parseDestinationIDs reads the model's JSON reply. A reply that is not JSON, or
// is missing the field, yields no ids rather than an error: the caller falls
// back to the deterministic path.
func parseDestinationIDs(text string) []string {
	trimmed := strings.TrimSpace(text)
	// Tolerate a fenced block, which some models add even under JSON mode.
	if i := strings.Index(trimmed, "{"); i > 0 {
		trimmed = trimmed[i:]
	}
	if j := strings.LastIndex(trimmed, "}"); j >= 0 && j < len(trimmed)-1 {
		trimmed = trimmed[:j+1]
	}
	var parsed struct {
		Destinations []string `json:"destinations"`
	}
	if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
		return nil
	}
	return parsed.Destinations
}

// validateDestinationIDs keeps only ids that are in the candidate set, in the
// model's order, capped so a chatty reply cannot produce a huge scope.
func validateDestinationIDs(ids []string, candidates []goalCandidate) []string {
	valid := make(map[string]bool, len(candidates))
	for _, c := range candidates {
		valid[c.ID] = true
	}
	out := make([]string, 0, maxGoalDestinations)
	seen := make(map[string]bool, maxGoalDestinations)
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if !valid[id] || seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
		if len(out) == maxGoalDestinations {
			break
		}
	}
	return out
}

// assistantDestinations converts the server's destination list into the public
// catalogue the assistant may name. It is a projection, not a filter: the
// assistant needs ids, names and descriptions, and nothing about a learner's
// progress, which is why this type cannot carry any.
func assistantDestinations(candidates []goalCandidate) []ai.Destination {
	if len(candidates) == 0 {
		return nil
	}
	out := make([]ai.Destination, 0, len(candidates))
	for _, c := range candidates {
		out = append(out, ai.Destination{ID: c.ID, Name: c.Name, Description: c.Description})
	}
	return out
}

// reserveAIUsage consumes one request from the caller's daily budget and writes
// the error response itself when the budget is gone or the store failed. It
// returns false when the request must stop.
func (s *Server) reserveAIUsage(w http.ResponseWriter, r *http.Request) bool {
	if s.aiQuota == nil {
		return true
	}
	ok, err := s.aiQuota.Reserve(r.Context(), s.aiQuotaKey(r))
	if err != nil {
		log.Printf("ai quota: %v", err)
		writeError(w, "ai unavailable", http.StatusServiceUnavailable)
		return false
	}
	if !ok {
		// A refusal is the one AI outcome that leaves no trace in a provider
		// log, because no provider was called.
		if s.aiChain != nil {
			s.aiChain.Stats().Denied()
		}
		writeError(w, "you have reached today's AI limit", http.StatusTooManyRequests)
		return false
	}
	return true
}

// aiQuotaKey is the identity a daily AI budget is spent against.
//
// An authenticated caller is keyed by student id. Everyone else is keyed by IP,
// not by the guest id they sent: guest ids are generated in the browser and are
// therefore freely mintable, so an id-based guest cap is bypassed by clearing
// localStorage.
func (s *Server) aiQuotaKey(r *http.Request) string {
	if sid, _ := r.Context().Value(authStudentKey{}).(string); sid != "" {
		return "u:" + sid
	}
	return "ip:" + clientIP(r, s.ipHeader)
}

// writeAIError maps a chain failure to a response. ErrNoProvider is not the
// client's fault and is reported as unavailable, never as a data problem.
func (s *Server) writeAIError(w http.ResponseWriter, label string, err error) {
	if errors.Is(err, ai.ErrNoProvider) {
		writeError(w, "ai unavailable", http.StatusServiceUnavailable)
		return
	}
	log.Printf("ai %s: %v", label, err)
	writeError(w, "ai request failed", http.StatusBadGateway)
}

// decodeAIJSON decodes a bounded JSON body, writing a 400 on failure.
func decodeAIJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxAIBody)
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return err
	}
	return nil
}
