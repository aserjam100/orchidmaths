# Derivatives of Hyperbolic Functions

**Hyperbolic functions:** $\sinh x = (e^x - e^{-x})/2$ and $\cosh x = (e^x + e^{-x})/2$, with $\tanh x = \sinh x/\cosh x$. Their derivatives follow from the exponentials and come in pairs without the sign changes of the circular functions.

## Defined from exponentials

Differentiating the definitions gives $(\sinh x)' = \cosh x$ and $(\cosh x)' = \sinh x$. The minus sign in $\sinh$ and the plus sign in $\cosh$ cancel the sign produced by differentiating $e^{-x}$, which is why the pair closes on itself. The identity $\cosh^2 x - \sinh^2 x = 1$ plays the role that $\sin^2 + \cos^2 = 1$ plays for the circular functions.

## Worked: the three basic derivatives

1. $(\sinh x)' = \cosh x$.
2. $(\cosh x)' = \sinh x$.
3. $(\tanh x)' = 1/\cosh^2 x$, since the quotient rule and the identity above leave $\cosh^2 - \sinh^2$ over $\cosh^2$.

## Chain rule with hyperbolic functions

The chain rule applies exactly as it does to the circular functions. For $f(x) = \sinh(3x)$, the derivative is $3\cosh(3x)$; for $f(x) = \cosh(x^2)$, it is $2x\sinh(x^2)$. For $\tanh(x/2)$, the outer derivative is $1/\cosh^2(x/2)$ and the inner derivative is $1/2$, so the result is $(1/2)/\cosh^2(x/2)$.
