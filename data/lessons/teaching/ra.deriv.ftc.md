# The Fundamental Theorem of Calculus

**Fundamental theorem of calculus:** differentiation and integration are inverse processes, so an integral with a variable upper limit differentiates back to the integrand, and a definite integral is the change in an antiderivative.

## Two halves

The theorem has two halves that face each other. The first says that if a function is continuous then the integral from a fixed base to a variable upper limit, as a function of the limit, has derivative equal to the integrand at that limit. The second says that a definite integral equals the difference of any antiderivative evaluated at the endpoints. Together they invert differentiation and integration, converting the hard limit of sums into the easy search for an antiderivative.

## Why it holds

The variable-limit integral changes by roughly the integrand at the upper limit times the small increment of the limit, because on a tiny interval the function is nearly constant. Dividing by the increment and letting it shrink makes the difference quotient tend to the integrand, which is the first half. The second half follows by writing the variable-limit integral as an antiderivative and using the mean value theorem to show that every antiderivative differs from it by a constant, so the change over an interval is the same for all of them.

## Worked: a definite integral

For a linear integrand the antiderivative is a quadratic plus a linear term, and evaluating it at the upper endpoint and subtracting its value at the lower endpoint gives the integral. Evaluating at the lower endpoint removes the constant of integration, so only the change in the antiderivative survives. This is the routine practice of the second half: find one antiderivative, evaluate it at both endpoints, and subtract. The value is exact, because the antiderivative already encodes the limit of the sums.
