# The Extreme Value Theorem

**Extreme value theorem:** a continuous function on a closed bounded interval attains a maximum and a minimum, so there are points of the interval where the function takes its largest and smallest values.

## The statement

A function continuous on a closed bounded interval has a maximum value and a minimum value that are actually attained at points of the interval. Both hypotheses are needed: the interval must be closed so the endpoints are included and bounded so the values cannot run away. The theorem turns a bound on the values into the existence of a place where the bound is reached, which is exactly what optimisation and many existence proofs require.

## Why the hypotheses matter

Drop either hypothesis and the conclusion fails. The function one over x is continuous on the half-open interval from nought to one but is unbounded there, so it has no maximum. The same function is continuous on the open interval from nought to one but again unbounded. On the whole line the identity function is continuous and bounded below but has no maximum. Each example removes exactly one hypothesis and loses exactly one part of the conclusion, which is the sharpest way to see why both are stated.

## Consequences

The theorem says the continuous image of a closed bounded interval is again a closed bounded interval, so it is bounded and attains its endpoints. It is the existence half of every maximum-minimum optimisation: the candidates found by calculus are compared knowing that a largest value exists. Together with the intermediate value theorem it shows that a continuous function on a closed bounded interval takes every value between its attained minimum and maximum, so the image is exactly that interval.
