# Histograms

**Histogram:** A graph of counts per interval. The range of the data is cut into bins, each value is counted in the bin that contains it, and each bar's height is the count. Unlike a bar graph, the bins are ordered numeric intervals of equal width.

## Bins and counts

The bins are usually half-open, such as $[10, 20)$, so each value belongs to exactly one bin. The bar height is the count, or the count divided by the total when relative frequencies are wanted. The shape that emerges, symmetric, skewed or bimodal, is the point of the display.

## Worked: reading a bin

1. Values 5, 12, 14, 19, 22, 25 with bins of width 10.
2. The bin $[10, 20)$ contains 12, 14 and 19.
3. So its bar has height 3, while $[0, 10)$ has height 1 and $[20, 30)$ has height 2.

## Choosing bin width

Too few bins hide the shape; too many turn noise into apparent structure. A common starting point is the square root of the sample size, adjusted so the bins land on round numbers. Comparing histograms across groups requires the same bins and, for unequal group sizes, relative frequencies rather than raw counts.
