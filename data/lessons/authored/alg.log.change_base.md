# Change of base

A logarithm with an awkward base can be rewritten with any other base. The change-of-base formula turns the logarithm into a ratio of two logarithms, and when both numbers are powers of the same base the ratio is exact.

## Write both numbers as powers

If the base and the argument are powers of a common number \(a\), write

\[
B=a^m,\qquad N=a^n.
\]

Then \(\log_B N\) asks which exponent on \(B\) produces \(N\). Since \(B=a^m\), the exponent \(m\) is applied to \(a\), and the target exponent is \(n\).

## Division subtracts exponents

The change-of-base formula is

\[
\log_B N=\frac{\log N}{\log B}.
\]

With \(B=a^m\) and \(N=a^n\), the ratio is

\[
\frac{\log a^n}{\log a^m}=\frac{n}{m}.
\]

The common factor \(\log a\) cancels, which is why the result does not depend on which logarithm the formula uses. Any base gives the same value.

## The answer is a ratio

The answer is the ratio of the two exponents, \(n/m\), not the difference and not the product. Reduce the ratio, and leave it as a fraction when it is not a whole number.

A useful check is to substitute: \(B^{n/m}=N\). If that equation holds, the ratio is right. The sign of the answer follows the sizes: an argument smaller than the base gives a negative logarithm, and here that cannot happen because both exponents are positive.
