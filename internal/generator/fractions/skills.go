package fractions

// Six more fraction concepts: a three-term sum, a fraction of a fraction, the least common
// denominator, decomposing a fraction into unit fractions, cancelling before multiplying, and
// dividing mixed numbers.
//
// These are the skills the domain's first thirty-two concepts left out: the sums that need more than
// one common denominator, the two-step "of" that a recipe calls for, the denominator-finding that
// addition depends on, and the cancellation that keeps the numbers small.
//
// Answers follow the package's convention, a reduced fraction written as `num/den`.
import (
	"fmt"
	"strconv"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerSkills(reg *generator.Registry) {
	reg.Register("frac.add.three", &addThreeGen{})
	reg.Register("frac.of_fraction", &ofFractionGen{})
	reg.Register("frac.lcd", &lcdGen{})
	reg.Register("frac.unit_decompose", &unitDecomposeGen{})
	reg.Register("frac.mult.cancel", &multCancelGen{})
	reg.Register("frac.mixed.div", &mixedDivGen{})
}

// addThreeGen adds three fractions with denominators that share factors, so the common denominator
// stays small.
type addThreeGen struct{}

func (g *addThreeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	b := smallInt(2, 4+int(ctx.Difficulty*5))
	d := b * smallInt(2, 3)
	f := smallInt(2, 4+int(ctx.Difficulty*4))
	a := smallInt(1, b-1)
	c := smallInt(1, d-1)
	e := smallInt(1, f-1)
	// a/b + c/d + e/f over the common denominator b*d*f
	num := a*d*f + c*b*f + e*b*d
	den := b * d * f
	n, dd := reduce(num, den)
	return generator.Problem{
		Question: fmt.Sprintf("Add \\(\\frac{%d}{%d} + \\frac{%d}{%d} + \\frac{%d}{%d}\\), giving your answer as a fraction in lowest terms.", a, b, c, d, e, f),
		Answer:   fmt.Sprintf("%d/%d", n, dd),
		Explanation: fmt.Sprintf("All three fractions go over the common denominator %d: %d + %d + %d = %d, which reduces to %d/%d. Adding three fractions is the same rule applied twice, not a new one.",
			den, a*d*f, c*b*f, e*b*d, num, n, dd),
		Facts: map[string]string{"out_num": strconv.Itoa(n), "out_den": strconv.Itoa(dd)},
	}
}

// ofFractionGen takes a fraction of a fraction, which is what scaling a recipe does twice.
type ofFractionGen struct{}

func (g *ofFractionGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := smallInt(1, 3+int(ctx.Difficulty*4))
	b := smallInt(2, 5+int(ctx.Difficulty*5))
	c := smallInt(1, 3+int(ctx.Difficulty*4))
	d := smallInt(2, 5+int(ctx.Difficulty*5))
	n, dd := reduce(a*c, b*d)
	return generator.Problem{
		Question: fmt.Sprintf("What is \\(\\frac{%d}{%d}\\) of \\(\\frac{%d}{%d}\\)? Give your answer as a fraction in lowest terms.", a, b, c, d),
		Answer:   fmt.Sprintf("%d/%d", n, dd),
		Explanation: fmt.Sprintf("The word of means multiply, so this is %d/%d times %d/%d = %d/%d, which reduces to %d/%d. Multiplying fractions multiplies the numerators and the denominators separately.",
			a, b, c, d, a*c, b*d, n, dd),
		Facts: map[string]string{"out_num": strconv.Itoa(n), "out_den": strconv.Itoa(dd)},
	}
}

// lcdGen asks for the least common denominator, which is the least common multiple of the two
// denominators.
type lcdGen struct{}

func (g *lcdGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	d1 := smallInt(2, 5+int(ctx.Difficulty*5))
	d2 := smallInt(2, 5+int(ctx.Difficulty*5))
	for d2 == d1 {
		d2 = smallInt(2, 5+int(ctx.Difficulty*5))
	}
	l := d1 * d2 / gcdInt(d1, d2)
	return generator.Problem{
		Question: fmt.Sprintf("What is the least common denominator of \\(\\frac{1}{%d}\\) and \\(\\frac{1}{%d}\\)?", d1, d2),
		Answer:   fmt.Sprintf("%d/%d", l, 1),
		Explanation: fmt.Sprintf("The least common denominator is the least common multiple of %d and %d, which is %d. Multiplying the two denominators also gives a common denominator, but not the smallest, and using it means more reducing at the end.",
			d1, d2, l),
		Facts: map[string]string{"out_num": strconv.Itoa(l), "out_den": "1",
			"left_den": strconv.Itoa(d1), "right_den": strconv.Itoa(d2)},
	}
}

// unitDecomposeGen gives a fraction as one half plus a unit fraction and asks for that unit.
type unitDecomposeGen struct{}

func (g *unitDecomposeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	k := smallInt(2, 5+int(ctx.Difficulty*5))
	// 1/2 + 1/k = (k + 2) / (2k)
	n, dd := reduce(k+2, 2*k)
	return generator.Problem{
		Question: fmt.Sprintf("The fraction \\(\\frac{%d}{%d}\\) can be written as \\(\\frac{1}{2} + \\frac{1}{k}\\). What is \\(k\\)?", n, dd),
		Answer:   fmt.Sprintf("%d/%d", k, 1),
		Explanation: fmt.Sprintf("Write the sum over a common denominator: 1/2 + 1/k = (k + 2)/(2k). Setting that equal to %d/%d and matching the denominators gives k = %d, which also satisfies the numerators.",
			n, dd, k),
		Facts: map[string]string{"k": strconv.Itoa(k), "out_num": strconv.Itoa(k), "out_den": "1",
			"target_num": strconv.Itoa(n), "target_den": strconv.Itoa(dd)},
	}
}

// multCancelGen pairs each numerator with a denominator it shares, so cancelling is visible before
// the multiplication.
type multCancelGen struct{}

func (g *multCancelGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := smallInt(1, 3+int(ctx.Difficulty*4))
	b := smallInt(2, 5+int(ctx.Difficulty*4))
	c := smallInt(1, 3+int(ctx.Difficulty*4))
	d := b * smallInt(1, 3)
	n, dd := reduce(a*c, b*d)
	return generator.Problem{
		Question: fmt.Sprintf("Multiply \\(\\frac{%d}{%d} \\times \\frac{%d}{%d}\\), cancelling common factors before multiplying if you can.", a, b, c, d),
		Answer:   fmt.Sprintf("%d/%d", n, dd),
		Explanation: fmt.Sprintf("The second numerator %d and the first denominator %d share a factor, so cancelling turns the product into %d/%d. Cancelling first keeps the numbers small, and the answer is the same as multiplying and reducing afterwards.",
			c, b, n, dd),
		Facts: map[string]string{"out_num": strconv.Itoa(n), "out_den": strconv.Itoa(dd)},
	}
}

// mixedDivGen divides a mixed number by a fraction.
type mixedDivGen struct{}

func (g *mixedDivGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	w := smallInt(1, 3+int(ctx.Difficulty*3))
	b := smallInt(2, 4+int(ctx.Difficulty*4))
	a := smallInt(1, b-1)
	c := smallInt(1, 3+int(ctx.Difficulty*4))
	d := smallInt(2, 4+int(ctx.Difficulty*5))
	// (w a/b) / (c/d) = ((w*b+a)/b) * (d/c)
	n, dd := reduce((w*b+a)*d, b*c)
	return generator.Problem{
		Question: fmt.Sprintf("Divide \\(%d\\frac{%d}{%d} \\div \\frac{%d}{%d}\\), giving your answer as a fraction in lowest terms.", w, a, b, c, d),
		Answer:   fmt.Sprintf("%d/%d", n, dd),
		Explanation: fmt.Sprintf("First turn the mixed number into an improper fraction: %d %d/%d = %d/%d. Then dividing by %d/%d is multiplying by %d/%d, giving %d/%d.",
			w, a, b, w*b+a, b, c, d, d, c, n, dd),
		Facts: map[string]string{"out_num": strconv.Itoa(n), "out_den": strconv.Itoa(dd),
			"improper_num": strconv.Itoa(w*b + a), "improper_den": strconv.Itoa(b)},
	}
}

// gcdInt is the greatest common divisor, used for the least common multiple.
func gcdInt(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}
