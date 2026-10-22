# Comparing growth rates

Two functions can both grow without bound and still be far apart. Comparing them as \(n \to \infty\) is what makes statements about algorithms meaningful, because the constant factors stop mattering and the shape of the growth decides.

## Comparing growth

Write \(f \prec g\) when \(f(n)/g(n) \to 0\): \(g\) eventually dwarfs \(f\), no matter how large the constants. Logarithms are the slowest of the common functions, then powers, then exponentials, then factorials; within each family the exponent decides.

The order is not about a particular \(n\). \(2^n\) starts below \(n^3\) and overtakes it at \(n = 10\), which is why comparing a few values is not a comparison of growth.

## Logs, powers and factorials

\(\log n \prec \sqrt{n} \prec n \prec n \log n \prec n^2 \prec n^3 \prec 2^n \prec n!\). Each step is a statement about a ratio: \(n/\log n \to \infty\), \(n^2/(n \log n) = n/\log n \to \infty\), and \(2^n/n^3 \to \infty\) because an exponential beats any polynomial.

The factorial is the fastest of the common sequences: \(n! \ge (n/2)^{n/2}\), which is already exponential in \(n \log n\). This is why brute-force enumeration over permutations is only viable for small \(n\).

## Limits that decide

The reliable test is the limit of the ratio. If \(f/g \to 0\) then \(f \prec g\); if it tends to a positive constant the two have the same growth; and if it diverges then the inequality runs the other way. L'Hôpital's rule and the standard limits handle most cases once the ratio is written down.

For \(n \log n\) against \(n^{1.5}\) the ratio is \(\log n / \sqrt{n}\), which tends to zero, so \(n \log n\) is the slower one. That single limit is why comparison sorting in \(n \log n\) beats a quadratic sort, and it is the kind of question the exam asks.
