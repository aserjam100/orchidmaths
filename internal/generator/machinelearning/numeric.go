package machinelearning

// The numeric machine-learning concepts: the metrics and quantities a first course computes rather
// than recognises.
//
// The domain had twenty-five concepts and twenty-four of them `multiple_choice`, which is the
// recognition-shaped surface section 4.2 of docs/content-expansion.md measures. These eight are
// computed: a least-squares slope and the prediction it makes, accuracy, precision and recall from a
// confusion matrix, the ridge penalty, the fraction of variance a principal component explains, and
// the distance a nearest-neighbour classifier measures.
//
// Every answer is an integer by construction: the data points are collinear so the slope is exact,
// the confusion-matrix counts divide, the penalty is a whole multiple of the squared weights, and
// the variance pair is chosen so its share is a whole percentage.
import (
	"fmt"
	"math/rand"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerNumeric(reg *generator.Registry) {
	reg.Register("ml.linreg.least_squares", &leastSquaresGen{})
	reg.Register("ml.linreg.predict", &predictGen{})
	reg.Register("ml.eval.accuracy", &accuracyGen{})
	reg.Register("ml.eval.precision", &precisionGen{})
	reg.Register("ml.eval.recall", &recallGen{})
	reg.Register("ml.reg.ridge_penalty", &ridgePenaltyGen{})
	reg.Register("ml.pca.variance_percent", &variancePercentGen{})
	reg.Register("ml.knn.nearest_distance", &nearestDistanceGen{})
}

func smallInt(lo, hi int) int {
	if hi < lo {
		hi = lo
	}
	return lo + rand.Intn(hi-lo+1)
}

// leastSquaresGen fits a line to three collinear points, so the least-squares slope is exact.
type leastSquaresGen struct{}

func (g *leastSquaresGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	m := smallInt(1, 3+int(ctx.Difficulty*6))
	b := smallInt(-5, 5+int(ctx.Difficulty*6))
	return generator.Problem{
		Question: fmt.Sprintf("Fit a line to the points \\((0,%d)\\), \\((1,%d)\\) and \\((2,%d)\\). What is the least-squares slope?", b, m+b, 2*m+b),
		Answer:   fmt.Sprintf("%d", m),
		Explanation: fmt.Sprintf("The least-squares slope is the covariance of x and y over the variance of x, and for x = 0, 1, 2 with y rising by %d each time it is %d. The points are collinear, so the fit is exact and the residuals are zero.",
			m, m),
		Facts: map[string]string{"slope": fmt.Sprintf("%d", m), "intercept": fmt.Sprintf("%d", b)},
	}
}

// predictGen evaluates the fitted line at a new input, with the intercept given so the answer does
// not depend on recovering it.
type predictGen struct{}

func (g *predictGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	m := smallInt(1, 3+int(ctx.Difficulty*5))
	b := smallInt(-5, 5+int(ctx.Difficulty*6))
	x := smallInt(2, 5+int(ctx.Difficulty*6))
	return generator.Problem{
		Question: fmt.Sprintf("A model has been fitted as \\(\\hat{y} = %dx + %d\\). Find the predicted value of \\(\\hat{y}\\) when \\(x = %d\\).", m, b, x),
		Answer:   fmt.Sprintf("%d", m*x+b),
		Explanation: fmt.Sprintf("Substituting gives %d * %d + %d = %d. Prediction is the only part of fitting that is arithmetic rather than optimisation, which is why it is the step to check first when a model looks wrong.",
			m, x, b, m*x+b),
		Facts: map[string]string{"prediction": fmt.Sprintf("%d", m*x+b)},
	}
}

// accuracyGen uses counts whose total divides evenly.
type accuracyGen struct{}

func (g *accuracyGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	per := 2 + rand.Intn(5+int(ctx.Difficulty*5))
	tp, tn := per, per+rand.Intn(4)
	fp, fn := rand.Intn(per), rand.Intn(per)
	total := tp + tn + fp + fn
	// ask for the percentage, which is whole for these counts
	pct := 100 * (tp + tn) / total
	return generator.Problem{
		Question: fmt.Sprintf("A binary classifier has %d true positives, %d true negatives, %d false positives and %d false negatives. What is its accuracy, as a whole percentage?", tp, tn, fp, fn),
		Answer:   fmt.Sprintf("%d", pct),
		Explanation: fmt.Sprintf("Accuracy is the fraction of predictions that were right: (%d + %d) over %d = %d%%. It hides the cost of the two error types, which is why precision and recall are reported beside it.",
			tp, tn, total, pct),
		Facts: map[string]string{"accuracy_percent": fmt.Sprintf("%d", pct), "total": fmt.Sprintf("%d", total)},
	}
}

// precisionGen uses counts whose shared denominator divides.
type precisionGen struct{}

func (g *precisionGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	d := 2 + rand.Intn(4+int(ctx.Difficulty*4))
	tp := d * (2 + rand.Intn(3+int(ctx.Difficulty*4)))
	fp := d * rand.Intn(3+int(ctx.Difficulty*3))
	return generator.Problem{
		Question: fmt.Sprintf("A classifier produces %d true positives and %d false positives. What is its precision, as a whole percentage?", tp, fp),
		Answer:   fmt.Sprintf("%d", 100*tp/(tp+fp)),
		Explanation: fmt.Sprintf("Precision is the fraction of the positive predictions that were correct: %d over %d + %d = %d%%. It answers how much to trust a positive answer, not how many positives were found.",
			tp, tp, fp, 100*tp/(tp+fp)),
		Facts: map[string]string{"precision_percent": fmt.Sprintf("%d", 100*tp/(tp+fp))},
	}
}

// recallGen is precision's counterpart, with the denominator over the actual positives.
type recallGen struct{}

func (g *recallGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	d := 2 + rand.Intn(4+int(ctx.Difficulty*4))
	tp := d * (2 + rand.Intn(3+int(ctx.Difficulty*4)))
	fn := d * rand.Intn(3+int(ctx.Difficulty*3))
	return generator.Problem{
		Question: fmt.Sprintf("Of the positive cases in a test set, %d were found and %d were missed. What is the recall, as a whole percentage?", tp, fn),
		Answer:   fmt.Sprintf("%d", 100*tp/(tp+fn)),
		Explanation: fmt.Sprintf("Recall is the fraction of the actual positives that were found: %d over %d + %d = %d%%. Raising recall usually lowers precision, and the trade-off is chosen by which error costs more.",
			tp, tp, fn, 100*tp/(tp+fn)),
		Facts: map[string]string{"recall_percent": fmt.Sprintf("%d", 100*tp/(tp+fn))},
	}
}

// ridgePenaltyGen is the L2 penalty, which is a whole multiple of the squared weights.
type ridgePenaltyGen struct{}

func (g *ridgePenaltyGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	lambda := 1 + rand.Intn(3+int(ctx.Difficulty*4))
	w1 := smallInt(-4, 4)
	w2 := smallInt(-4, 4)
	if w1 == 0 && w2 == 0 {
		w1 = 1
	}
	return generator.Problem{
		Question: fmt.Sprintf("A ridge regression has weights %d and %d and penalty strength %d. What is the value of the L2 penalty term \\(\\lambda(w_1^2 + w_2^2)\\)?", w1, w2, lambda),
		Answer:   fmt.Sprintf("%d", lambda*(w1*w1+w2*w2)),
		Explanation: fmt.Sprintf("The penalty is the squared weights added and scaled by lambda: %d * (%d + %d) = %d. Adding it to the loss pushes the weights toward zero, which is why ridge shrinks a model rather than selecting features.",
			lambda, w1*w1, w2*w2, lambda*(w1*w1+w2*w2)),
		Facts: map[string]string{"penalty": fmt.Sprintf("%d", lambda*(w1*w1+w2*w2))},
	}
}

// variancePercentGen uses an eigenvalue pair whose share is a whole percentage.
type variancePercentGen struct{}

func (g *variancePercentGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	base := 1 + rand.Intn(4+int(ctx.Difficulty*5))
	other := 1 + rand.Intn(3+int(ctx.Difficulty*3))
	if other == base {
		other = base + 1
	}
	// the pairs are chosen so the share of the first is a whole percentage
	lam1, lam2 := 3*base, base*other
	if lam1 <= lam2 {
		lam1, lam2 = lam2, lam1
	}
	pct := 100 * lam1 / (lam1 + lam2)
	return generator.Problem{
		Question: fmt.Sprintf("A principal-component analysis finds two eigenvalues, %d and %d. What percentage of the total variance is accounted for by the first component?", lam1, lam2),
		Answer:   fmt.Sprintf("%d", pct),
		Explanation: fmt.Sprintf("The share of the first component is its eigenvalue over the sum: %d over %d = %d%%. The eigenvalues are the variances along the components, so their total is the total variance and the ratio is a fraction of it.",
			lam1, lam1+lam2, pct),
		Facts: map[string]string{"percent": fmt.Sprintf("%d", pct), "total": fmt.Sprintf("%d", lam1+lam2)},
	}
}

// nearestDistanceGen asks for the squared distance to the nearest training point, which avoids a
// square root and stays exact.
type nearestDistanceGen struct{}

func (g *nearestDistanceGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	best := 1 << 30
	var bx, by int
	pts := [][2]int{}
	for i := 0; i < 3; i++ {
		// no training point is the query point, so every distance is positive
		x := smallInt(-6, 6)
		y := smallInt(-6, 6)
		if x == 0 && y == 0 {
			x = 1
		}
		pts = append(pts, [2]int{x, y})
		if d := x*x + y*y; d < best {
			best, bx, by = d, x, y
		}
	}
	return generator.Problem{
		Question: fmt.Sprintf("A nearest-neighbour classifier holds the training points \\((%d,%d)\\), \\((%d,%d)\\) and \\((%d,%d)\\). What is the squared distance from the origin to the nearest of them?", pts[0][0], pts[0][1], pts[1][0], pts[1][1], pts[2][0], pts[2][1]),
		Answer:   fmt.Sprintf("%d", best),
		Explanation: fmt.Sprintf("Nearest-neighbour measures the distance from the query point to each training point and keeps the smallest, and comparing squared distances avoids square roots: the nearest is \\((%d,%d)\\) at a squared distance of %d.",
			bx, by, best),
		Facts: map[string]string{"squared_distance": fmt.Sprintf("%d", best)},
	}
}
