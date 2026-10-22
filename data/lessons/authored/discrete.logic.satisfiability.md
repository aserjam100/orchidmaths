# Tautologies, contradictions and contingencies

Every formula in \(n\) variables has a truth table with \(2^n\) rows, and the formula falls into exactly one of three classes depending on which rows satisfy it: **tautology** if all of them do, **contradiction** if none does, and **contingent** if some do and some do not.

## The three classes

The classes are exhaustive and exclusive: a formula cannot be both satisfied by every row and by none, so classifying it is a matter of counting the satisfied rows. The empty distinction is worth stating: a contradiction is not the same as a false statement under one assignment, and a tautology is not simply true under the assignment a learner happened to try.

Two formulas are logically equivalent exactly when their biconditional is a tautology, which is why classification is the tool behind equivalence.

## Deciding by truth table

Write a column for each subformula, then fill the rows. For \((p \land q) \to p\), the antecedent \(p \land q\) is true only in the row where both are true, and \(p\) is true there, so the implication is true in all four rows: a tautology.

For \(p \land \lnot p\), no row makes both conjuncts true, so the formula is false in all four rows: a contradiction. For \(p \to q\), the implication fails only in the row \(p\) true, \(q\) false, so the other three satisfy it: contingent.

## Why the classification matters

Classification turns several questions into one. Asking whether a formula is a tautology, whether its negation is a contradiction, and whether it is implied by the empty set of premises are the same question about the same column of the table.

It also separates two claims that sound alike: "true in this instance" is a statement about a row, and "true in every row" is the classification. A learner who checks one row has checked an instance, not the formula.
