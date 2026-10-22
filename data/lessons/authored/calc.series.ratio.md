# The Ratio Test

## The test

For \(\sum a_{n}\) with \(a_{n} \neq 0\), compute

\[L = \lim_{n\to\infty}\left|\frac{a_{n+1}}{a_{n}}\right|\]

Then \(L < 1\) forces absolute convergence, \(L > 1\) (including \(\infty\)) forces divergence, and \(L = 1\) is inconclusive — the test says nothing and a comparison, integral, or root test is needed.

The ratio is the asymptotic growth factor of the terms. If \(|a_{n+1}|/|a_{n}|\) settles below some \(q < 1\), every term is eventually smaller than \(q\) times the one before it, so the tail is dominated by a geometric series \(q^{N}(1 + q + q^{2} + \cdots)\). That comparison is the proof: \(L<1\) means the terms decay geometrically, and geometric decay is what absolute convergence is.

The other direction is not a comparison but a term test. If \(L>1\), the ratio is eventually above 1, so the magnitudes increase and cannot tend to zero; a series whose terms do not tend to zero diverges. At \(L=\infty\) the same argument holds with no strengthening needed.

## Worked examples

For \(\sum 1/n!\): \(|a_{n+1}/a_{n}| = 1/(n+1) \to 0 < 1\), so the series converges. For \(\sum n!/n^{n}\): the ratio tends to \(1/e < 1\), so it converges too.

The second ratio is worth writing out, because the algebra is the whole method. With \(a_{n} = n!/n^{n}\),

\[\frac{a_{n+1}}{a_{n}} = \frac{(n+1)!}{(n+1)^{n+1}}\cdot\frac{n^{n}}{n!} = \frac{n^{n}}{(n+1)^{n}} = \left(\frac{n}{n+1}\right)^{n} = \frac{1}{(1+1/n)^{n}} \to \frac{1}{e}.\]

The limit is not obvious until the \(n^{n}/(n+1)^{n}\) step, which is the reason factorials and powers are the test's natural subjects: the ratio cancels almost everything.

A ratio above 1 looks like the same computation. For \(\sum n!/2^{n}\) the ratio is \((n+1)/2 \to \infty\), so the terms grow and the series diverges. For \(\sum 2^{n}/n!\) the ratio is \(2/(n+1) \to 0\), so it converges, which is the same pair of terms with the fraction inverted.

The test is useless for a \(p\)-series. With \(a_{n} = 1/n^{p}\),

\[\frac{a_{n+1}}{a_{n}} = \left(\frac{n}{n+1}\right)^{p} = \left(1+\frac{1}{n}\right)^{-p} \to 1,\]

so every \(p\) gives \(L=1\) and the test cannot separate the harmonic series from \(\sum 1/n^{2}\). The same collapse happens for any rational function of \(n\), because the ratio of two polynomials of the same degree tends to 1.

## When it is inconclusive

The case \(L=1\) is not a gap in the proof but a statement about the borderline: growth factors that tend to 1 are exactly the ones a geometric comparison cannot resolve, so the answer depends on how fast the approach to 1 happens. That is why the \(p\)-series sits on the line, converging for \(p>1\) and diverging for \(p\le1\).

The root test resolves some of the same cases and sometimes reports where the ratio cannot: with \(L' = \limsup |a_{n}|^{1/n}\), the conclusions \(L'<1\) and \(L'>1\) are the same as the ratio's, but the limit superior exists more often. For series whose terms contain \(n\)-th powers, the root is the shorter computation, because \(|a_{n}|^{1/n}\) removes the power while the ratio keeps it.

What is left after both is a genuinely borderline comparison. \(\sum 1/(n\log n)\) has ratio and root limit 1, and it diverges only because the integral test can see the logarithm; \(\sum 1/(n(\log n)^{2})\) converges for the same reason. No ratio can separate those two, which is the honest reading of \(L=1\): the test has told you that the terms decay, and nothing about how fast.
