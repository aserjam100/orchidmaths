# Oblique asymptotes

## When the degree is one higher

A rational function whose numerator's degree exceeds its denominator's by exactly one has an oblique asymptote: the graph approaches a slanted line rather than a horizontal one. If the difference is larger there is no asymptote at all, only higher-degree growth.

The distinction matters because the end behaviour of the function is then linear rather than constant.

## Finding the slope

Dividing the numerator by the denominator gives a quotient plus a proper remainder. The remainder becomes negligible for large \(x\), so the quotient is the asymptote, and its leading coefficient is the slope.

For \((2x^2 + 3x + 1)/(x + 4)\), the quotient is \(2x - 5\) with a remainder, so the asymptote is \(y = 2x - 5\) and the slope is 2.

## The polynomial division behind it

The division is the same long division that finds a remainder, and the identity \(f(x) = q(x) + r(x)/d(x)\) is what makes the asymptotic statement precise. As \(x\) grows, the fraction \(r(x)/d(x)\) tends to zero because the remainder has lower degree than the divisor, so the function and its quotient differ by less and less.

That identity also explains the horizontal case: when the degree difference is zero, the quotient is a constant and the asymptote is horizontal.
