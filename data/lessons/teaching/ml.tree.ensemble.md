# Bagging and Random Forests

**Bagging:** Train many copies of a model on bootstrap samples of the training data and average their predictions. **Random forests** add a random subset of features at each split, which decorrelates the trees further.

## Many trees, averaged

A single deep tree has low bias and high variance: small changes in the data move the splits and the predictions. Averaging many such trees keeps the low bias and cancels part of the variance, provided the trees' errors are not perfectly correlated.

## Worked: variance cancellation

1. Take $B$ trees, each with prediction variance $v$ and pairwise correlation $\rho$.
2. The variance of their average is $\rho v + (1-\rho)v/B$.
3. As $B$ grows the second term vanishes, but the first remains, so decorrelating the trees matters more than adding more of them.

## Random forests and boosting

A random forest is bagging plus feature subsampling, and it is trained in parallel. Boosting is sequential: each new weak learner fits the residual errors of the current ensemble, which reduces bias rather than only variance. Bagging is the safer default on noisy data; boosting often wins on clean tabular data but needs careful learning rates and validation.
