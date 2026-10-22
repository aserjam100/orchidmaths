package ai

import "testing"

// clearAIEnv blanks every environment variable FromEnv reads, so a variable set
// in the developer's shell or CI cannot change a test's outcome.
func clearAIEnv(t *testing.T) {
	t.Helper()
	for _, k := range []string{
		"AI_ENABLED", "AI_ENV", "AI_DEV_PROVIDERS", "AI_DAILY_USER_CAP", "AI_PROVIDER_CHAIN",
		"AI_GLOBAL_DAILY_CAP", "AI_ASSISTANT_MAX_TOKENS", "AI_GOAL_MAX_TOKENS",
		"AI_ZDR_CONFIRMED", "AI_TRUSTED_IP_HEADER",
		"GROQ_API_KEY", "GROQ_BASE_URL", "GROQ_MODEL", "GROQ_REASONING_EFFORT",
		"CLOUDFLARE_API_TOKEN", "CLOUDFLARE_ACCOUNT_ID", "CLOUDFLARE_MODEL",
		"NVIDIA_API_KEY", "NVIDIA_BASE_URL", "NVIDIA_MODEL",
		"CEREBRAS_API_KEY", "CEREBRAS_BASE_URL", "CEREBRAS_MODEL",
		"OPENROUTER_API_KEY", "OPENROUTER_BASE_URL", "OPENROUTER_MODEL",
		"OPENCODE_API_KEY", "OPENCODE_BASE_URL", "OPENCODE_MODEL",
	} {
		t.Setenv(k, "")
	}
}

func TestFromEnvDisabledByDefault(t *testing.T) {
	clearAIEnv(t)
	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if s.Registry.Enabled() {
		t.Error("AI must be disabled unless AI_ENABLED is set")
	}
	if s.DailyUserCap != DefaultDailyUserCap {
		t.Errorf("cap = %d, want %d", s.DailyUserCap, DefaultDailyUserCap)
	}
}

func TestFromEnvGroqIsProductionPseudonymous(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_ENV", "production")
	t.Setenv("AI_ZDR_CONFIRMED", "1")
	t.Setenv("GROQ_API_KEY", "k")

	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if !s.Registry.Enabled() {
		t.Fatal("groq should be enabled")
	}
	if got := s.Registry.Select(DataPseudonymous); len(got) != 1 || got[0].ID() != "groq" {
		t.Errorf("pseudonymous select = %v", got)
	}
	// Groq is not cleared for identifiable learner data.
	if got := s.Registry.Select(DataIdentified); len(got) != 0 {
		t.Errorf("identified select = %v, want none", got)
	}
}

func TestFromEnvDailyCapParsed(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_DAILY_USER_CAP", "7")
	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if s.DailyUserCap != 7 {
		t.Errorf("cap = %d, want 7", s.DailyUserCap)
	}
}

func TestFromEnvDailyCapRejectsGarbage(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_DAILY_USER_CAP", "-3")
	if _, err := FromEnv(); err == nil {
		t.Error("want error for negative cap")
	}
}

func TestFromEnvChainOrdering(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_ENV", "production")
	t.Setenv("AI_ZDR_CONFIRMED", "1")
	t.Setenv("GROQ_API_KEY", "k")
	t.Setenv("CLOUDFLARE_API_TOKEN", "t")
	t.Setenv("CLOUDFLARE_ACCOUNT_ID", "acct")

	// Default order: groq then cloudflare.
	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	got := ids(s.Registry.Select(DataPseudonymous))
	if !eq(got, []string{"groq", "cloudflare"}) {
		t.Fatalf("default order = %v", got)
	}

	// Reversed by the chain env.
	t.Setenv("AI_PROVIDER_CHAIN", "cloudflare,groq")
	s, err = FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	got = ids(s.Registry.Select(DataPseudonymous))
	if !eq(got, []string{"cloudflare", "groq"}) {
		t.Fatalf("chain order = %v", got)
	}
}

func TestFromEnvDevProviderNeedsModel(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_DEV_PROVIDERS", "1")
	t.Setenv("NVIDIA_API_KEY", "k")
	// No NVIDIA_MODEL: the provider must be skipped, not given an invented id.
	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if s.Registry.Enabled() {
		t.Errorf("dev provider without a model must be skipped, got %v", s.Registry.IDs())
	}

	// With a model and the dev flag it becomes selectable in development only.
	t.Setenv("NVIDIA_MODEL", "meta/llama")
	s, err = FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if got := ids(s.Registry.Select(DataPublic)); !eq(got, []string{"nvidia"}) {
		t.Errorf("dev select = %v", got)
	}
}

func TestFromEnvDevProviderInertInProduction(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_ENV", "production")
	t.Setenv("AI_ZDR_CONFIRMED", "1")
	t.Setenv("AI_DEV_PROVIDERS", "1")
	t.Setenv("NVIDIA_API_KEY", "k")
	t.Setenv("NVIDIA_MODEL", "meta/llama")
	t.Setenv("GROQ_API_KEY", "k")

	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if got := ids(s.Registry.Select(DataPublic)); !eq(got, []string{"groq"}) {
		t.Errorf("production must exclude dev-only providers, got %v", got)
	}
}

// TestFromEnvRefusesUnconfirmedZDRInProduction pins the graceful half of the
// gate: an unconfirmed zero-retention claim disables AI and records why, rather
// than stopping the process. The deterministic app must still boot.
func TestFromEnvRefusesUnconfirmedZDRInProduction(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_ENV", "production")
	t.Setenv("GROQ_API_KEY", "k")

	s, err := FromEnv()
	if err != nil {
		t.Fatalf("an AI configuration error must not fail the process: %v", err)
	}
	if s.Registry != nil && s.Registry.Enabled() {
		t.Fatalf("AI must be off while zero-data retention is unconfirmed, got %v", s.Registry.IDs())
	}
	if s.DisabledReason == "" {
		t.Error("the reason must be reported, not swallowed")
	}
}

// TestFromEnvGlobalAndTokenBudgets pins the other new knobs, including the
// defaults that bound a shared free tier.
func TestFromEnvGlobalAndTokenBudgets(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")

	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if s.GlobalDailyCap != DefaultGlobalDailyCap {
		t.Errorf("global cap = %d, want %d", s.GlobalDailyCap, DefaultGlobalDailyCap)
	}
	if s.AssistantMaxTokens != DefaultAssistantMaxTokens || s.GoalMaxTokens != DefaultGoalMaxTokens {
		t.Errorf("token budgets = %d/%d, want %d/%d",
			s.AssistantMaxTokens, s.GoalMaxTokens, DefaultAssistantMaxTokens, DefaultGoalMaxTokens)
	}

	t.Setenv("AI_GLOBAL_DAILY_CAP", "42")
	t.Setenv("AI_ASSISTANT_MAX_TOKENS", "900")
	t.Setenv("AI_GOAL_MAX_TOKENS", "300")
	s, err = FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if s.GlobalDailyCap != 42 || s.AssistantMaxTokens != 900 || s.GoalMaxTokens != 300 {
		t.Errorf("parsed budgets = %d/%d/%d, want 42/900/300", s.GlobalDailyCap, s.AssistantMaxTokens, s.GoalMaxTokens)
	}

	// A zero token budget is nonsense and is rejected rather than silently
	// truncating every reply to nothing.
	t.Setenv("AI_ASSISTANT_MAX_TOKENS", "0")
	if _, err := FromEnv(); err == nil {
		t.Error("want error for a zero token budget")
	}
}

// TestCloudflareDefaultModelIsTheProvenOne pins the fallback model id.
//
// The id was a documented guess until 2026-10-20, when the real leg failed
// with failovers=1 while Groq covered for it: the account serves Llama 3.1 8B
// quantised to FP8, under a different id. A guess that 404s is invisible while
// the first provider is healthy, and leaving the working id only in a Fly
// secret means a fresh deployment silently reintroduces the dead leg.
func TestCloudflareDefaultModelIsTheProvenOne(t *testing.T) {
	clearAIEnv(t)
	t.Setenv("AI_ENABLED", "1")
	t.Setenv("AI_ENV", "production")
	t.Setenv("AI_ZDR_CONFIRMED", "1")
	t.Setenv("CLOUDFLARE_API_TOKEN", "t")
	t.Setenv("CLOUDFLARE_ACCOUNT_ID", "acct")

	s, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	var caps Capabilities
	var found bool
	for _, p := range s.Registry.Select(DataPseudonymous) {
		if p.ID() == "cloudflare" {
			caps = p.Capabilities()
			found = true
		}
	}
	if !found {
		t.Fatal("cloudflare provider missing from the registry")
	}
	if len(caps.Models) != 1 || caps.Models[0] != "@cf/meta/llama-3.1-8b-instruct-fp8" {
		t.Errorf("default model = %v, want the fp8 id proven on 2026-10-20", caps.Models)
	}
	if caps.ContextWindow != 32768 {
		t.Errorf("context window = %d, want 32768 (the model's published size)", caps.ContextWindow)
	}
}
