# Degree sequences

The **degree sequence** of a graph is the list of its vertex degrees, usually written in non-increasing order. It records how many edges meet each vertex without recording which vertices they are.

## Reading a degree sequence

Writing the sequence in non-increasing order makes it canonical, so two graphs with the same sequence can be compared directly. For \(K_{2,3}\) the two vertices in the small part each have degree three and the three in the large part each have degree two, so the sequence is \(3, 3, 2, 2, 2\).

The sequence is an invariant, not a complete description: non-isomorphic graphs can share one, so a matching degree sequence does not prove two graphs are the same.

## The handshake lemma

Every edge contributes one to the degree of each of its two endpoints, so the degrees sum to twice the number of edges: \(\sum_v \deg(v) = 2\lvert E \rvert\). For the sequence above the sum is twelve, so the graph has six edges, which matches \(2 \cdot 3 = 6\).

A consequence is that the number of odd degrees is even, because the sum is even. A sequence with an odd number of odd entries cannot be the degree sequence of any graph, which is the first feasibility test.

## Feasibility

**Havel–Hakimi** decides whether a sequence is graphical: take the largest degree \(d\), connect that vertex to the \(d\) next-largest, remove it, and repeat; the sequence is graphical exactly when the process ends at all zeros.

For \(3, 3, 2, 2, 2\): connect the first vertex to the next three, leaving \(2, 1, 1, 2\); reorder to \(2, 2, 1, 1\); connect and continue, and the process reaches zeros. For \(3, 3, 3, 1\) the same procedure fails, and no graph has that sequence — the handshake test passes there, so feasibility needs the stronger algorithm.
