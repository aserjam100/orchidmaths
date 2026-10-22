package ai

import (
	"fmt"
	"strings"
)

// Registry is the ordered, policy-gated set of providers Orchid may call.
//
// It is built once at boot and never mutated, so a provider being reachable
// cannot change eligibility: only its declared Policy can, and only through
// Select. That is the difference between a registry and a hardcoded list — the
// AI caller never names a provider, and adding or removing one is a
// configuration change, not a code change to the caller.
type Registry struct {
	providers []Provider
	// production is true when the deployment serves real learner traffic. In
	// production, DevOnly providers and providers without ProductionPermitted
	// are never returned, so a preferred-provider failure cannot silently route
	// a request to an experimental endpoint.
	production bool
	// allowDev permits DevOnly providers (development builds only).
	allowDev bool
}

// Options configures a Registry's environment.
type Options struct {
	// Production marks a deployment that serves real learner traffic.
	Production bool
	// AllowDev permits DevOnly providers. It is ignored when Production is
	// true, so a development flag left set in production cannot widen the chain.
	AllowDev bool
	// ZDRConfirmed is the operator's assertion that zero-data retention is
	// switched on at every production-permitted provider that declares
	// RetentionNone.
	//
	// The declaration and the account setting are two different facts. Some
	// providers advertise zero retention but require an account-level switch
	// before it applies, so a policy that says RetentionNone is only true if
	// somebody enabled it. In production an unconfirmed zero-retention claim is
	// a configuration error: the code now trusts that claim in Select, and
	// trusting an unverified one is exactly the privacy promise broken by an
	// omission the Policy docs warn about.
	ZDRConfirmed bool
}

// NewRegistry validates a provider chain and returns it ready to Select from.
//
// An empty chain is valid and means AI is disabled; every method is nil-safe so
// callers need no special case.
func NewRegistry(providers []Provider, opts Options) (*Registry, error) {
	allowDev := opts.AllowDev && !opts.Production
	seen := make(map[string]bool, len(providers))
	for _, p := range providers {
		if p == nil {
			return nil, fmt.Errorf("ai: nil provider in chain")
		}
		id := p.ID()
		if strings.TrimSpace(id) == "" {
			return nil, fmt.Errorf("ai: provider with empty id")
		}
		if seen[id] {
			return nil, fmt.Errorf("ai: duplicate provider id %q", id)
		}
		seen[id] = true

		pol := p.Policy()
		if pol.MaxClass > DataPublic && pol.Source == "" {
			return nil, fmt.Errorf("ai: provider %q permits %s data but records no source", id, pol.MaxClass)
		}
		if pol.DevOnly && pol.ProductionPermitted {
			return nil, fmt.Errorf("ai: provider %q is both DevOnly and ProductionPermitted", id)
		}
		if opts.Production && pol.ProductionPermitted && pol.Retention == RetentionNone && !opts.ZDRConfirmed {
			return nil, fmt.Errorf("ai: provider %q claims zero retention but zero-data retention is unconfirmed; enable it in the provider's console, then set AI_ZDR_CONFIRMED=1", id)
		}
		// A dev-only provider with the dev flag off is kept but never Select-ed:
		// filtering is Select's job, so the same config boots in every
		// environment and only the eligible set differs.
	}
	return &Registry{providers: providers, production: opts.Production, allowDev: allowDev}, nil
}

// Enabled reports whether any provider is present. A nil registry is disabled,
// so callers can hold a nil *Registry without branching.
func (r *Registry) Enabled() bool { return r != nil && len(r.providers) > 0 }

// retentionPermits reports whether a provider's prompt-retention behaviour is
// acceptable for a request of the given class.
//
// Retention is a control, not a description. A provider that may train on
// prompts is never eligible for learner data, and "nobody checked" is treated as
// the worst case rather than the default: an unverified provider is denied
// anything that identifies a learner or their work. Only a provider with zero
// retention may receive identifiable data.
func retentionPermits(r Retention, class DataClass) bool {
	if class <= DataPublic {
		return true
	}
	switch r {
	case RetentionTrained, RetentionUnknown:
		return false
	case RetentionTemporary:
		return class < DataIdentified
	case RetentionNone:
		return true
	default:
		return false
	}
}

// Select returns the providers eligible for a request of the given class, in
// the configured order. An empty result means no provider may receive the
// request, and the caller must fail closed rather than downgrade to one that
// may not.
func (r *Registry) Select(class DataClass) []Provider {
	if r == nil {
		return nil
	}
	out := make([]Provider, 0, len(r.providers))
	for _, p := range r.providers {
		pol := p.Policy()
		if !pol.Permits(class) {
			continue
		}
		if !retentionPermits(pol.Retention, class) {
			continue
		}
		if pol.DevOnly && !r.allowDev {
			continue
		}
		if r.production && !pol.ProductionPermitted {
			continue
		}
		out = append(out, p)
	}
	return out
}

// IDs lists the configured provider ids in order. For boot logging only; it
// must never be written into an HTTP response.
func (r *Registry) IDs() []string {
	if r == nil {
		return nil
	}
	ids := make([]string, 0, len(r.providers))
	for _, p := range r.providers {
		ids = append(ids, p.ID())
	}
	return ids
}
