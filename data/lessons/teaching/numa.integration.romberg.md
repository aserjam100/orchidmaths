# Romberg Integration

**Romberg integration:** combine trapezoid estimates at successively halved widths by Richardson extrapolation to cancel the leading error terms and raise the order.

## Extrapolating the trapezoid ladder

The composite trapezoid error has an expansion in even powers of the width. Richardson extrapolation takes two estimates at widths in the ratio two and combines them to cancel the leading term, leaving an estimate of higher order. Applied repeatedly, this builds a triangular table whose first column is the trapezoid rule and whose later columns are increasingly accurate.

## Worked: one extrapolation step

If the trapezoid estimates at widths of one half and one quarter are combined as four times the finer minus the coarser, all over three, the leading error term cancels. The combination is exact for the next order up, which for the trapezoid rule means it behaves like Simpson's rule. One more level of extrapolation reaches the order of a higher rule again.

## When it works and when it does not

Romberg integration is most effective when the integrand is smooth, so that the error expansion is clean and the extrapolation cancels real terms. A discontinuity, a singularity or a noise floor breaks the expansion and the extrapolation can make things worse. The table is cheap once the trapezoid column is computed, because the finer widths reuse the earlier sample points.
