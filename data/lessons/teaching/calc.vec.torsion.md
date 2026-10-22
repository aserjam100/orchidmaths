# Torsion of a Space Curve

**Torsion:** For a regular space curve, the torsion is
$$ \tau = \frac{(r' \times r'') \cdot r'''}{\lVert r' \times r''\rVert^2}. $$
It measures how quickly the curve twists out of its osculating plane, and it is defined only where $r' \times r''$ is nonzero.

## How a curve leaves its osculating plane

At each point the first two derivatives span the osculating plane, and the binormal $B = (r' \times r'')/\lVert r' \times r''\rVert$ is normal to it. Torsion records how $B$ turns as the point moves. A planar curve has constant binormal, so its torsion is zero; a helix has constant nonzero torsion.

## Worked: the twisted cubic and the helix

1. For $r(t) = (t, t^2, t^3)$ at $t = 0$: $r' = (1,0,0)$, $r'' = (0,2,0)$, $r''' = (0,0,6)$.
2. Then $r' \times r'' = (0,0,2)$, its squared length is 4, and the dot product with $r'''$ is 12, so $\tau = 3$.
3. For the helix $r(t) = (\cos t, \sin t, t)$ at $t = 0$, the same computation gives $\tau = 1/2$.

## Torsion versus curvature

Curvature measures how fast the tangent turns and is always nonnegative; torsion measures how fast the curve leaves its plane and is signed. A circle has constant positive curvature and zero torsion, a helix has both constant, and a straight line has both zero. The pair $(\kappa, \tau)$ determines a space curve up to rigid motion, which is the fundamental theorem of space curves.
