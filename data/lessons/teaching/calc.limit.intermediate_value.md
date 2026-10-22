# The Intermediate Value Theorem in Practice

**Intermediate value theorem:** a continuous function on an interval takes every value between the values at the endpoints, so a sign change forces a root.

## What the theorem promises

A continuous function on a closed interval attains every value between the values at the two ends. The practical consequence is the root-existence test: if a continuous function is negative somewhere and positive somewhere, then it comes back through zero between those points, so a root exists. The theorem tells you that a root is there; it does not tell you where, and it does not apply to a function with a break in the interval. Continuity on the whole interval is the hypothesis that cannot be dropped, because a function with a jump can skip a value entirely.

## Turning the theorem into arithmetic

For the square-root equation, the function that subtracts a constant from the square of its input is continuous everywhere and is negative at zero and positive past the unknown root. The intermediate value theorem then says the root lies in the interval, and the equation itself says the root is the square root of the constant. Reading a root out this way is a single substitution once the endpoints and the target value are known, which is why the theorem is studied as a computation and not only as an existence statement.

## Bracketing and the practical use

Numerical solvers use the theorem as a licence to search. Every method that begins with a bracketing interval, where the function has opposite signs at the two ends, is relying on the theorem to guarantee that a root exists inside. Bisection then halves the interval and preserves the sign change, so the bracket always contains a root. The theorem does not guarantee uniqueness, and a function may cross several times inside one bracket, which is why a solver that needs a particular root is given a small bracket or a good first guess.
