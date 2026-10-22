# One-Way ANOVA

**One-way ANOVA:** Compare the means of $k$ groups by splitting the total variation into a between-group part and a within-group part. The test statistic is $F = \mathrm{MSB}/\mathrm{MSW}$, where each mean square is a sum of squares divided by its degrees of freedom.

## Comparing several means at once

Running many pairwise t tests inflates the chance of a false positive. ANOVA asks one question instead: is the variation between group means larger than the variation inside the groups? If the groups are all draws from the same population, both mean squares estimate the same variance and F is near 1.

## Worked: three groups

1. Groups 1, 3; 3, 5; 5, 7 have means 2, 4, 6 and grand mean 4.
2. Between: $2(2-4)^2 + 2(4-4)^2 + 2(6-4)^2 = 16$ on 2 degrees of freedom, so MSB = 8.
3. Within: each group contributes 2, so SSW = 6 on 3 degrees of freedom, MSW = 2, and $F = 8/2 = 4$.

## The F ratio

F near 1 means the between-group spread is what chance alone would produce; large F means the group means are farther apart than the within-group noise explains. The F distribution depends on both degrees of freedom, $k - 1$ and $N - k$. A significant F says at least one mean differs, not which one, so follow-up comparisons are needed.
