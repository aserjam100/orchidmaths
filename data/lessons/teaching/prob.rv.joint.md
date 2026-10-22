# Joint Probability Mass Functions

**Joint mass function:** for two discrete random variables, $p(x, y) = P(X = x, Y = y)$; the masses are nonnegative and sum to one over all pairs.

## Masses on pairs

The support of a joint distribution is the set of pairs with positive mass, and the masses over that support sum to one. A joint mass function is the discrete analogue of a two-dimensional density, and it is the object from which every question about the pair is answered. A table with rows indexed by $x$ and columns by $y$ is the usual way to display it.

## Worked: normalising a table

A joint mass function on the four pairs from zero to one is proportional to the weights one, two, three and four. The weights sum to ten, so the pair with weight three carries mass $3/10 = 0.3$. Dividing every weight by the total is the same normalisation a one-variable mass function needs, applied cell by cell.

## Marginals and independence

Summing the joint masses over all values of $y$ gives the marginal mass function of $X$, and the symmetric sum gives that of $Y$. The two variables are independent exactly when the joint factors into the product of the marginals for every pair. A joint table therefore shows dependence as a departure from the product of its own row and column totals.
