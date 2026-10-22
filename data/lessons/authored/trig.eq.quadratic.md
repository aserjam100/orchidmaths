# Quadratic equations in a trigonometric function

## A quadratic in disguise

An equation like \(2\sin^{2}x + \sin x - 1 = 0\) has three terms, one of which is the square of
another. Reading \(\sin x\) as a single unknown makes that visible: with \(u = \sin x\) the equation
is \(2u^{2} + u - 1 = 0\), an ordinary quadratic.

Nothing about the trigonometry has to happen before the factoring. The substitution is bookkeeping
that keeps the shape in view.

## Factoring and the roots

\(2u^{2} + u - 1\) factors as \((2u - 1)(u + 1)\), so the quadratic is solved by \(u = 1/2\) or
\(u = -1\). Substituting back gives \(\sin x = 1/2\) or \(\sin x = -1\), and the equation is now two
first-degree equations in a ratio.

A root outside \([-1, 1]\) would be discarded here rather than solved, because no angle has such a
sine. That is the one place where a quadratic in a ratio differs from a quadratic in a number.

## Counting the angles in one turn

Each surviving root contributes the angles in the first turn whose sine is that value. A sine of
\(1/2\) is reached twice, at \(30^{\circ}\) and at \(150^{\circ}\); a sine of \(-1\) is reached once,
at \(270^{\circ}\), because the turning points of the wave are the values a horizontal line touches
only at a peak.

So \(2\sin^{2}x + \sin x - 1 = 0\) has three solutions between \(0^{\circ}\) and \(360^{\circ}\). The
count follows from the roots and from which of them the sine reaches at a turning point.
