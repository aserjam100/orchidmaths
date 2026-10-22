# Arguments under multiplication

Multiplying complex numbers multiplies their moduli and adds their arguments. That single fact explains why multiplication rotates and scales the plane, and it is the bridge between the algebra and the geometry.

## Adding arguments

Write \(z = r_1(\cos\alpha + i\sin\alpha)\) and \(w = r_2(\cos\beta + i\sin\beta)\). Multiplying expands to \(r_1 r_2\) times the cosine and sine addition formulas, so the product has modulus \(r_1 r_2\) and argument \(\alpha + \beta\).

So \(z\) with argument \(300^\circ\) and \(w\) with argument \(90^\circ\) give a product with argument \(390^\circ\), which is the same direction as \(30^\circ\).

## Reducing modulo 360

Arguments are directions, so they are defined modulo \(360^\circ\) (or \(2\pi\)). Adding two arguments can leave the interval, and the answer is reduced: \(390^\circ\) becomes \(30^\circ\), and \(-450^\circ\) becomes \(270^\circ\).

The reduction is not a correction but part of the answer: the product's argument is the class of \(390^\circ\) modulo \(360^\circ\), and every representation of that class names the same direction.

## Powers and multiples

Applying the rule repeatedly gives \(\arg(z^n) = n\arg z\) modulo \(360^\circ\), which is De Moivre's theorem in geometric form. Squaring doubles the argument, cubing triples it, and the modulus is raised to the same power.

A consequence worth knowing: multiplying by a complex number of modulus one is exactly a rotation, and multiplying by one of modulus \(r\) is a rotation composed with a scaling by \(r\).
