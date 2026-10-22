# The Cross Product

**Cross product:** For $u = (u_1,u_2,u_3)$ and $v = (v_1,v_2,v_3)$ in $\mathbb{R}^3$, $u \times v = (u_2v_3 - u_3v_2,\; u_3v_1 - u_1v_3,\; u_1v_2 - u_2v_1)$. The result is perpendicular to both factors, and its length is the area of the parallelogram they span.

## A vector perpendicular to both factors

The formula is the determinant of a symbolic matrix whose first row is the standard basis. Swapping the factors swaps two rows, so $v \times u = -(u \times v)$: the product is anticommutative, and $u \times u = 0$.

## Worked: computing a cross product

1. $(1,2,3) \times (0,1,0) = (2\cdot0 - 3\cdot1,\; 3\cdot0 - 1\cdot0,\; 1\cdot1 - 2\cdot0) = (-3,0,1)$.
2. Check perpendicularity: $(-3,0,1) \cdot (1,2,3) = -3 + 0 + 3 = 0$ and $(-3,0,1) \cdot (0,1,0) = 0$.
3. $(1,0,1) \times (0,1,1) = (-1,-1,1)$.

## Length, area and orientation

$\|u \times v\| = \|u\|\|v\|\sin\theta$, the area of the parallelogram, so the cross product vanishes exactly when the vectors are parallel. The direction follows the right-hand rule, which is why the product changes sign when the order changes. In physics the same object gives torque and angular momentum.
