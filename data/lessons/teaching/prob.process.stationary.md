# Stationary Distributions

**Stationary distribution:** a probability vector $\pi$ that the transition matrix leaves unchanged, $\pi P = \pi$, so the law of the state does not vary in time.

## The balance equations

Stationarity says that the probability flowing into each state equals the probability flowing out, so the distribution is unchanged after one step. Writing the equations out and solving gives the stationary law. For a two-state chain the flow balance gives the stationary probability of state one as the upward rate divided by the sum of the two rates.

## Worked: a two-state chain

A chain moves from state zero to state one with probability nought point three and from state one to state zero with probability nought point seven. The stationary probability of state one is nought point three divided by the sum nought point three plus nought point seven, which is 0.3. At stationarity the chance of being in state one barely exceeds the chance of being in state zero because the downward rate is higher.

## Existence and uniqueness

A finite chain that is irreducible, meaning every state can be reached from every other, and aperiodic has a unique stationary distribution, and the distribution after many steps converges to it regardless of the starting point. If the chain is not irreducible there may be several stationary laws. These are the standard sufficient conditions for the long-run behaviour to be well defined.
