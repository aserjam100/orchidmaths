# The trace as a sum of eigenvalues

## The trace

The trace of a square matrix is the sum of its diagonal entries. For a two-by-two matrix it is the sum of the two numbers on the main diagonal, and it takes one glance to compute.

Despite being so easy to read off, it carries information about the eigenvalues that is not visible on the diagonal.

## Why it is the sum

The trace equals the sum of the eigenvalues, counting each with its algebraic multiplicity. The reason is that the characteristic polynomial's second-highest coefficient is the trace with a sign, and that coefficient is also the sum of the roots by the factor theorem.

So the two descriptions of the same coefficient give the identity, which is why knowing one eigenvalue of a two-by-two matrix determines the other from the trace alone.

## Using it

The trace and the determinant together determine the eigenvalues of a two-by-two matrix, since they give their sum and their product. That is often enough to answer a question without computing a characteristic polynomial at all.

For larger matrices the pair is no longer sufficient, and the whole polynomial is needed.
