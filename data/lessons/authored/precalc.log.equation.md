# Solving logarithmic equations

## Undoing the logarithm

An equation \(\log_a x = k\) is solved by undoing the logarithm with its inverse, the exponential: \(x = a^k\). For \(\log_2 x = 5\) the solution is \(x = 32\).

The domain must be checked, since a logarithm is only defined for positive arguments, and the exponential step produces a positive number automatically.

## Checking for extraneous solutions

An equation that combines two logarithms into one, such as \(\log_a(\text{first}) = \log_a(\text{second})\), is solved by equating the arguments. That step can produce a value for which one of the original arguments is negative or zero, and such a value must be discarded rather than kept.

This is the one place where solving a logarithmic equation differs from solving an exponential one: the algebraic step is safe, but the domain is not preserved by it.

## Equations with a sum of logarithms

A sum of logarithms becomes the logarithm of a product, and a difference becomes a quotient. Rewriting first is what makes the equation one logarithm on each side, after which the arguments can be equated and the domain checked.

For \(\log_2 x + \log_2(x - 2) = 3\), the product form gives \(x(x-2) = 8\), whose solutions are 4 and -2, and the negative one is discarded because the original logarithm required \(x > 2\).
