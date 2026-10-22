# The Poisson Process

**Poisson process:** arrivals occur so that the number in an interval of length $t$ is Poisson with mean $\lambda t$, and the counts in disjoint intervals are independent.

## Arrivals as a Poisson count

The count of arrivals in an interval of length $t$ has a Poisson distribution whose mean is the rate times the length. Its probability of exactly $k$ arrivals is $e^{-m} m^k / k!$ with $m = \lambda t$. The single parameter $m$ carries both the rate and the length of the window, which is why only the product matters.

## Worked: a rate over an interval

If the mean number of arrivals in the interval is two, the chance of exactly one arrival is $e^{-2} \cdot 2 = 0.270671$ to six places. The chance of no arrival at all is $e^{-2} = 0.135335$. Computing the whole distribution from the single mean is the routine part of a Poisson model.

## Waiting times are exponential

The time from one arrival to the next is exponential with rate $\lambda$, and the interarrival times are independent. The exponential has no memory, so the expected wait does not depend on how long the current gap has already lasted. This is the continuous counterpart of the geometric waiting time in a sequence of trials, and it links the process to the memoryless property.
