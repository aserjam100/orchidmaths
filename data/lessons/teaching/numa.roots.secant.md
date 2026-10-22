# The Secant Method

**Secant method:** replace the derivative in Newton's step by the slope of the secant line through the two most recent guesses, so no derivative is needed.

## A derivative-free Newton

Newton needs the derivative, which may be expensive or unavailable. The secant method estimates it by the slope of the line through the previous two points. The update is the same as Newton's with that estimated slope, so it keeps most of the speed while using only function values. It needs two starting guesses instead of one.

## Worked: one secant step

For the function whose square is c, two guesses give two function values, and the secant slope replaces the true derivative. With c equal to two and guesses of one and two, the next point is about one and one third, already much closer to the root than either guess. The method trades the exact slope for an approximation that improves as the points converge.

## Order and robustness

The secant method converges with order about one and six tenths, faster than bisection and slower than Newton, without ever needing a derivative. It does not keep a bracket, so it can stray or diverge when the guesses are poor. Once it is near a simple root it is reliable and cheap, which is why it is a common fallback.
