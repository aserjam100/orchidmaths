# Bézout's Identity

**Bézout's identity:** For integers $a$ and $b$ with $\gcd(a,b)=d$, there exist integers $x$ and $y$ such that

$$ax + by = d$$

Moreover, the set $\{ax+by : x,y \in \mathbb{Z}\}$ is exactly the set of multiples of $d$.

The identity is an existence statement with a constructive proof: the extended Euclidean algorithm produces $x$ and $y$, and it does so in the same number of steps the plain algorithm takes. The second half is the part that carries the weight in proofs. It says that no combination of $a$ and $b$ can be smaller than $d$ unless it is 0, so $d$ is the *smallest positive* value the expression can take, and that is why the coefficients of Bézout are the natural certificate of a gcd.

## The GCD as a Linear Combination

### Finding the Coefficients
The extended Euclidean algorithm finds $x$ and $y$ alongside $d$.
For $a=12$ and $b=8$: $\gcd(12,8)=4$. Indeed $12(1)+8(-1)=4$. Also $12(-1)+8(2)=4$.

The run for 12 and 8 is two steps: $12=8\cdot1+4$ and $8=4\cdot2+0$, so the last nonzero remainder is 4. Working back through the first line, $4=12-8\cdot1$, which reads off $x=1$, $y=-1$ directly. The pair is not unique: adding $8t$ to $x$ and subtracting $12t$ from $y$ leaves the expression unchanged, so every $(1+8t,\,-1-12t)$ works, and $(1+8(-1),-1+12)=( -7,\,11)$ is the same identity written with a smaller first coefficient.

### Smallest Positive Combination
The smallest positive value of $ax+by$ is $d$. Any larger combination is a multiple of $d$.

The proof is one division. Every combination is divisible by $d$, because $d$ divides both $a$ and $b$; so the positive values form a set of multiples of $d$, and it is enough to show that $d$ itself is one of them. Let $s$ be the smallest positive value of $ax+by$. Dividing $a$ by $s$ gives $a=qs+r$ with $0\le r<s$, and

$$r = a - qs = a - q(ax_0+by_0) = a(1-qx_0) + b(-qy_0),$$

so $r$ is itself a combination. If $r$ were positive it would be a smaller one than $s$, which contradicts minimality, so $r=0$ and $s\mid a$. The same division of $b$ by $s$ gives $s\mid b$, so $s$ is a common divisor of $a$ and $b$; and because $d$ is the greatest one, $d\le s$. Since every combination, including $s$, is a multiple of $d$, also $d\mid s$, hence $s=d$.

## Example

Find integers $x,y$ with $27x+18y=\gcd(27,18)=9$.
Since $27=18\cdot1+9$ and $18=9\cdot2$, working backwards: $9=27-18\cdot1$, so $x=1$, $y=-1$.

A longer run shows what the algorithm does with the quotients. Take $a=240$ and $b=46$:

$$240=46\cdot5+10,\qquad 46=10\cdot4+6,\qquad 10=6\cdot1+4,\qquad 6=4\cdot1+2,\qquad 4=2\cdot2+0.$$

The gcd is 2, the last nonzero remainder. Working upward one line at a time,

$$2=6-4,\qquad 4=10-6,\qquad 6=46-10\cdot4,\qquad 10=240-46\cdot5,$$

and substituting from the bottom,

$$2 = 6 - (10-6) = 2\cdot6 - 10 = 2(46-10\cdot4) - 10 = 2\cdot46 - 9\cdot10 = 2\cdot46 - 9(240-46\cdot5) = 47\cdot46 - 9\cdot240.$$

So $240(-9) + 46(47) = 2$. The coefficients can be larger than either input, which is why the identity is quoted for integers and not for naturals.

## Application: Inverses

If $\gcd(a,m)=1$, Bézout gives $ax+my=1$, so $ax \equiv 1 \pmod{m}$ and $x$ is the modular inverse of $a$.

That reduction is the reason the extended algorithm is the standard way to invert an element modulo a prime, and why the inverse exists exactly when the gcd is 1: if $d>1$ then every combination $ax+my$ is a multiple of $d$, so no such combination can equal 1.

The same identity is the engine behind the Chinese remainder theorem's construction. There, $m_1$ and $m_2$ are coprime, Bézout produces $u$ and $v$ with $um_1+vm_2=1$, and the pair $(um_1,\,vm_2)$ is the idempotent decomposition of 1 that the reconstruction formula uses.
