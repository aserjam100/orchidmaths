package machinelearning

// Tests for the ten converted concepts.
//
// Each recomputes the answer from the published facts instead of trusting the
// generator's arithmetic, and each is checked against its shipped grading type
// so a corpus row that drifts back to multiple choice fails here.

import (
	"strconv"
	"strings"
	"testing"
)

var mlConvertedIDs = []string{
	"ml.loss.mse",
	"ml.perceptron",
	"ml.forward_pass",
	"ml.gradient_descent",
	"ml.regularization",
	"ml.validation",
	"ml.data.onehot",
	"ml.arch.cnn",
	"ml.eval.calibration",
	"ml.optim.momentum",
}

func mlConvInt(t *testing.T, facts map[string]string, key string) int {
	t.Helper()
	raw, ok := facts[key]
	if !ok {
		t.Fatalf("missing fact %q", key)
	}
	v, err := strconv.Atoi(raw)
	if err != nil {
		t.Fatalf("fact %q is not an integer: %q", key, raw)
	}
	return v
}

func mlConvInts(t *testing.T, facts map[string]string, key string) []int {
	t.Helper()
	raw, ok := facts[key]
	if !ok {
		t.Fatalf("missing fact %q", key)
	}
	parts := strings.Split(raw, ",")
	out := make([]int, len(parts))
	for i, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			t.Fatalf("fact %q has a non-integer part %q", key, p)
		}
		out[i] = v
	}
	return out
}

// mlFactBin reads a calibration fact such as "20 examples, 60 accuracy, 80 confidence".
func mlConvBin(t *testing.T, facts map[string]string, key string) (int, int, int) {
	t.Helper()
	raw, ok := facts[key]
	if !ok {
		t.Fatalf("missing fact %q", key)
	}
	fields := strings.Fields(raw)
	if len(fields) != 6 {
		t.Fatalf("fact %q has unexpected shape: %q", key, raw)
	}
	n, _ := strconv.Atoi(fields[0])
	acc, _ := strconv.Atoi(fields[2])
	conf, _ := strconv.Atoi(fields[4])
	return n, acc, conf
}

func TestMLConvertedConceptsAreNumericInTheCorpus(t *testing.T) {
	types := mlShippedTypes(t)
	for _, id := range mlConvertedIDs {
		if types[id] != "numeric" {
			t.Errorf("%s: corpus says %q, want numeric", id, types[id])
		}
	}
}

func TestMLConvertedAnswersRecomputeFromFacts(t *testing.T) {
	reg := mlRegistry(t)

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.loss.mse", i)
		preds := mlConvInts(t, p.Facts, "predictions")
		targets := mlConvInts(t, p.Facts, "targets")
		sum := 0
		for j := range preds {
			d := preds[j] - targets[j]
			sum += d * d
		}
		want := sum / len(preds)
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("mse: preds=%v targets=%v answer=%s want %d", preds, targets, p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.perceptron", i)
		w := mlConvInts(t, p.Facts, "weights")
		b := mlConvInt(t, p.Facts, "bias")
		positives := 0
		for _, pt := range [][2]int{{1, 1}, {2, 1}, {1, 2}} {
			if w[0]*pt[0]+w[1]*pt[1]+b > 0 {
				positives++
			}
		}
		if p.Answer != strconv.Itoa(positives) {
			t.Fatalf("perceptron: w=%v b=%d answer=%s want %d", w, b, p.Answer, positives)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.forward_pass", i)
		w := mlConvInts(t, p.Facts, "weights")
		x := mlConvInts(t, p.Facts, "inputs")
		b := mlConvInt(t, p.Facts, "bias")
		want := w[0]*x[0] + w[1]*x[1] + b
		if want < 0 {
			want = 0
		}
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("forward pass: w=%v x=%v b=%d answer=%s want %d", w, x, b, p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.gradient_descent", i)
		want := mlConvInt(t, p.Facts, "weight") - mlConvInt(t, p.Facts, "learning_rate")*mlConvInt(t, p.Facts, "gradient")
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("gradient descent: answer=%s want %d", p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.regularization", i)
		w := mlConvInts(t, p.Facts, "weights")
		want := mlConvInt(t, p.Facts, "lambda") * (w[0]*w[0] + w[1]*w[1])
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("regularization: w=%v answer=%s want %d", w, p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.validation", i)
		samples := mlConvInt(t, p.Facts, "samples")
		folds := mlConvInt(t, p.Facts, "folds")
		if samples%folds != 0 {
			t.Fatalf("validation: %d samples do not divide into %d folds", samples, folds)
		}
		if want := samples / folds; p.Answer != strconv.Itoa(want) {
			t.Fatalf("validation: samples=%d folds=%d answer=%s want %d", samples, folds, p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.data.onehot", i)
		if want := mlConvInt(t, p.Facts, "categories") - 1; p.Answer != strconv.Itoa(want) {
			t.Fatalf("one-hot: answer=%s want %d", p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.arch.cnn", i)
		k := mlConvInt(t, p.Facts, "kernel")
		want := (k*k*mlConvInt(t, p.Facts, "in_channels") + 1) * mlConvInt(t, p.Facts, "out_channels")
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("cnn: kernel=%d answer=%s want %d", k, p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.eval.calibration", i)
		n1, acc1, conf1 := mlConvBin(t, p.Facts, "bin_one")
		n2, acc2, conf2 := mlConvBin(t, p.Facts, "bin_two")
		want := (n1*absInt(acc1-conf1) + n2*absInt(acc2-conf2)) / (n1 + n2)
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("calibration: bins=(%d,%d,%d),(%d,%d,%d) answer=%s want %d", n1, acc1, conf1, n2, acc2, conf2, p.Answer, want)
		}
	}

	for i := 0; i < 100; i++ {
		p := mlSample(t, reg, "ml.optim.momentum", i)
		velocity := mlConvInt(t, p.Facts, "velocity")
		nextVelocity := velocity/2 + mlConvInt(t, p.Facts, "gradient")
		want := mlConvInt(t, p.Facts, "weight") - mlConvInt(t, p.Facts, "learning_rate")*nextVelocity
		if p.Answer != strconv.Itoa(want) {
			t.Fatalf("momentum: answer=%s want %d", p.Answer, want)
		}
	}
}
