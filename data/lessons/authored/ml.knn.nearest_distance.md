# Nearest-neighbour distance

## Measuring to the neighbours

A nearest-neighbour classifier stores the training points and, for a new query, finds the ones closest to it. The label is then decided by those neighbours, so the distance computation is the whole of the method.

The distance is the ordinary Euclidean one, and comparing squared distances is enough because squaring does not change which is smallest.

## Squared distances

Squared distance keeps the arithmetic in whole numbers and avoids a square root per point, which matters when there are many points to compare. The nearest point is the same either way, so the square root is a cost with nothing to buy.

For the point \((1, 2)\), the squared distance is 5, and a training point at \((4, 6)\) is at squared distance 25, so the first is nearer.

## Choosing k

Using one neighbour makes the decision depend on a single training point, which is noisy. Using several and taking a vote is smoother, and the number to use is a parameter chosen by validation rather than by a formula.

The method needs no training at all, which is its virtue and its cost: the work moves entirely to prediction time, and every query is compared with every stored point.
