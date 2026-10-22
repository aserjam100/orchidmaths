# Algebraic and Transcendental Elements

**Algebraic element:** Let $F$ be a field and let $\alpha$ lie in an extension of $F$. Then $\alpha$ is algebraic over $F$ if it is a root of some nonzero polynomial with coefficients in $F$, and transcendental over $F$ otherwise.

## Elements that satisfy a polynomial

The test is existence: produce one nonzero polynomial that vanishes at $\alpha$. Every rational number is algebraic over $Q$, and so is every root of an integer polynomial. Transcendence is a negative statement and therefore harder, since it says that no nonzero rational polynomial has the element as a root. The numbers $\pi$ and $e$ are transcendental over $Q$, and so is $1/\pi$.

## Worked: sqrt(2) is algebraic, pi is not

1. $\sqrt{2}$ satisfies $x^2 - 2 = 0$, a nonzero polynomial over $Q$, so it is algebraic.
2. $\sqrt{3}$ satisfies $x^2 - 3 = 0$, so it is algebraic too.
3. $\pi$ satisfies no nonzero polynomial with rational coefficients, so it is transcendental.

## The minimal polynomial sets the degree

Among the polynomials vanishing at an algebraic $\alpha$, the monic one of least degree is the minimal polynomial, and it is irreducible over $F$. The extension $F(\alpha)$ consists of polynomials in $\alpha$ of degree below that minimum, so its degree over $F$ equals the degree of the minimal polynomial. For $\alpha = \sqrt{2}$ the minimal polynomial $x^2 - 2$ has degree 2, so $[Q(\sqrt{2}) : Q] = 2$ and $Q(\sqrt{2})$ is a finite extension of $Q$.
