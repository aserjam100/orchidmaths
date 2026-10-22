# Order of an Element

**Order:** The order of $g$ is the smallest $n > 0$ with $g^n = e$ (or infinite if none exists). In additive notation, the smallest $n > 0$ with $ng = 0$.

The definition is the same statement in both notations, and the additive form is the one to keep in mind for $\mathbb{Z}$ and for vector spaces, where the operation is written as addition and the identity as 0. The order measures how long the cyclic subgroup $\langle g\rangle$ is before it closes: the list $e, g, g^2, \dots$ repeats as soon as it returns to $e$, so the order is the size of the subgroup that one element generates.

## Worked: order of 2 in Z_6

Add 2 to itself mod 6 until hitting 0:
1. $2 + 2 = 4$, still nonzero.
2. $2 + 2 + 2 = 6 = 0$ mod 6.
3. Three copies suffice and fewer do not, so the order of 2 is 3.

So order is found by repeated operation, stopping at the first return to identity.

The same walk works for every element and produces a formula. In $\mathbb{Z}_n$ the element $k$ has order $n/\gcd(n,k)$, because $mk\equiv0\pmod n$ exactly when $n\mid mk$, and the smallest such $m$ is $n/\gcd(n,k)$. For $n=6$: the elements 1 and 5 have order 6 since they are coprime to 6, the elements 2 and 4 have order 3, the element 3 has order 2, and 0 has order 1. The orders that occur, $1, 2, 3, 6$, are the divisors of 6, which is not a coincidence: the order of an element always divides the size of the group.

That divisibility is Lagrange's theorem read from the other side. The cyclic subgroup $\langle g\rangle$ has exactly $\operatorname{ord}(g)$ elements and sits inside $G$, so its size divides $|G|$ by Lagrange, and the element's order divides the group's order as a consequence. It also gives the practical version of the test: raising $g$ to a power that is a multiple of its order returns $e$, so only the divisors of $|G|$ are worth trying.

## Worked: a transposition has order 2

Apply $(1\,2)$ twice:
1. First application swaps 1 and 2.
2. Second application swaps them back: identity.
3. So $(1\,2)$ has order 2, while the 3-cycle $(1\,2\,3)$ needs three applications: order 3.

So in $S_n$, cycle type reads off the order: disjoint cycle lengths, take the lcm.

The lcm rule is worth seeing on a case where the answer is not a cycle length. Take the permutation $(1\,2)(3\,4\,5)$ in $S_5$: the first cycle needs two applications to close and the second needs three, so the whole permutation closes when both have, which is at 6. The order is $\operatorname{lcm}(2,3)=6$, and no power smaller than 6 returns every point: after 2 the transposition is the identity but the 3-cycle has moved, after 3 the reverse.

Because the cycle lengths of an element of $S_n$ form a partition of $n$, the possible orders are the lcms of partitions of $n$. For $n=5$ those partitions give $1, 2, 3, 4, 5, 6$, and 6 is the largest order — reached by the cycle type $2+3$ above. The largest order in $S_n$ is therefore not $n$ once $n\ge5$, and grows like $e^{\sqrt{n\log n}}$, not like $n$.

The lcm is also the reason a group's exponent can exceed every element's order. In an abelian group the order of $gh$ divides $\operatorname{lcm}(\operatorname{ord}g,\operatorname{ord}h)$; in a non-abelian group it need not, and the product of two involutions can have any order at all.

## Infinite order

In $\mathbb Z$ under addition, $-1$ never returns to 0: $n \cdot (-1) = 0$ only for $n = 0$. Elements of infinite order generate infinite cyclic subgroups.

What infinite order means precisely is that the powers $g^{0}, g^{\pm1}, g^{\pm2}, \dots$ are all distinct, so the map $k\mapsto g^{k}$ is an injection from $\mathbb{Z}$ into $G$, and $\langle g\rangle$ is isomorphic to $\mathbb{Z}$. An element of finite order makes the same map periodic instead, with period its order. A group in which every element has finite order is called torsion; one with no element of infinite order is torsion-free; and a group can have both kinds of element without contradiction, as $\mathbb{Z}\times\mathbb{Z}_2$ does.

The distinction matters for the study of order because the two cases behave differently under every operation. A finite-order element can be inverted by a single power: if $\operatorname{ord}(g)=n$ then $g^{-1}=g^{n-1}$. An infinite-order element has no such expression, which is why the inverse is written as a symbol rather than computed in $\mathbb{Z}$ and in free groups. And a homomorphism can collapse an infinite-order element to a finite-order one but never the reverse: the order of $\varphi(g)$ divides the order of $g$ whenever the latter is finite.
