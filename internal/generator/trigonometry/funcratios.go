package trigonometry

// The six right-triangle ratios turned into computed questions.
//
// Each of these concepts used to be a fixed question with a fixed answer. The
// ratio itself is the definition, so the honest computed question is not "what
// is the ratio" but "given the ratio, what is the side": a triangle similar to
// one of the Pythagorean triples fixes the ratio and a scale factor fixes the
// side, so the answer is always an exact integer rather than a rounded decimal.
// Every generator publishes the ratio, the side it starts from and the answer,
// which is what the solution schema interpolates.

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/aserjam100/orchid/internal/generator"
)

type ratioTriple struct {
	a, b, hyp int
}

var ratioTriples = []ratioTriple{
	{3, 4, 5}, {5, 12, 13}, {8, 15, 17}, {7, 24, 25},
	{9, 40, 41}, {20, 21, 29}, {12, 35, 37}, {28, 45, 53},
}

var ratioScales = []int{3, 4, 5, 6, 8, 10}

// pickRightTriangle returns the leg opposite the angle, the leg adjacent to
// it, the hypotenuse of the smallest triple and a scale factor. The triangle
// actually presented is the triple multiplied by the scale.
func pickRightTriangle() (opposite, adjacent, hypotenuse, scale int) {
	t := ratioTriples[rand.Intn(len(ratioTriples))]
	if rand.Intn(2) == 0 {
		opposite, adjacent = t.a, t.b
	} else {
		opposite, adjacent = t.b, t.a
	}
	hypotenuse = t.hyp
	scale = ratioScales[rand.Intn(len(ratioScales))]
	return opposite, adjacent, hypotenuse, scale
}

func ratioProblem(question, answer, explanation string, facts map[string]string) generator.Problem {
	return generator.Problem{
		Question:    question,
		Answer:      answer,
		Explanation: explanation,
		Facts:       facts,
	}
}

// ----- sine: opposite over hypotenuse -----

type sineRatioGen struct{}

func (g *sineRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	opposite, _, hypotenuse, scale := pickRightTriangle()
	ratio := fmt.Sprintf("%d/%d", opposite, hypotenuse)
	given := hypotenuse * scale
	answer := opposite * scale
	return ratioProblem(
		fmt.Sprintf("A right triangle has \\(\\sin\\theta = \\frac{%d}{%d}\\). If its hypotenuse is %d, what is the length of the side opposite \\(\\theta\\)?",
			opposite, hypotenuse, given),
		strconv.Itoa(answer),
		fmt.Sprintf("The sine of an angle is the side opposite it over the hypotenuse, so the opposite side is that fraction of the hypotenuse: \\(\\frac{%d}{%d} \\times %d = %d\\).",
			opposite, hypotenuse, given, answer),
		map[string]string{
			"ratio":  ratio,
			"given":  strconv.Itoa(given),
			"answer": strconv.Itoa(answer),
		},
	)
}

// ----- cosine: adjacent over hypotenuse -----

type cosineRatioGen struct{}

func (g *cosineRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	_, adjacent, hypotenuse, scale := pickRightTriangle()
	ratio := fmt.Sprintf("%d/%d", adjacent, hypotenuse)
	given := hypotenuse * scale
	answer := adjacent * scale
	return ratioProblem(
		fmt.Sprintf("A right triangle has \\(\\cos\\theta = \\frac{%d}{%d}\\). If its hypotenuse is %d, what is the length of the side adjacent to \\(\\theta\\)?",
			adjacent, hypotenuse, given),
		strconv.Itoa(answer),
		fmt.Sprintf("The cosine of an angle is the side adjacent to it over the hypotenuse, so the adjacent side is that fraction of the hypotenuse: \\(\\frac{%d}{%d} \\times %d = %d\\).",
			adjacent, hypotenuse, given, answer),
		map[string]string{
			"ratio":  ratio,
			"given":  strconv.Itoa(given),
			"answer": strconv.Itoa(answer),
		},
	)
}

// ----- tangent: opposite over adjacent -----

type tangentRatioGen struct{}

func (g *tangentRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	opposite, adjacent, _, scale := pickRightTriangle()
	ratio := fmt.Sprintf("%d/%d", opposite, adjacent)
	given := adjacent * scale
	answer := opposite * scale
	return ratioProblem(
		fmt.Sprintf("A right triangle has \\(\\tan\\theta = \\frac{%d}{%d}\\). If the side adjacent to \\(\\theta\\) is %d, what is the length of the opposite side?",
			opposite, adjacent, given),
		strconv.Itoa(answer),
		fmt.Sprintf("The tangent of an angle is the side opposite it over the side adjacent to it, so the opposite side is that fraction of the adjacent side: \\(\\frac{%d}{%d} \\times %d = %d\\).",
			opposite, adjacent, given, answer),
		map[string]string{
			"ratio":  ratio,
			"given":  strconv.Itoa(given),
			"answer": strconv.Itoa(answer),
		},
	)
}

// ----- secant: hypotenuse over adjacent -----

type secantRatioGen struct{}

func (g *secantRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	_, adjacent, hypotenuse, scale := pickRightTriangle()
	ratio := fmt.Sprintf("%d/%d", hypotenuse, adjacent)
	given := adjacent * scale
	answer := hypotenuse * scale
	return ratioProblem(
		fmt.Sprintf("A right triangle has \\(\\sec\\theta = \\frac{%d}{%d}\\). If the side adjacent to \\(\\theta\\) is %d, what is the hypotenuse?",
			hypotenuse, adjacent, given),
		strconv.Itoa(answer),
		fmt.Sprintf("The secant is the reciprocal of the cosine, so it is the hypotenuse over the adjacent side. The hypotenuse is therefore that fraction of the adjacent side: \\(\\frac{%d}{%d} \\times %d = %d\\).",
			hypotenuse, adjacent, given, answer),
		map[string]string{
			"ratio":  ratio,
			"given":  strconv.Itoa(given),
			"answer": strconv.Itoa(answer),
		},
	)
}

// ----- cosecant: hypotenuse over opposite -----

type cosecantRatioGen struct{}

func (g *cosecantRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	opposite, _, hypotenuse, scale := pickRightTriangle()
	ratio := fmt.Sprintf("%d/%d", hypotenuse, opposite)
	given := opposite * scale
	answer := hypotenuse * scale
	return ratioProblem(
		fmt.Sprintf("A right triangle has \\(\\csc\\theta = \\frac{%d}{%d}\\). If the side opposite \\(\\theta\\) is %d, what is the hypotenuse?",
			hypotenuse, opposite, given),
		strconv.Itoa(answer),
		fmt.Sprintf("The cosecant is the reciprocal of the sine, so it is the hypotenuse over the opposite side. The hypotenuse is therefore that fraction of the opposite side: \\(\\frac{%d}{%d} \\times %d = %d\\).",
			hypotenuse, opposite, given, answer),
		map[string]string{
			"ratio":  ratio,
			"given":  strconv.Itoa(given),
			"answer": strconv.Itoa(answer),
		},
	)
}

// ----- cotangent: adjacent over opposite -----

type cotangentRatioGen struct{}

func (g *cotangentRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	opposite, adjacent, _, scale := pickRightTriangle()
	ratio := fmt.Sprintf("%d/%d", adjacent, opposite)
	given := opposite * scale
	answer := adjacent * scale
	return ratioProblem(
		fmt.Sprintf("A right triangle has \\(\\cot\\theta = \\frac{%d}{%d}\\). If the side opposite \\(\\theta\\) is %d, what is the length of the adjacent side?",
			adjacent, opposite, given),
		strconv.Itoa(answer),
		fmt.Sprintf("The cotangent is the reciprocal of the tangent, so it is the adjacent side over the opposite side. The adjacent side is therefore that fraction of the opposite side: \\(\\frac{%d}{%d} \\times %d = %d\\).",
			adjacent, opposite, given, answer),
		map[string]string{
			"ratio":  ratio,
			"given":  strconv.Itoa(given),
			"answer": strconv.Itoa(answer),
		},
	)
}

// registerFuncRatios registers the six ratio concepts that used to be fixed
// questions. The old registrations are removed from Register, because a concept
// id may be registered only once.
func registerFuncRatios(reg *generator.Registry) {
	reg.Register("trig.func.sine", &sineRatioGen{})
	reg.Register("trig.func.cosine", &cosineRatioGen{})
	reg.Register("trig.func.tangent", &tangentRatioGen{})
	reg.Register("trig.func.secant", &secantRatioGen{})
	reg.Register("trig.func.cosecant", &cosecantRatioGen{})
	reg.Register("trig.func.cotangent", &cotangentRatioGen{})
}
