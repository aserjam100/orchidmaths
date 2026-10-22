# Gini Impurity

**Gini impurity:** For a node with class probabilities $p_i$, the Gini impurity is $1 - \sum p_i^2$. It is the probability that two examples drawn at random from the node have different classes, so 0 means pure and larger values mean more mixed.

## Measuring how mixed a node is

For two classes with proportions $p$ and $1-p$, the impurity is $2p(1-p)$, which is largest at $p = 0.5$. A split is chosen to reduce the weighted impurity of the children, and the reduction is the information the split buys.

## Worked: Gini for 3 and 1

1. A node holds 3 examples of one class and 1 of another.
2. The proportions are $3/4$ and $1/4$.
3. The impurity is $1 - (9/16 + 1/16) = 6/16 = 0.375$.

A pure node of 4 examples of one class has impurity $1 - 1 = 0$, the best possible value.

## Choosing the split that lowers impurity

For each candidate split, compute the impurity of the left and right children, weight each by its share of the examples, and subtract from the parent's impurity. The split with the largest decrease wins. Gini and entropy usually choose similar splits; Gini avoids logarithms and is the default in many tree implementations.
