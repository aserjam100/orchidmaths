# The Bolzano-Weierstrass Theorem

**Bolzano-Weierstrass theorem:** every bounded sequence of real numbers has a convergent subsequence.

## Statement and sharpness

Every bounded real sequence has a convergent subsequence. The conclusion is a subsequential limit, not a limit of the whole sequence, so the theorem applies to wildly oscillating bounded sequences as well. Boundedness is essential: an unbounded sequence can escape every bounded region, and no subsequence of such an escape need converge. The theorem is one of the equivalent faces of completeness and is often taken as the bridge between the order picture and the topological picture of the real line.

## Proof by bisection

If some value occurs infinitely often, the constant subsequence at that value converges and there is nothing to prove. Otherwise bisect a bounded interval that contains infinitely many terms and keep a half that still contains infinitely many. Repeating gives nested closed bounded intervals whose lengths halve each time. By the nested interval property they have a common point, and choosing one term from each interval produces a subsequence converging to it. The proof is the bisection method of the nested interval lesson in action.

## The metric form

In a metric space the analogue is that every sequence in a compact space has a convergent subsequence, and a space with that property is called sequentially compact. For subsets of the real line the compact sets are exactly the closed bounded ones, so Bolzano-Weierstrass is the statement that a closed bounded set is sequentially compact. Recognising the theorem this way lets it be used in any metric space, not only on the line, and it is the form used to prove compactness criteria later in the subject.
