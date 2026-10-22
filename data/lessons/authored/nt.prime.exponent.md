# The exponent of a prime in a factorial

## Powers of the prime in the factorial

A factorial multiplies every number up to n, and each of them may contain several copies of a prime. The exponent of that prime in the whole product is the total number of copies across all the factors.

So the question is a counting one, and the count can be taken by powers rather than by factoring each number.

## Counting with floors

The multiples of the prime up to n each contribute one copy, giving the floor of n divided by the prime. Multiples of its square contribute a second copy, multiples of its cube a third, and so on until the power exceeds n.

Adding those floors gives the exponent, and the sum terminates because the powers eventually pass n.

## Why each power is counted once

A number divisible by the cube of the prime is counted at the first power, at the square and at the cube, which is exactly the three copies it contains. Counting by powers instead of by numbers is what makes the method short.

That is Legendre's formula, and it is the standard way to get the exponent of a prime in a factorial without any factorisation at all.
