# Rearrangement of Series

**Rearrangement:** a series formed by reordering the terms according to a bijection of the indices; absolutely convergent series keep their sum, while conditionally convergent series can be made to sum to anything.

## Rearranging the terms

A rearrangement of a series is a new series whose terms are the old terms in a different order, described by a bijection of the indices. Because addition of finitely many numbers does not depend on the order, every partial sum of the rearrangement is a sum of a finite set of the original terms, but the partial sums grow through different finite sets. For finitely many terms the order is irrelevant; the question is whether that irrelevance survives the passage to a limit, and it does not always.

## Absolutely convergent series are order-blind

If a series converges absolutely, every rearrangement converges to the same sum. The reason is that the total absolute mass can be bounded by a finite partial sum of the absolute series, so the difference between a partial sum of the rearrangement and a partial sum of the original can be pushed below any epsilon by taking enough terms. Absolute convergence lets the terms be added in any order because the tail of absolute values is itself small, so the reordering moves only a small amount of mass.

## Riemann's rearrangement theorem

If a series converges only conditionally, its terms can be reordered to produce any prescribed sum, or to diverge to plus infinity, minus infinity, or neither. The proof watches the positive and negative parts: each diverges, so the running sum can be pushed above the target by taking positive terms, then back below it by taking negative terms, closing in on any target. This is why conditional convergence is fragile and why the sum of a conditionally convergent series is inseparable from its order.
