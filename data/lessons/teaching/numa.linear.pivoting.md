# Partial Pivoting

**Partial pivoting:** at each stage of elimination, interchange rows so that the largest available entry in the current column becomes the pivot.

## Why the pivot size matters

Elimination divides by the pivot and multiplies by the multiplier, so a tiny pivot produces a large multiplier and magnifies any rounding error. Choosing the largest available entry in the column as the pivot keeps the multipliers at most one in absolute value, which controls the growth of error. The interchange is recorded and applied to the right-hand side as well.

## Worked: avoiding a small pivot

Suppose the diagonal entry is tiny while a lower entry is large. Without a swap the multiplier is huge and the elimination amplifies rounding error badly. Swapping the rows puts the large entry on the diagonal, the multiplier becomes small, and the same solution comes out with far less error. The swap changes nothing about the answer, only about how it is computed.

## The guarantee and its limits

Partial pivoting keeps the multipliers bounded by one, which in most practical cases keeps the error growth modest, though matrices can be constructed to defeat it. Complete pivoting, which searches the whole remaining submatrix, gives a stronger guarantee at more bookkeeping cost. A zero pivot column signals a singular matrix and needs a different treatment, such as least squares.
