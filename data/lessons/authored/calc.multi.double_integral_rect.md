# Double integrals over a rectangle

## Iterating the integral

An integral over a rectangle is computed as two ordinary integrals, one inside the other:

\[\int_0^a\!\int_0^b f(x,y)\,dy\,dx.\]

The inner integral treats \(x\) as a constant and produces a function of \(x\); the outer one then integrates that over the remaining interval.

## Choosing the order

For \(f(x, y) = x + y\) over \([0,a] \times [0,b]\), integrating in \(y\) first gives \(bx + b^2/2\), and integrating that in \(x\) gives \(ab(a+b)/2\). At \(a = 4\) and \(b = 3\) the value is \(4 \cdot 3 \cdot 7/2 = 42\).

Fubini's theorem says the order does not matter when the function is continuous on the rectangle, so the order is chosen for convenience rather than correctness. One order is often far easier, and for a function that separates as \(g(x)h(y)\) the two integrals can be done independently.

## What the value means

When the integrand is 1, the double integral is the area of the region. For a general integrand, it is the volume under the surface and above the rectangle, which is the same interpretation the single integral has as an area under a curve one dimension down.
