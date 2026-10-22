# Quadratic residues

## Squaring pairs the residues

Squaring the non-zero residues modulo a prime sends a residue and its negative to the same square. So the squares come in from pairs of residues, and the number of distinct squares is the number of those pairs.

That pairing is the reason the set of quadratic residues is exactly half of the non-zero residues.

## Why exactly half

The residues split into pairs, one pair for each distinct square, and no square is shared between two pairs because a difference of squares factors. So the count of squares is the count of residues divided by two, which for a prime is one less than the prime over two.

For a composite modulus the argument fails in places, which is why the theory is stated for primes.

## The Legendre symbol

Whether a particular residue is a square is recorded by the Legendre symbol, which is one, minus one or zero. The symbol is the local question and the count is the global one, and the two are related by the fact that half the symbols are one.

That is why the symbol is multiplicative and why Gauss's lemma can compute it without enumerating the squares.
