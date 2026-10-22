# Second-Order Taylor Approximation

**Second-order Taylor approximation:** At a point $p$, the quadratic model of $f$ is
$$ f(p + d) \approx f(p) + \nabla f(p)\cdot d + \frac12 d^T H_f(p) d, $$
where $d = (dx, dy)$ and $H$ is the Hessian of second partial derivatives. The first two terms are the tangent plane; the last adds curvature.

## From the tangent plane to a quadratic model

The gradient gives the linear term and the Hessian gives the quadratic correction. In components, the correction is $\frac12(f_{xx} dx^2 + 2 f_{xy} dx\,dy + f_{yy} dy^2)$. At a critical point the gradient vanishes, so the Hessian alone decides whether the point is a minimum, a maximum or a saddle.

## Worked: a quadratic expansion

1. At the expansion point, $f = 5$, $f_x = 2$, $f_y = 4$, $f_{xx} = 2$, $f_{xy} = 0$ and $f_{yy} = 2$.
2. For $dx = dy = 0.1$, the linear term is $0.2 + 0.4 = 0.6$.
3. The Hessian term is $\frac12(0.02 + 0 + 0.02) = 0.02$, so the model gives 5.62.

## The Hessian decides the shape

The approximation is exact for a quadratic function; for a general smooth function it is the best quadratic fit at the point. Its error is of third order in the step, so it is reliable only for small $d$. Definiteness of the Hessian classifies the critical point: positive definite is a local minimum, negative definite a local maximum, indefinite a saddle.
