# Stability of an Algorithm

**Backward stability:** an algorithm is backward stable when the answer it returns is the exact answer to a nearby problem whose data differ from the given data by a small relative amount.

## Pushing the error back into the data

Instead of asking how far the computed answer is from the true one, backward stability asks which nearby problem the answer solves exactly. If that problem is within a tiny relative perturbation of the original, the algorithm did the best a finite-precision method can do. The distance between the answer and the truth is then governed by the condition number.

## Worked: an accumulation of roundoff

Summing numbers in a different order gives slightly different answers, but each computed partial sum is the exact sum of a slightly perturbed list of the same numbers. The algorithm is backward stable; the variation between orders reflects the condition of the summation, which is mild when the terms have one sign and worse when they cancel.

## Why it is the right standard

Backward stability separates the part of the error an algorithm controls from the part the problem imposes. A backward-stable method on a well-conditioned problem has a relative error near machine epsilon, which is the best possible. On an ill-conditioned problem even a backward-stable method is inaccurate, and the blame lies with the problem.
