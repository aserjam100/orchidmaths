# Descartes' rule of signs

## Counting sign changes

Read the coefficients of a polynomial in order of descending degree and count how many times the sign flips. For \(x^3 - 2x^2 - x + 3\) the signs are \(+,-,-,+\), which change twice.

The count is a property of how the polynomial is written, so it must be in standard form with the zero coefficients kept out rather than written as zeros.

## What the count bounds

The number of positive real roots is at most the number of sign changes, and it has the same parity as that count: from the maximum you may subtract 2, 4, 6 and so on. Two sign changes therefore allow two positive roots or none.

Applying the rule to \(p(-x)\) bounds the negative roots the same way, which together with the candidates gives a short list to test.

## Using it with the candidates

The rule is a bound, not a list, so it is used with the rational-root theorem: the candidates say which numbers to try, and the sign changes say how many of them can succeed with which sign.

A polynomial of degree 3 with two sign changes has either two positive roots and one negative, or none positive and three negative, and the candidates decide which.
