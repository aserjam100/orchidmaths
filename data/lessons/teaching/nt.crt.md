# Chinese Remainder Theorem

**Chinese remainder theorem (CRT):** If $m_1$ and $m_2$ are coprime, the system

$$x \equiv r_1 \pmod{m_1}, \quad x \equiv r_2 \pmod{m_2}$$

has a unique solution modulo $M=m_1m_2$.

The theorem is a statement about counting. There are $m_1$ possible values of $x \pmod{m_1}$ and $m_2$ of $x \pmod{m_2}$, so there are $m_1m_2=M$ pairs of remainders; there are also exactly $M$ residue classes modulo $M$. CRT says the map from a class to its pair of remainders is a bijection, so every pair occurs exactly once. The hypothesis matters: modulo 4 and 6 the pairs are not independent, and the system $x\equiv1\pmod4$, $x\equiv3\pmod6$ has no solution at all.

The two congruences can be solved one at a time, and that is how the construction below works: each congruence contributes a term that is 1 modulo its own modulus and 0 modulo the other.

## Constructing the Solution

### Coprime Moduli
The hypothesis $\gcd(m_1,m_2)=1$ is essential. It guarantees Bézout coefficients $u,v$ with $um_1+vm_2=1$, which build the solution.

To see why, reduce the first congruence modulo any common divisor $d$ of $m_1$ and $m_2$: it forces $x\equiv r_1\pmod d$, and the second forces $x\equiv r_2\pmod d$. If $d>1$ and $r_1\not\equiv r_2\pmod d$, the two requirements contradict each other and nothing solves the system. Coprimality makes $d=1$ the only common divisor, so no such obstruction exists, and the uniqueness half of the theorem follows: if $x$ and $x'$ both solve the system then $m_1\mid x-x'$ and $m_2\mid x-x'$, and because $m_1,m_2$ are coprime their product divides $x-x'$.

### Combining Congruences
Write $M_1=M/m_1$, $M_2=M/m_2$. Let $y_1$ be the inverse of $M_1$ mod $m_1$ and $y_2$ the inverse of $M_2$ mod $m_2$. Then

$$x = r_1M_1y_1 + r_2M_2y_2 \pmod{M}$$

Each coefficient is a switch. The number $M_1$ is a multiple of $m_2$, so $M_1\equiv0\pmod{m_2}$; and $y_1$ exists precisely because $\gcd(M_1,m_1)=1$, so $M_1y_1\equiv1\pmod{m_1}$. Reading the formula modulo each modulus in turn,

$$x \equiv r_1\cdot1 + r_2\cdot0 \equiv r_1 \pmod{m_1}, \qquad x \equiv r_1\cdot0 + r_2\cdot1 \equiv r_2 \pmod{m_2},$$

which is the pair that was asked for. The identity is worth naming: the terms $M_1y_1$ and $M_2y_2$ are an *orthogonal idempotent* pair modulo $M$. Each squares to itself, they annihilate each other, and they sum to 1 modulo $M$.

## Example

Solve $x \equiv 2 \pmod{3}$ and $x \equiv 3 \pmod{5}$.
$M=15$, $M_1=5$, $M_2=3$. Inverses: $5^{-1}\equiv 2 \pmod{3}$, $3^{-1}\equiv 2 \pmod{5}$.
$x = 2\cdot5\cdot2 + 3\cdot3\cdot2 = 20+18=38 \equiv 8 \pmod{15}$. Check: $8\equiv2\pmod3$, $8\equiv3\pmod5$.

The check has to be run on both congruences, because a single one can be satisfied by accident. The answer is a class, not a number: every $x=8+15k$ solves the system, and no other class does.

Three or more congruences are handled by the same construction with more terms, or by merging two at a time: the pair above collapses to $x\equiv8\pmod{15}$, and a third congruence $x\equiv r_3\pmod{m_3}$ with $\gcd(15,m_3)=1$ is then solved against modulus 15. The moduli need only be pairwise coprime for the merge to work, not pairwise prime.

## Why It Is Used

CRT turns one computation modulo a large number into several computations modulo small ones, because arithmetic in each small modulus is independent and can be done in parallel. The classical application is to compute a value modulo $M$ from its residues modulo the factors of $M$; the reconstruction step is the formula above. It is also the standard way to prove that $\mathbb{Z}/M\mathbb{Z}\cong\mathbb{Z}/m_1\mathbb{Z}\times\mathbb{Z}/m_2\mathbb{Z}$ as rings when $M=m_1m_2$ with coprime factors, and that statement is the reason the residues can be combined at all.
