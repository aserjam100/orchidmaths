# Connected Sets

**Connected:** a set that cannot be written as the union of two nonempty pieces separated from each other; on the real line the connected sets are exactly the intervals.

## Separation and connectedness

A set is disconnected when it can be split into two nonempty pieces, each disjoint from the other and neither meeting the other's limit points. Connectedness is the denial of that: no such split exists. On the real line the connected sets are exactly the intervals, whether bounded or not and whether open, closed or half-open. The definition is topological rather than metric, so it survives the passage to general spaces, where connectedness is again the absence of a separation.

## Why intervals are connected

Suppose an interval were split into two separated nonempty pieces and let one point come from each. The segment between them lies inside the interval, because an interval contains everything between any two of its points. Moving along that segment from one chosen point to the other, the membership would have to change, and the point where it changes could be assigned to neither piece without contradicting the separation. This is the intermediate value idea in topological dress, and it is why the connected subsets of the line are the intervals.

## Continuity preserves connectedness

A continuous function carries a connected set to a connected set, because a separation of the image would pull back, through continuity, to a separation of the domain. This gives a clean proof of the intermediate value theorem: the interval is connected, its image under a continuous function is connected, hence an interval, so every value between the endpoint values is attained. The same argument shows that a continuous real function on a connected set has a connected image, which is exactly what the theorem asserts.
