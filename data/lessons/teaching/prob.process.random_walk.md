# Random Walks

**Random walk:** the partial sum $S_n$ of a sequence of independent steps, each plus one with some probability $p$ and minus one otherwise.

## Steps and partial sums

A random walk accumulates independent increments, so its mean and variance grow linearly in the number of steps. Each step contributes the same mean and the same variance, and independence makes the variances add. This linear accumulation is what distinguishes a walk from a stationary process.

## Worked: variance after several steps

A step is plus one with probability one half and minus one otherwise, so it has mean zero and variance one. After ten steps the sum has variance ten, because independent variances add. When the step probabilities are unequal the per-step variance is four times the product of the two probabilities, and the walk acquires a drift as well.

## Recurrence and transience

The simple symmetric walk in one dimension returns to its starting point almost surely and does so infinitely often. The same walk on the two-dimensional lattice is also recurrent, but in three dimensions it becomes transient and may never return. This dimensional change in behaviour is one of the most surprising facts in the subject and is not visible from the one-dimensional intuition alone.
