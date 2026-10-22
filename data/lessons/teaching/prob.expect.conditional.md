# Conditional Expectation

**Conditional expectation:** the mean of $X$ computed inside the slice where $Y$ takes a given value, $E[X \mid Y = y] = \sum_x x \, p(x \mid y)$.

## Conditioning on a value

Conditioning fixes the value of $Y$ and turns the joint law into the conditional law $p(x \mid y) = p(x, y)/p_Y(y)$. The conditional mean is then the weighted average of $X$ within that slice, using the conditional masses as weights. It reports the centre of the slice rather than of the whole distribution.

## Worked: a column of the table

A joint mass function on the four pairs from zero to one is proportional to the weights one, one, three and one. Conditioning on $Y$ equal to zero leaves the pairs with weights one and three, so $X$ is one with conditional probability $3/(1 + 3) = 0.75$ and the conditional mean is 0.75. Conditioning on the other value would use the other two weights.

## A function of the conditioning variable

The conditional expectation $E[X \mid Y]$ is itself a random variable, since it depends on the value that $Y$ turns out to take. Reading it as a function of $Y$ is what lets its own mean be taken, and that is the step behind the law of total expectation. Conditioning is therefore an operation on the whole distribution, not on a single number.
