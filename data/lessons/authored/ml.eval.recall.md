# Recall

## Of the positives that exist

Recall is the fraction of the actual positive cases that the model found:

\[
\frac{\text{true positives}}{\text{true positives} + \text{false negatives}}.
\]

With 40 found and 10 missed the recall is 40 over 50, which is 80 percent.

## The cost it measures

Recall answers how many of the things that mattered were caught. It is the metric to report when a false negative is expensive: a screening test that misses a case, a safety check that passes a fault, a search that returns nothing relevant.

A model can raise recall by predicting positive more often, which is exactly what lowers precision.

## The two together

Precision and recall are two views of the same confusion matrix, and neither alone describes a classifier. Reporting the pair, or the F1 score that combines them, is what keeps a result comparable between models and honest about the trade made.

The threshold that balances them is chosen from the costs of the two errors, which is a decision about the application rather than about the model, and it is worth stating explicitly when a result is reported.
