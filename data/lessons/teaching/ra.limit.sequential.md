# The Sequential Criterion

**Sequential criterion:** a function has limit L at a point exactly when, for every sequence approaching that point but never equal to it, the corresponding sequence of function values approaches L.

## Reduction to sequences

The sequential criterion converts a statement about a limit at a point into a statement about limits of sequences. It says the limit of the function is L exactly when every sequence of inputs approaching the point, with no term equal to the point, produces a sequence of outputs approaching L. This lets every theorem about limits of sequences be transferred to limits of functions, and it is the standard way to prove that a candidate is not the limit.

## Proving a limit fails

The contrapositive is the workhorse for showing a limit does not exist. Exhibit one sequence approaching the point whose outputs do not approach the candidate, or two sequences approaching the point whose output sequences have different limits. For instance, a function that equals minus one on one side of a point and one on the other has sequences approaching from the two sides with different output limits, so the two-sided limit does not exist even though each one-sided limit does.

## Continuity at a point

The same criterion describes continuity: a function is continuous at a point when the limit there equals the value of the function at the point, and by the sequential criterion this happens exactly when every sequence approaching the point is mapped to a sequence approaching the value. So continuity is the statement that the function commutes with the passage to the limit along sequences. This sequential view of continuity is often easier to apply than the epsilon-delta definition and is used to prove the algebra of continuity.
