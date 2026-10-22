# The ridge penalty

## The L2 term

Ridge regression adds a penalty to the loss:

\[
\lambda\sum_i w_i^2.
\]

The penalty grows with the square of each weight, so the training step is pulled toward smaller weights, and the strength \(\lambda\) decides how much. A weight of 3 and 4 with \(\lambda = 2\) contributes \(2(9 + 16) = 50\) to the loss.

The penalty is added to the loss being minimised, so it competes with fitting the data rather than replacing it.

## Why shrink the weights

Large weights make a model sensitive: a small change in an input moves the output a long way. Shrinking them trades a little training accuracy for a model that generalises, which is the same trade the whole of regularisation is about.

Ridge never sets a weight exactly to zero, because the penalty's slope near zero is small. That is the difference from lasso, whose absolute-value penalty has a corner at zero and therefore selects features by dropping them.

## Ridge against lasso

For correlated inputs, ridge tends to share the weight between them while lasso picks one and drops the other. Which is better depends on whether the question is how much each feature contributes or which features matter at all.

Both are one line added to the loss, and both are worth trying before anything more complicated.
