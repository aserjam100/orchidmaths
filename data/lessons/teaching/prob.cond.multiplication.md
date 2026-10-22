# The Multiplication Rule

**Multiplication rule:** For events with $P(A)$ positive, $P(A \cap B) = P(A) \cdot P(B \mid A)$.

## Conditioning step by step

The conditional probability $P(B \mid A)$ rescales the probability of $B$ by restricting attention to the outcomes in $A$. Multiplying back by $P(A)$ restores the correct scale and returns the joint probability. This is the definition of conditional probability read in reverse, and it is the first step of every chain-rule calculation on a tree of events.

## Worked: drawing without replacement

A bag has twenty marbles, ten of them red. Of the red marbles, six are large. Drawing one marble at random, the chance it is red is $10/20$, and given that it is red the chance it is also large is $6/10$. The multiplication rule gives $P(\text{red and large}) = (10/20)(6/10) = 6/20 = 0.3$, which is simply the six large red marbles out of twenty.

## The chain rule

For a list of events the rule extends to $P(A \cap B \cap C) = P(A)P(B \mid A)P(C \mid A \cap B)$. Each factor conditions on everything before it, so the product telescopes to the probability of the intersection. The chain rule is what lets a model factor a complicated joint distribution into a sequence of simpler conditional pieces, one decision at a time.
