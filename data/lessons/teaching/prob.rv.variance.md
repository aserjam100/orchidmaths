# Variance by the Second Moment

**Variance:** $\operatorname{Var}(X) = E[(X - \mu)^2] = E[X^2] - \mu^2$, where $\mu = E[X]$.

## Spread about the mean

Variance averages the squared distance from the mean, so it measures spread rather than location. Expanding the square and using linearity gives the computational form $E[X^2] - \mu^2$, which needs only the first two moments and avoids subtracting from every value. Squaring makes large deviations count disproportionately and keeps the two sides from cancelling.

## Worked: a three-point pmf

A variable takes zero, one and two with probabilities one quarter, one half and one quarter. The mean is one, and the second moment is $0 + 0.5 + 1 = 1.5$, so the variance is $1.5 - 1 = 0.5$. The same numbers from the definition give $0.25(1)^2 + 0.5(0)^2 + 0.25(1)^2 = 0.5$, which checks the two forms agree.

## Scaling and units

Variance is in squared units of the variable, so its square root, the standard deviation, is the quantity to compare with the values themselves. Shifting a variable does not change its variance, while scaling it by a constant multiplies the variance by the square of that constant. Those two rules are the whole effect of a linear transformation on spread.
