# Counting by complement

When the objects you want are awkward to describe and the ones you do not want are easy, count the whole and subtract. **Counting by complement** is the rule \(\lvert A \rvert = \lvert U \rvert - \lvert A^{c} \rvert\) for a universe \(U\) and a subset \(A\).

## The complement move

The move is only worth making when the complement is simpler. "How many of these strings contain at least one \(a\)?" forces you to separate the cases of one \(a\), two \(a\)'s and so on; "how many contain no \(a\)?" is one multiplication, because every position is restricted to the other letters.

That is the shape of every complement count: the condition "at least one" becomes "not none", and "not none" is one case instead of many.

## At least one

Take strings of length 4 over the three letters \(\{a, b, c\}\). There are \(3^4 = 81\) strings in all, and \(2^4 = 16\) of them use only \(b\) and \(c\). So \(81 - 16 = 65\) contain at least one \(a\).

The subtraction is exact because the two sets partition the universe: every string either contains an \(a\) or it does not. No inclusion-exclusion is needed, which is what makes this cheaper than counting the overlapping cases directly.

## Choosing the smaller side

The same rule applies to sets in a set problem: "how many subsets contain at least one of these two elements" is easier as the total minus the subsets containing neither. The count is \(2^n - 2^{n-2}\) when the two elements are distinct.

Deciding which side to count is a judgement, and the test is which side has fewer constraints. The complement usually wins when the wanted property is an existence claim and the unwanted one is an absence.
