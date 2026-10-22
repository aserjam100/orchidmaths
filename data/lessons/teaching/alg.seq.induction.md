# Proof by Induction

**Mathematical induction:** a proof technique that establishes a statement for all positive integers by proving a base case and a step from each integer to the next.

## The two parts

An induction proof has a base case, which shows the statement holds for the first integer, and an induction step, which shows that if the statement holds for some integer then it holds for the next. Together they form an unbroken chain: the base case starts it and the step carries it forward, so every integer at or above the base is covered. Both parts are required, and either one alone proves nothing.

## Structuring the step

The induction step begins with an assumption, called the induction hypothesis, that the statement is true for a particular n. The work is then to show the statement for n plus one using that assumption, usually by writing the n-plus-one case in a form that contains the n case and replacing it with the assumed value. Keeping the assumption visible while manipulating the larger case is the skill the technique teaches.

## Beyond the simple form

Variants relax the two parts without changing the idea. Strong induction assumes the statement for every integer up to n rather than only n, which is needed when the next case depends on several earlier ones. Starting the base case later allows a statement to hold from some integer onwards rather than from one, and it is exactly what a statement false at the start requires.
