# The Frenet Frame

**Frenet frame:** At a point of a regular space curve, the unit tangent $T = r'/\lVert r'\rVert$, the unit normal $N = T'/\lVert T'\rVert$, and the unit binormal $B = T \times N$ form an orthonormal frame $(T, N, B)$. The binormal can also be computed as $B = (r' \times r'')/\lVert r' \times r''\rVert$.

## Three perpendicular unit vectors

$T$ points along the motion, $N$ points toward the centre of curvature, and $B$ completes the right-handed triple. The frame changes along the curve, and the Frenet-Serret equations describe that change using curvature and torsion. The order of the triple matters: swapping two vectors changes the sign of the cross product.

## Worked: the frame of a circle

1. For $r(t) = (\cos t, \sin t, 0)$, $r' = (-\sin t, \cos t, 0)$ and $\lVert r'\rVert = 1$.
2. $T = (-\sin t, \cos t, 0)$; $T' = (-\cos t, -\sin t, 0)$, so $N = (-\cos t, -\sin t, 0)$.
3. $B = T \times N = (0, 0, 1)$, and the torsion is zero because the circle is planar.

At $t = 0$ this gives $T = (0,1,0)$, $N = (-1,0,0)$ and $B = (0,0,1)$.

## Why the binormal matters

$B$ is the normal to the osculating plane, so it is the vector the torsion turns. The generated questions ask for the whole frame — $T$, $N$ then $B$, entered as the nine components $(T_x, T_y, T_z, N_x, N_y, N_z, B_x, B_y, B_z)$ — and $B$ is defined only where $r' \times r''$ is nonzero. When the curve is planar, $B$ is constant; when it twists, $B$ moves and the torsion measures the rate.
