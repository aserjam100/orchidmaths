# Proof by contrapositive

The **contrapositive** of \(p \to q\) is \(\lnot q \to \lnot p\). The two are logically equivalent — they have the same truth table — so proving one proves the other. A proof by contrapositive assumes \(\lnot q\) and derives \(\lnot p\), and it is the right choice when the negated conclusion is easier to work with than the hypothesis.

## The contrapositive

For the statement "if \(n\) is even then \(n^2\) is even", the contrapositive is "if \(n^2\) is odd then \(n\) is odd". Neither is easier there. For "if \(n^2\) is even then \(n\) is even" the direct proof has nothing to grab — an even square says nothing directly about \(n\) — while the contrapositive assumes \(n\) odd, writes \(n = 2k+1\), and computes \(n^2 = 4k^2+4k+1\), which is odd. That is the standard use: the contrapositive turns a claim about a square back into a claim about its root.

## Converse and inverse

Two nearby statements are **not** equivalent to the original. The **converse** is \(q \to p\), and the **inverse** is \(\lnot p \to \lnot q\). They are equivalent to each other (each is the contrapositive of the other) but not to the original: "if it is a square then it is a rectangle" is true, while its converse "if it is a rectangle then it is a square" is false.

Keeping the three apart is the practical skill. A proof of the converse proves something else entirely, and a common error is to prove the converse while claiming the original.

## When to prove it instead

Choose the contrapositive when the hypothesis is a negative or an absence ("\(x\) is not a perfect square") and the conclusion gives you something concrete to compute with. Choose contradiction only when you need to derive a statement and its negation from the assumption, as in the irrationality of \(\sqrt{2}\); the contrapositive is the sharper tool whenever it applies, because it derives a specific claim rather than an impossibility.
