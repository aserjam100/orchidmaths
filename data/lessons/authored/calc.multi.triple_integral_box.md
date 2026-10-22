# Triple integrals over a box

## Three nested integrals

An integral over a rectangular box is three ordinary integrals applied in turn:

\[\int_0^a\!\int_0^b\!\int_0^c f(x,y,z)\,dz\,dy\,dx.\]

Each inner integral produces a function of the remaining variables, so the work is to choose an order that leaves the easiest integrand at each stage.

## Separating the integrand

When the integrand does not depend on a variable, that variable's integral is just the length of its interval. For \(f(x,y,z) = x\) over a box of side lengths \(a\), \(b\) and \(c\), the integrals in \(z\) and \(y\) contribute \(c\) and \(b\), leaving \(bc\int_0^a x\,dx = bca^2/2\). At \(a = 4\), \(b = 2\), \(c = 3\) that is \(2 \cdot 3 \cdot 16/2 = 48\).

A product integrand \(g(x)h(y)k(z)\) separates completely, so the triple integral is the product of three single integrals, which is the same factoring that makes a continuous joint distribution independent.

## Volumes from the constant integrand

With the integrand 1, the triple integral is the volume of the region, and over a box that is the product of the side lengths. Every other use is a weighted volume: density integrated over a solid gives its mass, and \(x\) integrated over it gives the first moment that locates the centre of mass.
