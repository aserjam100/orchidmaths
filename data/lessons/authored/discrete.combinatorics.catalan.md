# Catalan numbers

## Balanced parentheses

A string of parentheses is **balanced** when every prefix has at least as many opening brackets as closing ones, and the whole string has equally many. With 3 pairs there are \(\binom{6}{3} = 20\) strings in total, and only 5 of them are balanced.

That count is the third Catalan number, and the sequence begins \(1, 1, 2, 5, 14, 42, 132, 429\) for \(n = 0, 1, 2, \ldots\).

## The recurrence

The Catalan number \(C_n\) counts the balanced strings of \(n\) pairs and satisfies

\[C_{n+1} = \sum_{i=0}^{n} C_i\, C_{n-i},\]

with \(C_0 = 1\). The sum splits a balanced string at the bracket that closes its first opening bracket: what lies inside is a balanced string, and what follows it is another, which is why the product of two Catalan numbers appears under the sum.

A closed form is

\[C_n = \frac{1}{n+1}\binom{2n}{n}.\]

The division by \(n+1\) is exact for every \(n\), so the closed form computes an integer even though it is written as a fraction.

## Where they appear

Many problems that look unrelated share this count, because each one can be cut at a first return the way the parentheses are cut. Triangulations of a convex polygon with \(n+2\) sides, full binary trees with \(n+1\) leaves, monotone paths that do not cross the diagonal, and the ways to bracket a product of \(n+1\) factors are all \(C_n\).

Recognizing one of these shapes is the practical skill: the count itself is one formula, and choosing the formula is the work.
