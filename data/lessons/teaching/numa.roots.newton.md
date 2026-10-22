# Newton's Method

**Newton's method:** follow the tangent line to the function at the current guess down to the axis, using $x_{k+1} = x_k - f(x_k)/f'(x_k)$.

## Following the tangent

Newton replaces the function near the current guess by its tangent line and takes the point where that line crosses the axis as the next guess. Where the function is close to linear the tangent is a good model, so the guess improves quickly. The step is the function value divided by the derivative, subtracted from the current point.

## Worked: one step for a square root

To solve the equation that the square of x equals c, the derivative is twice x, and the Newton step simplifies to the average of the current guess and c divided by it. For c equal to two and a starting guess of one, the next guess is one and a half, and then about one and four tenths. Each step roughly doubles the number of correct digits.

## Quadratic convergence and its price

Near a simple root Newton's method doubles the number of correct digits each step, a rate called quadratic convergence. It pays for this with the derivative, which must exist and be nonzero, and with a starting guess close enough to the root. Far from the root the tangent can point badly, and the method can overshoot or diverge.
