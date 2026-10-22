# p-Groups

**p-group:** A finite group is a p-group if its order is a power of a prime $p$. These are the groups a finite group is assembled from: by Sylow, every finite group has subgroups of prime-power order.

## Groups whose order is a power of a prime

If $|G| = p^k$ with $k \ge 1$, every element has order a power of $p$, and Cauchy's theorem forbids an element of order $q \ne p$. The order of a p-group is therefore a power of one prime, and no other prime divides it.

## The center cannot be trivial

The class equation $|G| = |Z(G)| + \sum [G:C_G(x)]$ writes the order as the center plus the sizes of the noncentral conjugacy classes. Each of those sizes is a power of $p$ larger than 1, hence divisible by $p$, and $|G|$ is divisible by $p$, so $|Z(G)|$ is divisible by $p$. A nontrivial p-group therefore has a nontrivial center.

## Worked: centers of Q_8 and D_4

Both groups have order $8 = 2^3$.

1. In the quaternion group $Q_8$ the center is $\{1, -1\}$, of order 2.
2. In $D_4$ the center is $\{e, r^2\}$, also of order 2.
3. Both orders are divisible by 2, and 2 is the least possible value.

The least possible center order for a group of order $p^k$, once $k$ is at least 3, is $p$, and the upper unitriangular matrices over $F_p$ of order $p^3$ realize it.
