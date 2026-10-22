# Convergence of Sequences

**Convergence:** a sequence converges to L when, beyond some position, every term lies within any prescribed positive distance of L; L is then the limit of the sequence.

## The epsilon definition

A sequence converges to L when for every positive epsilon there is a positive integer N such that every term from position N onward is within epsilon of L. The order of the quantifiers matters: epsilon is chosen first and by an adversary, and N may depend on epsilon. The definition does not say how fast the terms approach L, only that they eventually stay close. A sequence that does not converge is divergent, and divergence includes both escape to infinity and endless oscillation.

## Worked: a ratio of polynomials

Consider the term formed by a n plus b over c n plus d, with c nonzero. Divide numerator and denominator by n, which does not change the fraction but reveals its tail: the constant terms contribute b over n and d over n, both of which shrink to zero, leaving the ratio a over c. So the limit is the ratio of the leading coefficients, whatever the constants. The constants b and d affect how soon the terms are close, but not where they settle.

## Uniqueness and boundedness

A sequence has at most one limit, because two different limits would each have to keep the tail close while the two neighbourhoods are disjoint. Every convergent sequence is bounded, since the tail stays near the limit and only finitely many earlier terms remain. The converse fails: alternating between minus one and one is bounded but not convergent, because the tail never settles. Boundedness is necessary for convergence but not sufficient, and the missing ingredient is that the terms must stop wandering.
