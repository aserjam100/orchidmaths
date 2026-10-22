# Power Series

**Power series:** a series of coefficients times powers of a variable; it converges inside a radius and diverges outside, and inside it may be differentiated and integrated term by term.

## Radius of convergence

A power series is a series of the form the sum of coefficients times powers of the variable. For every such series there is a radius of convergence, a number or the symbol infinity, such that the series converges absolutely strictly inside the interval of that radius about the centre and diverges strictly outside it. The radius is found from the coefficients by the root test or the ratio test, and it is the limit of the reciprocal of the n-th root of the coefficients. At the two endpoints the test is silent and the behaviour must be settled separately.

## Term-by-term operations

Inside its radius a power series can be added, multiplied, differentiated and integrated term by term, and the result is another power series with the same radius except possibly at the endpoints. Term-by-term differentiation and integration are justified by the uniform convergence of the series on compact subintervals strictly inside the radius. This is what makes power series the workhorse of analysis: a function given by a series can be manipulated as if it were a polynomial, and each operation produces a new series that represents the new function.

## Worked: a geometric radius

For the series whose general term is a constant to the n times the power, the ratio of successive terms is the constant times the variable, so the ratio test asks for the absolute value of that product to be below one. The series therefore converges when the variable is within the reciprocal of the constant and diverges beyond it, giving a radius equal to one over the constant. For a constant of ten the series converges inside a tenth of the centre, which is the geometric series rescaled.
