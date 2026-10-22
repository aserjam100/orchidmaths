# Pairwise and Mutual Independence

**Independence:** Events $A$ and $B$ are independent when $P(A \cap B) = P(A)P(B)$. A family is mutually independent when every subcollection satisfies the product rule.

## The product rule

Independence is a multiplicative condition, not the absence of a relationship in plain words. It says that learning one event occurred does not change the probability of the other. Two events with positive probability are independent exactly when $P(B \mid A) = P(B)$, which is the same equation written after dividing by $P(A)$.

## Worked: pairwise is not mutual

Throw two fair coins. Let $A$ be the first coin heads, $B$ the second heads, and $C$ the event that the two coins match. Each pair is independent, but the three are not mutually independent: once $A$ and $B$ are known, $C$ is determined, so $P(A \cap B \cap C) = 1/4$ while $P(A)P(B)P(C) = 1/8$. Pairwise independence is strictly weaker than mutual independence.

## Independence and disjointness

Two disjoint events with positive probability are dependent, because one occurring rules out the other. Independence and disjointness are easy to confuse and mean almost opposite things: disjoint events cannot both happen, while independent events carry no information about each other. A model that silently assumes one where it meant the other is wrong about the shape of its data.
