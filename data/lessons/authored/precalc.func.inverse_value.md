# Inverse functions

## Undoing the function

The inverse of \(f\), written \(f^{-1}\), undoes it: \(f^{-1}(f(x)) = x\) and \(f(f^{-1}(y)) = y\). Only one-to-one functions have inverses, because an inverse has to send each output back to a single input.

For \(f(x) = 3x + 2\), solving \(3x + 2 = y\) gives \(f^{-1}(y) = (y-2)/3\), and \(f^{-1}(11) = 3\).

The one-to-one condition is the whole requirement, and a counterexample makes it concrete. The function \(x^2\) on all of \(\mathbb{R}\) sends both 3 and \(-3\) to 9, so an inverse would have to send 9 to two different numbers at once, which no function does. Restricting the domain to \(x\ge0\) removes the ambiguity and gives the inverse \(\sqrt{y}\), defined for \(y\ge0\).

The procedure for finding an inverse is always the same: write \(y=f(x)\), solve for \(x\) in terms of \(y\), then swap the names. A rational example shows why the swap is bookkeeping rather than algebra: with

\[f(x) = \frac{2x+1}{x-3},\]

the equation \(y(x-3)=2x+1\) rearranges to \(x(y-2)=3y+1\), so \(x=(3y+1)/(y-2)\) and

\[f^{-1}(y) = \frac{3y+1}{y-2}, \qquad y\neq2.\]

The check is the definition, not a rearrangement: \(f(1)=(3)/(-2)=-3/2\) and \(f^{-1}(-3/2)=(-9/2+1)/(-3/2-2)=(-7/2)/(-7/2)=1\). Both compositions return the input, which is what the two equations in the definition demand.

## Undoing in reverse order

The inverse of a chain of operations undoes them in the opposite order: if \(f\) multiplies by 3 and then adds 2, its inverse subtracts 2 and then divides by 3. That reversal is why the inverse is not simply the expression read left to right but the operations performed backwards.

The rule behind it is a composition identity. If \(f=g\circ h\), meaning \(f(x)=g(h(x))\), then untying the outside function first gives \(f^{-1}=h^{-1}\circ g^{-1}\): apply \(g^{-1}\), then \(h^{-1}\). Putting on socks and then shoes is the usual illustration, and its inverse is shoes off before socks off.

A chain with a root shows the same reversal with a domain restriction attached. Take

\[f(x) = \sqrt{2x+5},\]

which is defined for \(x\ge-5/2\) and takes values \(y\ge0\). Squaring first and then undoing the inside operation gives \(y^2=2x+5\), so

\[f^{-1}(y) = \frac{y^2-5}{2}, \qquad y\ge0.\]

The restriction \(y\ge0\) is not decoration. Without it the expression \((y^2-5)/2\) would accept \(y=-3\) and return 2, but \(f(2)=\sqrt{9}=3\), not \(-3\), because the square root is defined to be the nonnegative one. The inverse of a function is a function only on the range of the original, and that is where the restriction comes from.

## Graphs and domains

The graph of \(f^{-1}\) is the graph of \(f\) reflected in the line \(y = x\), because swapping the roles of input and output swaps the axes. A restriction on the domain of \(f\) becomes a restriction on the range of \(f^{-1}\), and that restriction is often what makes the function one-to-one in the first place.

The reflection is exact rather than approximate: the point \((a,b)\) is on the graph of \(f\) precisely when \((b,a)\) is on the graph of \(f^{-1}\), and \((b,a)\) is the mirror image of \((a,b)\) in \(y=x\). That gives the horizontal line test its reading. A horizontal line \(y=b\) meets the graph of \(f\) once exactly when \(b\) is in the range and the preimage is unique, which is the same statement as the vertical line test applied to the reflected graph.

Domain and range trade places across the pair, so the swap is worth stating separately:

\[\operatorname{dom}(f^{-1}) = \operatorname{range}(f), \qquad \operatorname{range}(f^{-1}) = \operatorname{dom}(f).\]

For \(f(x)=x^2\) with domain \([0,\infty)\), the range is \([0,\infty)\) and \(f^{-1}(y)=\sqrt{y}\) has the same interval as its domain — the symmetry is a coincidence of this example, not a rule. For \(f(x)=e^{x}\), the domain is \(\mathbb{R}\) and the range is \((0,\infty)\), so the inverse \(\ln y\) has domain \((0,\infty)\) and range \(\mathbb{R}\).

One sufficient condition is worth carrying: a strictly increasing function is one-to-one, so it has an inverse on its range, and the same holds for strictly decreasing. That covers every example above without drawing a line. It does not cover every invertible function — \(f(x)=1/x\) is decreasing on each side of zero separately — so the horizontal line test remains the general statement.
