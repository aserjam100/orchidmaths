# Disproof by counterexample

A universal claim \(\forall x\, P(x)\) is disproved by a single \(x\) with \(\lnot P(x)\). That is the whole method: no argument about the general case is needed, and one witness is as complete as a thousand.

## One witness is enough

The claim "every odd number is prime" fails at \(x = 9\), since \(9 = 3 \cdot 3\). The claim "\(n^2 + n + 41\) is prime for every positive integer \(n\)" is more interesting, because the counterexample is far from the start.

A counterexample must also be checked against the claim's quantifiers: "for every positive integer" excludes 0, and a witness at 0 would not refute it. Reading the domain carefully is part of the method.

## Euler's polynomial

Euler's polynomial \(f(n) = n^2 + n + 41\) is prime for every \(n\) from 0 to 39. The values are small and the run is long, which is why the polynomial is famous as a trap: a learner who checks the first ten values sees primes and concludes too much.

At \(n = 40\) the value is \(1600 + 40 + 41 = 1681 = 41^2\), so the claim fails. The failure is not an accident: \(f(c-1) = (c-1)^2 + (c-1) + c = c^2\) for the constant \(c\), so a polynomial of this shape always fails at \(n = c-1\), and the interesting part is that every smaller value is prime.

## Searching for the first failure

Finding the smallest counterexample means scanning and stopping at the first failure, which is a computation rather than a proof. It is worth doing explicitly: test \(1, 2, 3, \dots\) in order, and the answer is the first index whose value factors.

For \(n^2 + n + 11\) the run ends at \(n = 10\), where \(100 + 10 + 11 = 121 = 11^2\); for \(n^2 + n + 17\) it ends at \(n = 16\), where the value is \(289 = 17^2\). The pattern \(n = c-1\) gives the factorisation, and the scan confirms that nothing failed earlier.
