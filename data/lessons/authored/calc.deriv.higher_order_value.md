# Evaluating a higher-order derivative

## The power rule applied k times

The k-th derivative of a power is found by applying the power rule k times, and each application brings the current exponent down as a factor and reduces it by one. So the result is a falling product of the exponents times the exponent that remains.

For a power of x the computation is therefore a product and a subtraction, and nothing about it needs the intermediate derivatives written out.

## The falling factorial

The factors that come down are the exponents counted downward, which is the falling factorial. It is why a derivative of a power eventually reaches zero: after as many steps as the exponent, one of the factors is zero, and every later derivative of a polynomial is zero too.

That is the reason a polynomial of degree n has derivatives of every order up to n and none above.

## Evaluating at a point

Substituting the point into the k-th derivative does not need the derivative written as a function: substitute into the product form directly. The exponent that remains is the original one minus the number of steps, so a value can be computed without any intermediate algebra.

The check is that the sign follows the parity of the point and the power in the way the product says, which catches a lost factor or a forgotten step.
