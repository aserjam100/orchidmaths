# Residues at simple poles

The **residue** of a function at a pole is the coefficient of \((z - z_0)^{-1}\) in its Laurent series there. At a **simple pole** — one where that term is the most singular — the coefficient has a one-line formula: \(\operatorname{Res}(f, z_0) = \lim_{z\to z_0} (z - z_0)f(z)\).

## The limit formula

If \(f(z) = \frac{g(z)}{z - z_0}\) with \(g\) analytic and nonzero at \(z_0\), then multiplying by \(z - z_0\) leaves \(g(z)\), and the limit is \(g(z_0)\). That is the residue.

The formula requires the pole to be simple: if \(f\) has \(1/(z-z_0)^2\), multiplying by \(z-z_0\) still leaves a pole and the limit does not exist.

## Why simple poles are easy

The residue is the coefficient that survives the residue theorem's integral, so computing it is the step that turns a contour integral into arithmetic. For a simple pole no series expansion is needed, only a substitution.

For a rational function \(\frac{P(z)}{Q(z)}\) with a simple zero of \(Q\) at \(z_0\), the residue is \(\frac{P(z_0)}{Q'(z_0)}\), which is the same limit in quotient form.

## A worked residue

Take \(f(z) = \frac{5z+i}{z - (1+i)}\), which has a simple pole at \(z_0 = 1+i\). The residue is the numerator at \(z_0\): \(5(1+i) + i = 5 + 6i\).

The check is the definition: \((z - (1+i))f(z) = 5z + i\) everywhere except at the pole, and its limit there is \(5 + 6i\).
