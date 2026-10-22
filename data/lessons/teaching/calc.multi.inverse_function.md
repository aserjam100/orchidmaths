# The Inverse Function Theorem

**Inverse function theorem:** a map with a nonzero Jacobian determinant at a point is invertible on a neighbourhood of it, and the inverse is differentiable with the inverse Jacobian.

## The determinant condition

A map from the plane to itself is locally invertible near a point exactly when its Jacobian determinant is not zero there. The determinant measures the local area magnification of the map: if it is nonzero the map neither collapses a direction nor folds distinct directions together near the point, so it can be undone. The partial derivatives written as a matrix are the linear approximation of the map, and the determinant being nonzero is the condition that this approximation is invertible.

## The inverse and its derivative

The theorem says not only that a local inverse exists but also that it is differentiable, with Jacobian the matrix inverse of the original Jacobian at the corresponding point. This is the matrix version of the one-variable rule that the derivative of the inverse is the reciprocal of the derivative. The inverse matrix is computed by the usual adjugate formula divided by the determinant, so the inverse map's linear approximation is immediate once the original Jacobian is known.

## The polar map and friends

The map from polar to Cartesian coordinates has Jacobian determinant equal to the radius, so it is locally invertible away from the origin, where the determinant vanishes and the angle loses meaning. The same computation gives the area element of a polar integral, which is the absolute value of that determinant. The complex squaring map and change-of-variables formulas for multiple integrals are further instances, and each one is the determinant condition doing the work.
