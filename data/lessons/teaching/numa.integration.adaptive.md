# Adaptive Quadrature

**Adaptive quadrature:** refine the mesh only where the estimated error is large, spending effort where the integrand is hard and none where it is easy.

## Spending effort where it is needed

A uniform mesh overestimates the work needed on smooth stretches and underestimates it near sharp features. An adaptive routine estimates the local error on each subinterval and splits only the ones that exceed their share of the tolerance. The result is accuracy near the difficult features without wasting samples on the flat parts.

## Worked: a local error estimate

A common estimate compares two rules of different orders on the same subinterval, such as Simpson and the trapezoid rule. Their difference is a practical bound on the error of the cheaper one. If that difference is small the subinterval is accepted; if it is large the subinterval is split and the two halves are tested in turn.

## Why it pays off

Adaptive routines are the standard tool for integrands with localised peaks, endpoint singularities or other features that a uniform mesh would smear. They deliver a requested tolerance with far fewer evaluations than a uniform rule on the same problem. Their cost is the bookkeeping and recursion, which is why simple problems still use a fixed composite rule.
