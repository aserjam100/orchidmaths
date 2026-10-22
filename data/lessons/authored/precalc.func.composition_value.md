# Composition of functions

## Inside out

The composition \((f \circ g)(x)\) means \(f(g(x))\): apply \(g\) first and then \(f\). The order of the letters is the reverse of the order of the operations, which is the single most common mistake in reading a composition.

For \(f(x) = 2x + 1\) and \(g(x) = 3x - 4\), the composition at \(x = 2\) is \(g(2) = 2\) and then \(f(2) = 5\), so \((f \circ g)(2) = 5\).

## Order matters

Composition is not commutative in general: \((f \circ g)\) and \((g \circ f)\) are different functions, and either may exist while the other does not, since the range of the inner function has to lie inside the domain of the outer one.

That condition on the domain is what makes composition harder than multiplication: two functions can be multiplied always, but not always composed.

## Composition and inverses

Composing a function with its inverse gives the identity, in either order: \(f(f^{-1}(x)) = x\) and \(f^{-1}(f(x)) = x\). That is the definition of an inverse stated as a composition, and it is how an inverse is checked.

A composition of two one-to-one functions is one-to-one, and its inverse is the composition of the inverses in the opposite order, which is the rule for undoing a chain of operations written in the language of functions.
