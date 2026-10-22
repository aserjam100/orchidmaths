# Metric Spaces

**Metric space:** a set with a distance function that is nonnegative, symmetric, vanishes exactly on equal points and obeys the triangle inequality.

## The axioms

A metric on a set is a function that assigns to each ordered pair of points a nonnegative number, called their distance, subject to four conditions: the distance is zero exactly when the points are equal, it is symmetric, and it satisfies the triangle inequality that the distance from one point to a third never exceeds the sum of the distances through any intermediate point. The triangle inequality is the one that does the work, because it controls how errors accumulate. Every argument on the line that used the absolute value generalises once these four properties are assumed.

## The examples

The real line with the absolute value, the plane with the usual straight-line distance, and any set with the discrete metric that assigns one to distinct points and zero to equal ones are all metric spaces. So are spaces of functions, where the distance between two functions is the supremum of their difference, provided the supremum is finite. The same abstract language describes all of them, so a theorem proved from the axioms holds for every example at once, which is the economy a metric space buys.

## Balls and the topology

The open ball of radius r around a point is the set of points whose distance from it is strictly less than r. Balls are the metric reading of open intervals, and they generate the open sets: a set is open when every point of it has a ball contained in the set. Two metrics on the same set are equivalent when they generate the same open sets, which is also the same notion of convergence. The whole topological vocabulary of neighbourhood, interior, closure and compactness is rebuilt from balls.
