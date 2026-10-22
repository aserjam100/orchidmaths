# Transition Matrices

**Transition matrix:** the matrix $P$ whose row $i$, column $j$ entry is the probability of moving from state $i$ to state $j$, with every row summing to one.

## Row stochastic matrices

A transition matrix is row stochastic: its entries are nonnegative and each row adds to one. The row sums to one because from a given state the chain moves somewhere with probability one. Matrices of this kind preserve the property of being a probability vector, which is why multiplying a distribution by the matrix produces another distribution.

## Worked: one step forward

A chain starts with probability one half in each state, and from state zero the chance of moving to state one is nought point four, while from state one the chance of staying is nought point eight. The probability of being in state one after one step is one half times nought point four plus one half times nought point eight, which is 0.6. The new distribution is the old one multiplied by the matrix.

## Powers give multi-step probabilities

The row $i$, column $j$ entry of the matrix power counts all the routes of that many steps from $i$ to $j$, weighted by their probabilities. Adding over the intermediate states is exactly matrix multiplication. Thus the whole multi-step structure of the chain is encoded in the powers of a single matrix.
