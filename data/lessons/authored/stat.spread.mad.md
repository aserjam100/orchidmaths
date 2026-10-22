# The mean absolute deviation

## Distance from the mean

The mean absolute deviation is the average distance between a data value and the mean, with the signs discarded. Each value's distance from the mean is taken, the distances are added, and the total is divided by the count.

A small value means the data sits close to its mean and a large one means it is spread out, which is the same thing the standard deviation says in different units.

## Why absolute values

Squaring the distances, as the variance does, gives extra weight to the values far from the mean, so an outlier moves the standard deviation much more than it moves the mean absolute deviation. Taking the absolute value instead treats every distance in proportion to its size.

The trade is that absolute values are harder to work with algebraically, which is why the standard deviation is the usual choice and this one is preferred when robustness matters.

## Comparing with the standard deviation

The two measures always rank data sets the same way on average but not always on the same pair, and neither is wrong. The mean absolute deviation is in the data's own units, so it is easier to read; the standard deviation has the algebra that statistics needs.

Both are useful, and a data set's spread is worth describing with one of them rather than only with a range, since a range depends on two values and ignores the rest.
