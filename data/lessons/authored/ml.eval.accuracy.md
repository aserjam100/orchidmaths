# Accuracy

## The fraction correct

Accuracy is the fraction of predictions the model got right:

\[
\frac{\text{true positives} + \text{true negatives}}{\text{all cases}}.
\]

With 40 true positives, 50 true negatives and 10 errors of each kind, the accuracy is 90 over 110, which is about 82 percent.

It is the first metric to compute and rarely the one to decide on.

## What accuracy hides

Accuracy treats the two kinds of error as equally costly, which is usually false. A test for a rare disease that always says no is very accurate and completely useless, because the errors it makes are all in the class that matters.

The same arithmetic problem appears whenever the classes are unbalanced: the majority class dominates the count, so a model that ignores the minority can still score well.

## When accuracy is the wrong metric

Accuracy is the right summary when the classes are balanced and the two errors cost about the same, which is the case in many well-posed classification problems and almost none of the interesting ones.

When they are not, precision and recall separate the two error types and let the trade be chosen deliberately rather than by default.
