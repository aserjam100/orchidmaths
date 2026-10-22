# Characteristic Functions of Sums

**Sums:** the characteristic function of a sum of independent variables is the product of their characteristic functions, so convolution of distributions becomes multiplication of transforms.

## The product rule

For independent variables the expectation of the exponential of the sum factors into the product of the two expectations, because the exponential of a sum is the product of the exponentials and independence factors expectations. The transform of the sum is therefore the product of the two transforms. This is the single most useful property of the transform.

## Convolution becomes multiplication

The distribution of a sum of independent variables is the convolution of their distributions. Under the transform the convolution becomes an ordinary product, so hard convolution integrals are replaced by algebra. This is why the transform is the standard tool for sums, and why it is a Fourier transform in disguise.

## Without independence

When the variables are dependent the product rule fails and the joint transform is needed, evaluated along the diagonal where the two parameters agree. Factorisation of the joint transform is in fact equivalent to independence, which is a stronger statement than the vanishing of the covariance. Independence is exactly the hypothesis that makes the product rule available.
