# Genus of a Surface

**Genus:** For a closed orientable surface, the genus is the number of handles: the sphere has genus 0, the torus genus 1, the double torus genus 2, and so on. The Euler characteristic and the genus determine each other by $\chi = 2 - 2g$.

## Handles and holes

Adding a handle to a surface lowers its Euler characteristic by 2. The sphere has $\chi = 2$; attach one handle and the torus has $\chi = 0$; attach another and the double torus has $\chi = -2$. The genus counts handles, not the number of holes in the informal sense.

## Worked: reading genus from Euler characteristic

1. A surface with $\chi = 0$ solves $2 - 2g = 0$, so $g = 1$.
2. A surface with $\chi = -4$ solves $2 - 2g = -4$, so $g = 3$.
3. A surface with genus 4 has $\chi = 2 - 2(4) = -6$.

So one number, $\chi$, carries the same information as the genus for closed orientable surfaces.

## Orientable versus non-orientable

For a closed non-orientable surface the formula becomes $\chi = 2 - k$, where $k$ is the number of crosscaps. The projective plane has $\chi = 1$ and $k = 1$; the Klein bottle has $\chi = 0$ and $k = 2$. Orientability is therefore part of the classification, not a consequence of $\chi$: the torus and the Klein bottle share $\chi = 0$.
