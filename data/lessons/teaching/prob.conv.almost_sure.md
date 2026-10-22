# Almost-Sure Convergence

**Almost-sure convergence:** a sequence converges to $X$ when the set of outcomes on which the numerical sequence $X_n$ converges to $X$ has probability one.

## Convergence on almost every path

Almost-sure convergence is first a statement about each outcome separately: fix an outcome and the numbers $X_n$ must settle down to $X$. The exceptional outcomes where this fails form a set of probability zero. This pointwise reading is what makes it stronger than convergence in probability, which only controls the size of large deviations.

## It implies convergence in probability

If the numbers converge almost everywhere then the tail probabilities of a fixed-size departure must vanish, so almost-sure convergence implies convergence in probability. For the converse, a sequence converging in probability always has a subsequence converging almost surely, which is often enough for applications. The two modes are linked by this subsequence theorem.

## Borel-Cantelli as the tool

The usual way to prove almost-sure convergence is the first Borel-Cantelli lemma, which turns a summable sequence of tail probabilities into the statement that only finitely many bad events occur. That gives the convergence on a probability-one set directly. Almost-sure results are therefore usually built from summability estimates.
