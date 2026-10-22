# Mixing Problems

**Mixing problem:** A tank of volume $V$ receives brine at rate $r$ with concentration $c$, and the well-mixed solution leaves at the same rate. The amount of salt $Q(t)$ satisfies $Q' = rc - rQ/V$.

## Rate in minus rate out

The inflow delivers salt at $rc$ kilograms per minute. The outflow carries the current concentration $Q/V$, so it removes salt at $rQ/V$. The difference is the rate of change of $Q$, a first-order linear equation with constant coefficients.

## Worked: a brine tank

1. A 100-litre tank receives 0.2 kg/L at 5 L/min and drains at 5 L/min.
2. The equation is $Q' = 1 - Q/20$.
3. The limiting amount solves $0 = 1 - Q/20$, so $Q = 20$ kg.

## The limiting amount

At equilibrium the concentration in the tank equals the inflow concentration, so $Q = Vc$ regardless of the flow rate. The solution approaches that value exponentially with time constant $V/r$: faster flow reaches the limit sooner but does not change the limit. If the inflow and outflow rates differ, the volume changes and the equation gains a time-dependent coefficient.
