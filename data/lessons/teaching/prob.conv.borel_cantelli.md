# The Borel-Cantelli Lemmas

**Borel-Cantelli lemmas:** if the probabilities of a sequence of events sum to a finite total then almost surely only finitely many occur, and if the events are independent with a divergent sum then almost surely infinitely many occur.

## The first lemma

If the total probability of the events is finite, the expected number of events that occur is finite, so infinitely many cannot occur with positive probability. No independence is assumed, which makes the lemma extremely flexible. It is the standard route from a summable estimate to an almost-sure statement.

## The second lemma

The converse needs independence. If the events are independent and their probabilities sum to infinity, then almost surely infinitely many of them occur. Independence is essential: a nested sequence of events with divergent probabilities can still have only the first few occur if they are not independent.

## Proving almost-sure results

Together the lemmas convert statements about sums of probabilities into statements about the limsup of a sequence of events. Almost-sure convergence of an estimator is usually shown by bounding the probability of a large error at each index and summing the bounds. The strong law of large numbers is the classical application.
