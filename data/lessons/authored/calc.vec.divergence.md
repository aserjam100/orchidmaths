# Divergence

## The diagonal partials

The divergence of a field \((P, Q)\) is

\[\frac{\partial P}{\partial x} + \frac{\partial Q}{\partial y}.\]

It is the trace of the matrix of partials, where the scalar curl is its antisymmetric part, and the two share the same four derivatives.

In three dimensions the definition adds the third diagonal entry, so for \(\mathbf{F}=(P,Q,R)\),

\[\operatorname{div}\mathbf{F} = \frac{\partial P}{\partial x} + \frac{\partial Q}{\partial y} + \frac{\partial R}{\partial z} = \operatorname{tr} J_{\mathbf{F}}.\]

The trace reading explains why divergence is invariant under rotations of the coordinate axes: the trace of a matrix is, so a formula built from it does not depend on which directions were called \(x\) and \(y\). The curl's components are the off-diagonal differences, which are not invariant individually but form a vector that transforms correctly.

A computed example fixes the mechanics. For \(\mathbf{F}(x,y)=(x^{2},\,xy)\),

\[\operatorname{div}\mathbf{F} = \frac{\partial}{\partial x}x^{2} + \frac{\partial}{\partial y}xy = 2x + x = 3x,\]

so the field spreads at the right of the \(y\)-axis and contracts at the left, and along the axis itself the two effects cancel exactly.

The operator has a meaning that the formula only summarises. Take a small square around a point, let the field carry its edges for a moment, and measure how much area the square gains: the rate of gain per unit area is the divergence. That is why the divergence of a velocity field is the rate at which a fluid expands, and why it appears multiplied by the density in the continuity equation — the divergence is the differential form of a conservation law, with the flux integral of the next section as its integral form.

## Sources and sinks

A positive divergence says the field spreads out of a point, so the point acts like a source; a negative one says the field converges into it. For \(\mathbf{F} = (x, y)\) the divergence is \(1 + 1 = 2\) everywhere, which is the field radiating evenly from the origin.

A field can have zero curl and non-zero divergence, or the reverse, so the two operators carry independent information.

The radial field has a counterpart with the same divergence and a different shape. \(\mathbf{F}(x,y)=(2x,\,0)\) also has divergence 2, but it stretches only horizontally, and its picture looks nothing like the radial one. Divergence does not see the difference, because it measures the *net* expansion of a small patch rather than the motion along any particular direction: a patch stretched horizontally and compressed vertically has zero divergence despite moving in both, which is why a shear flow is incompressible.

Zero divergence has its own name: a field with \(\operatorname{div}\mathbf{F}=0\) is divergence-free, or solenoidal, and its flow is incompressible. The rotation field \(\mathbf{F}(x,y)=(-y,\,x)\) is the standard example. Its divergence is \(-1+1=0\) everywhere, so it neither sources nor sinks, while the field is plainly moving at every point except the origin — incompressible is not the same as stationary.

The identity that ties the two operators together is

\[\nabla\cdot(\nabla\times\mathbf{G}) = 0\]

for any twice-differentiable field \(\mathbf{G}\). The divergence of a curl vanishes because each term in it is a difference of two mixed second partials of the same component, and mixed partials commute. The converse statement — that a divergence-free field on a simply connected region is a curl — is the harder half, and it is the one that needs the region's topology.

## Divergence versus curl

Divergence measures the change in area of a small patch as it is carried by the field, and curl measures its rotation. A flow that is incompressible has zero divergence, and one that is irrotational has zero curl, which is why the pair appears together in every physical field theory.

The two are the symmetric and antisymmetric parts of the same derivative matrix, and that decomposition is why a general field is described by both. The matrix of partials of \(\mathbf{F}=(P,Q)\) is

\[J = \begin{pmatrix}P_{x}&P_{y}\\Q_{x}&Q_{y}\end{pmatrix},\]

whose trace \(P_{x}+Q_{y}\) is the divergence and whose off-diagonal antisymmetric part \(Q_{x}-P_{y}\) is the scalar curl. The remaining symmetric off-diagonal part, \(P_{y}+Q_{x}\), is a shear that neither operator reports — the reason two different fields can agree in both divergence and curl.

Helmholtz's theorem is the statement that this much data determines the field: on a suitable region a field is fixed by its divergence, its curl and its behaviour at the boundary. Electrostatics and magnetostatics are that theorem applied twice. \(\operatorname{div}\mathbf{E}=\rho/\varepsilon_{0}\) makes the charge the source of the electric field, \(\nabla\times\mathbf{E}=0\) says the electrostatic field is irrotational, and the magnetic field reverses the pattern: it is divergence-free, and its curl is the current.

The mixed example worth remembering is the field of a point charge, \(\mathbf{F}=\mathbf{r}/r^{3}\). It is curl-free everywhere and divergence-free everywhere except the origin, yet its flux through any closed surface around the origin is the same positive number. The divergence integrates to a nonzero total while vanishing at every point it is defined, which is the finite version of the delta function, and it is why the integral theorems are stated with the point removed.
