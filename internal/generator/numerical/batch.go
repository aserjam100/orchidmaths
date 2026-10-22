package numerical

// The first numerical-analysis batch: twenty concepts that open the domain with
// error and conditioning, root finding, and the first half of interpolation.
//
// The domain is about the gap between an exact mathematical object and the
// finitely many digits a machine can hold and the finitely many steps a method
// can afford. This batch measures that gap (absolute and relative error,
// roundoff, cancellation, conditioning, stability, truncation), closes it by
// iteration (bisection, fixed-point, Newton, secant, false position, and the
// orders those methods achieve), and starts reconstructing a function from
// samples (Lagrange interpolation, divided differences, the error formula,
// Chebyshev nodes and splines).
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

func decimal6(v float64) string {
	s := strconv.FormatFloat(v, 'f', 6, 64)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if s == "" || s == "-0" {
		s = "0"
	}
	return s
}

// ---------------------------------------------------------------------------
// Absolute and relative error (numeric)
// ---------------------------------------------------------------------------

type absRelErrorGen struct{}

func (g *absRelErrorGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		true, approx int
		reason       string
	}
	claims := []claim{
		{true: 1000, approx: 999, reason: "The absolute error is 1 and the true value is 1000, so the relative error is 1/1000 = 0.001."},
		{true: 500, approx: 505, reason: "The absolute error is 5 and the true value is 500, so 5/500 = 0.01."},
		{true: 200, approx: 198, reason: "The absolute error is 2 and the true value is 200, so 2/200 = 0.01."},
		{true: 250, approx: 245, reason: "The absolute error is 5 and the true value is 250, so 5/250 = 0.02."},
		{true: 1000, approx: 1004, reason: "The absolute error is 4 and the true value is 1000, so 4/1000 = 0.004."},
		{true: 400, approx: 404, reason: "The absolute error is 4 and the true value is 400, so 4/400 = 0.01."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(math.Abs(float64(c.approx-c.true)) / math.Abs(float64(c.true)))
	return generator.Problem{
		Question: "A quantity has true value " + strconv.Itoa(c.true) + " and is approximated by " +
			strconv.Itoa(c.approx) + ". What is the relative error |approx − true|/|true|? (enter a decimal)",
		Answer:      answer,
		Explanation: "The relative error divides the absolute error by the size of the true value, which makes it independent of the units; " + c.reason,
		Facts: map[string]string{
			"true": strconv.Itoa(c.true), "approx": strconv.Itoa(c.approx), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Floating-point roundoff and machine epsilon (multiple choice)
// ---------------------------------------------------------------------------

type roundoffGen struct{}

func (g *roundoffGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Machine epsilon is...",
			correct:     "the gap between one and the next representable floating-point number",
			distractors: []string{"the smallest positive number a machine can hold", "the largest exponent", "the number of decimal digits printed"},
			reason:      "It measures the relative spacing of the floats near one, which is the natural unit of roundoff.",
		},
		{
			question:    "The number nought point one cannot be stored exactly in binary floating point because...",
			correct:     "its binary expansion does not terminate, so it must be rounded",
			distractors: []string{"it is larger than one", "it is negative", "the machine runs out of memory"},
			reason:      "Only fractions whose denominator is a power of two terminate in binary.",
		},
		{
			question:    "Roundoff error is unavoidable in floating-point arithmetic because...",
			correct:     "the machine keeps only finitely many significant digits",
			distractors: []string{"algorithms are written badly", "the input data are always noisy", "mathematics is inexact"},
			reason:      "Every real number is rounded to the nearest member of a finite set.",
		},
		{
			question:    "A single rounding of a real number to the nearest float produces a relative error no larger than...",
			correct:     "about half of machine epsilon",
			distractors: []string{"one", "the square root of machine epsilon", "the number itself"},
			reason:      "Rounding to nearest can never be off by more than half the gap to the next float.",
		},
		{
			question:    "Floating-point addition is not associative, so...",
			correct:     "the order of a sum can change the computed result",
			distractors: []string{"addition is undefined", "the sum is always exact", "subtraction is impossible"},
			reason:      "Each intermediate result is rounded, so regrouping changes which roundings occur.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Catastrophic cancellation (numeric)
// ---------------------------------------------------------------------------

type cancellationGen struct{}

func (g *cancellationGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		x, y   int
		reason string
	}
	claims := []claim{
		{x: 1001, y: 999, reason: "(1001 + 999)/|1001 − 999| = 2000/2 = 1000."},
		{x: 1002, y: 998, reason: "(1002 + 998)/4 = 500."},
		{x: 1005, y: 995, reason: "(1005 + 995)/10 = 200."},
		{x: 1010, y: 990, reason: "(1010 + 990)/20 = 100."},
		{x: 1100, y: 900, reason: "(1100 + 900)/200 = 10."},
		{x: 1001, y: 1000, reason: "(1001 + 1000)/1 = 2001."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(float64(absInt(c.x)+absInt(c.y)) / float64(absInt(c.x-c.y)))
	return generator.Problem{
		Question: "The subtraction x − y loses accuracy when x and y are close: the relative error of the difference is amplified by (|x| + |y|)/|x − y| relative to the inputs. For x = " +
			strconv.Itoa(c.x) + " and y = " + strconv.Itoa(c.y) + ", what is that amplification factor? (enter a decimal)",
		Answer:      answer,
		Explanation: "Cancellation is the loss of leading digits when nearly equal numbers are subtracted, and the amplification grows as the difference shrinks; " + c.reason,
		Facts: map[string]string{
			"x": strconv.Itoa(c.x), "y": strconv.Itoa(c.y), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The condition number of a problem (numeric)
// ---------------------------------------------------------------------------

type conditioningGen struct{}

func (g *conditioningGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		n      int
		reason string
	}
	claims := []claim{
		{n: 2, reason: "|x·2x|/|x²| = 2."},
		{n: 3, reason: "|x·3x²|/|x³| = 3."},
		{n: 4, reason: "|x·4x³|/|x⁴| = 4."},
		{n: 5, reason: "|x·5x⁴|/|x⁵| = 5."},
		{n: 6, reason: "|x·6x⁵|/|x⁶| = 6."},
		{n: 7, reason: "|x·7x⁶|/|x⁷| = 7."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(float64(c.n))
	return generator.Problem{
		Question: "The relative condition number of the function x raised to the power n at a point x is |x·f'(x)|/|f(x)|. For f(x) = x^" +
			strconv.Itoa(c.n) + ", what is that condition number? (enter a decimal)",
		Answer:      answer,
		Explanation: "Differentiating the power and simplifying cancels the x, leaving the exponent, so the evaluation is as well or badly conditioned as the exponent says; " + c.reason,
		Facts: map[string]string{
			"n": strconv.Itoa(c.n), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Stability of an algorithm (multiple choice)
// ---------------------------------------------------------------------------

type stabilityGen struct{}

func (g *stabilityGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "An algorithm is backward stable when...",
			correct:     "it returns the exact answer to a nearby problem",
			distractors: []string{"it is fast", "it uses little memory", "it returns the exact answer to the given problem"},
			reason:      "The rounding is pushed back into a small perturbation of the data, which the condition number then amplifies.",
		},
		{
			question:    "A stable algorithm applied to an ill-conditioned problem...",
			correct:     "can still return an inaccurate answer, because the problem amplifies the perturbation",
			distractors: []string{"always returns an accurate answer", "cannot be implemented", "has no rounding error"},
			reason:      "Stability controls the algorithm; the condition number controls the problem, and the two multiply.",
		},
		{
			question:    "The condition number of a problem measures...",
			correct:     "how much a relative change in the data can change the relative answer",
			distractors: []string{"the speed of the algorithm", "the memory it uses", "the order of convergence"},
			reason:      "It is a property of the problem alone, independent of the method chosen to solve it.",
		},
		{
			question:    "A backward-stable algorithm guarantees that the computed result is the exact result of...",
			correct:     "a problem whose data differ from the given data by a small relative amount",
			distractors: []string{"the given problem exactly", "any problem of the same size", "the problem in infinite precision"},
			reason:      "That is the definition of backward stability.",
		},
		{
			question:    "For a well-conditioned problem with condition number near one, a backward-stable algorithm...",
			correct:     "has a relative error of the order of machine epsilon",
			distractors: []string{"has a relative error near one", "has no error at all", "diverges"},
			reason:      "The small backward perturbation passes through a problem that barely amplifies it.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Taylor truncation error (multiple choice)
// ---------------------------------------------------------------------------

type taylorErrorGen struct{}

func (g *taylorErrorGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "The truncation error of a numerical method is the error made by...",
			correct:     "replacing an exact process with a finite approximation",
			distractors: []string{"finite-precision arithmetic", "noisy input data", "a badly chosen algorithm"},
			reason:      "It is the price of cutting off a series or a limiting process, and it is separate from roundoff.",
		},
		{
			question:    "For a method of order n the local truncation error behaves like...",
			correct:     "a constant times the step raised to the power n plus one",
			distractors: []string{"the step itself", "one over the step", "a constant, independent of the step"},
			reason:      "The leading term left after a Taylor expansion of order n is the next one.",
		},
		{
			question:    "Halving the step of a second-order method reduces its global error by about...",
			correct:     "a factor of four",
			distractors: []string{"a factor of two", "a factor of eight", "nothing"},
			reason:      "A second-order method has global error proportional to the square of the step.",
		},
		{
			question:    "Taylor's theorem with remainder is used in numerical analysis to...",
			correct:     "bound the truncation error of a difference formula",
			distractors: []string{"compute exact derivatives", "solve linear systems", "factor a matrix"},
			reason:      "The remainder term is exactly the error the formula makes.",
		},
		{
			question:    "The global error of a one-step method accumulates...",
			correct:     "over the whole interval, not just one step",
			distractors: []string{"over one step only", "over the number of digits", "only from the input data"},
			reason:      "The local errors of many steps add up, which is why the global order is one lower than the local order.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// The bisection method (numeric)
// ---------------------------------------------------------------------------

type bisectionGen struct{}

func (g *bisectionGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		width, tol float64
		widthStr   string
		tolStr     string
		reason     string
	}
	claims := []claim{
		{width: 8, tol: 0.5, widthStr: "8", tolStr: "0.5", reason: "log2(16) = 4, so four halvings bring the width to the tolerance."},
		{width: 1, tol: 0.001, widthStr: "1", tolStr: "0.001", reason: "log2(1000) = 9.97, so ten steps are needed."},
		{width: 2, tol: 0.01, widthStr: "2", tolStr: "0.01", reason: "log2(200) = 7.64, so eight steps are needed."},
		{width: 16, tol: 0.25, widthStr: "16", tolStr: "0.25", reason: "log2(64) = 6, so six steps."},
		{width: 1, tol: 0.0009765625, widthStr: "1", tolStr: "0.0009765625", reason: "1/0.0009765625 = 1024 = 2^10, so ten steps."},
		{width: 4, tol: 0.001, widthStr: "4", tolStr: "0.001", reason: "log2(4000) = 11.97, so twelve steps are needed."},
	}
	c := claims[rand.Intn(len(claims))]
	steps := int(math.Ceil(math.Log2(c.width / c.tol)))
	answer := strconv.Itoa(steps)
	return generator.Problem{
		Question: "Bisection halves the interval at each step. Starting from an interval of width " +
			c.widthStr + ", how many steps are needed to make the width at most " + c.tolStr + "? (enter a whole number)",
		Answer:      answer,
		Explanation: "The width after k steps is the starting width divided by 2^k, so k must be at least the base-two logarithm of the ratio; " + c.reason,
		Facts: map[string]string{
			"width": c.widthStr, "tol": c.tolStr, "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Fixed-point iteration (numeric)
// ---------------------------------------------------------------------------

type fixedPointGen struct{}

func (g *fixedPointGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, bn, bd, x0, n int
		reason           string
	}
	claims := []claim{
		{a: 1, bn: 1, bd: 2, x0: 0, n: 3, reason: "The iterates are 1, 1.5 and 1.75."},
		{a: 2, bn: 1, bd: 2, x0: 0, n: 2, reason: "The iterates are 2 and 3."},
		{a: 1, bn: 1, bd: 3, x0: 0, n: 2, reason: "1 then 1 + 1/3 = 1.333333."},
		{a: 0, bn: 1, bd: 2, x0: 4, n: 3, reason: "The iterates halve: 4, 2, 1, 0.5."},
		{a: 3, bn: 1, bd: 2, x0: 0, n: 2, reason: "The iterates are 3 and 4.5."},
		{a: 1, bn: 1, bd: 4, x0: 0, n: 2, reason: "1 then 1.25."},
	}
	c := claims[rand.Intn(len(claims))]
	x := float64(c.x0)
	b := float64(c.bn) / float64(c.bd)
	for i := 0; i < c.n; i++ {
		x = float64(c.a) + b*x
	}
	answer := decimal6(x)
	return generator.Problem{
		Question: "Fixed-point iteration applies x_{k+1} = a + b·x_k with a = " + strconv.Itoa(c.a) +
			", b = " + strconv.Itoa(c.bn) + "/" + strconv.Itoa(c.bd) + ", starting from x_0 = " +
			strconv.Itoa(c.x0) + ". What is x_" + strconv.Itoa(c.n) + " after " + strconv.Itoa(c.n) +
			" steps? (enter a decimal)",
		Answer:      answer,
		Explanation: "Substitute the previous iterate into the right-hand side repeatedly; the iteration converges when the slope of the right-hand side is smaller than one in absolute value; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "bn": strconv.Itoa(c.bn), "bd": strconv.Itoa(c.bd),
			"x0": strconv.Itoa(c.x0), "n": strconv.Itoa(c.n), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Newton's method (numeric)
// ---------------------------------------------------------------------------

type newtonGen struct{}

func (g *newtonGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		c, x0  float64
		cStr   string
		x0Str  string
		reason string
	}
	claims := []claim{
		{c: 2, x0: 1, cStr: "2", x0Str: "1", reason: "(1 + 2/1)/2 = 1.5."},
		{c: 2, x0: 2, cStr: "2", x0Str: "2", reason: "(2 + 2/2)/2 = 1.5."},
		{c: 9, x0: 2, cStr: "9", x0Str: "2", reason: "(2 + 9/2)/2 = 3.25."},
		{c: 16, x0: 5, cStr: "16", x0Str: "5", reason: "(5 + 16/5)/2 = 4.1."},
		{c: 4, x0: 1, cStr: "4", x0Str: "1", reason: "(1 + 4/1)/2 = 2.5."},
		{c: 25, x0: 4, cStr: "25", x0Str: "4", reason: "(4 + 25/4)/2 = 5.125."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6((c.x0 + c.c/c.x0) / 2)
	return generator.Problem{
		Question: "One Newton step for f(x) = x² − c is x_1 = (x_0 + c/x_0)/2. For c = " + c.cStr +
			" and x_0 = " + c.x0Str + ", what is x_1? (enter a decimal)",
		Answer:      answer,
		Explanation: "Newton replaces the function by its tangent line, and the tangent for the square touches the axis at the average of the current guess and c divided by it; " + c.reason,
		Facts: map[string]string{
			"c": c.cStr, "x0": c.x0Str, "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The secant method (numeric)
// ---------------------------------------------------------------------------

type secantGen struct{}

func (g *secantGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		c, x0, x1 float64
		reason    string
	}
	claims := []claim{
		{c: 2, x0: 1, x1: 2, reason: "(2 − 2)·(2 − 1)/(2 − (−1)) is subtracted from 2, giving 1.333333."},
		{c: 9, x0: 1, x1: 2, reason: "x_2 = 2 − (−5)(1)/3 = 3.666667."},
		{c: 4, x0: 1, x1: 3, reason: "x_2 = 3 − 5·2/8 = 1.75."},
		{c: 5, x0: 2, x1: 3, reason: "x_2 = 3 − 4·1/5 = 2.2."},
		{c: 10, x0: 2, x1: 4, reason: "x_2 = 4 − 6·2/12 = 3."},
		{c: 3, x0: 1, x1: 2, reason: "x_2 = 2 − 1·1/3 = 1.666667."},
	}
	c := claims[rand.Intn(len(claims))]
	f := func(x float64) float64 { return x*x - c.c }
	x2 := c.x1 - f(c.x1)*(c.x1-c.x0)/(f(c.x1)-f(c.x0))
	answer := decimal6(x2)
	return generator.Problem{
		Question: "One step of the secant method for f(x) = x² − c uses x_2 = x_1 − f(x_1)(x_1 − x_0)/(f(x_1) − f(x_0)). For c = " +
			decimal6(c.c) + ", x_0 = " + decimal6(c.x0) + " and x_1 = " + decimal6(c.x1) + ", what is x_2? (enter a decimal)",
		Answer:      answer,
		Explanation: "The secant method replaces the derivative by the slope of the line through the two most recent points; " + c.reason,
		Facts: map[string]string{
			"c": decimal6(c.c), "x0": decimal6(c.x0), "x1": decimal6(c.x1), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The method of false position (multiple choice)
// ---------------------------------------------------------------------------

type falsePositionGen struct{}

func (g *falsePositionGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "The method of false position keeps a bracket of the root and updates...",
			correct:     "one endpoint, by the intercept of the line joining the two endpoints",
			distractors: []string{"both endpoints by halving", "the midpoint of the bracket", "the derivative at the midpoint"},
			reason:      "It is the secant method constrained to keep the sign change inside the bracket.",
		},
		{
			question:    "False position usually converges...",
			correct:     "faster than bisection but slower than Newton",
			distractors: []string{"faster than Newton", "more slowly than bisection", "at exactly the bisection rate"},
			reason:      "It uses the function values to place the new endpoint, which beats halving but lacks the derivative's speed.",
		},
		{
			question:    "A known weakness of false position is that...",
			correct:     "one endpoint can stay fixed for many iterations when the function is very curved",
			distractors: []string{"it needs the derivative", "it cannot start without a bracket", "it always diverges"},
			reason:      "The interpolation can land very close to one side, so the other endpoint moves little.",
		},
		{
			question:    "Compared with the secant method, false position...",
			correct:     "keeps the root bracketed",
			distractors: []string{"does not bracket the root", "requires the derivative", "cannot converge superlinearly"},
			reason:      "Dropping the secant point that would leave the bracket is what distinguishes the two methods.",
		},
		{
			question:    "False position needs to begin with...",
			correct:     "two points where the function has opposite signs",
			distractors: []string{"one point and the derivative", "two points with the same sign", "no points at all"},
			reason:      "The sign change is the bracket that guarantees a root inside.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Order of convergence (numeric)
// ---------------------------------------------------------------------------

type orderGen struct{}

func (g *orderGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b, c int
		reason  string
	}
	claims := []claim{
		{a: 1, b: 2, c: 4, reason: "(4 − 2)/(2 − 1) = 2, the quadratic order."},
		{a: 1, b: 2, c: 3, reason: "(3 − 2)/(2 − 1) = 1, the linear order."},
		{a: 1, b: 2, c: 6, reason: "(6 − 2)/(2 − 1) = 4."},
		{a: 1, b: 3, c: 9, reason: "(9 − 3)/(3 − 1) = 3."},
		{a: 2, b: 4, c: 8, reason: "(8 − 4)/(4 − 2) = 2."},
		{a: 1, b: 2, c: 5, reason: "(5 − 2)/(2 − 1) = 3."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(float64(c.c-c.b) / float64(c.b-c.a))
	return generator.Problem{
		Question: "The errors of an iterative method satisfy e_{k+2} ≈ e_{k+1}^p / e_k^{p−1}. From three successive errors of 10^-" +
			strconv.Itoa(c.a) + ", 10^-" + strconv.Itoa(c.b) + " and 10^-" + strconv.Itoa(c.c) +
			", estimate the order p as log(e_{k+2}/e_{k+1}) divided by log(e_{k+1}/e_k). (enter a decimal)",
		Answer:      answer,
		Explanation: "Taking logarithms turns the error relation into a linear equation for p, and the differences of the exponents solve it; " + c.reason,
		Facts: map[string]string{
			"a": strconv.Itoa(c.a), "b": strconv.Itoa(c.b), "c": strconv.Itoa(c.c), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Multiple roots and modified Newton (multiple choice)
// ---------------------------------------------------------------------------

type multipleRootGen struct{}

func (g *multipleRootGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "At a root of multiplicity m, the ordinary Newton iteration converges...",
			correct:     "linearly rather than quadratically",
			distractors: []string{"quadratically", "not at all", "cubically"},
			reason:      "The zero of the derivative at the root removes the quadratic term that normally drives the fast convergence in the root",
		},
		{
			question:    "The modified Newton iteration for a root of multiplicity m...",
			correct:     "multiplies the Newton step by m",
			distractors: []string{"divides the Newton step by m", "drops the derivative", "uses the second derivative only"},
			reason:      "Restoring the quadratic rate needs a step m times the ordinary one.",
		},
		{
			question:    "A function with a double root has...",
			correct:     "both the function and its derivative equal to zero at the root",
			distractors: []string{"only the function equal to zero", "only the derivative equal to zero", "neither equal to zero"},
			reason:      "Multiplicity counts how many derivatives vanish with the function.",
		},
		{
			question:    "Near a simple root, Newton's method converges...",
			correct:     "quadratically",
			distractors: []string{"linearly", "cubically", "not at all"},
			reason:      "A nonzero derivative at the root gives the doubling of correct digits each step.",
		},
		{
			question:    "Detecting a multiple root numerically is hard because...",
			correct:     "the function is flat near the root, so small errors blur the multiplicity",
			distractors: []string{"the derivative is infinite there", "the function is discontinuous", "there is no root"},
			reason:      "Near a repeated root the graph is tangent to the axis, and the tiny function values are dominated by roundoff.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Bracketing versus open methods (multiple choice)
// ---------------------------------------------------------------------------

type bracketingGen struct{}

func (g *bracketingGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "A bracketing method needs...",
			correct:     "two points where the function has opposite signs",
			distractors: []string{"one point and the derivative", "two points with the same sign", "no starting points"},
			reason:      "The sign change guarantees a root inside the interval by continuity.",
		},
		{
			question:    "An open method needs...",
			correct:     "one or two starting guesses and may diverge",
			distractors: []string{"a sign change to begin", "an interval that contains the root", "a guaranteed bracket"},
			reason:      "Without a bracket there is no guarantee, so the method can wander or cycle.",
		},
		{
			question:    "Bisection is guaranteed to converge...",
			correct:     "whenever the initial interval brackets a root",
			distractors: []string{"for any starting points", "only for polynomials", "only when the derivative is known"},
			reason:      "Halving a bracket that contains a sign change keeps it containing one.",
		},
		{
			question:    "Newton's method can fail when...",
			correct:     "the derivative is zero or the iterate leaves the region of attraction",
			distractors: []string{"the function is continuous", "a bracket is given", "the root is simple"},
			reason:      "A vanishing derivative makes the step infinite, and a distant iterate can converge elsewhere or diverge.",
		},
		{
			question:    "The secant method differs from false position in that...",
			correct:     "it does not keep the root bracketed",
			distractors: []string{"it needs the derivative", "it halves the interval", "it cannot converge superlinearly"},
			reason:      "The secant method uses the newest two points even if the bracket is lost, which speeds it up and weakens the guarantee.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Horner's method (numeric)
// ---------------------------------------------------------------------------

type hornerGen struct{}

func (g *hornerGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		coeffs []int
		x      int
		reason string
	}
	claims := []claim{
		{coeffs: []int{1, 2, 3}, x: 2, reason: "1 + 2·2 + 3·4 = 17."},
		{coeffs: []int{1, -1, 1}, x: 3, reason: "1 − 3 + 9 = 7."},
		{coeffs: []int{2, 0, 1}, x: 4, reason: "2 + 0 + 16 = 18."},
		{coeffs: []int{1, 1, 1, 1}, x: 2, reason: "1 + 2 + 4 + 8 = 15."},
		{coeffs: []int{3, 2}, x: 5, reason: "3 + 10 = 13."},
		{coeffs: []int{1, 0, 0, 1}, x: 3, reason: "1 + 27 = 28."},
	}
	c := claims[rand.Intn(len(claims))]
	val := 0
	for i := len(c.coeffs) - 1; i >= 0; i-- {
		val = val*c.x + c.coeffs[i]
	}
	answer := strconv.Itoa(val)
	parts := make([]string, len(c.coeffs))
	for i, a := range c.coeffs {
		parts[i] = strconv.Itoa(a)
	}
	return generator.Problem{
		Question: "Horner's method evaluates a polynomial by nested multiplication. For the polynomial with coefficients (constant first) " +
			strings.Join(parts, ", ") + " evaluated at x = " + strconv.Itoa(c.x) + ", what is the value p(x)? (enter a decimal)",
		Answer:      answer,
		Explanation: "Horner's scheme starts from the leading coefficient and repeatedly multiplies by x and adds the next coefficient, using one multiplication per term; " + c.reason,
		Facts: map[string]string{
			"coeffs": strings.Join(parts, ","), "x": strconv.Itoa(c.x), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Lagrange interpolation (numeric)
// ---------------------------------------------------------------------------

type lagrangeGen struct{}

func (g *lagrangeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, b, c, q int
		reason     string
	}
	claims := []claim{
		{a: 1, b: 0, c: 0, q: 3, reason: "The quadratic is x² and its value at 3 is 9."},
		{a: 1, b: 1, c: 1, q: 2, reason: "x² + x + 1 evaluates to 7 at 2."},
		{a: 2, b: -1, c: 0, q: 2, reason: "2x² − x is 6 at 2."},
		{a: 1, b: 2, c: 3, q: 2, reason: "x² + 2x + 3 is 11 at 2."},
		{a: 2, b: 0, c: 1, q: 3, reason: "2x² + 1 is 19 at 3."},
		{a: 1, b: -1, c: 2, q: 3, reason: "x² − x + 2 is 8 at 3."},
	}
	cl := claims[rand.Intn(len(claims))]
	y0 := cl.c
	y1 := cl.a + cl.b + cl.c
	y2 := 4*cl.a + 2*cl.b + cl.c
	val := cl.a*cl.q*cl.q + cl.b*cl.q + cl.c
	answer := strconv.Itoa(val)
	return generator.Problem{
		Question: "The Lagrange interpolant through the three points (0, " + strconv.Itoa(y0) + "), (1, " +
			strconv.Itoa(y1) + ") and (2, " + strconv.Itoa(y2) + ") is the quadratic through them. What is the interpolated value at x = " +
			strconv.Itoa(cl.q) + "? (enter a decimal)",
		Answer:      answer,
		Explanation: "Each Lagrange basis polynomial is one at its own node and zero at the others, so the interpolant is the corresponding weighted sum of the three values; " + cl.reason,
		Facts: map[string]string{
			"y0": strconv.Itoa(y0), "y1": strconv.Itoa(y1), "y2": strconv.Itoa(y2),
			"q": strconv.Itoa(cl.q), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// Newton's divided differences (numeric)
// ---------------------------------------------------------------------------

type dividedDiffGen struct{}

func (g *dividedDiffGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		x0, y0, x1, y1 int
		reason         string
	}
	claims := []claim{
		{x0: 1, y0: 2, x1: 4, y1: 11, reason: "(11 − 2)/(4 − 1) = 3."},
		{x0: 0, y0: 3, x1: 2, y1: 7, reason: "(7 − 3)/2 = 2."},
		{x0: 2, y0: 1, x1: 5, y1: 16, reason: "(16 − 1)/3 = 5."},
		{x0: 0, y0: 0, x1: 3, y1: 6, reason: "6/3 = 2."},
		{x0: 1, y0: 4, x1: 3, y1: 8, reason: "(8 − 4)/2 = 2."},
		{x0: 2, y0: 2, x1: 6, y1: 14, reason: "(14 − 2)/4 = 3."},
	}
	c := claims[rand.Intn(len(claims))]
	answer := decimal6(float64(c.y1-c.y0) / float64(c.x1-c.x0))
	return generator.Problem{
		Question: "The first divided difference of the points (" + strconv.Itoa(c.x0) + ", " + strconv.Itoa(c.y0) +
			") and (" + strconv.Itoa(c.x1) + ", " + strconv.Itoa(c.y1) + ") is (y_1 − y_0)/(x_1 − x_0). What is it? (enter a decimal)",
		Answer:      answer,
		Explanation: "Divided differences are the slopes that build the Newton form of the interpolant, and the first one is the average slope over the interval; " + c.reason,
		Facts: map[string]string{
			"x0": strconv.Itoa(c.x0), "y0": strconv.Itoa(c.y0),
			"x1": strconv.Itoa(c.x1), "y1": strconv.Itoa(c.y1), "answer": answer,
		},
	}
}

// ---------------------------------------------------------------------------
// The interpolation error (multiple choice)
// ---------------------------------------------------------------------------

type interpErrorGen struct{}

func (g *interpErrorGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "The error of polynomial interpolation at a point is proportional to...",
			correct:     "the product of the distances from that point to the interpolation nodes",
			distractors: []string{"the spacing of the nodes only", "the number of nodes", "the derivative at the nodes"},
			reason:      "The node product vanishes at the nodes and grows between them, and it multiplies the derivative term.",
		},
		{
			question:    "By the interpolation error formula, the error involves the (n+1)-th derivative...",
			correct:     "evaluated at some unknown point of the interval",
			distractors: []string{"at the left endpoint", "at a node", "at infinity"},
			reason:      "The remainder theorem places the unknown point strictly inside the interval.",
		},
		{
			question:    "Adding more equally spaced nodes can...",
			correct:     "make the error worse near the ends, as Runge's example shows",
			distractors: []string{"always reduce the maximum error", "never change the error", "remove the need for a derivative"},
			reason:      "Runge's phenomenon is the blow-up of the interpolant's edges when the nodes are equally spaced.",
		},
		{
			question:    "The interpolation error is zero at...",
			correct:     "the interpolation nodes",
			distractors: []string{"the midpoint", "every point", "no point"},
			reason:      "The interpolant matches the function exactly at the data.",
		},
		{
			question:    "The error formula shows that accuracy depends on...",
			correct:     "both the node spacing and the size of the derivatives",
			distractors: []string{"only the node spacing", "only the function values", "neither"},
			reason:      "The node product and the derivative bound enter as a product.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Chebyshev nodes (multiple choice)
// ---------------------------------------------------------------------------

type chebyshevGen struct{}

func (g *chebyshevGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Chebyshev nodes are...",
			correct:     "the roots of a Chebyshev polynomial, clustered toward the ends of the interval",
			distractors: []string{"equally spaced points", "the endpoints only", "random points"},
			reason:      "Their clustering tames the edge growth that equally spaced nodes provoke.",
		},
		{
			question:    "Using Chebyshev nodes helps because they...",
			correct:     "minimise the maximum absolute value of the node product in the error formula",
			distractors: []string{"make the function linear", "remove the derivative term", "use fewer nodes"},
			reason:      "The node product is exactly what the Chebyshev construction minimises.",
		},
		{
			question:    "Chebyshev nodes cluster near...",
			correct:     "the endpoints of the interval",
			distractors: []string{"the centre", "the left endpoint only", "the data values"},
			reason:      "The density rises where the equally spaced error is worst.",
		},
		{
			question:    "Runge's phenomenon on equally spaced nodes is...",
			correct:     "the error blowing up near the interval ends",
			distractors: []string{"the error at the centre", "a machine-precision effect", "a monotone convergence"},
			reason:      "It is the failure of equally spaced interpolation to converge for smooth functions with large higher derivatives.",
		},
		{
			question:    "Among all choices of nodes, the Chebyshev nodes give the node product...",
			correct:     "the smallest possible maximum absolute value",
			distractors: []string{"the largest possible maximum", "a vanishing derivative", "no zeros"},
			reason:      "This is the minimax property of the Chebyshev polynomial.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// Cubic splines (multiple choice)
// ---------------------------------------------------------------------------

type splinesGen struct{}

func (g *splinesGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "A cubic spline is...",
			correct:     "a piecewise cubic that matches the values and the first two derivatives at the knots",
			distractors: []string{"a single high-degree polynomial", "a piecewise linear function", "a rational function"},
			reason:      "The derivative matching makes the curve smooth across each knot.",
		},
		{
			question:    "Splines are preferred to one interpolating polynomial because...",
			correct:     "they avoid the wild oscillation of high-degree polynomials",
			distractors: []string{"they are exact", "they need no data", "they are linear"},
			reason:      "Low-degree pieces keep the interpolant close to the data between the knots.",
		},
		{
			question:    "The natural cubic spline has...",
			correct:     "zero second derivative at the two endpoints",
			distractors: []string{"zero first derivative at the endpoints", "a prescribed slope everywhere", "no boundary conditions"},
			reason:      "Two extra conditions are needed to pin down the otherwise free end pieces.",
		},
		{
			question:    "The interpolation error of a cubic spline behaves like...",
			correct:     "a constant times the fourth power of the knot spacing",
			distractors: []string{"the square of the knot spacing", "the knot spacing", "a constant"},
			reason:      "The fourth-order behaviour comes from bounding the fourth derivative on each piece.",
		},
		{
			question:    "A spline of degree one is...",
			correct:     "a piecewise linear interpolant",
			distractors: []string{"a quadratic", "a cubic", "a constant"},
			reason:      "Degree counts the highest power used on each piece.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
