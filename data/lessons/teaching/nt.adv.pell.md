# Pell's Equation

**Pell's equation:** For a positive nonsquare integer $d$, the equation $x^2 - d y^2 = 1$ has infinitely many positive integer solutions, generated from the smallest one. The smallest solution is the fundamental solution, and it can be found from the continued fraction of $\sqrt d$.

## The smallest nontrivial solution

The trivial solution $(1,0)$ always exists, so "smallest" means smallest $x$ with $y > 0$. For small $d$ it can be found by trying $x = 2, 3, \dots$ and checking whether $(x^2-1)/d$ is a perfect square.

## Worked: d = 2 and d = 5

1. For $d = 2$, $x = 3$ gives $9 - 2\cdot4 = 1$, so the fundamental solution is $(3,2)$.
2. For $d = 5$, $x = 9$ gives $81 - 5\cdot16 = 1$, so the fundamental solution is $(9,4)$.
3. For $d = 3$, $x = 2$ already works: $4 - 3\cdot1 = 1$.

## Continued fractions generate the solution

The continued fraction of $\sqrt d$ is periodic, and the fundamental solution comes from a convergent just before the period repeats. Multiplying a solution by itself in the ring $\mathbb{Z}[\sqrt d]$ produces the next one, so all solutions form a sequence. Pell's equation is the classical example of a Diophantine equation with infinitely many solutions produced by a unit group.
