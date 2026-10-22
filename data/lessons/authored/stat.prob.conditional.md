# Conditional probability

## Restricting to a condition

Conditional probability is the probability of one event given that another has happened. The condition restricts attention to the part of the sample space where it holds, so every probability is recomputed within that smaller space.

Writing the condition as a division by its probability is what performs the restriction: the outcomes where the condition fails are divided out of consideration.

## The ratio of two probabilities

The conditional probability of A given B is the probability of both divided by the probability of the condition:

\[
P(A \mid B) = \frac{P(A \text{ and } B)}{P(B)}.
\]

The numerator is the part of the condition's space where A also holds, and the denominator rescales so the condition's space becomes the whole space.

## Why the order matters

Conditioning on one event is not the same as conditioning on the other: the chance of rain given that the ground is wet is not the chance the ground is wet given that it rains. The two are related by Bayes' theorem, and confusing them is the most common error in reasoning about evidence.

The conditional probability is therefore not symmetric, which is why it is written with the condition after the bar rather than as an unordered pair.
