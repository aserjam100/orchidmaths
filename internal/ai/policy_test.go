package ai

import "testing"

func TestPolicyPermitsIsOrdered(t *testing.T) {
	cases := []struct {
		max   DataClass
		class DataClass
		want  bool
	}{
		{DataPublic, DataPublic, true},
		{DataPublic, DataPseudonymous, false},
		{DataPublic, DataIdentified, false},
		{DataPseudonymous, DataPublic, true},
		{DataPseudonymous, DataPseudonymous, true},
		{DataPseudonymous, DataIdentified, false},
		{DataIdentified, DataPublic, true},
		{DataIdentified, DataPseudonymous, true},
		{DataIdentified, DataIdentified, true},
	}
	for _, c := range cases {
		p := Policy{MaxClass: c.max}
		if got := p.Permits(c.class); got != c.want {
			t.Errorf("Policy{MaxClass:%s}.Permits(%s) = %v, want %v", c.max, c.class, got, c.want)
		}
	}
}

func TestPolicyZeroValueDeniesLearnerData(t *testing.T) {
	var p Policy
	if !p.Permits(DataPublic) {
		t.Error("zero policy should permit public/synthetic content")
	}
	if p.Permits(DataPseudonymous) || p.Permits(DataIdentified) {
		t.Error("zero policy must deny pseudonymous and identified learner data")
	}
	if p.ProductionPermitted {
		t.Error("zero policy must not be production-permitted")
	}
}

func TestRetentionString(t *testing.T) {
	cases := map[Retention]string{
		RetentionUnknown:   "unknown",
		RetentionNone:      "none",
		RetentionTemporary: "temporary",
		RetentionTrained:   "trained",
		Retention(99):      "unknown",
	}
	for r, want := range cases {
		if got := r.String(); got != want {
			t.Errorf("Retention(%d).String() = %q, want %q", r, got, want)
		}
	}
}

func TestDataClassString(t *testing.T) {
	cases := map[DataClass]string{
		DataPublic:       "public",
		DataPseudonymous: "pseudonymous",
		DataIdentified:   "identified",
		DataClass(99):    "unknown",
	}
	for d, want := range cases {
		if got := d.String(); got != want {
			t.Errorf("DataClass(%d).String() = %q, want %q", d, got, want)
		}
	}
}
