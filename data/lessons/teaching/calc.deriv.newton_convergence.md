# Convergence of Newton's Method

**Newton's method:** the iteration that steps from each guess to where its tangent line meets the axis, converging quadratically near a simple root when the function is smooth.

## The step and the plant

Newton's method replaces the function near a guess by its tangent line and jumps to the point where that line crosses the axis. The step is the value of the function divided by its derivative, subtracted from the current guess. Geometrically the guess is projected along the tangent, and near a simple root the tangent points almost directly at the root, so each step roughly squares the distance. The method is the standard fast solver for equations, and the derivative condition that the slope not vanish at the root is what keeps the step from blowing up.

## Why the error squares

The error after one step is the error before, squared, times a factor built from the second and first derivatives. When the function is twice differentiable and the slope at the root is not zero, that factor is bounded near the root, so the error is roughly squared at each step. Squaring a small number makes it drastically smaller, which is the quadratic convergence that makes Newton's method worth its cost. The factor is the curvature-to-slope ratio, so a nearly straight function converges fastest.

## When it fails

Newton's method can cycle, run away, or converge to a different root. A zero derivative near the guess makes the step enormous. A poor starting point can send the iteration far from the intended root, and a function with several roots may converge to one the solver did not want. The usual fixes are a good starting bracket obtained from the intermediate value theorem, a damped step that refuses to leave the bracket, and a check on the derivative. The convergence proof is local, so a bad start voids its guarantee.
