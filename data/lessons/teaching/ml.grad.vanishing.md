# Vanishing and Exploding Gradients

**The problem:** Backpropagation multiplies one local derivative per layer. If those factors are consistently below 1 the product vanishes; if they are consistently above 1 it explodes, and either way the early layers learn badly.

## A product of many derivatives

For a stack of layers the gradient of the loss with respect to an early weight carries a factor from every layer after it. With factors of 0.25, ten layers give $0.25^{10}$, about one millionth; with factors of 2, ten layers give 1024. Depth turns small per-layer effects into exponential ones.

## Worked: a stack of sigmoids

1. Each sigmoid has derivative at most 0.25.
2. A chain of five such layers contributes at most $0.25^5 \approx 0.001$ to the gradient.
3. The earliest layers therefore receive almost no update while the last layers keep learning.

## Remedies: activations, initialization, clipping

ReLU keeps a derivative of 1 on its positive side and 0 elsewhere, so it does not shrink the product the way the sigmoid does. Residual connections give the gradient a short path back to early layers, careful initialization keeps the per-layer scale near 1, normalization stabilises it, and gradient clipping caps the norm when the product explodes.
