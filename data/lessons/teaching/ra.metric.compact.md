# Compactness

**Compact:** a metric space in which every open cover has a finite subcover; compactness upgrades local information to global, and on the line it means closed and bounded.

## The finite subcover property

A collection of open sets covers a space when every point lies in at least one of them. The space is compact when every such cover, however large or irregular, contains finitely many members that already cover the whole space. The definition is abstract but exactly right: it is the property that lets a local statement at each point be combined into a global one, because the finitely many members of the subcover can be handled one at a time. The empty set is compact, and the closed bounded intervals of the line are the model examples.

## Compact and sequentially compact

In a metric space, compactness is equivalent to sequential compactness: every sequence has a subsequence converging to a point of the space. The proof that compactness gives a subsequence uses the finite subcover to show a sequence cannot keep a positive distance from every point, and the converse assembles a finite subcover from the failure of sequential compactness. This equivalence is what makes compactness usable, because a convergent subsequence is usually easier to produce than a finite subcover, by the Bolzano-Weierstrass theorem on the line.

## Continuity and closed bounded

A continuous real-valued function on a compact space is bounded and attains its maximum and minimum, because the continuous image of a compact space is compact and a compact subset of the line is closed and bounded. On the line the Heine-Borel theorem identifies the compact sets as exactly the closed bounded ones, so the extreme value theorem of the calculus is the statement that a closed bounded interval is compact. In a general metric space compact sets are always closed and bounded, but the converse can fail without the extra structure of Euclidean space.
