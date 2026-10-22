# Sigmoid Activation

**Sigmoid:** The sigmoid is $\sigma(x) = 1/(1+e^{-x})$. It maps the whole real line into the interval $(0, 1)$, is increasing, and equals 0.5 at zero. In a network it turns a linear score into a probability.

## A smooth step from 0 to 1

Unlike the perceptron's step, the sigmoid is differentiable everywhere, so a gradient can pass through it. Its output is never exactly 0 or 1, which keeps logarithms of predictions finite, and it is symmetric about the point $(0, 0.5)$: $\sigma(-x) = 1 - \sigma(x)$.

## Worked: sigma of exact inputs

1. $\sigma(0) = 1/(1+e^0) = 1/2 = 0.5$.
2. $\sigma(\ln 3) = 1/(1+e^{-\ln 3}) = 1/(1 + 1/3) = 3/4 = 0.75$.
3. $\sigma(-\ln 3) = 1/(1 + 3) = 1/4 = 0.25$.

The pair $\ln 3$ and $-\ln 3$ gives probabilities that sum to 1, which is the symmetry above in numbers.

## Why the derivative is small

Differentiating gives $\sigma'(x) = \sigma(x)(1 - \sigma(x))$, whose maximum is $1/4$ at $x = 0$. A chain of sigmoid layers therefore multiplies factors no larger than $1/4$, which is the mechanism behind vanishing gradients in deep networks.
