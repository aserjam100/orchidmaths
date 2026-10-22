# The Law of Total Probability

**Law of total probability:** If the events $B_1, \dots, B_n$ partition the sample space, then $P(A) = \sum P(A \mid B_i) P(B_i)$.

## Partitioning the sample space

A partition is a family of disjoint events whose union is the whole space, so every outcome lies in exactly one part. The event $A$ is then the disjoint union of its pieces $A \cap B_i$, and countable additivity makes the pieces add. Writing each piece as $P(A \mid B_i)P(B_i)$ turns that sum into a weighted average of the conditional probabilities, with the prior probabilities as the weights.

## Worked: two boxes

Two boxes are equally likely. Box one yields the event with probability three quarters and box two with probability one quarter. The two priors are one half each, so the total probability is $(1/2)(3/4) + (1/2)(1/4) = 0.5$. The answer is exactly between the two conditional chances because the boxes are equally likely, and unequal priors would pull it toward the more likely box.

## The denominator of Bayes

The quantity $\sum P(E \mid H_i)P(H_i)$ is the denominator of Bayes' theorem, so the law of total probability is what makes the posterior over a partition sum to one. Computing the unconditional probability of the evidence first, then dividing each weighted likelihood by it, is the standard two-step route from priors to posteriors.
