package geometry

// Similarity, polygon angles, transformations, and the circle theorems a school geometry course is
// built on.
//
// The domain had the basics, circles, coordinates, quadrilaterals, two solids and five triangle
// concepts -- but nothing on similar figures at all, nothing on the angle sums of polygons, no
// transformation, and only two of the circle theorems. Those are the four blocks here.
//
// Every answer is an integer by construction: the similarity and dilation factors are whole, the
// polygon angles come from exact divisions, the chord uses a Pythagorean triple as offsets from the
// centre, and the special right triangle is the one whose hypotenuse doubles the short leg.
import (
	"fmt"
	"math/rand"

	"github.com/aserjam100/orchid/internal/generator"
)

func smallInt(lo, hi int) int {
	if hi < lo {
		hi = lo
	}
	return lo + rand.Intn(hi-lo+1)
}

func registerSimilarity(reg *generator.Registry) {
	reg.Register("geo.similar.ratio", &similarRatioGen{})
	reg.Register("geo.similar.missing_side", &similarMissingSideGen{})
	reg.Register("geo.transform.dilation", &dilationGen{})
	reg.Register("geo.polygon.interior_angles", &interiorAnglesGen{})
	reg.Register("geo.polygon.exterior_angle", &exteriorAngleGen{})
	reg.Register("geo.circle.inscribed_angle", &inscribedAngleGen{})
	reg.Register("geo.circle.chord", &chordGen{})
	reg.Register("geo.triangle.special_right", &specialRightGen{})
}

// similarRatioGen asks for a side of the larger of two similar figures, given the scale factor.
type similarRatioGen struct{}

func (g *similarRatioGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	k := 2 + rand.Intn(2+int(ctx.Difficulty*3))
	a := 2 + rand.Intn(4+int(ctx.Difficulty*6))
	return generator.Problem{
		Question: fmt.Sprintf("Two triangles are similar, and the ratio of a side of the larger to the corresponding side of the smaller is %d. If that side of the smaller triangle is %d, how long is the corresponding side of the larger?", k, a),
		Answer:   fmt.Sprintf("%d", k*a),
		Explanation: fmt.Sprintf("Similar figures have every pair of corresponding sides in the same ratio, so the side is %d times %d = %d. The angles are unchanged, which is what makes the figures the same shape rather than the same size.",
			k, a, k*a),
		Facts: map[string]string{"side": fmt.Sprintf("%d", k*a), "factor": fmt.Sprintf("%d", k)},
	}
}

// similarMissingSideGen sets the two similar triangles' sides so the unknown is whole.
type similarMissingSideGen struct{}

func (g *similarMissingSideGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	k := 2 + rand.Intn(2+int(ctx.Difficulty*3))
	a := 2 + rand.Intn(4+int(ctx.Difficulty*5))
	b := 2 + rand.Intn(4+int(ctx.Difficulty*5))
	return generator.Problem{
		Question: fmt.Sprintf("Two similar triangles have corresponding sides in the ratio %d to 1. In the smaller triangle, one side is %d and another is %d. What is the length, in the larger triangle, of the side corresponding to the %d?", k, a, b, a),
		Answer:   fmt.Sprintf("%d", k*a),
		Explanation: fmt.Sprintf("Corresponding sides are proportional, so the larger side is %d times the smaller: %d * %d = %d. The other side given, %d, corresponds to %d in the larger triangle and is not needed.",
			k, k, a, k*a, b, k*b),
		Facts: map[string]string{"side": fmt.Sprintf("%d", k*a), "factor": fmt.Sprintf("%d", k),
			"given": fmt.Sprintf("%d", a)},
	}
}

// dilationGen reads a dilation as a scale factor applied to every length.
type dilationGen struct{}

func (g *dilationGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	k := 2 + rand.Intn(3+int(ctx.Difficulty*4))
	l := 2 + rand.Intn(5+int(ctx.Difficulty*6))
	shrink := rand.Intn(3) == 0
	if shrink {
		return generator.Problem{
			Question: fmt.Sprintf("A segment of length %d is dilated by a factor of 1/%d about a point. How long is its image?", k*l, k),
			Answer:   fmt.Sprintf("%d", l),
			Explanation: fmt.Sprintf("A dilation multiplies every length by the factor, here 1/%d: %d / %d = %d. A factor below 1 shrinks the figure but leaves its shape and its angles alone.",
				k, k*l, k, l),
			Facts: map[string]string{"length": fmt.Sprintf("%d", l), "factor": fmt.Sprintf("1/%d", k)},
		}
	}
	return generator.Problem{
		Question: fmt.Sprintf("A segment of length %d is dilated by a factor of %d about a point. How long is its image?", l, k),
		Answer:   fmt.Sprintf("%d", k*l),
		Explanation: fmt.Sprintf("A dilation multiplies every length by the factor, so the image is %d times %d = %d. The image is similar to the original, and its orientation is unchanged.",
			k, l, k*l),
		Facts: map[string]string{"length": fmt.Sprintf("%d", k*l), "factor": fmt.Sprintf("%d", k)},
	}
}

// interiorAnglesGen asks for the interior-angle sum of an n-gon.
type interiorAnglesGen struct{}

func (g *interiorAnglesGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	n := 3 + rand.Intn(3+int(ctx.Difficulty*5))
	return generator.Problem{
		Question: fmt.Sprintf("What is the sum of the interior angles of a polygon with %d sides, in degrees?", n),
		Answer:   fmt.Sprintf("%d", 180*(n-2)),
		Explanation: fmt.Sprintf("Triangulating the polygon from one vertex gives %d - 2 = %d triangles, and each contributes 180 degrees, so the sum is 180 * %d = %d.",
			n, n-2, n-2, 180*(n-2)),
		Facts: map[string]string{"sum": fmt.Sprintf("%d", 180*(n-2)), "sides": fmt.Sprintf("%d", n)},
	}
}

// exteriorAngleGen asks for one exterior angle of a regular polygon, which is a whole number of
// degrees only when the side count divides 360.
type exteriorAngleGen struct{}

func (g *exteriorAngleGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	divisors := []int{3, 4, 5, 6, 8, 9, 10, 12}
	n := divisors[rand.Intn(len(divisors))]
	return generator.Problem{
		Question: fmt.Sprintf("What is the measure, in degrees, of one exterior angle of a regular polygon with %d sides?", n),
		Answer:   fmt.Sprintf("%d", 360/n),
		Explanation: fmt.Sprintf("The exterior angles of any polygon add to 360 degrees, and a regular polygon has all of them equal, so each is 360 / %d = %d.",
			n, 360/n),
		Facts: map[string]string{"angle": fmt.Sprintf("%d", 360/n), "sides": fmt.Sprintf("%d", n)},
	}
}

// inscribedAngleGen uses the inscribed-angle theorem, with an even central angle so the answer is
// whole.
type inscribedAngleGen struct{}

func (g *inscribedAngleGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	central := 2 * (2 + rand.Intn(4+int(ctx.Difficulty*5)))
	return generator.Problem{
		Question: fmt.Sprintf("An inscribed angle in a circle subtends an arc whose central angle is %d degrees. What is the measure of the inscribed angle, in degrees?", central),
		Answer:   fmt.Sprintf("%d", central/2),
		Explanation: fmt.Sprintf("An inscribed angle is half the central angle that subtends the same arc, so it is %d / 2 = %d degrees. That is also why an angle inscribed in a semicircle is a right angle: the central angle there is 180 degrees.",
			central, central/2),
		Facts: map[string]string{"angle": fmt.Sprintf("%d", central/2), "central": fmt.Sprintf("%d", central)},
	}
}

// chordGen uses a Pythagorean distance from the centre, so half the chord is whole.
type chordGen struct{}

func (g *chordGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	triples := [][3]int{{3, 4, 5}, {6, 8, 10}, {5, 12, 13}, {8, 15, 17}}
	tr := triples[rand.Intn(len(triples))]
	// distance from the centre is tr[0], half the chord is tr[1], radius is tr[2]
	return generator.Problem{
		Question: fmt.Sprintf("A chord of a circle lies %d units from the centre, and the circle's radius is %d. How long is the chord?", tr[0], tr[2]),
		Answer:   fmt.Sprintf("%d", 2*tr[1]),
		Explanation: fmt.Sprintf("The perpendicular from the centre bisects the chord, so half the chord, the distance %d and the radius %d form a right triangle: sqrt(%d - %d) = %d, and the chord is twice that, %d.",
			tr[0], tr[2], tr[2]*tr[2], tr[0]*tr[0], tr[1], 2*tr[1]),
		Facts: map[string]string{"chord": fmt.Sprintf("%d", 2*tr[1]), "radius": fmt.Sprintf("%d", tr[2]),
			"distance": fmt.Sprintf("%d", tr[0])},
	}
}

// specialRightGen uses the 30-60-90 triangle, whose hypotenuse is twice the short leg.
type specialRightGen struct{}

func (g *specialRightGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := 2 + rand.Intn(5+int(ctx.Difficulty*6))
	if rand.Intn(2) == 0 {
		return generator.Problem{
			Question: fmt.Sprintf("A 30-60-90 triangle has its shortest leg %d units long. How long is the hypotenuse?", a),
			Answer:   fmt.Sprintf("%d", 2*a),
			Explanation: fmt.Sprintf("In a 30-60-90 triangle the hypotenuse is twice the shortest leg, because the short leg is opposite the 30 degree angle: 2 * %d = %d. The other leg would be %d times the square root of 3.",
				a, 2*a, a),
			Facts: map[string]string{"hypotenuse": fmt.Sprintf("%d", 2*a), "short_leg": fmt.Sprintf("%d", a)},
		}
	}
	b := 2 * (1 + rand.Intn(4+int(ctx.Difficulty*5)))
	return generator.Problem{
		Question: fmt.Sprintf("A 30-60-90 triangle has a hypotenuse %d units long. How long is its shortest leg?", b),
		Answer:   fmt.Sprintf("%d", b/2),
		Explanation: fmt.Sprintf("The shortest leg is opposite the 30 degree angle and is half the hypotenuse: %d / 2 = %d.",
			b, b/2),
		Facts: map[string]string{"hypotenuse": fmt.Sprintf("%d", b), "short_leg": fmt.Sprintf("%d", b/2)},
	}
}
