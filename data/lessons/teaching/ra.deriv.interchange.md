# Interchange of Limits

**Interchange of limits:** the passage of a limit through a derivative, an integral or a second limit; it is valid under uniform convergence and fails in general without it.

## Which interchanges need care

Many operations in analysis are themselves limits, and asking whether two limits can be taken in either order is asking whether they commute. Differentiation is the delicate one: the derivative of the limit need not be the limit of the derivatives, even when the functions converge everywhere. Integration and continuity are more forgiving under uniform convergence. The general lesson is that an interchange of limits needs a hypothesis stronger than pointwise convergence, and uniform convergence supplies it for integrals, continuity and, with the derived series, for differentiation.

## Interchanging the integral

If a sequence of integrable functions converges uniformly on a closed bounded interval, then the integrals converge to the integral of the limit. The reason is that the error between each function and the limit is bounded by the uniform error over the whole interval, and the integral of a constant over a bounded interval is the constant times the length. Pointwise convergence gives no such bound, so the limit of the integrals can differ from the integral of the limit, and the standard counterexample is a tall narrow spike sliding toward an endpoint.

## Differentiating a series

To differentiate a series of functions term by term, it is not enough that the original series converge uniformly; the series of derivatives must converge uniformly. When it does, the derivative of the sum is the sum of the derivatives, because the mean value theorem lets the difference quotient of the original series be controlled by the derived series. This is the cleanest hypothesis for differentiating a power series inside its radius, and it is why power series behave so well.
