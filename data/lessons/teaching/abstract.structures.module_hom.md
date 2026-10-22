# Module Homomorphisms

**Module homomorphism:** Let $M$ and $N$ be modules over the same ring $R$. A function $f: M \to N$ is an $R$-module homomorphism if $f(x + y) = f(x) + f(y)$ and $f(rx) = r f(x)$ for all $x, y \in M$ and all $r \in R$.

## Maps that preserve addition and scaling

The first condition makes $f$ a homomorphism of the additive groups, and the second makes it respect the ring action. The scalar $r$ is not an argument of $f$, so it appears unchanged on the right of the equation. A bijective module homomorphism is an isomorphism, and its inverse is automatically a module homomorphism as well.

## Kernels, images and injectivity

The kernel $\{x : f(x) = 0\}$ is a submodule of the domain, and the image is a submodule of the codomain. Because $f(x) = f(y)$ exactly when $f(x - y) = 0$, the map is injective precisely when its kernel is $\{0\}$. The first isomorphism theorem then gives $M/\ker f \cong \operatorname{im} f$ as modules.

## Worked: homomorphisms from Z_2 to Z_4

As $Z$-modules these are abelian groups, and a homomorphism is fixed by the image of the class of 1.

1. In $Z_2$ we have $2 \cdot 1 = 0$, so the image $a$ must satisfy $2a = 0$ in $Z_4$.
2. The solutions are $a = 0$ and $a = 2$.
3. So there are exactly two $Z$-module homomorphisms from $Z_2$ to $Z_4$.

The count is the size of the annihilator of 2 in $Z_4$, which is the subgroup $\{0, 2\}$.
