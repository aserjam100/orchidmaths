# k-Means Clustering

**k-means:** Choose $k$ centres, assign each point to the nearest centre, move each centre to the mean of its points, and repeat until the assignments stop changing. It partitions the data into $k$ clusters without using labels.

## Assign, then move the centroid

Each iteration has two steps. Assignment uses squared Euclidean distance, so a point belongs to the centre it is closest to. The update replaces each centre by the coordinate-wise mean of its assigned points. Both steps can only decrease the within-cluster sum of squares, so the loop terminates.

## Worked: one centroid update

1. Points $(1, 1)$, $(3, 5)$ and $(5, 3)$ are assigned to one cluster.
2. The x-coordinate is $(1 + 3 + 5)/3 = 3$.
3. The y-coordinate is $(1 + 5 + 3)/3 = 3$, so the new centroid is $(3, 3)$.

## Choosing k and when k-means fails

The number of clusters is a hyperparameter chosen by validation or an elbow plot, not by the algorithm. k-means assumes clusters are roughly spherical and similarly sized, and it converges to a local optimum that depends on initialization, which is why k-means++ spreads the starting centres.
