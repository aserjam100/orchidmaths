# Random Variables

**Random variable:** A random variable is a measurable function $X$ from the sample space $\Omega$ to the real numbers. Measurability means that every set of the form $\{X \le x\}$ is an event.

## A function, not a number

The word random hides the definition: a random variable is an ordinary function, and its randomness comes entirely from the probability on its domain. It assigns a real number to each outcome. Measurability ensures that the sets a probability is asked about are members of the sigma-algebra, so quantities such as $P(X \le x)$ are well defined.

## Worked: the indicator variable

For an event $A$, the indicator $\mathbf{1}_A$ takes the value one on $A$ and zero outside. It is a random variable, and its expectation is $P(A)$. Indicators turn set statements into arithmetic and are the smallest pieces from which more complicated random variables are built by summing and scaling.

## Discrete and continuous

A random variable is discrete when it takes countably many values, and continuous when its distribution is given by a density. The two cases need different bookkeeping: masses for the discrete case and integrals for the continuous one. Some variables are mixed, with both a continuous part and atoms of positive probability, which is why the distribution function is the general description that covers every case.
