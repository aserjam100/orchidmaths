# Counting Paths in a DAG

**Counting paths in a DAG:** Process the vertices in topological order and, for each vertex, set its path count to the sum of the path counts of its predecessors. The count at the target is the number of directed paths from the source.

## Dynamic programming on a DAG

Set the source count to one and every other count to zero. In topological order, each edge $u \to v$ adds the count at $u$ to the count at $v$, so a vertex's count becomes the number of directed paths from the source to it. Because every predecessor appears earlier, a vertex is finished before any successor reads it, and the whole computation takes time linear in the vertices plus the edges.

## Worked: a diamond

Consider the DAG with vertices 0, 1, 2, 3 and edges $0 \to 1$, $0 \to 2$, $1 \to 3$, $2 \to 3$. The order 0, 1, 2, 3 is topological. The counts begin at 1, 0, 0, 0. Processing 0 gives counts 1, 1, 1, 0, and processing 1 and 2 gives the count at 3 the value $1 + 1 = 2$. The two paths are $0 \to 1 \to 3$ and $0 \to 2 \to 3$, so the answer 2 is correct.

## Why acyclicity matters

A directed cycle would make the path count infinite, because a path could loop around the cycle any number of times before reaching the target, so the problem is only well posed on a DAG. Acyclicity is also what makes the topological order possible: it guarantees that each vertex is processed once, after all its predecessors. For an undirected graph the same counting problem is #P-complete, so the directed acyclic case is the tractable one.
