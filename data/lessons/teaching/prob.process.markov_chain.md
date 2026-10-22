# Markov Chains

**Markov chain:** a sequence of states in which the next state depends on the present state only and not on the earlier history.

## The Markov property

Conditionally on the present state, the future is independent of the past. This single assumption is what makes long-run behaviour tractable: the whole transition structure is captured by the one-step probabilities between states. Chains are the simplest models of dependence over time and the starting point for the theory of stochastic processes.

## Worked: a two-step transition

A two-state chain has one-step probabilities from state zero of nought point six to itself and nought point four to state one, and from state one of nought point seven to itself. The chance of being in state one two steps after starting in state zero is the sum over the intermediate state, nought point six times nought point four plus nought point four times nought point seven, which is 0.52. Every two-step probability is a sum of such products.

## The transition matrix

Collecting the one-step probabilities into a matrix, with rows indexed by the current state and columns by the next, gives a matrix whose rows each sum to one. Row stochasticity is exactly the statement that from any state the chain must go somewhere. Powers of this matrix give the multi-step transition probabilities.
