# The Cauchy Product

**Cauchy product:** the series whose terms are the sums of all pairwise products with matching index, the correct way to multiply two series term by term.

## Multiplying series

The product of two series is not formed by multiplying them term by term; it is formed by collecting all pairs of terms whose indices sum to the same value. The coefficient of the n-th power in the product is the sum of the products of the coefficients whose indices add to n. This is exactly what happens when two polynomials are multiplied and then the finite sums are replaced by series, so the Cauchy product is the natural extension of polynomial multiplication to infinite series.

## Convergence of the product

When both series converge absolutely, the Cauchy product converges absolutely and its sum is the product of the two sums. The proof groups the double sum by the index sum, which absolute convergence lets you do without changing the value. If only one series is absolutely convergent and the other merely convergent, the product still converges to the product of the sums, by the Mertens theorem. If neither is absolutely convergent, the product may fail to converge at all, so the hypothesis is not cosmetic.

## The geometric example

Multiplying the geometric series by itself is the cleanest example: the coefficient of the n-th power is the number of index pairs that add to n, which is n plus one. The product is therefore the series of n plus one times the n-th power, whose radius matches the geometrically determined radius of the factors. Reading the coefficient as a count of pairs is the intuition behind every Cauchy product and shows why the coefficient grows rather than staying constant.
