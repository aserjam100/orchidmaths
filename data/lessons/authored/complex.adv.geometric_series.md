# Complex geometric series

The series \(1 + r + r^2 + \cdots\) converges exactly when \(|r| < 1\), and then its sum is \(\frac{1}{1-r}\). Complex ratios are no different: the convergence test is on the modulus, and the sum is found by the same algebra as in the real case.

## Convergence by modulus

The partial sums are \(\frac{1 - r^{n+1}}{1-r}\), and \(r^{n+1}\to 0\) exactly when \(|r| < 1\). That is a statement about distance from the origin, so a complex ratio with modulus one half converges even though it rotates at every step.

For \(r = i/2\) the modulus is \(1/2 < 1\), so the series converges, and the sum is a complex number rather than a real one.

## The sum formula

The formula \(S = \frac{1}{1-r}\) is valid for complex \(r\) with \(|r|<1\), and it is the analytic continuation of the polynomial identity \(1 + r + \cdots + r^n = \frac{1-r^{n+1}}{1-r}\).

For \(r = i/2\): \(1 - r = 1 - i/2\), and the sum is \(\frac{1}{1 - i/2} = \frac{2}{2 - i}\).

## Dividing by a conjugate

To write the sum as \(a + bi\), multiply above and below by the conjugate of the denominator: \(\frac{2(2+i)}{(2-i)(2+i)} = \frac{4 + 2i}{5} = \frac{4}{5} + \frac{2}{5}i\).

The denominator is the squared modulus of \(1-r\), which is never zero when \(|r|<1\). The answer has rational parts whenever \(r\) does, which is why these questions are exact rather than approximate.
