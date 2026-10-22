# The slope of a polar curve

## The derivative in polar form

A polar curve has \(x = r\cos\theta\) and \(y = r\sin\theta\) with \(r\) a function of the angle, so it is a parametrised curve in disguise. Differentiating both and dividing gives

\[
\frac{dy}{dx} = \frac{r'\sin\theta + r\cos\theta}{r'\cos\theta - r\sin\theta}.
\]

The formula is not worth memorising as a formula: it is worth recognising that the curve is parametrised and that the derivatives are the two in the quotient.

## Where it simplifies

For a circle the radius is constant, so \(r' = 0\) and the slope reduces to \(-\cot\theta\), which is the tangent to a circle at the point angle \(\theta\) measures. That special case is a good check on the general formula, since the geometry of a circle is known independently.

At a point where the denominator vanishes the ratio is undefined, which corresponds to a vertical tangent: the same failure the parametric form has, and the honest answer there.

## Checking against the geometry

Every polar derivative can be checked by drawing the point and asking which way the curve is heading. A slope of \(-1\) at 45 degrees says the tangent runs down and to the right, which is what a circle's tangent does at that point.

That check is worth doing because the two terms in the quotient are easy to transpose, and the drawing catches the error in a second.
