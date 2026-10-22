# Bell numbers

## Partitions of a set

The Bell number \(B(n)\) counts all the ways to split \(n\) distinct items into non-empty groups, with no fixed number of groups and the groups unlabelled. For three items the partitions are one group of three and the three ways to make a pair plus a single, giving \(B(3) = 5\).

The sequence begins \(1, 2, 5, 15, 52, 203, 877\) for \(n = 1, 2, \ldots\).

## Summing the Stirling numbers

\(B(n)\) is the sum of \(S(n,k)\) over every \(k\) from 1 to \(n\), because a partition has some number of blocks and the Stirling number counts the partitions with exactly that many. There is no closed form and no recurrence in a single previous value, so the numbers are usually tabulated.

## Why it grows so fast

The Bell numbers grow faster than exponentially but slower than factorially: by \(n = 8\) there are already 4140 partitions of a set of eight items. That growth is what makes them the natural count for questions about grouping, clustering and equivalence relations, since a partition of a set and an equivalence relation on it are the same object.
