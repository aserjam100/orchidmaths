# Vandermonde's identity

## A sum of products

Choosing \(r\) items from two groups of sizes \(a\) and \(b\) can be done by splitting the choice according to how many come from the first group:

\[\binom{a+b}{r} = \sum_{i=0}^{r} \binom{a}{i}\binom{b}{r-i}.\]

Each term picks \(i\) from the first group and the remaining \(r-i\) from the second, and the terms are disjoint because they differ in \(i\), so they add.

## Evaluating the sum

The identity turns a convolution of two binomial rows into a single coefficient. For \(a = 3\), \(b = 4\) and \(r = 2\) the sum is \(\binom{7}{2} = 21\), which is faster than adding the four products.

When an index falls outside its range the corresponding coefficient is zero, which is what makes the sum run only over the pairs that make sense.

## Where it appears

The identity says the rows of Pascal's triangle convolve to later rows, and it underlies the convolution of two generating functions: the coefficients of a product are sums of products of the two sequences' coefficients.

Recognising a convolution of binomial coefficients is usually faster than expanding it term by term.
