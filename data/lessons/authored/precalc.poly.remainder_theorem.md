# The remainder and factor theorems

## The remainder theorem

Dividing a polynomial \(p(x)\) by \(x - c\) leaves a remainder, and the remainder theorem says that remainder is exactly \(p(c)\). The division is exact when \(p(c) = 0\), which is the factor theorem: \(x - c\) is a factor precisely when \(c\) is a zero.

For \(p(x) = x^3 + 2x^2 - x + 3\) divided by \(x - 2\), the remainder is \(8 + 8 - 2 + 3 = 17\), which is one substitution instead of a long division.

## Why it works

Dividing gives \(p(x) = (x-c)q(x) + r\), where \(r\) is a constant because the divisor is linear. Substituting \(x = c\) kills the first term and leaves \(p(c) = r\), which is the whole proof.

## Finding zeros

A zero of \(p\) is a value where the remainder is zero, so the theorem turns the search for factors into evaluation. Testing the candidates the rational-root theorem produces and keeping the ones that give zero is how a short list becomes the actual roots.

The theorem does not say how to guess a candidate, and it says nothing about how many roots exist, which is the fundamental theorem of algebra.
