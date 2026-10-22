# Probability Axioms

**Kolmogorov's axioms:** A probability on a sample space $\Omega$ is a function $P$ on the events of $\Omega$ with three properties: $P(A) \ge 0$ for every event, $P(\Omega) = 1$, and $P(\bigcup A_i) = \sum P(A_i)$ for any countable family of pairwise disjoint events.

## The three axioms

Nonnegativity says no event has a negative probability. Normalisation fixes the whole sample space at probability one. Countable additivity makes disjoint unions add, and it is the axiom that separates a genuine measure from a merely finitely additive set function. Every other rule of probability is a consequence of these three, which is why the axioms can be stated so briefly.

## Worked: the complement rule

The events $A$ and its complement are disjoint and their union is $\Omega$. Countable additivity applied to this two-event family gives $P(A) + P(\text{not } A) = P(\Omega)$, and normalisation turns the right side into one. Rearranging gives $P(\text{not } A) = 1 - P(A)$. The same argument with the empty event shows that $P(\emptyset)$ is zero, because $\Omega$ and the empty set partition $\Omega$.

## Consequences and independence

Monotonicity, the inclusion-exclusion identity and the union bound all follow from the axioms. Independence does not: it is an extra modelling assumption, the equation $P(A \cap B) = P(A)P(B)$ imposed on top of the measure. Keeping the two ideas separate matters because a probability can satisfy the axioms and still fail every independence claim a careless model makes.
