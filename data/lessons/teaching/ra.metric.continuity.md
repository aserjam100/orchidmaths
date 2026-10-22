# Continuity in Metric Spaces

**Continuity:** a map between metric spaces is continuous when it preserves the way points approach each other, equivalently when the preimage of every open set is open.

## The epsilon-delta and open-set forms

A map between metric spaces is continuous at a point when, for every positive epsilon, there is a positive delta such that any point within delta of the given point is mapped within epsilon of the image. Continuity on the whole space means continuity at every point. The same condition is equivalent to the preimage of every open set being open, which is the form that carries over to general topological spaces and that makes continuity a purely topological notion, independent of the particular distance functions chosen.

## Equivalences and composition

A map is continuous exactly when it carries convergent sequences to convergent sequences with the expected limits, and exactly when the preimage of every closed set is closed. Compositions of continuous maps are continuous, and the identity is continuous, so continuity behaves as a structure-preserving notion. Uniform continuity, which fixes one delta for the whole space, is stronger; continuity plus compactness of the domain gives it, which is the Heine-Cantor theorem in metric form.

## Homeomorphisms

A continuous bijection whose inverse is also continuous is a homeomorphism, and it identifies the two spaces as topologically the same: every property defined from open sets, such as compactness and connectedness, is preserved in both directions. Continuity alone, without continuity of the inverse, does not suffice, because a continuous bijection can still collapse a discrete structure. Homeomorphism is the correct notion of sameness for metric spaces when the distances themselves are not part of the question.
