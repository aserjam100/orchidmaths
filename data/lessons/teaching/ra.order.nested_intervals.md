# Nested Intervals

**Nested interval property:** a decreasing sequence of nonempty closed bounded intervals has a nonempty intersection, and when their lengths tend to zero the intersection is a single point.

## The property

Take intervals that shrink inside one another: each is contained in the previous, all are closed and bounded, and none is empty. Their intersection is not empty. The proof looks at the left endpoints, which increase, and the right endpoints, which decrease. The increasing left endpoints are bounded above by every right endpoint, so by completeness they have a supremum, and that supremum lies in every interval. When the lengths tend to zero the supremum is forced to be the only common point.

## Why closedness matters

The word closed in the statement is doing real work. The open intervals from nought to one over n are nested, nonempty, bounded and shrinking, yet their intersection is empty, because no positive number is below every one over n. Closedness keeps the limiting point inside each interval; dropping it lets the common point escape. Boundedness is doing work too: nested rays such as the intervals from n to infinity have empty intersection.

## The bisection method

Nested intervals are the machine behind several existence theorems. To locate a point with a desired property, halve the current interval and keep the half that still has the property; the halves form a nested sequence whose unique common point is the answer. This is how the intermediate value theorem finds a root and how the Bolzano-Weierstrass theorem extracts a convergent subsequence. The method is also practical: it bounds the position of the answer to within the current length.
