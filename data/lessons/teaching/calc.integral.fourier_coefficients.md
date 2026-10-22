# Fourier Coefficients

**Fourier coefficients:** the numbers that weight the sine and cosine terms in a periodic function's series, computed as integrals of the function against each basis function.

## The coefficients as projections

A periodic function is written as a sum of sines and cosines, and the coefficient of each basis function is the integral of the function against that basis function over a full period, divided by the integral of the basis function squared. The division makes the coefficients the coordinates of the function in an orthogonal basis, exactly as a vector's coordinates are its dot products with the basis vectors. The sines and cosines are orthogonal over the period, so each coefficient is computed independently of the others.

## Even and odd functions

If the function is odd, the cosine coefficients vanish because the product of an odd and an even function is odd and integrates to zero over a symmetric period, and only the sine terms remain. If the function is even, the sine coefficients vanish instead. The function equal to its own input on a symmetric interval is odd, so its series has only sine terms, and its n-th sine coefficient is a simple alternating multiple of the reciprocal of the index. This parity shortcut removes half the integrals before any work is done.

## Convergence and the series

Under mild conditions the Fourier series converges to the function at points of continuity and to the average of the one-sided limits at a jump. The coefficients tend to zero, because a bounded function times a bounded oscillating basis integrates to something small. Squaring the coefficients and summing recovers the mean square of the function, by Parseval's identity, which is the energy statement of the series. The series is the standard way to solve heat and wave problems with periodic data.
