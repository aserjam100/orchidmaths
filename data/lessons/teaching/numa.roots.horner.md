# Horner's Method and Polynomial Roots

**Horner's method:** evaluate a polynomial by nested multiplication, from the highest coefficient down, using one multiplication and one addition per term.

## Nested multiplication

Writing the polynomial in the nested form factors out the powers of x, so each step multiplies the running total by x and adds the next coefficient. Starting from the leading coefficient and working down evaluates the polynomial with the fewest possible operations. The same scheme, read as synthetic division, also gives the value of the derivative.

## Worked: a small evaluation

The polynomial with coefficients one, two and three, read from the constant term up, is one plus two x plus three x squared. Evaluated at two by nesting, the scheme forms three, then three times two plus two is eight, then eight times two plus one is seventeen. Direct substitution gives the same seventeen with more multiplications.

## From evaluation to roots

Horner's scheme is the engine of synthetic division, which deflates a polynomial once a root is known. Newton's method for polynomials uses Horner to get both the value and the derivative cheaply at each step. This is why polynomial root finding is fast in practice: each Newton step costs only a pair of nested passes.
