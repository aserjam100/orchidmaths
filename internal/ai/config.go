package ai

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DefaultDailyUserCap is the per-user daily request budget when unconfigured.
//
// It is deliberately small. Free tiers are shared across the whole
// organization, not per user (Groq's docs say limits apply at the org level,
// and its response headers report 1000 requests and 8000 tokens per window), so
// the product default must not let one learner spend the day for everyone.
const DefaultDailyUserCap = 5

// DefaultGlobalDailyCap is the application-wide daily request ceiling.
//
// It exists because a per-user cap does not bound the total: many learners at
// once still spend the whole organization's free tier. The default sits below
// the 1000 requests/day the Groq free tier reports, leaving headroom for
// retries and for the Cloudflare failover leg.
const DefaultGlobalDailyCap = 800

// Default token budgets. They are not constants in the handlers because a
// reasoning model (openai/gpt-oss-120b) spends part of every budget on hidden
// reasoning tokens before it emits an answer: a 200-token budget was measured
// spending 135 tokens on reasoning and then failing outright.
const (
	DefaultAssistantMaxTokens = 600
	DefaultGoalMaxTokens      = 600
)

// Settings is the AI configuration resolved from the environment.
type Settings struct {
	// Registry is nil when AI is disabled, which every consumer handles.
	Registry *Registry
	// DailyUserCap is the per-identity daily request budget.
	DailyUserCap int
	// GlobalDailyCap is the application-wide daily request ceiling.
	GlobalDailyCap int
	// AssistantMaxTokens and GoalMaxTokens bound one completion each.
	AssistantMaxTokens int
	GoalMaxTokens      int
	// Production marks a deployment serving real learner traffic.
	Production bool
	// DisabledReason is non-empty when AI was configured but refused to start,
	// and Registry is then nil. It carries a configuration error — most often an
	// unconfirmed zero-retention claim — rather than a runtime failure: the
	// deterministic app boots either way, because AI is an enhancement and a
	// mistake in it must not stop a learner from studying.
	DisabledReason string
}

// FromEnv builds the provider registry from the environment. It never dials a
// provider: a key that is present simply adds a provider to the chain, and a
// key that is absent leaves it out. Unset AI_ENABLED returns an empty Settings,
// so the feature is off unless it is asked for.
func FromEnv() (Settings, error) {
	settings := Settings{
		Production:         os.Getenv("AI_ENV") == "production",
		DailyUserCap:       DefaultDailyUserCap,
		GlobalDailyCap:     DefaultGlobalDailyCap,
		AssistantMaxTokens: DefaultAssistantMaxTokens,
		GoalMaxTokens:      DefaultGoalMaxTokens,
	}
	intEnv := func(key string, dst *int, allowZero bool) error {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			return nil
		}
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 || (!allowZero && n == 0) {
			kind := "a positive"
			if allowZero {
				kind = "a non-negative"
			}
			return fmt.Errorf("ai: %s must be %s integer, got %q", key, kind, v)
		}
		*dst = n
		return nil
	}
	if err := intEnv("AI_DAILY_USER_CAP", &settings.DailyUserCap, true); err != nil {
		return settings, err
	}
	// A global cap of zero means "no ceiling", which is a real choice, but it is
	// the one setting that can spend a shared free tier without limit, so it is
	// opted into explicitly rather than reached by leaving a default in place.
	if err := intEnv("AI_GLOBAL_DAILY_CAP", &settings.GlobalDailyCap, true); err != nil {
		return settings, err
	}
	if err := intEnv("AI_ASSISTANT_MAX_TOKENS", &settings.AssistantMaxTokens, false); err != nil {
		return settings, err
	}
	if err := intEnv("AI_GOAL_MAX_TOKENS", &settings.GoalMaxTokens, false); err != nil {
		return settings, err
	}
	if !envBool("AI_ENABLED") {
		return settings, nil
	}

	allowDev := envBool("AI_DEV_PROVIDERS")
	providers := make([]Provider, 0, 4)

	for _, spec := range productionSpecs() {
		p, ok, err := buildSpec(spec)
		if err != nil {
			return settings, err
		}
		if ok {
			providers = append(providers, p)
		}
	}
	if p, ok, err := buildCloudflare(); err != nil {
		return settings, err
	} else if ok {
		providers = append(providers, p)
	}
	for _, spec := range devSpecs() {
		p, ok, err := buildSpec(spec)
		if err != nil {
			return settings, err
		}
		if ok {
			providers = append(providers, p)
		}
	}

	providers = orderProviders(providers, os.Getenv("AI_PROVIDER_CHAIN"))
	reg, err := NewRegistry(providers, Options{
		Production:   settings.Production,
		AllowDev:     allowDev,
		ZDRConfirmed: envBool("AI_ZDR_CONFIRMED"),
	})
	if err != nil {
		// A configuration error in an optional feature disables that feature; it
		// does not stop the process. The reason is returned so the caller can say
		// why, rather than leaving an operator to guess from a silent "disabled".
		settings.DisabledReason = err.Error()
		return settings, nil
	}
	settings.Registry = reg
	return settings, nil
}

// providerSpec describes one env-configurable OpenAI-compatible provider.
type providerSpec struct {
	id       string
	keyEnv   string
	baseEnv  string
	modelEnv string
	baseURL  string
	model    string
	// effort is "reasoning_effort" for this provider, and is left empty when
	// Orchid has not verified that the provider accepts the field.
	effort    string
	effortEnv string
	pol       Policy
	caps      Capabilities
}

func buildSpec(s providerSpec) (Provider, bool, error) {
	key := strings.TrimSpace(os.Getenv(s.keyEnv))
	if key == "" {
		return nil, false, nil
	}
	base := s.baseURL
	if v := strings.TrimSpace(os.Getenv(s.baseEnv)); v != "" {
		base = v
	}
	model := s.model
	if v := strings.TrimSpace(os.Getenv(s.modelEnv)); v != "" {
		model = v
	}
	effort := s.effort
	if s.effortEnv != "" {
		if v := strings.TrimSpace(os.Getenv(s.effortEnv)); v != "" {
			effort = v
		}
	}
	// A dev/eval provider with no model has nothing safe to default to, since
	// its model ids are not ones Orchid has verified; leave it out rather than
	// invent one.
	if strings.TrimSpace(model) == "" {
		return nil, false, nil
	}
	p, err := NewOpenAIProvider(OpenAIConfig{
		ID: s.id, BaseURL: base, APIKey: key, Model: model, Policy: s.pol,
		Capabilities: s.caps, ReasoningEffort: effort,
	})
	if err != nil {
		return nil, false, err
	}
	return p, true, nil
}

// productionSpecs are providers whose terms permit real learner traffic.
func productionSpecs() []providerSpec {
	return []providerSpec{
		{
			id:       "groq",
			keyEnv:   "GROQ_API_KEY",
			baseEnv:  "GROQ_BASE_URL",
			modelEnv: "GROQ_MODEL",
			baseURL:  "https://api.groq.com/openai/v1",
			model:    "openai/gpt-oss-120b",
			// Verified against the live API on 2026-10-18: without this, a
			// 200-token request spent 135 tokens on hidden reasoning and then
			// failed with json_validate_failed; with it, reasoning fell to ~32
			// tokens and the same request succeeded.
			effort:    "low",
			effortEnv: "GROQ_REASONING_EFFORT",
			pol: Policy{
				MaxClass:            DataPseudonymous,
				ProductionPermitted: true,
				Retention:           RetentionNone,
				// No retention: zero data retention is enabled on this
				// organisation, and the registry refuses a production-permitted
				// RetentionNone provider until AI_ZDR_CONFIRMED=1 records that
				// the setting was checked (console.groq.com/docs/your-data,
				// verified 2026-10-16).
				Source: "https://console.groq.com/docs/your-data (2026-10-16)",
			},
			caps: Capabilities{ContextWindow: 131072, Models: []string{"openai/gpt-oss-120b"}, JSON: true, Streaming: true},
		},
	}
}

// devSpecs are providers restricted to development and synthetic content. They
// are inert in production regardless of any flag, because their policies are
// DevOnly.
func devSpecs() []providerSpec {
	devPol := func(what string) Policy {
		return Policy{
			MaxClass:  DataPublic,
			DevOnly:   true,
			Retention: RetentionUnknown,
			Source:    "dev-only: " + what,
		}
	}
	return []providerSpec{
		{
			id: "nvidia", keyEnv: "NVIDIA_API_KEY", baseEnv: "NVIDIA_BASE_URL", modelEnv: "NVIDIA_MODEL",
			baseURL: "https://integrate.api.nvidia.com/v1",
			pol:     devPol("https://integrate.api.nvidia.com (trial terms: no personal or confidential data)"),
		},
		{
			id: "cerebras", keyEnv: "CEREBRAS_API_KEY", baseEnv: "CEREBRAS_BASE_URL", modelEnv: "CEREBRAS_MODEL",
			baseURL: "https://api.cerebras.ai/v1",
			pol:     devPol("https://inference-docs.cerebras.ai/support/pricing (not intended for production)"),
		},
		{
			id: "openrouter", keyEnv: "OPENROUTER_API_KEY", baseEnv: "OPENROUTER_BASE_URL", modelEnv: "OPENROUTER_MODEL",
			baseURL: "https://openrouter.ai/api/v1",
			pol:     devPol("https://openrouter.ai/docs/api-reference/limits (per-model data policy)"),
		},
		{
			id: "opencode", keyEnv: "OPENCODE_API_KEY", baseEnv: "OPENCODE_BASE_URL", modelEnv: "OPENCODE_MODEL",
			baseURL: "https://opencode.ai/inference/v1",
			pol:     devPol("opencode.ai console (free models, limited time)"),
		},
	}
}

// buildCloudflare constructs the Workers AI provider. Cloudflare needs an
// account id and a token and carries the base path in full, so it does not fit
// the generic spec.
func buildCloudflare() (Provider, bool, error) {
	token := strings.TrimSpace(os.Getenv("CLOUDFLARE_API_TOKEN"))
	account := strings.TrimSpace(os.Getenv("CLOUDFLARE_ACCOUNT_ID"))
	if token == "" || account == "" {
		return nil, false, nil
	}
	model := strings.TrimSpace(os.Getenv("CLOUDFLARE_MODEL"))
	if model == "" {
		// Workers AI's Llama 3.1 8B is served quantised to FP8 under this id;
		// the unquantised id it replaced is the one that was failing with
		// failovers=1 while Groq covered for it. Pinning the working id here
		// means a deployment cannot depend on a Fly secret for it: the earlier
		// default was a guess, and a guess that 404s is invisible while the
		// first provider is healthy. Verified against the live API on
		// 2026-10-20 (200 in 0.56s, content "ok").
		model = "@cf/meta/llama-3.1-8b-instruct-fp8"
	}
	p, err := NewOpenAIProvider(OpenAIConfig{
		ID:      "cloudflare",
		BaseURL: "https://api.cloudflare.com/client/v4/accounts/" + account + "/ai/v1",
		APIKey:  token,
		Model:   model,
		Policy: Policy{
			MaxClass:            DataPseudonymous,
			ProductionPermitted: true,
			Retention:           RetentionNone,
			// Cloudflare does not train on or improve services with Customer
			// Content (verified 2026-10-16).
			Source: "https://developers.cloudflare.com/workers-ai/platform/data-usage/ (2026-10-16)",
		},
		Capabilities: Capabilities{ContextWindow: 32768, Models: []string{model}, JSON: true, Streaming: true},
	})
	if err != nil {
		return nil, false, err
	}
	return p, true, nil
}

// orderProviders reorders providers to match AI_PROVIDER_CHAIN (a comma
// separated list of ids). Providers named in the chain come first, in that
// order; the rest keep their original order after them. An unknown id in the
// chain is ignored rather than fatal: a typo must not down the feature.
func orderProviders(providers []Provider, chain string) []Provider {
	if strings.TrimSpace(chain) == "" || len(providers) == 0 {
		return providers
	}
	byID := make(map[string]Provider, len(providers))
	for _, p := range providers {
		byID[p.ID()] = p
	}
	used := make(map[string]bool, len(providers))
	out := make([]Provider, 0, len(providers))
	for _, id := range strings.Split(chain, ",") {
		id = strings.TrimSpace(id)
		if id == "" || used[id] {
			continue
		}
		if p, ok := byID[id]; ok {
			out = append(out, p)
			used[id] = true
		}
	}
	for _, p := range providers {
		if !used[p.ID()] {
			out = append(out, p)
		}
	}
	return out
}

// envBool parses a conventional truthy env var. Anything other than an explicit
// true value is false, so an unset or malformed flag cannot enable a feature.
func envBool(key string) bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(key))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}
