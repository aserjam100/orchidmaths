# Evaluating a Laplace transform

## The transform of a power

The Laplace transform of a power of the variable is a factorial over a power of the transform variable. So the transform of the power turns an exponent into a factorial, and evaluating at a point turns the expression into a number.

The transform is defined by an integral, and for a power that integral evaluates in closed form.

## Evaluating at a point

Substituting a value of the transform variable is the last step, and choosing the value one makes the answer the factorial alone. The value must be positive for the transform to converge, which is a condition on the integral rather than on the substitution.

The transform is a function of a new variable, so evaluating it is not the same as transforming a number.

## Why the factorial appears

Integrating a power by parts repeatedly brings the exponent down one at a time until it reaches zero, and the product of the descending exponents is the factorial. The transform of the zeroth power is the base case, and the recursion is the reason the factorial shows up at all.

That is the same recursion the transform of a derivative inverts, which is why the transform turns a differential equation into an algebraic one.
