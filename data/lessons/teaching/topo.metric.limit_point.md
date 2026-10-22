# Limit Points

**Limit point:** A point $x$ is a limit point of a set $A$ when every neighbourhood of $x$ contains a point of $A$ other than $x$. The point $x$ itself need not belong to $A$, which is why limits can sit outside the set.

## Every neighbourhood meets the set again

The definition is local and repeatable: no matter how small the ball around $x$, it still catches some point of $A$. In a metric space this is the same as saying some sequence of points of $A$, all different from $x$, converges to $x$. A point of $A$ that is not a limit point is isolated.

## Worked: the limit point of 1/n

1. Take $A = \{1/n : n \ge 1\}$.
2. Every ball around 0 has radius $r > 0$ and contains some $1/n$ once $n$ is large enough.
3. No $1/n$ is a limit point: a ball of radius smaller than the gap to its neighbours contains no other point of $A$.

So the only limit point is 0, and 0 is not in $A$.

## Closed sets contain their limit points

A set is closed exactly when it contains all its limit points. The interval $[0,1]$ is closed because the only limits of sequences inside it are points of $[0,1]$; the interval $(0,1]$ is not closed because 0 is a limit point it omits. This is why closedness can be checked with sequences in a metric space.
