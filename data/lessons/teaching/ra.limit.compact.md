# Compactness and Heine-Borel

**Compactness:** a set is compact when every collection of open sets that covers it contains a finite subcollection that already covers it; on the real line this is the same as being closed and bounded.

## Open covers

An open cover of a set is a collection of open sets whose union contains the set. The set is compact when every such cover, however large or awkward, can be trimmed to finitely many of its sets that still cover the set. This finite-subcover property is the abstract form of the closed bounded condition and it is what makes compact sets behave like finite sets: a continuous function on a compact set is bounded and attains its bounds, and sequences in a compact set have convergent subsequences.

## Heine-Borel

On the real line the abstract definition collapses to something concrete. The Heine-Borel theorem says a subset of the reals is compact exactly when it is closed and bounded. The closed bounded intervals are the model compact sets, and this equivalence is why the extreme value theorem asks for a closed bounded interval: that is precisely the hypothesis of compactness in disguise. In higher-dimensional spaces the same equivalence holds, though not in every metric space.

## Sequential compactness

Compactness has a sequential face: a set is sequentially compact when every sequence drawn from it has a subsequence converging to a point still inside the set. For subsets of the real line, compact, sequentially compact, and closed and bounded all describe the same sets, and the Bolzano-Weierstrass theorem is the statement that closed bounded sets are sequentially compact. The limit must stay inside the set, which is what closedness supplies, and the existence of a convergent subsequence is what boundedness supplies.
