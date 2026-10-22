# The Chain Rule

**Chain rule:** the derivative of a composition is the derivative of the outer function at the inner value times the derivative of the inner function, so rates multiply along a chain.

## Rates multiply

If a quantity depends on an intermediate quantity which in turn depends on the input, the overall rate of change is the product of the two rates. Written for a composition, the derivative of the outer function at the inner value times the derivative of the inner function at the input gives the derivative of the composition. The rule is not an approximation: dividing the change in the outer value by the change in the input and inserting the intermediate change in numerator and denominator turns the two quotients into the two derivatives in the limit.

## Worked: a power of a linear function

For the function that raises a linear expression to a whole power, the chain rule applies the power rule to the outer function and multiplies by the slope of the linear expression. If the inner value at the chosen point is some number, the derivative is the exponent times the coefficient times that inner value raised to the exponent minus one. Every step is arithmetic once the inner value is computed, and the result is the slope of the tangent to the composite curve at the point.

## Chains of three and more

The rule extends to longer compositions: each link contributes its derivative, evaluated at the value it receives, and the derivatives multiply. For a composition of three functions the derivative is the outer derivative at the middle value times the middle derivative at the inner value times the inner derivative at the input. In differential notation the intermediate variables cancel formally, which is the bookkeeping that makes the chain rule look like ordinary fraction multiplication without pretending the derivatives are fractions.
