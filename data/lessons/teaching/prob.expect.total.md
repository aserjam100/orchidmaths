# The Law of Total Expectation

**Law of total expectation:** the overall mean is the average of the conditional means, $E[X] = E[E[X \mid Y]]$, where the outer average uses the distribution of $Y$.

## Averaging the conditional means

The law says that computing a mean in slices and then averaging the slice means with the probabilities of the slices recovers the original mean. No independence is needed, only that the slices partition the outcomes. This is often the easiest route when the conditional problem is simple but the joint one is not.

## Worked: two groups

A variable $Y$ is one with probability one half, and the conditional mean of $X$ is one when $Y$ is zero and three when $Y$ is one. The overall mean is $(1/2)(1) + (1/2)(3) = 2$, the plain midpoint. Unequal group probabilities would pull the answer toward the larger group.

## The tower property

Writing the law as $E[E[X \mid Y]] = E[X]$ emphasises that conditioning and then averaging is the identity. Applying it twice, first on $Y$ and then on a finer variable, telescopes to the same mean. The property is what makes conditional expectation a projection of one variable onto the information carried by another.
