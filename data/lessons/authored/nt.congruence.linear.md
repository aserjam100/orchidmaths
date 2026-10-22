# Linear congruences

## Invertible coefficients

A congruence whose coefficient is coprime to the modulus can be solved by multiplying both sides by the coefficient's inverse. That inverse exists because of the coprimality, and it turns the congruence into one whose coefficient is one.

Finding the inverse is the extended Euclidean algorithm, and it is the only real work in the problem.

## When the coefficient is not a unit

A coefficient sharing a factor with the modulus may still be solvable, but the solutions are not unique: they form a set spaced by the modulus divided by the shared factor. The congruence has solutions exactly when the shared factor also divides the right-hand side.

Recognising that case matters, because there is no inverse to multiply by and the answer is a family rather than a residue.

## Least non-negative solutions

The solutions are residues, so the answer is usually given as the least one that is not negative. Two residues differing by the modulus are the same solution, and quoting the smallest keeps the answer canonical.

The check is a substitution: putting the answer back into the congruence must make the two sides congruent, and that verification is one multiplication.
