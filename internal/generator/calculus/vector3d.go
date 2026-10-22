package calculus

import (
	"fmt"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerVector3D(reg *generator.Registry) {
	reg.Register("calc.vec.divergence_theorem", &divergenceTheoremGen{})
	reg.Register("calc.vec.flux_uniform", &fluxUniformGen{})
	reg.Register("calc.vec.potential_value", &potentialValueGen{})
	reg.Register("calc.multi.coordinates_jacobian", &coordinatesJacobianGen{})
}

// divergenceTheoremGen turns a flux through a box into a volume integral of the divergence, which
// for a linear field is constant.
type divergenceTheoremGen struct{}

func (g *divergenceTheoremGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := small(1, 2+int(ctx.Difficulty*3))
	b := small(1, 2+int(ctx.Difficulty*3))
	c := small(1, 2+int(ctx.Difficulty*3))
	p := small(2, 3+int(ctx.Difficulty*3))
	q := small(2, 3+int(ctx.Difficulty*3))
	r := small(2, 3+int(ctx.Difficulty*3))
	flux := (a + b + c) * p * q * r
	return generator.Problem{
		Question: fmt.Sprintf("Use the divergence theorem to find the outward flux of \\(\\mathbf{F} = (%dx,\\ %dy,\\ %dz)\\) through the boundary of the box \\(0 \\le x \\le %d\\), \\(0 \\le y \\le %d\\), \\(0 \\le z \\le %d\\).", a, b, c, p, q, r),
		Answer:   fmt.Sprintf("%d", flux),
		Explanation: fmt.Sprintf("The divergence is %d + %d + %d = %d, constant, so the flux is the divergence times the volume: %d times %d*%d*%d = %d.",
			a, b, c, a+b+c, a+b+c, p, q, r, flux),
		Facts: map[string]string{"volume": fmt.Sprintf("%d", p*q*r), "flux": fmt.Sprintf("%d", flux)},
	}
}

// fluxUniformGen is the same flux without the theorem, for a field that only pushes through one
// pair of faces, so the answer is the field strength times the face area.
type fluxUniformGen struct{}

func (g *fluxUniformGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	c := small(1, 3+int(ctx.Difficulty*4))
	p := small(2, 3+int(ctx.Difficulty*3))
	q := small(2, 3+int(ctx.Difficulty*3))
	return generator.Problem{
		Question: fmt.Sprintf("Find the flux of the constant field \\(\\mathbf{F} = (0, 0, %d)\\) through the rectangle \\(0 \\le x \\le %d\\), \\(0 \\le y \\le %d\\) in the plane \\(z = 0\\), oriented upward.", c, p, q),
		Answer:   fmt.Sprintf("%d", c*p*q),
		Explanation: fmt.Sprintf("The field is perpendicular to the rectangle and has no component through the other faces, so the flux is the field strength times the area: %d * %d * %d = %d.",
			c, p, q, c*p*q),
		Facts: map[string]string{"area": fmt.Sprintf("%d", p*q), "flux": fmt.Sprintf("%d", c*p*q)},
	}
}

// potentialValueGen recovers a potential and evaluates it, which avoids asking for a symbolic
// antiderivative while still testing that the potential is the thing computed.
type potentialValueGen struct{}

func (g *potentialValueGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := small(1, 3+int(ctx.Difficulty*3))
	p := small(1, 4+int(ctx.Difficulty*4))
	q := small(1, 4+int(ctx.Difficulty*4))
	return generator.Problem{
		Question: fmt.Sprintf("The field \\(\\mathbf{F} = (%dx,\\ %dy)\\) is conservative. Its potential \\(f\\) satisfies \\(f(0,0) = 0\\); evaluate \\(f(%d, %d)\\).", 2*a, 2*a, p, q),
		Answer:   fmt.Sprintf("%d", a*(p*p+q*q)),
		Explanation: fmt.Sprintf("Integrating %dx in x gives %dx^2, and the other component confirms there is no y-term, so f = %d(x^2 + y^2). At (%d, %d) that is %d * (%d + %d) = %d.",
			2*a, a, a, p, q, a, p*p, q*q, a*(p*p+q*q)),
		Facts: map[string]string{"value": fmt.Sprintf("%d", a*(p*p+q*q))},
	}
}

// coordinatesJacobianGen evaluates the volume element of a three-dimensional coordinate system: r
// for cylindrical, and rho^2 sin(phi) for spherical, which is whole at the angles that appear here.
type coordinatesJacobianGen struct{}

func (g *coordinatesJacobianGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	if pickMV(2) == 0 {
		r := small(2, 3+int(ctx.Difficulty*4))
		return generator.Problem{
			Question:    fmt.Sprintf("For cylindrical coordinates \\((r, \\theta, z)\\), evaluate the volume element factor \\(r\\) at \\(r = %d\\) (so that \\(dV = r\\,dr\\,d\\theta\\,dz\\)).", r),
			Answer:      fmt.Sprintf("%d", r),
			Explanation: fmt.Sprintf("The matrix of partials has determinant r, so the volume element carries a factor of r and at r = %d the factor is %d.", r, r),
			Facts:       map[string]string{"factor": fmt.Sprintf("%d", r), "system": "cylindrical"},
		}
	}
	rho := small(2, 3+int(ctx.Difficulty*4))
	// sin(pi/2) = 1 and sin(pi/6) = 1/2 are the only angles here that keep the value whole, and
	// the half needs an even radius.
	phi, factor := "\\pi/2", rho*rho
	if pickMV(2) == 0 && rho%2 == 0 {
		phi, factor = "\\pi/6", rho*rho/2
	}
	return generator.Problem{
		Question: fmt.Sprintf("For spherical coordinates \\((\\rho, \\theta, \\varphi)\\), evaluate the volume element factor \\(\\rho^2 \\sin\\varphi\\) at \\(\\rho = %d\\) and \\(\\varphi = %s\\).", rho, phi),
		Answer:   fmt.Sprintf("%d", factor),
		Explanation: fmt.Sprintf("The factor is rho^2 sin(varphi), evaluated at rho = %d and varphi = %s, which is %d. It is largest at the equator and vanishes at the poles, which is why a spherical volume element carries that sine.",
			rho, phi, factor),
		Facts: map[string]string{"factor": fmt.Sprintf("%d", factor), "system": "spherical"},
	}
}
