# Zero Divisors and Cancellation

**Zero divisor:** In a ring, a nonzero element $a$ is a zero divisor if some nonzero $b$ satisfies $ab = 0$ or $ba = 0$. The definition excludes 0 itself: the point is that a nonzero element is annihilated.

## When a product of nonzero elements is zero

In the integers a product of nonzero numbers is never zero, but that property is special. In $Z_n$ the product $ab$ is zero mod $n$ whenever $n$ divides the ordinary product, even when neither $a$ nor $b$ is zero mod $n$. That happens exactly when $a$ and $b$ share a factor with $n$ between them.

## Worked: zero divisors in Z_4 and Z_6

1. In $Z_4$, $2 \cdot 2 = 4 = 0$, so 2 is a zero divisor. The elements 1 and 3 are units, so 2 is the only nonzero zero divisor.
2. In $Z_6$, $2 \cdot 3 = 6 = 0$, so both 2 and 3 are zero divisors, and so is 4.
3. In a field such as $Z_7$ no such product exists, because every nonzero element has an inverse.

## No zero divisors means cancellation

The cancellation law, $ab = ac$ with $a \neq 0$ implies $b = c$, holds precisely when there are no zero divisors. Subtract to get $a(b - c) = 0$; either $b - c = 0$ or $a$ annihilates a nonzero element. So a ring where cancellation is valid is exactly a ring with no zero divisors, which is why integral domains are the settings where familiar arithmetic is legal.
