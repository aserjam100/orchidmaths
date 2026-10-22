# LU Decomposition

**LU decomposition:** $A = LU$ with $L$ unit lower-triangular and $U$ upper-triangular. It solves $Ax = b$ as two triangular systems: $Ly = b$ forward, then $Ux = y$ backward. Needs nonzero leading pivots (else permute: $PA = LU$).

The factorisation is Gaussian elimination with the multipliers kept. Each elimination step subtracts a multiple of a pivot row from the rows below it, and those multiples are exactly the entries of $L$: the matrix records what elimination did, so $LU$ reconstructs $A$. The reward is that the expensive part is done once. Eliminating costs about $n^{3}/3$ operations, and solving a triangular system costs about $n^{2}$; so a factorisation reused for $m$ right-hand sides costs $n^{3}/3 + mn^{2}$ where inverting would cost $n^{3}$ per side.

## Worked: solve via Ly=b, Ux=y

Take $L = [[1,0],[2,1]]$, $b = [1,3]$:
1. Forward: $y_1 = 1$, then $2\cdot 1 + y_2 = 3$ gives $y_2 = 1$.
2. So $y = [1, 1]$: forward substitution reads top-down.
3. Then $Ux = y$ solves bottom-up the same way.

So triangular systems never need elimination — substitution walks straight through.

A $3\times3$ case shows where the multipliers come from. With

$$A = \begin{pmatrix}1&2&3\\4&5&6\\7&8&10\end{pmatrix},$$

elimination subtracts 4 times row 1 from row 2 and 7 times row 1 from row 3, then 2 times the new row 2 from row 3. The multipliers are read off as they are used, so

$$L = \begin{pmatrix}1&0&0\\4&1&0\\7&2&1\end{pmatrix}, \qquad U = \begin{pmatrix}1&2&3\\0&-3&-6\\0&0&1\end{pmatrix},$$

and multiplying back returns $A$ row by row: row 2 of $LU$ is $4\cdot(1,2,3) + (0,-3,-6) = (4,5,6)$, and row 3 is $7\cdot(1,2,3) + 2\cdot(0,-3,-6) + (0,0,1) = (7,8,10)$.

For $b = (1,2,3)$ the two solves are short. Forward, $Ly=b$:

$$y_1 = 1, \qquad 4y_1 + y_2 = 2 \Rightarrow y_2 = -2, \qquad 7y_1 + 2y_2 + y_3 = 3 \Rightarrow y_3 = 0.$$

Backward, $Ux = y$:

$$x_3 = 0, \qquad -3x_2 - 6x_3 = -2 \Rightarrow x_2 = \tfrac{2}{3}, \qquad x_1 + 2x_2 + 3x_3 = 1 \Rightarrow x_1 = -\tfrac{1}{3}.$$

Substituting back into $A$ confirms it: the first row gives $-1/3 + 4/3 = 1$, the second $-4/3 + 10/3 = 2$, the third $-7/3 + 16/3 = 3$. The advantage is not that the answer is prettier than elimination's but that the same $L$ and $U$ now solve every other $b$ for the price of two substitutions.

## Worked: det from U

With unit $L$ and $U$ diagonal $(2, 3)$:
1. $\det(L) = 1$ (unit diagonal).
2. $\det(U) = 2 \cdot 3 = 6$.
3. So $\det(A) = 6$: the determinant is the $U$-diagonal product, no expansion needed.

So LU factors the determinant computation into a diagonal read-off.

The identity behind it is $\det(A)=\det(L)\det(U)$. A triangular matrix's determinant is the product of its diagonal, and $L$'s diagonal is all ones by construction, so $\det(A)$ is the product of $U$'s diagonal. The $3\times3$ example above has $U$ diagonal $(1,-3,1)$, so $\det(A) = -3$; expanding along the first row gives $1(50-48) - 2(40-42) + 3(32-35) = 2 + 4 - 9 = -3$, which agrees.

One warning: a determinant of zero shows up as a zero on $U$'s diagonal, and that is exactly the case where the factorisation without pivoting failed. The determinant is still defined — the matrix is singular — but the elimination stopped, so the value has to come from a factorisation with a permutation, where the sign of $\det(P)$ enters: $\det(PA)=\det(P)\det(A)$.

## Pivoting

Zero leading pivots need row swaps: $PA = LU$ with a permutation $P$. Partial pivoting (largest available pivot) is the stable default.

A zero pivot does not mean the matrix is singular. For $A = [[0,1],[1,0]]$ the first pivot is 0, yet $A$ is its own inverse; swapping the rows first gives the factorisation, and recording the swap in $P$ keeps the identity $PA=LU$ true.

Partial pivoting does more than avoid the zero: it keeps the multipliers at most 1 in absolute value, by choosing the largest available entry in the column as the pivot. That bound is what controls the error. Without it, a small pivot produces a large multiplier, and the subtraction that follows cancels most of the significant digits of the rows it touches. The classical illustration is $A=\begin{pmatrix}10^{-20}&1\\1&1\end{pmatrix}$: eliminating without a swap multiplies by $10^{20}$ and destroys the second row in double precision, while swapping first keeps every multiplier at 1.

The residual risk is growth: partial pivoting bounds the multipliers but not the size of the entries that elimination creates, and the growth factor is what a worst-case error bound depends on. In practice it is small for almost every matrix, which is why partial pivoting rather than full pivoting is the default in every numerical library; full pivoting searches both rows and columns and is reserved for matrices known to be adversarial.
