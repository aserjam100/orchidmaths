# The conjugate-root theorem

## Why the pairs appear

If a polynomial has real coefficients and \(a + bi\) is a root, then \(a - bi\) is also a root. The reason is that conjugation commutes with the arithmetic: conjugating every coefficient of an equation whose right-hand side is zero leaves the right-hand side at zero, and the coefficients do not change because they are real.

So the non-real roots of a real polynomial always arrive in pairs.

## Building the quadratic factor

The two roots together give a quadratic factor with real coefficients:

\[
(x - (a+bi))(x - (a-bi)) = x^2 - 2ax + (a^2 + b^2).
\]

For the roots \(3 + 4i\) and \(3 - 4i\) that is \(x^2 - 6x + 25\), so a coefficient of -6 on \(x\) and a constant of 25 come from them.

The factor is irreducible over the reals, which is why it cannot be split further by factoring.

## What it implies for odd degrees

A polynomial of odd degree with real coefficients always has at least one real root, because the non-real roots come in pairs and an odd count cannot be made of pairs alone. That is the argument behind the graph of every odd-degree polynomial crossing the axis at least once.

For even degrees no such guarantee exists: \(x^2 + 1\) has no real root at all, and its two roots form one conjugate pair.
