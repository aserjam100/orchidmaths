# Line integrals

## Integrating along a path

A line integral adds up a quantity along a curve rather than across an interval. Writing the path as \(x(t), y(t)\) over \(a \le t \le b\), the integral of a field \((P, Q)\) is

\[\int_C P\,dx + Q\,dy = \int_a^b \left(P\,\frac{dx}{dt} + Q\,\frac{dy}{dt}\right) dt.\]

The path's direction matters: reversing the parameter interval changes the sign.

## Gradient fields

When \((P, Q)\) is the gradient of a function \(f\), the integral is \(f\) evaluated at the endpoint minus its value at the start. For \((x, y)\), which is the gradient of \((x^2 + y^2)/2\), the integral from the origin to \((4, 6)\) is \((16 + 36)/2 = 26\).

The path never enters the calculation, which is the whole content of the fundamental theorem for line integrals.

## Path independence

A field is a gradient exactly when it is path independent, and by Clairaut that happens when \(\partial P/\partial y = \partial Q/\partial x\). When the test fails, a closed loop can return a non-zero value, and that number is what a curl measures.
