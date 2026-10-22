# Trigonometric functions from exponentials

Euler's formula \(e^{i\theta} = \cos\theta + i\sin\theta\) gives the exponential form of a complex number. Adding and subtracting the formula at \(\theta\) and at \(-\theta\) isolates the cosine and the sine.

## Cosine as an average

Adding the two formulas: \(e^{i\theta} + e^{-i\theta} = 2\cos\theta\), because the sine terms cancel and the cosines add. So \(\cos\theta = \frac{e^{i\theta} + e^{-i\theta}}{2}\).

That is the identity to know: cosine is the even part of the exponential, its average under \(\theta \mapsto -\theta\).

## Sine as a difference

Subtracting: \(e^{i\theta} - e^{-i\theta} = 2i\sin\theta\), so \(\sin\theta = \frac{e^{i\theta} - e^{-i\theta}}{2i}\). The \(2i\) in the denominator is the part that is easy to lose.

Sine is the odd part of the exponential, and the extra \(i\) is what makes the result real for real \(\theta\).

## Euler's formula behind both

Both identities are Euler's formula read twice, and they are the standard route from trigonometry to exponentials: a product of sines and cosines becomes a sum of exponentials, which is why trig identities simplify so well in this form.

For complex arguments the same formulas define \(\cos z\) and \(\sin z\), and they are then unbounded, unlike their real restrictions.
