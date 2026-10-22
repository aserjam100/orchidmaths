package geometry

// The angle facts, the midsegment, and the prism.
//
// These are the school-course workhorses that were still missing: the exterior angle theorem, the
// midsegment's half, the base angles of an isosceles triangle, the opposite angles of a cyclic
// quadrilateral, and the prism volume that a pyramid's third is measured against.
//
// Every answer is an integer by construction, and each ships an authored solution schema because
// geometry is one of the six learner domains.
import (
	"fmt"
	"math/rand"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerAngleFacts(reg *generator.Registry) {
	reg.Register("geo.triangle.exterior_angle", &triangleExteriorAngleGen{})
	reg.Register("geo.triangle.midsegment", &midsegmentGen{})
	reg.Register("geo.triangle.isosceles", &isoscelesGen{})
	reg.Register("geo.circle.cyclic_quad", &cyclicQuadGen{})
	reg.Register("geo.solid.prism_volume", &prismVolumeGen{})
}

// triangleExteriorAngleGen asks for an exterior angle, which is the sum of the two remote interior angles.
type triangleExteriorAngleGen struct{}

func (g *triangleExteriorAngleGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	// two interior angles that leave a sensible third; the second is clamped rather than redrawn,
	// so the generator cannot spin
	a := 20 + rand.Intn(30+int(ctx.Difficulty*40))
	b := 20 + rand.Intn(30+int(ctx.Difficulty*40))
	if a+b > 150 {
		b = 150 - a
	}
	return generator.Problem{
		Question: fmt.Sprintf("In a triangle, one side is extended to form an exterior angle. The two interior angles not adjacent to it are %d and %d degrees. What is the measure of the exterior angle, in degrees?", a, b),
		Answer:   fmt.Sprintf("%d", a+b),
		Explanation: fmt.Sprintf("The three interior angles sum to 180, so the third one is 180 - %d - %d = %d. The exterior angle and that interior angle are supplementary, so it is 180 - %d = %d, which is the same as the sum of the two remote interior angles.",
			a, b, 180-a-b, 180-a-b, a+b),
		Facts: map[string]string{"exterior": fmt.Sprintf("%d", a+b), "interior_a": fmt.Sprintf("%d", a),
			"interior_b": fmt.Sprintf("%d", b), "third": fmt.Sprintf("%d", 180-a-b)},
	}
}

// midsegmentGen uses the midsegment's half of the third side.
type midsegmentGen struct{}

func (g *midsegmentGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	// an even third side, so the half is whole
	third := 2 * (3 + rand.Intn(8+int(ctx.Difficulty*10)))
	return generator.Problem{
		Question: fmt.Sprintf("A midsegment of a triangle joins the midpoints of two of its sides. The third side is %d units long. How long is the midsegment?", third),
		Answer:   fmt.Sprintf("%d", third/2),
		Explanation: fmt.Sprintf("A midsegment is parallel to the third side and half its length, so it is %d / 2 = %d. That is the same similarity ratio that makes the small triangle at the vertex similar to the whole one with factor 1/2.",
			third, third/2),
		Facts: map[string]string{"midsegment": fmt.Sprintf("%d", third/2),
			"third_side": fmt.Sprintf("%d", third)},
	}
}

// isoscelesGen asks for a base angle of an isosceles triangle, given the vertex angle.
type isoscelesGen struct{}

func (g *isoscelesGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	// an even vertex angle, so the two base angles are whole
	vertex := 2 * (5 + rand.Intn(40+int(ctx.Difficulty*30)))
	if vertex > 178 {
		vertex = 178
	}
	return generator.Problem{
		Question: fmt.Sprintf("An isosceles triangle has a vertex angle of %d degrees. What is the measure of each base angle, in degrees?", vertex),
		Answer:   fmt.Sprintf("%d", (180-vertex)/2),
		Explanation: fmt.Sprintf("The two base angles of an isosceles triangle are equal, and the three angles sum to 180, so each base angle is (180 - %d) / 2 = %d.",
			vertex, (180-vertex)/2),
		Facts: map[string]string{"base_angle": fmt.Sprintf("%d", (180-vertex)/2),
			"vertex_angle": fmt.Sprintf("%d", vertex)},
	}
}

// cyclicQuadGen uses the fact that opposite angles of a cyclic quadrilateral are supplementary.
type cyclicQuadGen struct{}

func (g *cyclicQuadGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := 30 + rand.Intn(50+int(ctx.Difficulty*50))
	if a > 150 {
		a = 150
	}
	return generator.Problem{
		Question: fmt.Sprintf("A quadrilateral is inscribed in a circle, so its four vertices lie on the circle. One of its angles is %d degrees. What is the measure of the opposite angle, in degrees?", a),
		Answer:   fmt.Sprintf("%d", 180-a),
		Explanation: fmt.Sprintf("Opposite angles of a cyclic quadrilateral are supplementary, because they subtend the two arcs that together make the whole circle and each is half of its arc: 180 - %d = %d.",
			a, 180-a),
		Facts: map[string]string{"opposite": fmt.Sprintf("%d", 180-a), "angle": fmt.Sprintf("%d", a)},
	}
}

// prismVolumeGen is the comparison figure a pyramid's third is taken against.
type prismVolumeGen struct{}

func (g *prismVolumeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	base := 2 + rand.Intn(10+int(ctx.Difficulty*15))
	h := 2 + rand.Intn(6+int(ctx.Difficulty*8))
	return generator.Problem{
		Question: fmt.Sprintf("A prism has a base of area %d square units and a height of %d units. What is its volume, in cubic units?", base, h),
		Answer:   fmt.Sprintf("%d", base*h),
		Explanation: fmt.Sprintf("A prism's volume is the base area times the height, because every cross-section parallel to the base is the same: %d * %d = %d. The shape of the base does not matter, only its area.",
			base, h, base*h),
		Facts: map[string]string{"volume": fmt.Sprintf("%d", base*h), "base": fmt.Sprintf("%d", base),
			"height": fmt.Sprintf("%d", h)},
	}
}
