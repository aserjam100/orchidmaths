# Continuity of Probability

**Continuity from below:** If the events $A_1 \subseteq A_2 \subseteq \cdots$ increase to a union $A$, then $P(A_n)$ tends to $P(A)$.

## Increasing sequences

The union of an increasing sequence can be written as a disjoint union of the increments: the first event, then each new event with the previous one removed. Countable additivity applies to those disjoint increments, and the resulting series has partial sums exactly $P(A_n)$. Letting $n$ grow gives the limit, which is why continuity is not an extra axiom but a theorem about the original ones.

## Worked: a geometric union

Let the events be nested with $P(A_n) = 1 - 2^{-n}$. The sequence increases to probability one, so the union has probability one even though no single $A_n$ ever attains it. Continuity makes the limiting statement meaningful: the probability of the whole union is the limit of the finite probabilities, not a value assumed at some finite stage.

## Continuity from above

For a decreasing sequence $A_1 \supseteq A_2 \supseteq \cdots$ with intersection $A$, the same argument on complements gives $P(A_n) \to P(A)$, provided the first event has finite probability. The finiteness condition is not decorative: with an infinite starting measure, subtracting limits can produce an indeterminate form. Both directions together are called the continuity of probability.
