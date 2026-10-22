# Counting the same set two ways

## The method

Double counting means counting one collection in two different ways and setting the results equal. The equation is the result, and it is often the only way to prove an identity about binomial coefficients.

## A worked count

How many subsets of an \(n\)-element set contain a particular element? Every subset of the other \(n-1\) elements can be extended by that element, so the answer is \(2^{n-1}\).

The same number comes from counting pairs of the form (element, subset containing it). Each element lies in \(2^{n-1}\) subsets, so there are \(n \cdot 2^{n-1}\) pairs; and a subset of size \(k\) contributes \(k\) pairs, so the total is also the sum of the sizes of all subsets. The two expressions are equal, which is an identity about the rows of Pascal's triangle.

## When to reach for it

Double counting is the method to try when a question asks how many objects have a property and one way of grouping them is obvious while the other is what makes the arithmetic work.

The handshake count is the standard example: \(n\) people each shake \(n-1\) hands, which counts each handshake twice, so there are \(n(n-1)/2\) handshakes.
