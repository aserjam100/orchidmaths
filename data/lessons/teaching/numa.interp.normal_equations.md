# The Normal Equations

**Normal equations:** the linear system obtained by setting the derivatives of the least-squares sum to zero, which has one equation for each unknown coefficient.

## Setting the derivatives to zero

Minimising a smooth function of several unknowns means setting every partial derivative to zero. For a linear fit those equations are linear in the coefficients, so they form a square system called the normal equations. Their solution is the least-squares fit, and the matrix of the system is built from the inner products of the basis functions.

## Worked: the intercept

Once the slope is known, the intercept is whatever makes the line pass through the centre of mass of the data. With the points one, four and four over x values nought, one and two, the slope is one and a half, the mean of y is three and the mean of x is one, so the intercept is three minus one and a half, which is 1.5. The line always passes through the centroid.

## The conditioning warning

The normal equations square the condition number of the design matrix, so with nearly dependent columns the system becomes hard to solve accurately. This is why a QR factorization is often preferred even though the normal equations are simpler. The equations are still the right way to describe the fit; they are just a poor way to compute it.
