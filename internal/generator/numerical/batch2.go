package numerical

// The second numerical-analysis batch: the twenty concepts that close the
// domain with the rest of interpolation and approximation, numerical
// integration, and numerical linear algebra.
//
// The first batch measured error, found roots and started interpolating. This
// batch finishes the reconstruction of a function from data (least squares, the
// normal equations, orthogonal polynomials and the dangers of extrapolation),
// replaces the exact integral with a weighted sum of function values
// (trapezoid, midpoint, Simpson, composite rules, Romberg, Gaussian and
// adaptive quadrature, and improper integrals), and solves the linear system
// (Gaussian elimination, pivoting, LU, Jacobi, Gauss-Seidel, the power method
// and least squares by QR).
//
// Nine concepts are multiple choice and eleven are numeric. Every numeric
// generator publishes the parameters its answer follows from, so the tests
// recompute the answer rather than trusting the author's arithmetic.

import (
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/aserjam100/orchid/internal/generator"
)

// ---------------------------------------------------------------------------
// Least squares (numeric)
// ---------------------------------------------------------------------------

type leastSquaresGen struct{}

func (g *leastSquaresGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		xs, ys []int
		reason string
	}
	claims := []claim{
		{xs: []int{0, 1, 2}, ys: []int{1, 3, 4}, reason: "The centred sums give sxy = 3 and sxx = 2, so the slope is 1.5."},
		{xs: []int{0, 1, 2}, ys: []int{1, 2, 6}, reason: "sxy = 5 and sxx = 2, so the slope is 2.5."},
		{xs: []int{0, 1, 2}, ys: []int{2, 3, 3}, reason: "sxy = 1 and sxx = 2, so the slope is 0.5."},
		{xs: []int{1, 2, 3}, ys: []int{1, 2, 4}, reason: "sxy = 3 and sxx = 2, so the slope is 1.5."},
		{xs: []int{0, 1, 2}, ys: []int{3, 4, 7}, reason: "sxy = 4 and sxx = 2, so the slope is 2."},
		{xs: []int{0, 1, 2, 3}, ys: []int{0, 1, 3, 4}, reason: "sxy = 7 and sxx = 5, so the slope is 1.4."},
	}
	c := claims[rand.Intn(len(claims))]
	n := float64(len(c.xs))
	var sx, sy float64
	for i := range c.xs {
		sx += float64(c.xs[i])
		sy += float64(c.ys[i])
	}
	xbar, ybar := sx/n, sy/n
	var sxx, sxy float64
	for i := range c.xs {
		dx := float64(c.xs[i]) - xbar
		sxx += dx * dx
		sxy += dx * (float64(c.ys[i]) - ybar)
	}
	answer := decimal6(sxy / sxx)
	xs := make([]string, len(c.xs))
	ys := make([]string, len(c.ys))
	for i := range c.xs {
		xs[i] = strconv.Itoa(c.xs[i])
		ys[i] = strconv.Itoa(c.ys[i])
	}
	return generator.Problem{
		Question: "The least-squares line y = a + b x through the points with x values " + strings.Join(xs, ", ") +
			" and y values " + strings.Join(ys, ", ") + " has slope b equal to the centred sum of x times y divided by the centred sum of x squared. What is b? (enter a decimal)",
		Answer:      answer,
		Explanation: "Subtracting the means removes the intercept from the slope, so the numerator and denominator are the centred cross-sum and the centred sum of squares; " + c.reason,
		Facts: map[string]string{
			"xs": strings.Join(xs, ","), "ys": strings.Join(ys, ","), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The normal equations (numeric)
// ---------------------------------------------------------------------------

type normalEquationsGen struct{}

func (g *normalEquationsGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		xs, ys []int
		reason string
	}
	claims := []claim{
		{xs: []int{0, 1, 2}, ys: []int{1, 4, 4}, reason: "b = 1.5, so a = 3 − 1.5·1 = 1.5."},
		{xs: []int{0, 1, 2}, ys: []int{0, 3, 3}, reason: "b = 1.5, so a = 2 − 1.5·1 = 0.5."},
		{xs: []int{0, 1, 2}, ys: []int{2, 5, 5}, reason: "b = 1.5, so a = 4 − 1.5·1 = 2.5."},
		{xs: []int{0, 1, 2}, ys: []int{4, 4, 7}, reason: "b = 1.5, so a = 5 − 1.5·1 = 3.5."},
		{xs: []int{1, 2, 3}, ys: []int{2, 3, 4}, reason: "b = 1, so a = 3 − 1·2 = 1."},
		{xs: []int{0, 1, 2}, ys: []int{3, 4, 5}, reason: "b = 1, so a = 4 − 1·1 = 3."},
	}
	c := claims[rand.Intn(len(claims))]
	n := float64(len(c.xs))
	var sx, sy float64
	for i := range c.xs {
		sx += float64(c.xs[i])
		sy += float64(c.ys[i])
	}
	xbar, ybar := sx/n, sy/n
	var sxx, sxy float64
	for i := range c.xs {
		dx := float64(c.xs[i]) - xbar
		sxx += dx * dx
		sxy += dx * (float64(c.ys[i]) - ybar)
	}
	b := sxy / sxx
	answer := decimal6(ybar - b*xbar)
	xs := make([]string, len(c.xs))
	ys := make([]string, len(c.ys))
	for i := range c.xs {
		xs[i] = strconv.Itoa(c.xs[i])
		ys[i] = strconv.Itoa(c.ys[i])
	}
	return generator.Problem{
		Question: "The normal equations for the least-squares line y = a + b x through the points with x values " + strings.Join(xs, ", ") +
			" and y values " + strings.Join(ys, ", ") + " give the intercept a = mean(y) − b·mean(x). What is a? (enter a decimal)",
		Answer:      answer,
		Explanation: "The intercept is the one that makes the line pass through the centre of mass of the data once the slope is known; " + c.reason,
		Facts: map[string]string{
			"xs": strings.Join(xs, ","), "ys": strings.Join(ys, ","), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Orthogonal polynomials (multiple choice)
// ---------------------------------------------------------------------------

type orthogonalGen struct{}

func (g *orthogonalGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Least squares is easier with orthogonal basis functions because the normal equations...",
			correct:     "become diagonal, so each coefficient can be computed on its own",
			distractors: []string{"become singular", "need no data", "give the exact fit with one term"},
			reason:      "Orthogonality kills the off-diagonal inner products that couple the unknowns.",
		},
		{
			question:    "With an orthogonal basis, the coefficient of each basis function is...",
			correct:     "the inner product of the data with that function divided by the squared norm of the function",
			distractors: []string{"the sum of the residuals", "the derivative of the function", "the largest data value"},
			reason:      "The diagonal normal equation gives exactly that quotient.",
		},
		{
			question:    "Legendre polynomials are orthogonal with respect to...",
			correct:     "the unweighted integral on the interval from minus one to one",
			distractors: []string{"a discrete set of points only", "the Gaussian weight", "the unit circle"},
			reason:      "Their weight function is the constant one on the standard interval.",
		},
		{
			question:    "Chebyshev polynomials are orthogonal with respect to the weight...",
			correct:     "one over the square root of one minus x squared",
			distractors: []string{"one over one plus x squared", "a constant weight", "the exponential weight"},
			reason:      "The reciprocal square-root weight is what makes the cosine substitution an isometry.",
		},
		{
			question:    "Using the successive powers of x as a basis is a poor choice because...",
			correct:     "the powers are highly correlated, so the normal equations are ill conditioned",
			distractors: []string{"they are not functions", "they cannot be evaluated", "they have no coefficients"},
			reason:      "On a short interval the monomials look alike, which is exactly the ill conditioning orthogonal bases avoid.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Extrapolation (multiple choice)
// ---------------------------------------------------------------------------

type extrapolationGen struct{}

func (g *extrapolationGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Extrapolation means estimating the function...",
			correct:     "outside the range covered by the data",
			distractors: []string{"between the data points", "by fitting a line", "by removing noise"},
			reason:      "It is interpolation taken beyond the last node, where nothing constrains the fit.",
		},
		{
			question:    "Extrapolation is risky because...",
			correct:     "the interpolant is unconstrained beyond the data, so its error can grow quickly",
			distractors: []string{"the data are always wrong", "the interpolant does not exist outside the nodes", "there are never enough points"},
			reason:      "The node product grows fast outside the interval, and the derivative term is unbounded.",
		},
		{
			question:    "The farther the extrapolation point is from the data, the...",
			correct:     "larger the error bound generally becomes",
			distractors: []string{"smaller the error becomes", "more exact the answer becomes", "less the derivatives matter"},
			reason:      "The error bound is proportional to those very distances.",
		},
		{
			question:    "Richardson extrapolation is different because it...",
			correct:     "combines a numerical sequence using its known error expansion, rather than continuing a fitted curve",
			distractors: []string{"fits a polynomial beyond the data", "ignores the error terms", "extrapolates at random"},
			reason:      "It extrapolates the limit of a sequence, not a function off the end of its data.",
		},
		{
			question:    "When asked for a value far outside the data, a careful response is to...",
			correct:     "report the estimate with a much wider uncertainty",
			distractors: []string{"report it as exact", "refuse under all circumstances", "quietly double the data"},
			reason:      "The prediction may still be useful, but its uncertainty grows with the distance.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// The trapezoid rule (numeric)
// ---------------------------------------------------------------------------

type trapezoidGen struct{}

func (g *trapezoidGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b   int
		reason string
	}
	claims := []claim{
		{a: 0, b: 1, reason: "(1 − 0)(0 + 1)/2 = 0.5."},
		{a: 1, b: 2, reason: "(2 − 1)(1 + 4)/2 = 2.5."},
		{a: 0, b: 2, reason: "(2 − 0)(0 + 4)/2 = 4."},
		{a: 1, b: 3, reason: "(3 − 1)(1 + 9)/2 = 10."},
		{a: 2, b: 3, reason: "(3 − 2)(4 + 9)/2 = 6.5."},
		{a: 0, b: 3, reason: "(3 − 0)(0 + 9)/2 = 13.5."},
	}
	c := claims[rand.Intn(len(claims))]
	f := func(x int) int { return x * x }
	answer := decimal6(float64(c.b-c.a) * float64(f(c.a)+f(c.b)) / 2)
	return generator.Problem{
		Question: "The trapezoid rule estimates the integral of f over the interval from a to b by (b − a)(f(a) + f(b))/2. For f(x) = x², a = " +
			strconv.Itoa(c.a) + " and b = " + strconv.Itoa(c.b) + ", what is the estimate? (enter a decimal)",
		Answer:      answer,
		Explanation: "The rule replaces the curve by the straight line joining the endpoints and takes the area of the resulting trapezoid; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The midpoint rule (numeric)
// ---------------------------------------------------------------------------

type midpointGen struct{}

func (g *midpointGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b   int
		reason string
	}
	claims := []claim{
		{a: 0, b: 1, reason: "(1 − 0)·(0.5)² = 0.25."},
		{a: 0, b: 2, reason: "(2 − 0)·1² = 2."},
		{a: 1, b: 3, reason: "(3 − 1)·2² = 8."},
		{a: 2, b: 4, reason: "(4 − 2)·3² = 18."},
		{a: 0, b: 4, reason: "(4 − 0)·2² = 16."},
		{a: 1, b: 2, reason: "(2 − 1)·1.5² = 2.25."},
	}
	c := claims[rand.Intn(len(claims))]
	m := float64(c.a+c.b) / 2
	answer := decimal6(float64(c.b-c.a) * m * m)
	return generator.Problem{
		Question: "The midpoint rule estimates the integral of f over the interval from a to b by (b − a)·f((a + b)/2). For f(x) = x², a = " +
			strconv.Itoa(c.a) + " and b = " + strconv.Itoa(c.b) + ", what is the estimate? (enter a decimal)",
		Answer:      answer,
		Explanation: "The rule replaces the curve by the horizontal line at the midpoint value and takes the area of the rectangle; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Simpson's rule (numeric)
// ---------------------------------------------------------------------------

type simpsonGen struct{}

func (g *simpsonGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b   int
		reason string
	}
	claims := []claim{
		{a: 0, b: 1, reason: "(1/6)(0 + 4·0.25 + 1) = 0.333333."},
		{a: 0, b: 2, reason: "(2/6)(0 + 4·1 + 4) = 2.666667."},
		{a: 1, b: 2, reason: "(1/6)(1 + 9 + 4) = 2.333333."},
		{a: 0, b: 3, reason: "(3/6)(0 + 9 + 9) = 9."},
		{a: 2, b: 3, reason: "(1/6)(4 + 25 + 9) = 6.333333."},
		{a: 1, b: 3, reason: "(2/6)(1 + 16 + 9) = 8.666667."},
	}
	c := claims[rand.Intn(len(claims))]
	f := func(x float64) float64 { return x * x }
	m := float64(c.a+c.b) / 2
	answer := decimal6((float64(c.b-c.a) / 6) * (f(float64(c.a)) + 4*f(m) + f(float64(c.b))))
	return generator.Problem{
		Question: "Simpson's rule estimates the integral of f over the interval from a to b by (b − a)(f(a) + 4 f(mid) + f(b))/6. For f(x) = x², a = " +
			strconv.Itoa(c.a) + " and b = " + strconv.Itoa(c.b) + ", what is the estimate? (enter a decimal)",
		Answer:      answer,
		Explanation: "The rule fits a parabola through the two endpoints and the midpoint and integrates it exactly; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Composite rules (numeric)
// ---------------------------------------------------------------------------

type compositeGen struct{}

func (g *compositeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		b, n   int
		reason string
	}
	claims := []claim{
		{b: 1, n: 2, reason: "h = 0.5, giving 0.5(0 + 0.25 + 0.5) = 0.375."},
		{b: 2, n: 4, reason: "h = 0.5, giving 0.5(0 + 0.25 + 1 + 2.25 + 2) = 2.75."},
		{b: 1, n: 4, reason: "h = 0.25, giving 0.25(0 + 0.0625 + 0.25 + 0.5625 + 0.5) = 0.34375."},
		{b: 3, n: 3, reason: "h = 1, giving 1(0 + 1 + 4 + 4.5) = 9.5."},
		{b: 2, n: 2, reason: "h = 1, giving 1(0 + 1 + 2) = 3."},
		{b: 4, n: 4, reason: "h = 1, giving 1(0 + 1 + 4 + 9 + 8) = 22."},
	}
	c := claims[rand.Intn(len(claims))]
	h := float64(c.b) / float64(c.n)
	sum := 0.0
	for i := 0; i <= c.n; i++ {
		x := float64(i) * h
		w := 1.0
		if i == 0 || i == c.n {
			w = 0.5
		}
		sum += w * x * x
	}
	answer := decimal6(h * sum)
	return generator.Problem{
		Question: "The composite trapezoid rule with n equal subintervals of width h on the interval from 0 to b is h(f_0/2 + f_1 + ... + f_{n−1} + f_n/2). For f(x) = x², b = " +
			strconv.Itoa(c.b) + " and n = " + strconv.Itoa(c.n) + ", what is the estimate? (enter a decimal)",
		Answer:      answer,
		Explanation: "The single-interval rule is applied to each subinterval and the interior values are shared by two neighbours, which is why the end values carry weight one half; " + c.reason,
		Facts: map[string]string{
			"b": strconv.Itoa(c.b), "n": strconv.Itoa(c.n), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Error bounds for quadrature (multiple choice)
// ---------------------------------------------------------------------------

type quadErrorGen struct{}

func (g *quadErrorGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "The trapezoid rule error on a single interval is proportional to...",
			correct:     "the cube of the width times the second derivative",
			distractors: []string{"the square of the width", "the width itself", "the fourth power of the width"},
			reason:      "Because the rule is exact for straight lines, the leading error term is the second derivative times the cube of the width.",
		},
		{
			question:    "Simpson's rule is exact for polynomials up to degree...",
			correct:     "three",
			distractors: []string{"one", "two", "four"},
			reason:      "Although it is built from a parabola, the symmetry cancels the cubic error term.",
		},
		{
			question:    "Simpson's rule error is proportional to the fourth power of the width times...",
			correct:     "the fourth derivative",
			distractors: []string{"the second derivative", "the function value", "the first derivative"},
			reason:      "The first nonzero error term needs the fourth derivative.",
		},
		{
			question:    "Halving every subinterval in the composite trapezoid rule reduces the error by about...",
			correct:     "a factor of four",
			distractors: []string{"a factor of two", "a factor of eight", "nothing"},
			reason:      "The global error is proportional to the square of the width.",
		},
		{
			question:    "Doubling the number of subintervals in the composite Simpson rule reduces the error by about...",
			correct:     "a factor of sixteen",
			distractors: []string{"a factor of four", "a factor of two", "nothing"},
			reason:      "The global error is proportional to the fourth power of the width.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Romberg integration (multiple choice)
// ---------------------------------------------------------------------------

type rombergGen struct{}

func (g *rombergGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Romberg integration combines trapezoid estimates to...",
			correct:     "cancel the leading error terms by Richardson extrapolation",
			distractors: []string{"average the estimates", "increase the step size", "fit a straight line"},
			reason:      "Each extrapolation removes the next power of the width from the error expansion.",
		},
		{
			question:    "The first column of the Romberg table holds...",
			correct:     "the composite trapezoid rule at successively halved widths",
			distractors: []string{"the midpoint rule", "the Simpson rule", "the exact integral"},
			reason:      "Everything else in the table is extrapolated from that column.",
		},
		{
			question:    "Each Richardson step raises the order of accuracy by...",
			correct:     "two",
			distractors: []string{"one", "three", "four"},
			reason:      "The trapezoid error expansion runs in even powers of the width.",
		},
		{
			question:    "Romberg integration is most effective when the integrand is...",
			correct:     "smooth, so the error expansion is regular",
			distractors: []string{"discontinuous", "noisy", "constant"},
			reason:      "A jump or a noise floor breaks the expansion that the extrapolation assumes.",
		},
		{
			question:    "The Romberg table is built from...",
			correct:     "trapezoid estimates at compatible, nested step sizes",
			distractors: []string{"Gaussian nodes", "random samples", "the derivative of the integrand"},
			reason:      "Nested widths are what let one column be combined into the next.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Gaussian quadrature (numeric)
// ---------------------------------------------------------------------------

type gaussGen struct{}

func (g *gaussGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		k      int
		reason string
	}
	claims := []claim{
		{k: 1, reason: "Each node contributes (1/3)^1 = 1/3, so the sum is 2/3 = 0.666667."},
		{k: 2, reason: "Each node contributes (1/3)^2 = 1/9, so the sum is 2/9 = 0.222222."},
		{k: 3, reason: "Each node contributes 1/27, so the sum is 2/27 = 0.074074."},
		{k: 4, reason: "Each node contributes 1/81, so the sum is 2/81 = 0.024691."},
		{k: 5, reason: "Each node contributes 1/243, so the sum is 2/243 = 0.008230."},
		{k: 6, reason: "Each node contributes 1/729, so the sum is 2/729 = 0.002743."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(2 * math.Pow(1.0/3.0, float64(c.k)))
	return generator.Problem{
		Question: "Two-point Gauss quadrature on the interval from minus one to one uses the nodes at plus and minus one over the square root of three, with equal weights, so it estimates the integral of f by f(node) + f(−node). For f(x) = x raised to the power " +
			strconv.Itoa(2*c.k) + ", what is the estimate? (enter a decimal)",
		Answer:      answer,
		Explanation: "The two nodes are symmetric, so each contributes the same power of one over the square root of three; " + c.reason,
		Facts: map[string]string{
			"k": strconv.Itoa(c.k), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Adaptive quadrature (multiple choice)
// ---------------------------------------------------------------------------

type adaptiveGen struct{}

func (g *adaptiveGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Adaptive quadrature refines the mesh...",
			correct:     "where the estimated error is large",
			distractors: []string{"everywhere uniformly", "at random", "where the function is smallest"},
			reason:      "Effort is spent only where it reduces the error.",
		},
		{
			question:    "The main benefit of adaptivity is...",
			correct:     "accuracy near sharp features without wasting work on the smooth parts",
			distractors: []string{"a fixed number of samples", "simpler code", "exact answers"},
			reason:      "It concentrates subintervals where the integrand is hard.",
		},
		{
			question:    "Error estimates for adaptivity often come from...",
			correct:     "comparing two rules of different orders on the same subinterval",
			distractors: []string{"the function value alone", "a random sample", "the sign of the derivative"},
			reason:      "The difference between the two estimates is a practical bound on the error.",
		},
		{
			question:    "Adaptive methods are well suited to integrands with...",
			correct:     "localised peaks or endpoint singularities",
			distractors: []string{"constant values", "low-degree polynomial form", "linear shape"},
			reason:      "Those are exactly the cases where a uniform mesh is wasteful or inaccurate.",
		},
		{
			question:    "A recursive adaptive scheme subdivides a subinterval when...",
			correct:     "its local error estimate exceeds its share of the tolerance",
			distractors: []string{"the function is positive", "the interval is large", "the derivative vanishes"},
			reason:      "The tolerance is divided among the pieces, and only the pieces that miss it are split.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Integrating to infinity (multiple choice)
// ---------------------------------------------------------------------------

type improperGen struct{}

func (g *improperGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "To integrate over an infinite interval numerically, a standard approach is...",
			correct:     "a change of variable that maps the infinite interval to a finite one",
			distractors: []string{"to truncate and ignore the tail", "to sample at infinity", "to use the derivative"},
			reason:      "A finite interval can then be handled by an ordinary rule.",
		},
		{
			question:    "A Gauss-type rule for a half-infinite interval uses...",
			correct:     "Laguerre or Hermite weights that decay in the tail",
			distractors: []string{"equal weights", "Chebyshev weights", "no weights"},
			reason:      "The weights absorb the exponential decay of the integrand.",
		},
		{
			question:    "The danger of naive truncation of an infinite integral is...",
			correct:     "the neglected tail may still carry substantial mass",
			distractors: []string{"the integrand is undefined", "the interval is not finite", "the nodes repeat"},
			reason:      "A slowly decaying integrand can leave a large, silent error beyond the cut.",
		},
		{
			question:    "An endpoint singularity is often handled by...",
			correct:     "a substitution that removes the singularity",
			distractors: []string{"ignoring the endpoint", "halving the interval forever", "extrapolating at random"},
			reason:      "Once the integrand is bounded, an ordinary rule applies.",
		},
		{
			question:    "Improper integrals of the second kind have...",
			correct:     "an integrand that becomes unbounded inside or at the edge of the interval",
			distractors: []string{"an infinite interval only", "a bounded integrand", "no integral"},
			reason:      "The difficulty is the unbounded integrand, not the length of the interval.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Gaussian elimination (numeric)
// ---------------------------------------------------------------------------

type eliminationGen struct{}

func (g *eliminationGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b, c, d, x, y int
	}
	claims := []claim{
		{a: 2, b: 1, c: 1, d: 3, x: 1, y: 3},
		{a: 1, b: 1, c: 1, d: -1, x: 3, y: 1},
		{a: 3, b: 2, c: 1, d: 2, x: 2, y: 1},
		{a: 1, b: 2, c: 3, d: 1, x: 1, y: 2},
		{a: 2, b: 3, c: 1, d: 1, x: 4, y: 1},
		{a: 1, b: 1, c: 2, d: 3, x: 5, y: 2},
	}
	c := claims[rand.Intn(len(claims))]
	e := c.a*c.x + c.b*c.y
	f := c.c*c.x + c.d*c.y
	det := c.a*c.d - c.b*c.c
	answer := decimal6(float64(c.a*f-c.c*e) / float64(det))
	return generator.Problem{
		Question: "Solve the linear system a x + b y = e and c x + d y = f by Gaussian elimination. For a = " +
			strconv.Itoa(c.a) + ", b = " + strconv.Itoa(c.b) + ", c = " + strconv.Itoa(c.c) + ", d = " + strconv.Itoa(c.d) +
			", e = " + strconv.Itoa(e) + " and f = " + strconv.Itoa(f) + ", what is y? (enter a decimal)",
		Answer:      answer,
		Explanation: "Elimination forms a single equation for y by subtracting a multiple of the first row from the second; the determinant below is the pivot product; " + "y = (a f − c e)/(a d − b c).",
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "c": strconv.Itoa(c.c), "d": strconv.Itoa(c.d),
			"e": strconv.Itoa(e), "f": strconv.Itoa(f), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Partial pivoting (multiple choice)
// ---------------------------------------------------------------------------

type pivotingGen struct{}

func (g *pivotingGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Partial pivoting chooses the pivot as...",
			correct:     "the largest available entry in the current column",
			distractors: []string{"the smallest entry", "the diagonal entry always", "a random entry"},
			reason:      "A large pivot keeps the multipliers small and the elimination stable.",
		},
		{
			question:    "Partial pivoting improves...",
			correct:     "numerical stability by avoiding small pivots",
			distractors: []string{"the speed only", "the memory use", "the exactness of the arithmetic"},
			reason:      "Dividing by a tiny pivot would magnify rounding error.",
		},
		{
			question:    "Meeting a zero pivot during elimination...",
			correct:     "forces a row interchange, unless the matrix is singular at that stage",
			distractors: []string{"is impossible", "is harmless", "means the system has no solution"},
			reason:      "A zero in the pivot position needs a swap to bring a nonzero entry up.",
		},
		{
			question:    "Complete pivoting searches for the largest entry...",
			correct:     "in the whole remaining submatrix",
			distractors: []string{"in the first row only", "on the diagonal only", "in the last column only"},
			reason:      "It also allows column swaps, at more bookkeeping and with little extra practical benefit.",
		},
		{
			question:    "Partial pivoting usually keeps the growth factor...",
			correct:     "small, though matrices can be built to defeat it",
			distractors: []string{"always exactly one", "unbounded in practice", "negative"},
			reason:      "Worst-case growth is possible, but it is rare in practice.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// LU factorization (numeric)
// ---------------------------------------------------------------------------

type luGen struct{}

func (g *luGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b, c, d int
		reason     string
	}
	claims := []claim{
		{a: 2, b: 1, c: 4, d: 3, reason: "4/2 = 2."},
		{a: 3, b: 1, c: 6, d: 2, reason: "6/3 = 2."},
		{a: 1, b: 2, c: 3, d: 4, reason: "3/1 = 3."},
		{a: 4, b: 1, c: 6, d: 5, reason: "6/4 = 1.5."},
		{a: 2, b: 3, c: 5, d: 7, reason: "5/2 = 2.5."},
		{a: 5, b: 2, c: 10, d: 3, reason: "10/5 = 2."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(float64(c.c) / float64(c.a))
	return generator.Problem{
		Question: "For the matrix whose first row is a and b and second row is c and d, the elimination multiplier in the LU factorization is c divided by a. For a = " +
			strconv.Itoa(c.a) + ", b = " + strconv.Itoa(c.b) + ", c = " + strconv.Itoa(c.c) + " and d = " + strconv.Itoa(c.d) +
			", what is the multiplier? (enter a decimal)",
		Answer:      answer,
		Explanation: "The multiplier is the factor by which the first row is subtracted from the second to clear the entry below the pivot; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "c": strconv.Itoa(c.c), "d": strconv.Itoa(c.d), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The Jacobi method (numeric)
// ---------------------------------------------------------------------------

type jacobiGen struct{}

func (g *jacobiGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b, e, y0 int
		reason      string
	}
	claims := []claim{
		{a: 4, b: 1, e: 9, y0: 1, reason: "(9 − 1·1)/4 = 2."},
		{a: 3, b: 2, e: 8, y0: 1, reason: "(8 − 2·1)/3 = 2."},
		{a: 5, b: 1, e: 12, y0: 2, reason: "(12 − 1·2)/5 = 2."},
		{a: 4, b: 3, e: 10, y0: 2, reason: "(10 − 3·2)/4 = 1."},
		{a: 2, b: 1, e: 7, y0: 3, reason: "(7 − 1·3)/2 = 2."},
		{a: 5, b: 2, e: 13, y0: 4, reason: "(13 − 2·4)/5 = 1."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(float64(c.e-c.b*c.y0) / float64(c.a))
	return generator.Problem{
		Question: "One step of the Jacobi iteration solves each equation for its own unknown using the previous values. For the equation " +
			strconv.Itoa(c.a) + " x + " + strconv.Itoa(c.b) + " y = " + strconv.Itoa(c.e) + " with the previous value y = " +
			strconv.Itoa(c.y0) + ", what is the new x? (enter a decimal)",
		Answer:      answer,
		Explanation: "Jacobi uses only values from the previous sweep, so the other unknowns in the equation keep their old values; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "e": strconv.Itoa(c.e),
			"y0": strconv.Itoa(c.y0), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The Gauss-Seidel method (multiple choice)
// ---------------------------------------------------------------------------

type gaussSeidelGen struct{}

func (g *gaussSeidelGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Gauss-Seidel differs from Jacobi in that it...",
			correct:     "uses each updated value as soon as it is available",
			distractors: []string{"uses the previous values throughout", "needs a better guess", "does not iterate"},
			reason:      "The fresher values feed immediately into the later equations of the same sweep.",
		},
		{
			question:    "For a diagonally dominant matrix, Gauss-Seidel...",
			correct:     "converges, often about twice as fast as Jacobi",
			distractors: []string{"diverges", "converges more slowly than Jacobi always", "converges in one step"},
			reason:      "Diagonal dominance gives convergence, and the immediate reuse of values speeds it up.",
		},
		{
			question:    "Both Jacobi and Gauss-Seidel are guaranteed to converge when the matrix is...",
			correct:     "strictly diagonally dominant",
			distractors: []string{"singular", "symmetric only", "triangular only"},
			reason:      "The dominance of the diagonal controls the iteration matrix.",
		},
		{
			question:    "Gauss-Seidel can be written as a splitting in which the iteration matrix...",
			correct:     "involves the lower triangle of the matrix",
			distractors: []string{"involves only the diagonal", "involves only the upper triangle", "is the identity"},
			reason:      "Solving forward within a sweep uses the lower triangle already updated in that sweep.",
		},
		{
			question:    "A practical advantage of Gauss-Seidel is...",
			correct:     "it usually needs fewer iterations than Jacobi for the same problem",
			distractors: []string{"it needs no diagonal dominance", "it is a direct method", "it works on any matrix"},
			reason:      "The faster information flow across the mesh is the usual reason it is preferred.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// The power method (numeric)
// ---------------------------------------------------------------------------

type powerMethodGen struct{}

func (g *powerMethodGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b, d int
		reason  string
	}
	claims := []claim{
		{a: 2, b: 1, d: 2, reason: "The midpoint is 2 and the square root of b² is 1, giving 2 + 1 = 3."},
		{a: 3, b: 1, d: 3, reason: "3 + 1 = 4."},
		{a: 2, b: 2, d: 2, reason: "2 + 2 = 4."},
		{a: 1, b: 2, d: 1, reason: "1 + 2 = 3."},
		{a: 4, b: 1, d: 4, reason: "4 + 1 = 5."},
		{a: 3, b: 2, d: 3, reason: "3 + 2 = 5."},
	}
	c := claims[rand.Intn(len(claims))]
	mid := float64(c.a+c.d) / 2
	rad := math.Sqrt(math.Pow(float64(c.a-c.d)/2, 2) + float64(c.b*c.b))
	answer := decimal6(mid + rad)
	return generator.Problem{
		Question: "The power method converges to the dominant eigenvalue of a matrix. For the symmetric matrix with diagonal entries a = " +
			strconv.Itoa(c.a) + " and d = " + strconv.Itoa(c.d) + " and off-diagonal entry b = " + strconv.Itoa(c.b) +
			", the dominant eigenvalue is (a + d)/2 + sqrt(((a − d)/2)² + b²). What is it? (enter a decimal)",
		Answer:      answer,
		Explanation: "The eigenvalues of a two-by-two symmetric matrix are the midpoint of the diagonal plus or minus a square-root term, and the power method returns the larger; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "d": strconv.Itoa(c.d), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Least squares by QR (multiple choice)
// ---------------------------------------------------------------------------

type leastSquaresQRGen struct{}

func (g *leastSquaresQRGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Solving least squares by QR avoids...",
			correct:     "forming the normal equations and squaring the condition number",
			distractors: []string{"using any data", "the need for a matrix", "the residual"},
			reason:      "Orthogonal transformations preserve the two-norm, so the problem is not made worse conditioned.",
		},
		{
			question:    "In least squares by QR, the coefficients solve...",
			correct:     "an upper-triangular system from the first columns of R",
			distractors: []string{"the normal equations only", "a diagonal system", "no system at all"},
			reason:      "The orthogonal factor drops out of the normal equations, leaving a triangular system.",
		},
		{
			question:    "Compared with QR, the normal equations square...",
			correct:     "the condition number of the design matrix",
			distractors: []string{"the number of unknowns", "the residual norm", "the data values"},
			reason:      "That squaring is why the normal equations lose accuracy on ill-conditioned problems.",
		},
		{
			question:    "QR is preferred over the normal equations when...",
			correct:     "the design matrix is ill conditioned",
			distractors: []string{"the matrix is square and well conditioned", "there is no data", "the system is underdetermined"},
			reason:      "The more sensitive the problem, the more the squaring of the condition number hurts.",
		},
		{
			question:    "Householder reflections are used in QR because they...",
			correct:     "zero out entries stably using orthogonal transformations",
			distractors: []string{"scale the rows", "invert the matrix", "randomise the data"},
			reason:      "An orthogonal reflection does not change lengths, so it does not amplify rounding error.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}
