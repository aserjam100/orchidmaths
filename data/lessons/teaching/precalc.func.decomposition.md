# Decomposing a Function

**Decomposition:** Writing a function $h$ as a composition $h(x) = f(g(x))$, choosing an outer function $f$ and an inner function $g$. It is composition in reverse, and the decomposition is not unique.

## Composition in reverse

To decompose, look for the outermost operation: the last thing done to the input. That operation becomes $f$, and everything inside it becomes $g$. For $h(x) = (x+1)^2$, squaring is outermost, so $f(u) = u^2$ and $g(x) = x+1$.

## Worked: (x + 1)^2

1. Identify the outer operation: squaring.
2. Set $f(u) = u^2$ and $g(x) = x + 1$.
3. Check: $f(g(x)) = (x+1)^2 = h(x)$.

## More than one valid decomposition

The same $h$ can be decomposed in several ways. $(x+1)^2$ is also $f(g(x))$ with $f(u) = u$ and $g(x) = (x+1)^2$, or with $f(u) = (u+1)^2$ and $g(x) = x$. The useful decomposition names a genuinely simpler outer function, which is why the derivative and integral rules are stated in terms of an inner function.

A good check is to compose the proposed pair and compare domains, not only formulas. Cancelling or simplifying inside $g$ can lose a value that the original $h$ excluded, so the decomposition is valid only where both the inner expression and the outer function are defined.
