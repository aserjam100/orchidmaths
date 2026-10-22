# Heaviside Steps and Impulses

**Unit step:** $H(t-a)$ is 0 for $t < a$ and 1 for $t \ge a$. Its distributional derivative is the Dirac delta $\delta(t-a)$, a spike at $a$ with total integral 1. Both model inputs that switch on or strike at a definite time.

## Switches and spikes

The step models a voltage turned on, a force applied from time $a$, or a valve opened. The delta models an instantaneous blow, such as a hammer strike or a very short pulse. Ordinary functions cannot represent the delta, so the Laplace transform handles it as a distribution.

## Worked: transforms of the step and delta

1. $\mathcal{L}\{H(t-a)\} = e^{-as}/s$, by the time-shift rule.
2. $\mathcal{L}\{\delta(t-a)\} = e^{-as}$, because the delta samples the kernel at $t = a$.
3. The delta at 0 transforms to 1, the constant function in the $s$-domain.

## Why distributions are needed

The delta is not a function in the ordinary sense: it is zero everywhere except at one point, yet its integral is 1. Treating it as a limit of narrow tall pulses makes the formulas correct, and the Laplace transform turns convolution with a delta into multiplication by $e^{-as}$. This is why the impulse response and the transfer function are the same object.
