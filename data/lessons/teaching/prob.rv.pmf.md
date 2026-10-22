# The Probability Mass Function

**Mass function:** For a discrete random variable, the mass function is $p(x) = P(X = x)$. The masses are nonnegative and sum to one over the values of the variable.

## Masses sum to one

The values of a discrete random variable form a countable set, and the events that the variable equals each value are disjoint. Countable additivity makes their probabilities sum to the probability of the whole space, which is one. A list of numbers is a valid mass function exactly when it is nonnegative and its total is one, and normalising any list of positive weights produces such a function.

## Worked: normalising weights

A random variable takes four outcomes with weights two, four, six and eight. Their total is twenty, so the mass of the last outcome is $8/20 = 0.4$. Each weight divided by the total gives a probability, and the four probabilities now sum to one. Normalising by the total is the general recipe for turning counts or weights into a mass function.

## From masses to the CDF

The distribution function of a discrete variable is the running total of the masses: $F(x)$ is the sum of $p(v)$ over all values $v$ at most $x$. It is flat between values and jumps by exactly one mass at each value, so the mass function is recoverable from the jumps. This is why a discrete distribution can be described by either object, and each is built from the other.
