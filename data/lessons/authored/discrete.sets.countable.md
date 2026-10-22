# Countable and uncountable sets

Two sets have the same **cardinality** when there is a bijection between them. That definition is what makes infinite cardinalities comparable at all: the even integers and the integers are matched by \(n \mapsto 2n\), so they have the same size despite one being a proper subset of the other.

## Cardinality as a bijection

A set is **countable** when it is finite or in bijection with the natural numbers, which means its elements can be listed in a sequence that eventually reaches every one of them. The list is the witness; a bijection is the same information written as a function.

This is the first place where "same size" stops meaning "same number of elements". The bijection is the definition, and the proper-subset intuition is the thing the definition overrides.

## Countable sets

The integers are countable: list \(0, 1, -1, 2, -2, \dots\). The rational numbers are countable too, by listing the fractions by the sum of numerator and denominator, which walks the grid diagonally and reaches every fraction.

The algebraic numbers — roots of polynomials with integer coefficients — are countable, because there are countably many such polynomials and each has finitely many roots. The finite binary strings are countable: list them by length, then lexicographically within each length.

## Cantor's diagonal argument

The real numbers are not countable. Suppose the reals in \((0,1)\) could be listed as decimals \(0.d_{n1}d_{n2}\dots\). Build a number whose \(n\)-th digit differs from \(d_{nn}\); it differs from the \(n\)-th listed number at position \(n\), so it is not on the list, and the list was supposed to be complete.

The same argument shows that the power set of any set is larger than the set, and that the set of functions from the naturals to two values is uncountable. The pattern is always the same: assume a complete list, then build the object that escapes it.
