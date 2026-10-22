# Iterated Limits

**Iterated limits:** the limits taken by holding one variable fixed and then letting the other move; they can disagree with each other and with the double limit.

## Two ways to take a limit

A function of two variables has three limits of interest at a point: the double limit, where both variables approach together, and the two iterated limits, where one variable is fixed and the other moves first and then the first moves. Each is a genuine limit, and they need not agree. The double limit is the strongest: it controls the function along every path to the point, so if it exists it forces both iterated limits to exist and to equal it.

## The classic counterexample

The function that is the product of the two variables divided by the sum of their squares shows the gap. Along either axis the quotient is zero, so both iterated limits are zero and they agree with each other. Along the diagonal the quotient is one half, so different paths to the origin give different limits and the double limit does not exist. Agreement of the iterated limits is therefore not enough to conclude that the double limit exists, which is the trap this example exposes.

## The order of integration

The same issue explains why changing the order of integration in an improper double integral can change the answer. Fubini's theorem says that for a nonnegative integrand, or for an absolutely integrable one, the order does not matter, so when the two iterated integrals disagree the integrand must not be absolutely integrable. The iterated limits and the iterated integrals raise the same question, and the answer in both cases is that the strong condition buys the right to rearrange.
