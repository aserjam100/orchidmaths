# Logistic Regression

**Logistic regression:** Model the probability of a binary label as $\sigma(w \cdot x + b)$, where $\sigma$ is the sigmoid. The linear score decides the direction and the sigmoid turns it into a probability, so the model is a linear classifier with a probabilistic output.

## A linear score turned into a probability

Training chooses $w$ and $b$ to make the observed labels likely, which is the same as minimising cross-entropy. The model is not a regression on the label itself: fitting a line to 0/1 targets with squared error gives predictions outside the unit interval, while the sigmoid keeps them inside it.

## Worked: classifying at the 0.5 threshold

1. A point has score $w \cdot x + b = 2$.
2. The sigmoid gives $\sigma(2) \approx 0.88$, above the threshold.
3. So the point is classified positive; a score of $-2$ gives about 0.12 and the negative class.

The threshold 0.5 is the score zero, because $\sigma(0) = 0.5$.

## Why the boundary stays linear

The predicted class flips exactly where $\sigma(w \cdot x + b) = 0.5$, which is where $w \cdot x + b = 0$: a hyperplane. The probability is nonlinear in the inputs, but the decision boundary is not, which is why logistic regression is called a linear classifier.
