# Conjugate roots of real polynomials

A polynomial with real coefficients that has \(3 - 2i\) as a root must also have \(3 + 2i\) as a root. The reason is that conjugation commutes with real arithmetic, and it forces the non-real roots to appear in pairs.

## Why pairs appear

For a polynomial \(p\) with real coefficients, \(p(\bar z) = \overline{p(z)}\): conjugating the input conjugates the output, because the coefficients are unchanged by conjugation. So if \(p(z_0) = 0\), then \(\overline{p(z_0)} = 0\), which is \(p(\bar z_0) = 0\).

The pair is therefore not a coincidence of the example but a theorem. Only the pair with the same real part and the opposite imaginary part qualifies; \(-3 + 2i\) is not forced by \(3 - 2i\).

## Multiplicity and degree

The pairing preserves multiplicity: if \(z_0\) is a root of multiplicity \(k\), so is \(\bar z_0\), because the factorisation over the reals pairs \((z - z_0)(z - \bar z_0) = z^2 - 2\operatorname{Re}(z_0) z + |z_0|^2\), a real quadratic.

So a real polynomial factors into linear factors for its real roots and irreducible real quadratics for its conjugate pairs, which is the version of the fundamental theorem of algebra that stays inside the reals.

## Counting roots

A real polynomial of odd degree has at least one real root, because the non-real roots come in pairs and cannot account for an odd count. A polynomial of degree four with two non-real roots therefore has two real roots, counted with multiplicity.

That counting is often enough to answer a question without finding any root: degree minus the number of non-real roots, each counted twice, is the number of real roots.
