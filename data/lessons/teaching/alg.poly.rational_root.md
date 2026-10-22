# The Rational Root Theorem

**Rational root theorem:** every rational root of a polynomial with integer coefficients is a fraction whose numerator divides the constant term and whose denominator divides the leading coefficient.

## The candidates

The theorem narrows the search for rational roots to a finite list. Each candidate is a fraction formed from a divisor of the constant term over a divisor of the leading coefficient, in every combination and both signs. For a monic polynomial the denominator is always one, so the only rational roots possible are integer divisors of the constant term, which makes the list very short.

## Testing and reducing

Once the candidates are listed, each is tested by substitution or by synthetic division until one gives zero. Finding one root reduces the polynomial by that factor and lowers the degree, and the process repeats on the quotient. The theorem does not guarantee that a rational root exists; it only says that any rational root must come from the list, so a list with no successes means the roots are irrational or complex.

## Why the theorem works

Writing a rational root in lowest terms and substituting it into the polynomial, then multiplying through by the denominator's highest power, shows that the numerator must divide the constant term and the denominator must divide the leading coefficient. The argument is a divisibility count, not a computation, and it is why the theorem is a search tool rather than a solution method. For a monic polynomial the divisibility condition collapses to the integer divisors of the constant.
