# Stirling numbers of the second kind

## Partitions into blocks

The Stirling number of the second kind \(S(m,n)\) counts the ways to split a set of \(m\) distinct items into exactly \(n\) non-empty groups, where the groups are not labelled. Splitting three items into two groups gives \(S(3,2) = 3\).

The groups being unlabelled is the only difference from counting onto functions: an onto function to \(n\) images labels them by which image is hit.

## The recurrence

Every partition of \(m\) items into \(n\) blocks either puts the last item in a block by itself or adds it to one of the existing blocks:

\[S(m,n) = S(m-1,n-1) + n\,S(m-1,n).\]

The seeds are \(S(m,1) = 1\) and \(S(m,m) = 1\), and the recurrence computes every value with integers alone.

## Relation to onto functions

Labelling the \(n\) blocks can be done in \(n!\) ways, so the number of onto functions is \(n!\,S(m,n)\), which is another way to write the inclusion-exclusion sum of the previous concept. The two formulas agree because they count the same objects with the labels kept and dropped.
