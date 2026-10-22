# Primitive Pythagorean triples

## Euclid's formula

Whole-number solutions of the Pythagorean equation can be produced from two integers: the triple is the difference of their squares, twice their product, and the sum of their squares. Every primitive triple arises this way from exactly one pair.

So generating triples is choosing a pair and evaluating three expressions, with no search over the numbers involved.

## Why coprime and opposite parity

The conditions on the pair are what make the triple primitive rather than a multiple of a smaller one. If the two share a factor, all three entries share it; if they have the same parity, the entries are all even.

Both conditions are therefore about dividing out a common factor, and dropping either produces a triple that is a scaled copy of a smaller one.

## Checking a triple

A triple is verified by squaring and adding, which is one line. A primitive triple can also be recognised by the fact that its hypotenuse is odd and exactly one leg is even, and that the three entries have no common factor.

Those checks are worth running because a formula with two conditions is easy to apply with one of them missed.
