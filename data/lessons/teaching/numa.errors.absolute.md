# Absolute and Relative Error

**Error:** the absolute error is the size of the difference between an approximation and the true value, and the relative error divides it by the size of the true value.

## Two measures of error

The absolute error is $|\hat{x} - x|$, the plain distance between the approximation and the truth. The relative error is that distance divided by $|x|$, the true value, which removes the units and makes the number comparable across scales. A kilometre of absolute error is huge when measuring a computer chip and small when measuring a planet, so the relative error is usually the one that matters.

## Worked: a rounded measurement

A quantity with true value one thousand is approximated by nine hundred ninety nine. The absolute error is one, and the relative error is one divided by a thousand, which is 0.001. Approximating five hundred by five hundred and five gives an absolute error of five and a relative error of five over five hundred, which is 0.01, ten times larger than the first.

## Why relative error is the default

Rounding to a fixed number of significant digits keeps the relative error bounded and lets the absolute error grow with the size of the number. Numerical analysis therefore measures accuracy in relative terms almost everywhere. An algorithm is judged good when its relative error stays close to the unit roundoff, not when its absolute error is small in some fixed unit.
