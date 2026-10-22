# Nested quantifiers and their negations

A quantified statement with two quantifiers has an order, and the order changes the meaning. \(\forall x\, \exists y\, (x < y)\) says every \(x\) has some larger \(y\); \(\exists y\, \forall x\, (x < y)\) says one \(y\) is larger than every \(x\). Over the integers the first is true and the second is false, so they are not interchangeable.

## Quantifier order

Read the quantifiers left to right, and let each depend only on the variables already fixed. In \(\forall x\, \exists y\), the \(y\) may be chosen after seeing \(x\), which is why it can be a function of \(x\). In \(\exists y\, \forall x\), the \(y\) is chosen first and must work for every \(x\) afterwards.

The implication is one-directional: \(\exists y\, \forall x\, P(x,y)\) implies \(\forall x\, \exists y\, P(x,y)\), and the converse fails. Statements in the "there is a uniform choice" form are the stronger ones.

## Negating a quantified statement

Negation flips each quantifier and negates the inner statement: \(\lnot \forall x\, P(x)\) is \(\exists x\, \lnot P(x)\), and \(\lnot \exists x\, P(x)\) is \(\forall x\, \lnot P(x)\). For nested quantifiers the rule applies left to right, and the relation itself is negated only at the end.

So the negation of \(\forall x\, \exists y\, (x < y)\) is \(\exists x\, \forall y\, (x \ge y)\): some \(x\) is at least as large as every \(y\). Note that the negation of an inequality is not always a reversal — \(\ge\) is the negation of \(<\), and \(>\) is not.

## What the order costs

The two orders differ exactly in whether a witness may depend on the previous variable. A proof of \(\forall x\, \exists y\) is allowed to construct \(y\) from \(x\), so a formula such as \(y = x+1\) is a valid witness. A proof of \(\exists y\, \forall x\) must exhibit one value before any \(x\) appears, which is why such statements are rare outside finite settings and why their negations are easy to prove when they fail.
