# Derangements

## Definition

A **derangement** of \(n\) objects is a rearrangement in which no object stays where it started. Written \(!n\), it counts arrangements with no fixed point at all.

For three objects the six permutations contain two derangements, so \(!3 = 2\). For four, \(!4 = 9\), and the sequence continues \(2, 9, 44, 265, 1854\).

## The recurrence

Derangements satisfy

\[!n = (n-1)\left(!(n-1) + !(n-2)\right)\]

with \(!0 = 1\) and \(!1 = 0\). The reasoning is to follow one object, say the first: it moves to one of the other \(n-1\) positions, and then either the object it displaced takes the first position, leaving \(!(n-2)\) derangements of the rest, or it does not, leaving \(!(n-1)\).

The recurrence computes \(!n\) with integers alone, which is why it is the safe way to evaluate the count for a large \(n\).

## Inclusion and exclusion

The same number comes from counting what to exclude. From \(n!\) permutations, subtract those fixing at least one object, add back those fixing at least two, and continue: the alternating sum over subsets of fixed objects gives

\[!n = n!\sum_{k=0}^{n} \frac{(-1)^k}{k!}.\]

The two descriptions agree, and the sum shows something the recurrence hides: for \(n \ge 1\) the value of \(!n\) is the nearest whole number to \(n!/e\). That is why the ratio \(!n/n!\) tends to \(1/e \approx 0.3679\) as \(n\) grows.
