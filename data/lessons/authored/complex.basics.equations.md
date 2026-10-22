# Linear equations over the complex numbers

An equation like \((a+bi)z = c+di\) is solved the same way as its real analogue: divide both sides by the coefficient. The division is multiplication by the reciprocal, which means multiplying above and below by the conjugate.

## Isolating z

The coefficient is a complex number, not a symbol to be moved across the equals sign. Dividing by it is a computation: \(z = \frac{c+di}{a+bi} = \frac{(c+di)(a-bi)}{a^2+b^2}\).

For \((1+i)z = 5+i\): the conjugate of \(1+i\) is \(1-i\) and the denominator is 2, so \(z = \frac{(5+i)(1-i)}{2} = \frac{5 - 5i + i + 1}{2} = \frac{6 - 4i}{2} = 3 - 2i\).

## Dividing by a complex number

The pattern is always the same: multiply the numerator and denominator by the coefficient's conjugate, expand the numerator, and divide both parts by the real denominator. The denominator is the squared modulus and is positive unless the coefficient is zero.

A coefficient of zero makes the equation either impossible (if the right side is nonzero) or true for every \(z\), which is worth naming rather than dividing by zero.

## Checking the answer

Substitute the answer back: \((1+i)(3-2i) = 3 - 2i + 3i - 2i^2 = 3 + i + 2 = 5 + i\). The check is a multiplication, and it catches a sign error in either part where a single component check would not.

Because the solution is a single complex number, "two solutions" is a sign that a quadratic was solved by mistake, or that the equation was not linear.
