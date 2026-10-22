# Characteristic Functions

**Characteristic function:** the function $\varphi(t) = E[e^{itX}]$ of the real parameter $t$, the Fourier transform of the distribution of $X$.

## Definition and existence

The modulus of the exponential is one for every real $t$, so the expectation is finite for every distribution, discrete or continuous, light-tailed or heavy-tailed. This universal existence is the great advantage over the moment generating function, which may fail to be finite. The function is bounded by one and uniformly continuous.

## Reading off familiar laws

A Bernoulli variable taking zero and one with probability one half each has transform one half plus one half times the exponential. A normal variable with mean zero and variance one has transform the exponential of minus half the square. Recognising these transforms is how the limit theorems identify their limit laws.

## Why they are used

Independence turns the transform of a sum into a product, so a sum of many independent terms has a transform that is a high power of a single one. Taking logarithms turns that power into a simple factor of the index. The central limit theorem and the continuity theorem both exploit exactly this simplification.
