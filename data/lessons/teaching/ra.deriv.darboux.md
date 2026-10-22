# Darboux's Theorem

**Darboux's theorem:** a derivative has the intermediate value property even though it need not be continuous, so it cannot jump from one value to another.

## Derivatives cannot jump

Darboux's theorem says that if a function is differentiable then its derivative takes every value between any two of its values. The derivative need not be continuous, but the intermediate value property survives: unlike a general function, a derivative cannot have a jump discontinuity. The property continuity implies, taking intermediate values, is therefore strictly weaker than continuity, and the derivative is the standard example that separates them. Darboux's theorem is why the derivative is never discontinuous in the stepwise sense.

## Proof idea

The proof fixes two points and a value between the two derivative values, then studies an auxiliary function that subtracts the linear function with that slope. The auxiliary function has a derivative of the required value exactly where the original derivative equals the target, so it suffices to show the auxiliary function has a critical point. Since the derivative of the auxiliary function has one sign at each end, the auxiliary function cannot be monotone across the interval, and a differentiable function that is not monotone has a critical point inside, which is the point required.

## The oscillation example

The function equal to the square of the input times the sine of one over the input, with the value zero at zero, is differentiable at zero and its derivative oscillates without a limit as the input tends to zero. The derivative is therefore discontinuous at zero, which would forbid it under continuity but is permitted under Darboux's theorem because the oscillations do not jump: between any two values the derivative assumes every intermediate value. The example is the sharp witness that the intermediate value property does not imply continuity.
