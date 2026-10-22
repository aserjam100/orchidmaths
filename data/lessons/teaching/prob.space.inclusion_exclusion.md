# Inclusion-Exclusion

**Inclusion-exclusion:** For two events, $P(A \cup B) = P(A) + P(B) - P(A \cap B)$.

## Counting the overlap twice

Adding $P(A)$ and $P(B)$ counts the outcomes in both events twice. Subtracting the probability of the intersection removes the extra copy and leaves exactly the probability of the union. This is the probability form of the counting identity $|A \cup B| = |A| + |B| - |A \cap B|$, and with equally likely outcomes the two are the same statement.

## Worked: two events

Suppose a sample space has one hundred equally likely outcomes, event $A$ has forty of them, event $B$ has thirty, and their intersection has ten. The union then has $40 + 30 - 10 = 60$ outcomes, so $P(A \cup B) = 60/100 = 0.6$. Checking against the complement is a useful habit: the union misses $100 - 60 = 40$ outcomes, which is the complement of $A \cup B$.

## Three events and beyond

For three events the identity is $P(A \cup B \cup C) = \sum P(\text{singles}) - \sum P(\text{pairs}) + P(A \cap B \cap C)$, with the signs alternating. In general a union of $n$ events is the alternating sum over all nonempty subcollections. The same identity gives the union bound: dropping the negative terms leaves $P(\bigcup A_i) \le \sum P(A_i)$.
