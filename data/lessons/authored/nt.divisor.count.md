# Counting divisors

## From a factorisation

The number of positive divisors follows from the prime factorisation alone: each divisor is built by choosing, for every prime, how many of its copies to include. A divisor is therefore one choice per prime.

So counting divisors is counting the combinations of those choices, and no divisor has to be listed.

## Choosing an exponent for each prime

For a prime raised to a power, the available choices run from including none of its copies to including all of them, which is one more option than the exponent. The counts multiply across the primes because the choices for different primes are independent.

For a number whose factorisation is a square, that gives an odd divisor count, which is why a perfect square is the only kind of number with a square's worth of pairs.

## Why the exponents multiply

Choosing an exponent for one prime does not restrict the choices for another, so the number of combined choices is the product of the individual counts. That is the multiplication principle, and it is the same reasoning behind counting outcomes of independent events.

The formula is worth reading that way rather than memorising: the divisor count is a product of options, and each option is an exponent between zero and the prime's power.
