# Bayes' Theorem with Priors

**Bayes' theorem:** For a hypothesis $H$ and evidence $E$ with positive probability, $P(H \mid E) = P(E \mid H)P(H)/P(E)$. The factor $P(H)$ is the prior and the result is the posterior.

## Prior, likelihood, posterior

The prior encodes what is believed before the evidence. The likelihood $P(E \mid H)$ says how well the hypothesis explains the evidence. Bayes' theorem combines them into the posterior, the revised belief. The denominator $P(E)$ is the total probability of the evidence, computed by summing the likelihood times the prior over the possible hypotheses.

## Worked: two boxes

Two boxes are equally likely. Box one holds three red marbles out of four, box two holds one out of four. A red marble is drawn. The two equal priors cancel, so the posterior for box one is $(3/4)/((3/4) + (1/4)) = 3/4$. The evidence has shifted the belief from one half to three quarters.

## The law of total probability

The denominator of Bayes' theorem is the law of total probability: if the hypotheses partition the space, then $P(E)$ is the sum of $P(E \mid H_i)P(H_i)$ over the partition. This identity turns a hard unconditional probability into a weighted average of easy conditional ones, and it is the reason the posterior over a partition automatically sums to one.
