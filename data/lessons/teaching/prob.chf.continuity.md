# The Continuity Theorem

**Continuity theorem:** a sequence of distributions converges weakly if and only if its characteristic functions converge pointwise to a function that is continuous at zero.

## Transforms identify laws

The theorem says that weak convergence of distributions is equivalent to pointwise convergence of their transforms, provided the limit is continuous at the origin. Because a characteristic function determines its distribution uniquely, recognising the limit transform identifies the limit law. This is the bridge between the analytic computation and the probabilistic conclusion.

## The role of continuity at zero

Continuity of the limit at the origin is the technical condition that prevents probability from escaping to infinity along the sequence. Without it the transforms could converge to a function that is not a characteristic function, and no limiting distribution need exist. The condition is exactly what tightness reduces to in the transform language.

## The central limit theorem

The main use of the theorem is to prove the central limit theorem, by showing that the transforms of the standardised sums converge to the normal transform. Once the limit transform is recognised, the theorem turns that analytic convergence into weak convergence of the distributions. The continuity theorem is the machinery that closes the argument.
