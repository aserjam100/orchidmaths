# Predator-Prey Systems

**Lotka-Volterra system:** $x' = ax - bxy$ and $y' = -cy + dxy$, where $x$ is the prey and $y$ the predator. The prey grows on its own and is reduced by encounters; the predator dies without prey and grows by encounters.

## Two coupled populations

The interaction terms $-bxy$ and $+dxy$ couple the equations. The system is nonlinear, but its nonzero equilibrium is easy: set $x' = 0$ to get $y = a/b$, and set $y' = 0$ to get $x = c/d$.

## Worked: the nonzero equilibrium

1. Take $a=2$, $b=1$, $c=3$, $d=1$.
2. Setting $y' = 0$ gives $x = c/d = 3$.
3. Setting $x' = 0$ gives $y = a/b = 2$, so the nonzero equilibrium is $(3,2)$.

## Cycles and linearization

Linearizing at the nonzero equilibrium gives a center: the eigenvalues are purely imaginary, so the linearization predicts closed orbits. The nonlinear system also has a conserved quantity, and its orbits are closed level curves around the equilibrium. Predator-prey cycles are therefore robust in this model, unlike the spirals a stable focus would produce.
