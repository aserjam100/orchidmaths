# The Exponential Family

**Exponential family:** a family of distributions whose mass or density can be written $h(x) \exp(\eta(\theta) T(x) - A(\theta))$, with $h$ an $x$-only factor, $T$ the sufficient statistic, $\eta$ the natural parameter and $A$ the log-partition function.

## The canonical form

The form separates the density into a factor that depends only on the data and an exponential whose argument is linear in a statistic and a parameter. The statistic $T(x)$ summarises the data, the natural parameter $\eta(\theta)$ multiplies it, and the log-partition function $A(\theta)$ is the constant that makes the integral over the support equal one. Any family with this shape shares the same machinery.

## Worked: reading a Bernoulli

A Bernoulli variable with success probability $p$ has mass $p^x (1-p)^{1-x}$ on the values zero and one. Factoring it gives $(1-p)\exp(x \log(p/(1-p)))$, so the sufficient statistic is $x$, the natural parameter is the log-odds, and the log-partition function is $-\log(1-p)$. Reading those three pieces off a familiar mass function is the standard exercise.

## Members and non-members

The Poisson, binomial, exponential, normal, gamma and beta families all belong, because each can be put into the canonical form. The Cauchy does not, because its logarithmic term does not separate into a statistic times a parameter. A uniform whose support depends on the parameter, or a mixture of two distributions, generally does not belong either, and the boundary is exactly where the factorisation fails.
