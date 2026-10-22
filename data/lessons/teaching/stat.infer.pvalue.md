# p-values and Significance

**p-value:** The probability, computed under the null hypothesis, of observing a test statistic at least as extreme as the one obtained. It measures compatibility with the null; it is not the probability that the null is true.

## The probability of data this extreme

A small p-value says that data like ours would be unusual if the null held. A two-sided test counts both tails, a one-sided test counts only the direction of the alternative. The p-value is a summary of the evidence, not a decision: the decision comes from comparing it with the significance level fixed in advance.

## Worked: comparing p with alpha

1. A test gives p = 0.03 and alpha is 0.05.
2. Since 0.03 < 0.05, the result is statistically significant and the null is rejected.
3. With p = 0.20 the same test would fail to reject, because the data are not unusual enough under the null.

## What a p-value does not say

It does not give the probability that the null is true, the size of the effect, or the probability that a replication will succeed. A tiny effect in a huge sample can produce a tiny p-value, and a large effect in a small sample can produce a large one. Report the effect size and interval alongside the p-value.
