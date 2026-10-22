# Tournaments

**Tournament:** A tournament is an orientation of the edges of a complete graph: every pair of vertices is joined by exactly one directed edge. Each pair plays exactly one game.

## Orienting a complete graph

The underlying graph is $K_n$, with $\binom{n}{2} = n(n-1)/2$ edges and the same number of games. Each edge can point either way, so the number of tournaments on labelled vertices is $2^{\binom{n}{2}}$. A transitive tournament orders the vertices so that every edge points from earlier to later, and it is the only acyclic tournament up to relabelling.

## Worked: counting small tournaments

For $n = 4$ there are $\binom{4}{2} = 6$ games and hence $2^6 = 64$ tournaments. For $n = 5$ there are $\binom{5}{2} = 10$ games and $2^{10} = 1024$ tournaments. The count grows very fast, which is why tournaments are studied through their properties, such as the score sequence, rather than by enumeration. For $n = 3$ the eight tournaments are the transitive ones and the two cyclic orders, up to relabelling one cyclic tournament.

## Scores and Hamiltonian paths

The score of a vertex is its out-degree, and the score sequence records the scores in non-decreasing order. A tournament is transitive exactly when its score sequence is $0, 1, \dots, n-1$, and every other tournament contains a directed triangle. Rédei's theorem says every tournament on $n$ vertices has a Hamiltonian directed path, through all vertices, even though it need not have a Hamiltonian cycle. A tournament has a Hamiltonian cycle exactly when it is strongly connected.
