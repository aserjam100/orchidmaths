# The sum and difference formulas

## Two angles at once

Knowing the sine of \(\alpha\) and the cosine of \(\beta\) does not tell you the sine of
\(\alpha + \beta\) by adding the two numbers. The sine of a sum is not the sum of the sines:
\(\sin(30^{\circ} + 60^{\circ}) = 1\), while \(\sin 30^{\circ} + \sin 60^{\circ}\) is about 1.37.

What the formulas give is the true relationship, and it is a product of the two angles' ratios
rather than a sum of them.

## The four formulas

Adding an angle to itself is the case the double-angle formulas handle, and the general case is
written the same way. For sine, \(\sin(\alpha + \beta) = \sin\alpha\cos\beta + \cos\alpha\sin\beta\),
and for cosine, \(\cos(\alpha + \beta) = \cos\alpha\cos\beta - \sin\alpha\sin\beta\).

The difference versions have the same shape with one sign turned. Sine's difference subtracts:
\(\sin(\alpha - \beta) = \sin\alpha\cos\beta - \cos\alpha\sin\beta\). Cosine's difference adds:
\(\cos(\alpha - \beta) = \cos\alpha\cos\beta + \sin\alpha\sin\beta\).

That flip is the part worth remembering. The sine keeps the sign of the sign in the angle, and the
cosine reverses it, which is why \(\cos(\alpha - \beta)\) is the one that expands with two plus
signs.

## Combining two triangles

Each term in the expansion is a ratio from one angle times a ratio from the other, so two right
triangles are enough to evaluate it. If \(\sin\alpha = 3/5\) and \(\cos\beta = 12/13\), the first
triangle has opposite 3 and hypotenuse 5, giving \(\cos\alpha = 4/5\), and the second has adjacent
12 and hypotenuse 13, giving \(\sin\beta = 5/13\).

Then \(\sin(\alpha + \beta) = 3/5 \cdot 12/13 + 4/5 \cdot 5/13 = 36/65 + 20/65 = 56/65\). With both
angles acute every ratio is positive, so no sign needs repairing and the answer is the fraction as
it comes.
