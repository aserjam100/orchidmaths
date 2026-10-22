# Power sets and counting subsets

The **power set** of \(A\), written \(\mathcal{P}(A)\), is the set of all subsets of \(A\). A set with \(n\) elements has \(2^n\) subsets, because each element is independently either in or out, and the \(n\) independent choices give \(2^n\) outcomes.

## The power set

For \(A = \{1,2,3\}\) the power set has eight members: the empty set, the three singletons, the three pairs, and \(A\) itself. The empty set and \(A\) are always members, which is why the count starts at two rather than at one.

The power set is a set whose elements are sets, and that is the first place where the distinction between a member and a subset matters: \(\{1\}\) is an element of \(\mathcal{P}(A)\), while 1 is an element of \(A\).

## Counting subsets

A subset of size \(k\) is a choice of \(k\) elements from \(n\), so the number of \(k\)-element subsets is \(\binom{n}{k}\). Summing over all \(k\) recovers the total: \(\sum_k \binom{n}{k} = 2^n\), which is the binomial theorem at \(x = y = 1\).

For a set with six elements, the number of two-element subsets is \(\binom{6}{2} = 15\). The total of 64 splits into \(1 + 6 + 15 + 20 + 15 + 6 + 1\) by subset size.

## Choosing which side to count

Some questions are easier by counting the complement. The number of subsets of an \(n\)-element set that **avoid** a fixed element \(a\) is \(2^{n-1}\), since the remaining elements are free; the number that **contain** \(a\) is the same, since the choice of \(a\) is forced and the rest are free.

So "subsets containing \(a\)" and "subsets avoiding \(a\)" are both \(2^{n-1}\), and the two counts partition the \(2^n\) subsets into equal halves. Recognising which side of a constraint is the smaller count is the same move as counting by complement, one level up.
