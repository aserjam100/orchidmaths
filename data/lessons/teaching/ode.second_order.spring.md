# The Spring-Mass System

**Spring-mass equation:** An undamped mass on a spring satisfies $m y'' + k y = 0$, or $y'' + \omega^2 y = 0$ with $\omega = \sqrt{k/m}$. The solution is $y = A\cos(\omega t) + B\sin(\omega t)$, with $A = y(0)$ and $B = y'(0)/\omega$.

## Frequency from the equation

The coefficient $\omega^2$ is the square of the natural frequency: a stiffer spring or a lighter mass raises $\omega$ and speeds up the oscillation. The amplitude and phase are set by the initial displacement and velocity, not by the equation alone.

## Worked: an undamped oscillation

1. For $y'' + 4y = 0$ with $y(0)=0$ and $y'(0)=2$, $\omega = 2$ and $y = \sin(2t)$.
2. Then $y(\pi/4) = \sin(\pi/2) = 1$.
3. For $y'' + 9y = 0$ with $y(0)=1$ and $y'(0)=0$, $y = \cos(3t)$, so $y(\pi/6) = 0$.

## Damping and forcing

Adding a damping term gives $y'' + 2\zeta\omega y' + \omega^2 y = 0$: underdamped for $\zeta < 1$, critically damped at $\zeta = 1$, overdamped above it. A forcing term adds a particular solution, and when the forcing frequency matches $\omega$ the amplitude grows linearly, which is resonance.
