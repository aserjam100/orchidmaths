# Order Statistics

**Order statistics:** For observations $X_1, \dots, X_n$, the sorted values $X_{(1)} \le \cdots \le X_{(n)}$ are the order statistics. The maximum is $X_{(n)}$ and the minimum is $X_{(1)}$.

## Sorting the sample

Order statistics record the positions of the observations rather than their labels. The maximum and minimum are the most-used members, and their distributions are simple for independent observations. The indicator that the maximum is at most a threshold becomes the intersection of the indicators that every observation is at most that threshold, which is where the product structure comes from.

## Worked: the maximum of uniforms

If the observations are independent and uniform on the unit interval, then the maximum is at most $x$ exactly when all of them are, so its distribution function is $x^n$. For three observations the probability that the maximum is at most one half is $(1/2)^3 = 0.125$. The maximum is driven toward one as the sample grows, because it must exceed every observation.

## The minimum and the density

The minimum exceeds a threshold exactly when every observation does, so its distribution function is $(1 - x)^n$ on the unit interval. Differentiating an order statistic's distribution function gives its density, which for the maximum is $n x^{n-1}$. These formulas are the basis of tests and estimators that rely on extremes rather than on averages.
