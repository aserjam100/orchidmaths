# Catastrophic Cancellation

**Cancellation:** the loss of leading significant digits when two nearly equal numbers are subtracted, which amplifies the relative error already present in the inputs.

## Subtracting close numbers

If two numbers agree in their first several digits, the subtraction cancels those digits and leaves a result made mostly of the digits that were uncertain. The absolute error is unchanged but the result is small, so the relative error blows up. The amplification is roughly the sum of the sizes divided by the size of the difference.

## Worked: a close difference

Subtracting nine hundred ninety nine from one thousand and one gives two, a result with one significant digit, while each input carried at least four. The amplification factor is two thousand divided by two, which is a thousand, so an input relative error of one part in a million becomes an output error of one part in a thousand.

## Avoiding cancellation

Cancellation is often avoidable by rearranging the algebra before computing. The quadratic formula is the classic case: the smaller root is computed through a difference of nearly equal quantities, but rewriting it as the product of the roots divided by the larger root removes the subtraction of close numbers. Good algorithms choose the stable form of a formula.
