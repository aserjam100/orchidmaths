# Arc Length in Polar Coordinates

**Polar arc length:** For a polar curve $r = f(\theta)$ traced from $\theta = a$ to $\theta = b$, the arc length is
$$ L = \int_a^b \sqrt{r^2 + \left(\frac{dr}{d\theta}\right)^2}\, d\theta. $$
The $r^2$ term replaces the 1 of the Cartesian formula because the radial direction itself moves as the angle changes.

## The polar arc element

A small change $d\theta$ moves the point along a direction with radial component $dr$ and angular component $r\,d\theta$, which are perpendicular. Pythagoras gives $ds = \sqrt{dr^2 + r^2 d\theta^2}$, and pulling out $d\theta$ produces the integrand above.

## Worked: a secant curve and a cardioid

1. For $r = \sec\theta$, $r' = \sec\theta\tan\theta$, so $r^2 + (r')^2 = \sec^2\theta(1 + \tan^2\theta) = \sec^4\theta$. The integrand is $\sec^2\theta$, and the length from 0 to $\pi/4$ is $\tan(\pi/4) = 1$.
2. For $r = 1 + \cos\theta$, the integrand simplifies to $2\cos(\theta/2)$ on the first half-turn, so the half-cardioid has length 4 and the full cardioid has length 8.

## When the integral is elementary

The polar integrand is often a square root that simplifies only after a trigonometric identity, as in the two examples above. When it does not simplify, the length is left as an integral or evaluated numerically. Symmetry can halve or quarter the work: the full cardioid is twice its upper half, and a circle about the origin has the elementary length $r(\theta_2 - \theta_1)$.
