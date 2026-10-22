// Package ai is Orchid's provider registry for optional AI augmentation.
//
// Nothing here is on the learning critical path: grading, scheduling, mastery
// and the onboarding fallback are deterministic and stay so. This package only
// decides which inference backends may be called, with what data, and in what
// order — so a provider being reachable can never by itself make it eligible.
package ai

// DataClass is how sensitive a request's payload is, ordered least to most
// sensitive. The ordering is the whole point: a provider permitted to receive a
// more sensitive class may also receive anything less sensitive, and never the
// reverse. Eligibility is therefore one comparison and cannot be read two ways.
type DataClass int

const (
	// DataPublic is public or synthetic content: course text, a generated
	// question, or a learner's typed goal with no identifier attached.
	DataPublic DataClass = iota
	// DataPseudonymous is learning context with identifiers removed: concept
	// ids, mastery states, attempt shapes. The learner cannot be named from it.
	DataPseudonymous
	// DataIdentified is identifiable learner data: names, emails, student ids,
	// or free text a learner wrote.
	DataIdentified
)

func (d DataClass) String() string {
	switch d {
	case DataPublic:
		return "public"
	case DataPseudonymous:
		return "pseudonymous"
	case DataIdentified:
		return "identified"
	}
	return "unknown"
}

// Retention states how a provider treats prompts after the request completes.
type Retention int

const (
	// RetentionUnknown means nobody has established it. Unknown is treated as
	// the worst case: a provider whose retention is unverified is not eligible
	// for anything above public content. Defaulting to "trusted" is how a
	// privacy promise gets broken by an omission.
	RetentionUnknown Retention = iota
	// RetentionNone means zero data retention: prompts are not stored. Some
	// providers require an explicit account setting to reach this; the Source
	// field must name that the setting is enabled, not merely available.
	RetentionNone
	// RetentionTemporary means prompts are retained for a bounded operational
	// window (for example abuse monitoring) and then discarded.
	RetentionTemporary
	// RetentionTrained means prompts may be used to train models or improve the
	// provider's products.
	RetentionTrained
)

func (r Retention) String() string {
	switch r {
	case RetentionNone:
		return "none"
	case RetentionTemporary:
		return "temporary"
	case RetentionTrained:
		return "trained"
	}
	return "unknown"
}

// Policy is what a provider's terms permit.
//
// Every field defaults to the deny value, so a provider nobody has documented
// is ineligible until someone writes the policy down. The type deliberately has
// no single "safe"/"trial" string: that collapses independent facts — who may
// see data, whether the terms allow production, whether prompts may train a
// model — into one value that can only ever be wrong in one direction.
type Policy struct {
	// MaxClass is the most sensitive DataClass this provider may receive.
	MaxClass DataClass
	// ProductionPermitted is whether the terms allow a public application's
	// real learner traffic. False for evaluation, preview or promotional tiers.
	ProductionPermitted bool
	// DevOnly restricts the provider to development and synthetic content. It is
	// never eligible in production, even under failover.
	DevOnly bool
	// Retention is the provider's prompt-retention behaviour.
	Retention Retention
	// Source records where the policy was established — a URL and the date it
	// was verified — or is empty when unverified. A provider that permits more
	// than public data with no Source is a configuration error the registry
	// rejects, so an undocumented policy cannot reach a learner.
	Source string
}

// Permits reports whether the policy allows a request of the given class,
// ignoring the dev/production split (which the Registry applies on top).
func (p Policy) Permits(class DataClass) bool {
	return class <= p.MaxClass
}
