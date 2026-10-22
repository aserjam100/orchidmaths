# Bracketing versus Open Methods

**Bracketing versus open methods:** bracketing methods keep a sign change and are guaranteed to converge, while open methods use one or two guesses and can converge faster but may diverge.

## The two families

A bracketing method needs two points where the function has opposite signs, which guarantees a root between them by continuity, and it never loses that guarantee. An open method needs only starting guesses and follows a formula, so it can move anywhere. The bracket is a promise; the open method is a gamble that usually pays.

## Worked: a failure of the open method

Newton's method fails when the derivative at a guess is zero, because the step becomes infinite. It can also jump far from the root when the tangent line is nearly horizontal. A bracketing method applied to the same problem would simply keep halving a valid interval and would still arrive, more slowly but surely.

## Choosing a method

In practice one often starts with a bracket to locate a root roughly and then switches to an open method to polish it quickly. The bracket supplies the guarantee and the starting guess; the open method supplies the speed. Hybrid methods formalise this by falling back to bisection whenever a Newton or secant step would leave the bracket.
