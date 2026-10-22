# Extrapolation and Its Dangers

**Extrapolation:** estimating a function outside the range of the data, where nothing constrains the fit and the error can grow quickly.

## Beyond the data

Interpolation is safe because the data surround the query point and hold the curve in place. Extrapolation has no data beyond the last node, so the interpolant is free to swing away. The error formula makes this precise: the product of the distances to the nodes grows without bound outside the interval, and the derivative term is unbounded too.

## Worked: the growing error bound

A polynomial that fits the data well inside its interval can bend sharply just beyond the last point. Predicting twice as far out can multiply the error bound many times over, because the bound is proportional to the distances to every node. A prediction one step past the data is risky; a prediction ten steps out is usually worthless.

## The careful exceptions

Richardson extrapolation is not the same thing: it extrapolates the limit of a numerical sequence using the known form of the error, not a fitted curve beyond its data. Where a physical model constrains the behaviour outside the data, a principled extrapolation can work. Absent such a model, a prediction far outside the data should be reported with a much wider uncertainty.
