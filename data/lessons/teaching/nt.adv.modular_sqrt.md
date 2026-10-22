# Square Roots Modulo a Prime

**Modular square root:** For an odd prime $p$ and a quadratic residue $a$, the congruence $x^2 \equiv a \pmod p$ has exactly two solutions, $x$ and $p-x$. The smaller one lies in $1,\dots,(p-1)/2$.

## Two roots when one exists

The polynomial $x^2-a$ has at most two roots over a field, and if $x$ is one then $-x = p-x$ is the other. When $a$ is a nonresidue there are no roots at all, which Euler's criterion detects by computing $a^{(p-1)/2} \bmod p$: the value is 1 for a residue and $-1$ for a nonresidue.

## Worked: roots modulo 11 and 13

1. Modulo 11, $4^2 = 16 \equiv 5$ and $7^2 = 49 \equiv 5$, so the roots of 5 are 4 and 7.
2. Modulo 13, $4^2 = 16 \equiv 3$ and $9^2 = 81 \equiv 3$, so the roots of 3 are 4 and 9.
3. The smaller root is the one below $p/2$, here 4 in both cases.

## Finding a root in practice

For small $p$, search the residues. For large $p$ with $p \equiv 3 \pmod 4$, the root is $a^{(p+1)/4}$; for $p \equiv 1 \pmod 4$, the Tonelli-Shanks algorithm handles the general case. Both run in polynomial time once a quadratic nonresidue is available.
