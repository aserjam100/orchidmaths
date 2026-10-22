# Marginal Distributions

**Marginal distribution:** the distribution of one variable of a pair, obtained by summing the joint mass function over all values of the other variable, or integrating the joint density over it.

## Summing out a variable

For a discrete pair the marginal of $X$ is $p_X(x) = \sum_y p(x, y)$, and the marginal of $Y$ is the symmetric sum. In the continuous case the sum becomes an integral over the other coordinate. The operation simply forgets the second variable, collapsing a table onto one of its margins, and what remains is a genuine one-variable distribution.

## Worked: a joint table

A joint mass function on the four pairs from zero to one is proportional to the weights one, two, three and four. The weights sum to ten, and the pairs with $X = 1$ carry the last two weights, so the marginal probability that $X$ is one is $(3 + 4)/10 = 0.7$. Adding a column of the table is the whole computation.

## Marginals do not determine the joint

Two different joint distributions can share the same pair of marginals, so the margins alone do not describe the dependence. For example, a joint law concentrated on the diagonal and an independent law with the same margins both have those margins, yet they assign very different mass to the off-diagonal cells. The joint law carries strictly more information than the two margins together.
