# Fixed-Point Iteration

**Fixed-point iteration:** rewrite the equation as $x = g(x)$ and generate the sequence $x_{k+1} = g(x_k)$, which converges to a fixed point when the slope of $g$ is small enough.

## Turning a root into a fixed point

Any equation can be rearranged into the form x equals some function of x, and a value that equals its own image is a fixed point. Iterating the function from a starting guess walks toward such a value. The rearrangement matters: the same root can be reached by a convergent rearrangement or missed by a divergent one.

## Worked: a geometric iteration

With $g(x) = 1 + x/2$ and a starting value of zero, the iterates are one, one and a half, and one and three quarters. Each step halves the distance to the fixed point, which is two, so the iteration converges steadily. The slope of the right-hand side is one half, and its small size is what drives the convergence.

## The contraction condition

The iteration converges when the absolute slope of the right-hand side is less than one near the fixed point, because each step shrinks the error by that factor. A slope close to one gives slow convergence and a slope above one in absolute value causes divergence. Finding a rearrangement whose slope is small is the practical art of the method.
