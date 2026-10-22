package geometry

// Triangle facts and the solid volumes.
//
// The triangle inequality is what decides whether three lengths can close, the midsegment and the
// area ratio are the two consequences of similarity a school course uses most, the tangent length is
// the circle fact that turns into a right triangle, and the pyramid is the solid whose volume
// carries the one-third that a prism's does not.
//
// Every answer is an integer by construction, and each concept ships an authored solution schema
// because geometry is one of the six learner domains.
import (
	"fmt"
	"math/rand"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerTriangles(reg *generator.Registry) {
	reg.Register("geo.triangle.inequality", &triangleInequalityGen{})
	reg.Register("geo.similar.area_ratio", &areaRatioGen{})
	reg.Register("geo.circle.tangent_length", &tangentLengthGen{})
	reg.Register("geo.solid.pyramid_volume", &pyramidVolumeGen{})
}

// triangleInequalityGen asks for the longest or shortest whole third side of a triangle given two
// sides, which is the inequality read as a bound.
type triangleInequalityGen struct{}

func (g *triangleInequalityGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := 2 + rand.Intn(5+int(ctx.Difficulty*6))
	b := 2 + rand.Intn(5+int(ctx.Difficulty*6))
	if a > b {
		a, b = b, a
	}
	if rand.Intn(2) == 0 {
		// the largest whole side, which must be strictly less than the other two added
		return generator.Problem{
			Question: fmt.Sprintf("Two sides of a triangle are %d and %d. What is the largest whole number the third side can be?", a, b),
			Answer:   fmt.Sprintf("%d", a+b-1),
			Explanation: fmt.Sprintf("The triangle inequality requires each side to be less than the sum of the other two, so the third side must be below %d + %d = %d. The largest whole number below that is %d.",
				a, b, a+b, a+b-1),
			Facts: map[string]string{"side": fmt.Sprintf("%d", a+b-1), "given_a": fmt.Sprintf("%d", a),
				"given_b": fmt.Sprintf("%d", b), "bound": fmt.Sprintf("%d", a+b),
				"difference": fmt.Sprintf("%d", b-a)},
		}
	}
	// the smallest whole side, which must exceed the difference
	return generator.Problem{
		Question: fmt.Sprintf("Two sides of a triangle are %d and %d. What is the smallest whole number the third side can be?", a, b),
		Answer:   fmt.Sprintf("%d", b-a+1),
		Explanation: fmt.Sprintf("The third side must be longer than the difference of the other two, which is %d - %d = %d, so the smallest whole number above that is %d.",
			b, a, b-a, b-a+1),
		Facts: map[string]string{"side": fmt.Sprintf("%d", b-a+1), "given_a": fmt.Sprintf("%d", a),
			"given_b": fmt.Sprintf("%d", b), "difference": fmt.Sprintf("%d", b-a),
			"bound": fmt.Sprintf("%d", a+b)},
	}
}

// areaRatioGen asks for the area ratio of similar figures, which is the square of the length ratio.
type areaRatioGen struct{}

func (g *areaRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	k := 2 + rand.Intn(3+int(ctx.Difficulty*4))
	return generator.Problem{
		Question: fmt.Sprintf("Two figures are similar, and a side of the larger is %d times the corresponding side of the smaller. How many times larger is the larger figure's area?", k),
		Answer:   fmt.Sprintf("%d", k*k),
		Explanation: fmt.Sprintf("Areas scale by the square of the length ratio, because a figure has two dimensions: %d^2 = %d. A ratio of volumes would scale by the cube for the same reason.",
			k, k*k),
		Facts: map[string]string{"factor": fmt.Sprintf("%d", k), "area_ratio": fmt.Sprintf("%d", k*k)},
	}
}

// tangentLengthGen uses the right angle between a tangent and its radius, with a Pythagorean
// distance from the external point.
type tangentLengthGen struct{}

func (g *tangentLengthGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	triples := [][3]int{{3, 4, 5}, {6, 8, 10}, {5, 12, 13}, {8, 15, 17}}
	tr := triples[rand.Intn(len(triples))]
	// radius tr[0], tangent length tr[1], distance from the point to the centre tr[2]
	return generator.Problem{
		Question: fmt.Sprintf("A point lies %d units from the centre of a circle of radius %d. A tangent is drawn from the point, touching the circle. How long is the tangent segment?", tr[2], tr[0]),
		Answer:   fmt.Sprintf("%d", tr[1]),
		Explanation: fmt.Sprintf("A tangent meets the radius at a right angle, so the radius %d, the tangent and the distance %d form a right triangle: sqrt(%d - %d) = %d.",
			tr[0], tr[2], tr[2]*tr[2], tr[0]*tr[0], tr[1]),
		Facts: map[string]string{"tangent": fmt.Sprintf("%d", tr[1]), "radius": fmt.Sprintf("%d", tr[0]),
			"distance": fmt.Sprintf("%d", tr[2])},
	}
}

// pyramidVolumeGen uses the one-third in the pyramid volume, with a base area divisible by three.
type pyramidVolumeGen struct{}

func (g *pyramidVolumeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	h := 1 + rand.Intn(4+int(ctx.Difficulty*4))
	// base area is 3h' so the third is exact
	base := 3 * (1 + rand.Intn(4+int(ctx.Difficulty*4)))
	return generator.Problem{
		Question: fmt.Sprintf("A pyramid has a base of area %d square units and a height of %d units. What is its volume, in cubic units?", base, h),
		Answer:   fmt.Sprintf("%d", base*h/3),
		Explanation: fmt.Sprintf("A pyramid's volume is one third of the prism with the same base and height: %d * %d / 3 = %d. The third is the whole difference from a prism, and it holds for any pyramid, not just the square-based ones.",
			base, h, base*h/3),
		Facts: map[string]string{"volume": fmt.Sprintf("%d", base*h/3), "base": fmt.Sprintf("%d", base),
			"height": fmt.Sprintf("%d", h)},
	}
}
