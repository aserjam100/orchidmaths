# Dirichlet's Test

**Dirichlet's test:** a series whose factors are bounded partial sums and a monotone sequence tending to zero converges, extending the alternating-series test to non-alternating cancellation.

## The two hypotheses

Dirichlet's test applies to a series whose term is the product of two factors. The first factor must have bounded partial sums, which supplies cancellation; the second must decrease monotonically to zero, which supplies the shrinking size. Together they force convergence even when the first factor has no limit at all, as the sine of the index shows. The alternating-series test is the special case where the first factor alternates between one and minus one, whose partial sums are bounded by one.

## The sine example

The series of sine of the index divided by the index converges by Dirichlet's test, even though the numerator does not tend to zero and the series is not absolutely convergent. The partial sums of the sine stay bounded by a constant, because the sum of a fixed number of sines telescopes through the sine half-angle formula, and the factor one over the index decreases to zero. This example is the reason the test exists: no other standard test reaches a conclusion here, because the terms are not eventually positive and the absolute series diverges.

## Abel's test and the family

Abel's test is the companion result for a monotone factor that decreases to a positive limit rather than to zero. Subtracting that limit leaves a factor tending to zero, and since the partial sums of the first factor stay bounded by adding a constant, Dirichlet's test applies to the remainder. The two tests, together with the alternating-series test, form the family of cancellation tests, and choosing between them is a matter of whether the shrinking factor tends to zero or to a positive number.
