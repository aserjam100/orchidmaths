# Precision

## Of the positives predicted

Precision is the fraction of the cases the model called positive that really are positive:

\[
\frac{\text{true positives}}{\text{true positives} + \text{false positives}}.
\]

With 60 true positives and 20 false positives it is 60 over 80, which is 75 percent.

## The cost it measures

Precision answers how much a positive prediction can be trusted. It is the metric to report when a false positive is expensive: a spam filter that discards real mail, a fraud flag that freezes a legitimate account, a medical test that leads to an unnecessary procedure.

A model can raise precision by predicting positive only when it is very sure, which is exactly what lowers recall.

## Trading against recall

Precision and recall move in opposite directions as the decision threshold moves, and the choice between them is a statement about which error costs more rather than a property of the model.

The trade is summarised by the F1 score, the harmonic mean of the two, which is high only when both are high. Quoting one without the other is the most common way a classification result is made to look better than it is.
