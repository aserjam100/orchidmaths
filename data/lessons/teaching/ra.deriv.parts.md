# Integration by Parts

**Integration by parts:** the product rule read backwards turns an integral of a product into a boundary term minus another integral, trading a product for something easier.

## The rule and its origin

Integrating both sides of the product rule gives the integration by parts formula: the integral of u times the derivative of v equals the product of u and v minus the integral of v times the derivative of u, all understood on a definite interval with the product evaluated at the endpoints. The point is to move the derivative from one factor to the other. Choosing which factor to differentiate is the whole art, and the usual guide is to differentiate the factor that simplifies, such as a polynomial, and integrate the factor that stays simple.

## Worked: a polynomial times an exponential

For the integrand that multiplies the input by the exponential, differentiating the input turns it into a constant and integrating the exponential leaves it essentially unchanged. Applying the formula gives a boundary term of the product evaluated at the endpoints minus the integral of the exponential alone. The remaining integral is elementary, so the whole computation reduces to one boundary evaluation and one easy antiderivative, which is why this integrand is the standard first example of the technique.

## Repeated and tabular use

When the factor being differentiated is a polynomial of higher degree, the rule is applied repeatedly until the derivative becomes zero, and the boundary terms accumulate. Organising the repetitions in a table, differentiating down one column and integrating across the other, keeps track of the alternating signs and the decreasing degree. The same formula is the source of the reduction formulas for powers of trigonometric and logarithmic functions, where each application lowers an exponent by one instead of killing a factor outright.
