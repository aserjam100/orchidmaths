# The Matrix Exponential

**Matrix exponential:** For a square matrix $A$, $e^{At} = \sum_{n=0}^{\infty} (At)^n/n!$. The series converges for every matrix, and it is the matrix analogue of the scalar exponential.

## The series definition

The definition copies the scalar series, with the matrix product in place of multiplication. For a diagonal matrix the powers act entrywise, so the exponential is diagonal with the entries exponentiated. For a nilpotent matrix the series terminates, because some power of the matrix is zero.

## Worked: diagonal and nilpotent matrices

1. For $A = \operatorname{diag}(2,3)$, $e^A = \operatorname{diag}(e^2,e^3)$.
2. For $A = [[0,1],[0,0]]$, $A^2 = 0$, so $e^A = I + A = [[1,1],[0,1]]$.
3. In both cases $e^{At}$ solves $x' = Ax$ with $x(0) = I$.

## Solving x' = Ax

Differentiating the series term by term gives $\frac{d}{dt}e^{At} = A e^{At}$, and at $t = 0$ the exponential is the identity. So $x(t) = e^{At}x_0$ solves the initial value problem. When $A$ is diagonalizable, $e^{At} = P e^{Dt} P^{-1}$; the matrix exponential is the clean way to write the solution whether or not the eigenbasis is complete.
