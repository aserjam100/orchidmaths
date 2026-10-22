# Fixed points of Möbius transformations

A **Möbius transformation** is \(f(z) = \frac{az+b}{cz+d}\) with \(ad - bc \neq 0\), and a **fixed point** is a \(z\) with \(f(z) = z\). Setting the two equal clears the denominator into a quadratic, so a Möbius transformation has at most two fixed points.

## Setting z equal to f(z)

The equation \(z = \frac{az+b}{cz+d}\) becomes \(z(cz+d) = az+b\) once the denominator is cleared, and that is \(cz^2 + (d-a)z - b = 0\).

The coefficient of \(z^2\) is \(c\). When \(c = 0\) the map is affine and the equation is linear, with one fixed point unless the map is the identity.

## The quadratic that follows

For \(f(z) = \frac{2z+3}{z+4}\): \(z(z+4) = 2z+3\), so \(z^2 + 2z - 3 = 0\), and the roots are 1 and \(-3\).

The two roots are the fixed points, and they can be real, complex, or a repeated pair. A repeated fixed point is called parabolic, and the map is then conjugate to a translation.

## Choosing the root

When a question asks for one fixed point, the rule that selects it has to be stated: "the one with the larger real part" is a total order on two distinct roots unless they share a real part, and then the imaginary part decides.

For the example above the larger real part belongs to 1, so the answer is 1, not \(-3\). The choice matters because the two points are different answers to the same equation.
