# The Mean Value Theorem

**Mean value theorem:** a differentiable function on a closed interval has a point where the instantaneous rate equals the average rate over the interval, so the tangent is parallel to the secant.

## The statement

If a function is continuous on a closed interval and differentiable on its interior, then there is at least one interior point where the derivative equals the average slope over the interval: the change in the values divided by the length of the interval. Rolle's theorem is the special case where the endpoint values are equal, so the average slope is zero and the derivative vanishes somewhere inside. The theorem is the bridge between local information, the derivative at a point, and global information, the change over an interval.

## Proof from Rolle's theorem

The mean value theorem follows from Rolle's theorem by tilting. Subtract the linear function whose graph is the secant line from the original function; the difference vanishes at both endpoints, so Rolle's theorem gives an interior point where its derivative is zero. That derivative is the original derivative minus the constant slope of the secant, so the original derivative equals the average slope at that point. This is why the mean value theorem is never proved from first principles in practice: it is Rolle's theorem with the secant removed.

## Consequences

The mean value theorem is the source of most inequalities between function values and derivatives. It shows that a function with zero derivative on an interval is constant there, that a positive derivative makes a function increasing, and that the difference between two values is bounded by the largest derivative times the distance. It also underlies the error term of the linear approximation and the convergence proofs for Newton's method, because both compare the actual change with the derivative times the step.
