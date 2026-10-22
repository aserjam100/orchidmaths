# Partial orders and comparability

A **partial order** is a relation that is reflexive, antisymmetric and transitive. Divisibility on the positive integers is the standard example: \(a \mid a\), \(a \mid b\) and \(b \mid a\) force \(a = b\), and \(a \mid b\) with \(b \mid c\) gives \(a \mid c\).

## Partial orders

"Partial" is the point of the name: two elements need not be comparable. On the positive integers, 2 and 3 divide neither each other, so the order leaves them side by side. A **total order** is the special case where every pair is comparable, and the usual \(\le\) on the integers is one.

The Hasse diagram draws a partial order without the edges that transitivity forces: elements are placed by rank, and a line runs between \(a\) and \(b\) when \(a \prec b\) and nothing sits between them. Reading it upward recovers the relation.

## Comparability

Two elements are **comparable** when one is below the other. Counting comparable pairs is a way to measure how far an order is from being total: a total order on \(n\) elements has \(\binom{n}{2}\) comparable pairs, and a partial order has fewer.

For divisibility on \(\{1,2,3,4,6,12\}\), the comparable pairs with a strict inequality are \(1 \mid 2, 1 \mid 3, 1 \mid 4, 1 \mid 6, 1 \mid 12\), then \(2 \mid 4, 2 \mid 6, 2 \mid 12\), then \(3 \mid 6, 3 \mid 12\), then \(4 \mid 12\) and \(6 \mid 12\): twelve in all, against the fifteen pairs a total order would give.

## Counting comparable pairs

Enumerating the pairs is the reliable method for a small order: list the elements in increasing order, and for each one count its multiples that are also in the set. The total counts each comparable pair once, because the smaller element is the one enumerated.

The comparison with \(\binom{n}{2}\) is a check rather than a shortcut: equality means the order is total, and a deficit measures the incomparable pairs. For the divisor set above, the three missing pairs are exactly the incomparable ones.
