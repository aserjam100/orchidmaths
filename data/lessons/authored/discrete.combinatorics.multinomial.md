# Arrangements of a multiset

## When letters repeat

Arranging the letters of \(ABC\) gives \(3! = 6\) orders, but arranging the letters of \(AAB\) gives 3: the two \(A\)s are interchangeable, so each order appears twice in the list of \(3! = 6\) permutations and the doubling must be divided out.

The general rule divides by the factorial of every group of identical items. A multiset with \(k\) distinct letters, appearing \(n_1, n_2, \ldots, n_k\) times, has

\[\frac{n!}{n_1!\, n_2! \cdots n_k!}\]

distinct arrangements, where \(n = n_1 + n_2 + \cdots + n_k\).

## The multinomial formula

That expression is the **multinomial coefficient**, written \(\binom{n}{n_1, n_2, \ldots, n_k}\). It equals 1 only when some group has every item, and it equals \(n!\) when every group has one item.

The word \(AABBC\) has \(n = 5\) letters, with two \(A\)s, two \(B\)s and one \(C\), so it has

\[\frac{5!}{2!\,2!\,1!} = \frac{120}{4} = 30\]

distinct arrangements.

## Worked example

To count the rearrangements of the word \(LEVEL\), first tally the letters: two \(L\)s, two \(E\)s and one \(V\), so \(n = 5\) and

\[\frac{5!}{2!\,2!\,1!} = 30.\]

A useful check is to list the multiset in two different orders of the groups: the denominator does not depend on the order the groups are written in, so 30 is the only sensible answer, and an answer that changed with the order would show a mistake in the tally rather than in the formula.
