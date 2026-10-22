# Linear independence

## Independent or dependent

A set of vectors is linearly independent when no one of them can be written as a combination of the others. Two vectors in the plane are independent exactly when they do not lie along the same line, and three vectors in space are independent exactly when they do not lie in the same plane.

Dependence means the set contains a redundancy: one vector adds no direction that the others did not already provide.

## Counting the independent ones

The number of independent vectors in a set is the dimension of what they span, and it is the rank of the matrix formed by putting them as rows. Three vectors with a rank of two contain two independent directions and one that is a combination of them.

The count is what matters rather than which ones, since any maximal independent subset spans the same space.

## Why the count is the dimension

A space has a dimension when every independent set inside it has at most the same size, and every spanning set has at least that size. The independent sets that reach the maximum are the bases, and they all have the same number of vectors.

That equality is what makes dimension well defined, and it is the reason a basis is a choice rather than a property of the space.
