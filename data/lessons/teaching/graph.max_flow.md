# Max-Flow Min-Cut

**Max-flow min-cut theorem:** In a directed network with integer capacities, the maximum value of a flow from a source to a sink equals the minimum total capacity of a set of edges whose removal separates the source from the sink.

## Flows and cuts

A flow assigns a value to every directed edge, never exceeding its capacity, and conserves the amount at every vertex except the source and the sink. A cut is a partition of the vertices into a source side and a sink side, and its capacity is the sum of the capacities of the edges that cross from the source side to the sink side. Every flow is at most every cut, because the flow must cross the cut. The theorem says the best flow and the smallest cut meet.

## Worked: a small network

Take source 0, sink 3, and edges $0 \to 1$ with capacity 3, $0 \to 2$ with capacity 3, $1 \to 3$ with capacity 2, and $2 \to 3$ with capacity 3. Send 2 units along $0 \to 1 \to 3$ and 3 units along $0 \to 2 \to 3$, for a total of 5. The cut that separates the source from the sink at the two edges into 3 has capacity $2 + 3 = 5$, so no flow can exceed 5. The flow and the cut agree at 5.

## Residual graphs and integrality

An augmenting path is a path from the source to the sink in the residual graph, which records unused capacity forward and the flow that can be cancelled backward. The Ford-Fulkerson method augments along such paths until none remains, at which point the reachable vertices form a minimum cut and the theorem is proved constructively. If the capacities are integers, every augmentation adds an integer, so an integral maximum flow exists.
