# The Delta Method

**Delta method:** when a sample mean is asymptotically normal, a differentiable transformation of it is also asymptotically normal, with variance equal to the square of the derivative times the original variance.

## The first-order expansion

Expanding the transformation in a Taylor series about the true mean gives the transformed value as the transform of the mean plus the derivative times the small error, to first order. The error is asymptotically normal, so a linear function of it is too. The derivative therefore scales the asymptotic variance by its square.

## Worked: the square

The sample mean has mean two and variance one, and the transformation is the square. The derivative at two is four, so the asymptotic variance of the square is the derivative squared, sixteen. Squaring a noisy mean magnifies its spread by the square of the slope, which is the whole content of the method.

## When it applies

The transformation must be differentiable at the mean with a nonzero derivative. If the derivative vanishes the linear term disappears and the second-order term dominates, so a different limit law is needed. The method is a first-order tool, and its hypotheses are exactly the ones that make the first-order term the leading one.
