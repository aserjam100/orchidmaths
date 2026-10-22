package precalculus

// Polynomials, rational functions and the exponential family.
//
// The domain had conics, parametric and polar curves, sequences, vectors and a couple of special
// functions, but nothing on what a polynomial's zeros are or how they can be found, nothing on a
// rational function's asymptotes, and no exponential or logarithmic concept at all -- which is the
// bridge into calculus that the domain is named for. These are those three blocks.
//
// Every answer is an integer by construction: the polynomial evaluations are integers, the count of
// rational-root candidates is a product of divisor counts, the exponential and logarithmic equations
// use exact powers, and the circle's radius comes from a Pythagorean triple.
import (
	"fmt"
	"math/rand"

	"github.com/aserjam100/orchid/internal/generator"
)

// chooseSmall is the binomial coefficient, computed multiplicatively so it stays exact.
func chooseSmall(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	c := 1
	for i := 0; i < r; i++ {
		c = c * (n - i) / (i + 1)
	}
	return c
}

// divisors counts the positive divisors of n, which is what the rational-root theorem multiplies.
func divisors(n int) int {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return 0
	}
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func smallInt(lo, hi int) int {
	if hi < lo {
		hi = lo
	}
	return lo + rand.Intn(hi-lo+1)
}

// remainderTheoremGen evaluates a cubic at a point, which is what the remainder theorem asks for.
type remainderTheoremGen struct{}

func (g *remainderTheoremGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := smallInt(-3, 2+int(ctx.Difficulty*4))
	b := smallInt(-3, 2+int(ctx.Difficulty*4))
	c := smallInt(-4, 4+int(ctx.Difficulty*4))
	r := smallInt(-2, 2+int(ctx.Difficulty*3))
	v := r*r*r + a*r*r + b*r + c
	return generator.Problem{
		Question: fmt.Sprintf("What is the remainder when \\(p(x) = x^3 + %dx^2 + %dx + %d\\) is divided by \\(x - %d\\)?", a, b, c, r),
		Answer:   fmt.Sprintf("%d", v),
		Explanation: fmt.Sprintf("The remainder theorem says the remainder is p(%d), and substituting gives %d^3 + %d*%d^2 + %d*%d + %d = %d.",
			r, r, a, r, b, r, c, v),
		Facts: map[string]string{"remainder": fmt.Sprintf("%d", v)},
	}
}

// rationalRootsGen asks how many candidates the rational-root theorem produces, which is a product
// of two divisor counts rather than a list.
type rationalRootsGen struct{}

func (g *rationalRootsGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	lead := 1 + rand.Intn(2+int(ctx.Difficulty*4))
	constant := 1 + rand.Intn(4+int(ctx.Difficulty*6))
	n := divisors(constant) * divisors(lead)
	return generator.Problem{
		Question: fmt.Sprintf("A polynomial has leading coefficient %d and constant term %d, both non-zero. How many possible rational roots does the rational-root theorem give, before any are discarded?", lead, constant),
		Answer:   fmt.Sprintf("%d", n),
		Explanation: fmt.Sprintf("Every rational root has the form (divisor of the constant)/(divisor of the leading coefficient). The constant %d has %d positive divisors and %d has %d, so there are %d * %d = %d candidates, each of which may appear with either sign.",
			constant, divisors(constant), lead, divisors(lead), divisors(constant), divisors(lead), n),
		Facts: map[string]string{"candidates": fmt.Sprintf("%d", n)},
	}
}

// rationalAsymptoteGen asks for the vertical asymptote of a simple rational function.
type rationalAsymptoteGen struct{}

func (g *rationalAsymptoteGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := 1 + rand.Intn(3+int(ctx.Difficulty*3))
	b := smallInt(-4, 4+int(ctx.Difficulty*4))
	d := 1 + rand.Intn(2)
	e := d * smallInt(1, 4+int(ctx.Difficulty*4))
	return generator.Problem{
		Question: fmt.Sprintf("Where is the vertical asymptote of \\(f(x) = \\frac{%dx + %d}{%dx - %d}\\)?", a, b, d, e),
		Answer:   fmt.Sprintf("%d", e/d),
		Explanation: fmt.Sprintf("The vertical asymptote is where the denominator vanishes and the numerator does not: %dx - %d = 0 gives x = %d.",
			d, e, e/d),
		Facts: map[string]string{"x": fmt.Sprintf("%d", e/d)},
	}
}

// expEquationGen solves a base raised to an unknown power, with the right-hand side an exact power
// of that base.
type expEquationGen struct{}

func (g *expEquationGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	base := 2 + rand.Intn(3) // 2, 3 or 4
	k := 1 + rand.Intn(3+int(ctx.Difficulty*4))
	v := 1
	for i := 0; i < k; i++ {
		v *= base
	}
	return generator.Problem{
		Question: fmt.Sprintf("Solve \\(%d^x = %d\\) for \\(x\\).", base, v),
		Answer:   fmt.Sprintf("%d", k),
		Explanation: fmt.Sprintf("%d is %d raised to the %d, so the exponents must be equal: x = %d.",
			v, base, k, k),
		Facts: map[string]string{"x": fmt.Sprintf("%d", k)},
	}
}

// logEvaluateGen evaluates a logarithm whose argument is an exact power of the base.
type logEvaluateGen struct{}

func (g *logEvaluateGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	base := 2 + rand.Intn(3)
	k := 1 + rand.Intn(3+int(ctx.Difficulty*4))
	v := 1
	for i := 0; i < k; i++ {
		v *= base
	}
	return generator.Problem{
		Question: fmt.Sprintf("Evaluate \\(\\log_{%d} %d\\).", base, v),
		Answer:   fmt.Sprintf("%d", k),
		Explanation: fmt.Sprintf("A logarithm asks what exponent produces the argument, and %d = %d^%d, so the value is %d.",
			v, base, k, k),
		Facts: map[string]string{"value": fmt.Sprintf("%d", k)},
	}
}

// growthModelGen uses exact doubling or tripling, so the population stays whole.
type growthModelGen struct{}

func (g *growthModelGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	start := 1 + rand.Intn(4+int(ctx.Difficulty*6))
	hours := 1 + rand.Intn(2+int(ctx.Difficulty*3))
	periods := 1 + rand.Intn(3+int(ctx.Difficulty*3))
	factor := 2
	word := "doubles"
	if rand.Intn(2) == 0 {
		factor, word = 3, "triples"
	}
	v := start
	for i := 0; i < periods; i++ {
		v *= factor
	}
	return generator.Problem{
		Question: fmt.Sprintf("A culture starts with %d bacteria and %s every %d hours. How many are there after %d hours?", start, word, hours, hours*periods),
		Answer:   fmt.Sprintf("%d", v),
		Explanation: fmt.Sprintf("The population multiplies by %d once per %d hours, and %d hours is %d such periods, so it is %d * %d^%d = %d.",
			factor, hours, hours*periods, periods, start, factor, periods, v),
		Facts: map[string]string{"population": fmt.Sprintf("%d", v)},
	}
}

// inverseValueGen evaluates an inverse function at a point, which avoids asking for a symbolic
// inverse while testing that the inverse is what was found.
type inverseValueGen struct{}

func (g *inverseValueGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := 2 + rand.Intn(4)
	b := smallInt(-6, 6+int(ctx.Difficulty*6))
	k := smallInt(-3, 3+int(ctx.Difficulty*4))
	// f(x) = a x + b, so f^-1(y) = (y - b)/a; k is the offset so the value is whole.
	y := a*k + b
	return generator.Problem{
		Question: fmt.Sprintf("The function \\(f(x) = %dx + %d\\) is one-to-one, so it has an inverse. Evaluate \\(f^{-1}(%d)\\).", a, b, y),
		Answer:   fmt.Sprintf("%d", k),
		Explanation: fmt.Sprintf("Solve %dx + %d = %d: the inverse undoes the operations in reverse order, so x = (%d - %d)/%d = %d.",
			a, b, y, y, b, a, k),
		Facts: map[string]string{"value": fmt.Sprintf("%d", k)},
	}
}

// circleRadiusGen gives a circle in general form whose radius is whole, built from a Pythagorean
// triple and a second square.
type circleRadiusGen struct{}

func (g *circleRadiusGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type triple struct{ p, q, r int }
	triples := []triple{{3, 4, 5}, {6, 8, 10}, {5, 12, 13}, {8, 15, 17}}
	t := triples[rand.Intn(len(triples))]
	radius := 1 + rand.Intn(2+int(ctx.Difficulty*4))
	f := t.p*t.p + t.q*t.q - radius*radius
	return generator.Problem{
		Question: fmt.Sprintf("The circle \\(x^2 + y^2 + %dx + %dy + %d = 0\\) has radius \\(r\\). Find \\(r\\).", 2*t.p, 2*t.q, f),
		Answer:   fmt.Sprintf("%d", radius),
		Explanation: fmt.Sprintf("Complete the square: (x + a)^2 + (y + b)^2 = r^2 with a = %d and b = %d, so r^2 = %d + %d - %d = %d and r = %d.",
			t.p, t.q, t.p*t.p, t.q*t.q, f, radius*radius, radius),
		Facts: map[string]string{"radius": fmt.Sprintf("%d", radius)},
	}
}
