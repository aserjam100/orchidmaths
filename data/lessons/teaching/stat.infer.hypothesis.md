# Hypothesis Tests

**Hypothesis test:** A procedure that starts with a null hypothesis $H_0$ (a default value or no effect) and an alternative $H_1$, then asks whether the data are compatible with $H_0$. The significance level $\alpha$ is fixed before the data arrive and caps the chance of rejecting a true null.

## Two claims, one decision

The null is the baseline that the test tries to discredit; the alternative is what replaces it if the evidence is strong enough. A one-sided alternative puts the whole rejection region in one tail, while a two-sided alternative splits it between both. The choice belongs to the question, not to the data.

## Worked: writing the hypotheses

1. A factory claims bags weigh 500 grams. A sample average far below 500 is the concern.
2. The null is $\mu = 500$; the alternative is $\mu < 500$, a one-sided test.
3. If either overfilling or underfilling mattered, the alternative would be $\mu \ne 500$ and the test two-sided.

## Significance is a rule fixed in advance

Alpha is the maximum false-alarm rate the analyst is willing to accept, commonly 0.05. Fixing it before looking at the data prevents moving the goalposts after the result. The test then rejects $H_0$ exactly when the p-value falls below alpha, and "fail to reject" never means that $H_0$ has been proved.
