# Moments from the Characteristic Function

**Moments:** if the characteristic function is differentiable enough at zero, its derivatives give the moments, $\varphi'(0) = i E[X]$ and $\varphi''(0) = -E[X^2]$.

## Differentiating under the expectation

Differentiating the transform with respect to the parameter brings down a factor of $iX$ each time, so the derivatives at zero encode the moments. The first derivative at zero is $i$ times the mean and the second is minus the second moment. This is the same idea as the moment generating function, made safe by the bounded exponential.

## Worked: the Bernoulli variance

A variable taking zero and one with probability one half each has transform one half plus one half times the exponential. Differentiating twice at zero gives the second moment one half and the mean one half, so the variance is one half minus one quarter, which is 0.25. The transform carries the full moment sequence in its derivatives at the origin.

## Second moment and variance

The second derivative gives the second moment directly, and the variance is the second moment minus the square of the first. Because the transform always exists, these moments can be read whenever the derivatives exist, even when the moment generating function does not converge. The only price is the imaginary unit, which disappears in the final real answer.
