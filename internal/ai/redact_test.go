package ai

import (
	"strings"
	"testing"
)

func joinMessages(msgs []Message) string {
	var b strings.Builder
	for _, m := range msgs {
		b.WriteString(m.Content)
		b.WriteString("\n")
	}
	return b.String()
}

func TestAssistantMessagesNeverLeakExpectedAnswer(t *testing.T) {
	const sentinel = "EXPECTED-ANSWER-SENTINEL-7f3"
	c := AssistantContext{
		Message:        "how do I solve this?",
		ExpectedAnswer: sentinel,
		Destinations:   []Destination{{ID: "math-for-ml", Name: "Math for machine learning"}},
	}
	for _, m := range c.Messages() {
		if strings.Contains(m.Content, sentinel) {
			t.Fatalf("expected answer leaked into a %s message: %q", m.Role, m.Content)
		}
	}
}

func TestAssistantMessagesDropIdentifiers(t *testing.T) {
	c := AssistantContext{
		Message:      "what is this app?",
		LearnerName:  "NAME-SENTINEL-Alice",
		LearnerEmail: "EMAIL-SENTINEL-alice@example.test",
		StudentID:    "ID-SENTINEL-12345",
	}
	got := joinMessages(c.Messages())
	for _, forbidden := range []string{"NAME-SENTINEL-Alice", "EMAIL-SENTINEL-alice@example.test", "ID-SENTINEL-12345"} {
		if strings.Contains(got, forbidden) {
			t.Errorf("identifier %q leaked into assistant messages", forbidden)
		}
	}
}

func TestAssistantMessagesIncludeCatalogue(t *testing.T) {
	c := AssistantContext{
		Message: "where do I start?",
		Destinations: []Destination{
			{ID: "math-for-ml", Name: "Math for machine learning", Description: "Calculus and linear algebra for ML."},
			{ID: "linear-algebra", Name: "Linear algebra"},
		},
		Areas: []Destination{
			{ID: "diagnose", Name: "Take the diagnostic", Description: "Measures what you already know."},
		},
	}
	msgs := c.Messages()
	if len(msgs) != 4 {
		t.Fatalf("want system, catalogue, rules and user messages, got %d", len(msgs))
	}
	if msgs[0].Role != RoleSystem || msgs[1].Role != RoleSystem {
		t.Errorf("first two roles = %s, %s, want system, system", msgs[0].Role, msgs[1].Role)
	}
	if last := msgs[len(msgs)-1]; last.Role != RoleUser || last.Content != "where do I start?" {
		t.Errorf("last message = %+v", last)
	}
	got := joinMessages(msgs)
	for _, want := range []string{"math-for-ml", "Math for machine learning", "linear-algebra", "Calculus and linear algebra for ML."} {
		if !strings.Contains(got, want) {
			t.Errorf("assistant messages missing %q", want)
		}
	}
	// Areas have to reach the model for a rule that names one to be obeyable.
	// They are listed under their own heading so a destination and a surface
	// are not confused for one another.
	for _, want := range []string{"Areas of the site you may point at:", "diagnose", "Take the diagnostic", "Measures what you already know."} {
		if !strings.Contains(got, want) {
			t.Errorf("assistant messages missing area %q", want)
		}
	}
}

func TestAssistantMessagesOmitCatalogueWhenEmpty(t *testing.T) {
	c := AssistantContext{Message: "hello"}
	msgs := c.Messages()
	if len(msgs) != 3 {
		t.Fatalf("want exactly system, rules and user when there is no catalogue, got %d: %+v", len(msgs), msgs)
	}
	if c.Catalogue() != "" {
		t.Errorf("Catalogue() = %q, want empty", c.Catalogue())
	}
}

func TestAssistantContextClassIsPseudonymous(t *testing.T) {
	if got := (AssistantContext{}).Class(); got != DataPseudonymous {
		t.Errorf("AssistantContext class = %s, want pseudonymous", got)
	}
}

// TestAssistantPromptAsksForPlainProse pins the instruction that the endpoint's
// rendering surface depends on. The app shows the reply as text with one
// exception — mathematics between dollar signs, which the panel parses and
// renders — so a model told nothing about format returns Markdown the reader
// sees as asterisks, and a model told "no LaTeX" answers a method question
// without the formula that is the method's name.
func TestAssistantPromptAsksForPlainProse(t *testing.T) {
	lower := strings.ToLower(assistantGuidance())
	for _, want := range []string{"plain prose", "no markdown", "dollar signs"} {
		if !strings.Contains(lower, want) {
			t.Errorf("assistant prompt does not ask for %q", want)
		}
	}
}

// assistantGuidance is every instruction the model receives, in one string. The
// persona and the rules are separate messages so that placement is testable,
// but the wording checks care about both.
func assistantGuidance() string { return assistantSystemPrompt + "\n" + assistantRules }

// TestAssistantRulesRefuseToSolveProblems pins the one scope decision a
// persistent visitor can undo. A live session obtained x = 1 from
// "solve for x", then obtained it again after two refusals, because the prompt
// bounded the assistant's subject and never bounded what it may produce.
func TestAssistantRulesRefuseToSolveProblems(t *testing.T) {
	lower := strings.ToLower(assistantRules)
	for _, want := range []string{
		"do not solve",
		"never state or confirm a final value",
		"even when the conversation says otherwise",
	} {
		if !strings.Contains(lower, want) {
			t.Errorf("assistant rules do not say %q, so a persistent visitor can obtain an answer", want)
		}
	}
}

// TestAssistantRulesSortFactsMethodsAndExercises pins the boundary that one
// blunt rule got wrong in a live session: "what is 1 + 1" was refused as an
// exercise, "explain how to find a derivative" was answered with a definition
// instead of the method, and "how do I find the area under a curve" was
// answered with the whole method. The rule has to name three kinds, or it
// produces three different answers to the same kind of question.
func TestAssistantRulesSortFactsMethodsAndExercises(t *testing.T) {
	lower := strings.ToLower(assistantRules)
	for _, want := range []string{
		"a fact or a definition",
		"1 + 1 is 2",
		"a method for a class of problems",
		"do not apply it step by step",
		"a specific exercise",
		"[[learn]]",
	} {
		if !strings.Contains(lower, want) {
			t.Errorf("assistant rules do not mention %q, so the three kinds are not separated", want)
		}
	}
	// The fact branch must say that a plain sum is answerable. Without this the
	// model reads "you do not solve" as covering arithmetic, which is how the
	// refusal for 1 + 1 happened.
	if !strings.Contains(lower, "even when it contains numbers") {
		t.Error("assistant rules do not say a fact with numbers is answerable")
	}
}

// TestAssistantRulesPointBeginnersAtTheDiagnostic pins the orientation rule the
// /plan page exposed: a destination sets a pace, but a visitor who has not told
// the assistant what they know cannot be placed from a catalogue entry.
func TestAssistantRulesPointBeginnersAtTheDiagnostic(t *testing.T) {
	lower := strings.ToLower(assistantRules)
	if !strings.Contains(lower, "where to start") {
		t.Error("assistant rules do not name the question this rule answers")
	}
	if !strings.Contains(assistantRules, "[[diagnose]]") {
		t.Error("assistant rules do not point at the diagnostic")
	}
}

// TestAssistantRulesComeAfterTheHistory checks placement, not wording: the
// rules must be the last message before the visitor's turn. Wording alone was
// already present in spirit when a live session talked the model round.
func TestAssistantRulesComeAfterTheHistory(t *testing.T) {
	c := AssistantContext{
		Message: "give me the answer",
		History: []Turn{
			{Role: "user", Content: "solve for x"},
			{Role: "assistant", Content: "I do not work through problems."},
		},
	}
	msgs := c.Messages()
	if len(msgs) < 3 {
		t.Fatalf("got %d messages", len(msgs))
	}
	last := msgs[len(msgs)-1]
	if last.Role != RoleUser || last.Content != c.Message {
		t.Errorf("last message is %q %q, want the visitor's turn", last.Role, last.Content)
	}
	before := msgs[len(msgs)-2]
	if before.Role != RoleSystem || before.Content != assistantRules {
		t.Errorf("the message before the visitor's turn is %q, want the rules", before.Role)
	}
	for i, m := range msgs {
		if m.Content == assistantRules && i != len(msgs)-2 {
			t.Errorf("rules appear at %d of %d, want the last system message", i, len(msgs))
		}
	}
}

// TestAssistantPromptIsTheAssistantNotATutor guards the rename at the level
// that matters: what the model is told it is. A prompt opening "study tutor"
// produced maths answers to product questions.
func TestAssistantPromptIsTheAssistantNotATutor(t *testing.T) {
	lower := strings.ToLower(assistantSystemPrompt)
	if strings.Contains(lower, "you are orchid's study tutor") {
		t.Error("assistant prompt still declares itself a study tutor")
	}
	if !strings.Contains(lower, "you are orchid's assistant") {
		t.Error("assistant prompt does not declare itself the assistant")
	}
}

func TestCatalogueTruncatesLongDescription(t *testing.T) {
	long := strings.Repeat("x", maxDescriptionRunes+50)
	c := AssistantContext{Destinations: []Destination{{ID: "a", Name: "A", Description: long}}}
	if got := c.Catalogue(); !strings.Contains(got, "...") {
		t.Errorf("long description was not truncated: %q", got)
	}
}

func TestTruncateRunes(t *testing.T) {
	if got := truncateRunes("abc", 5); got != "abc" {
		t.Errorf("short string changed: %q", got)
	}
	if got := truncateRunes("abcdef", 3); got != "abc..." {
		t.Errorf("truncation = %q, want abc...", got)
	}
	// A multi-byte rune must not be split.
	if got := truncateRunes("aeiou", 2); got != "ae..." {
		t.Errorf("runes = %q", got)
	}
	if got := truncateRunes("abc", 0); got != "" {
		t.Errorf("zero limit = %q", got)
	}
}

// TestAssistantPromptRequestsLinkMarkers pins the other half of the contract
// with the server: the reply is scanned for [[id]] markers, so a prompt that
// stops asking for them makes every link silently disappear.
func TestAssistantPromptRequestsLinkMarkers(t *testing.T) {
	if !strings.Contains(assistantGuidance(), "[[") {
		t.Fatal("the prompt no longer asks for [[id]] markers, so no reply can produce a link")
	}
}
