# Prediction from a fitted model

## Substituting into the model

Once a model has been fitted, prediction is arithmetic: put the new input into the formula and evaluate. For the line \(\hat{y} = 3x + 2\), the prediction at an input of 4 is 14.

That is the whole computation, and it is the step worth checking first when a model's predictions look wrong, because the fitting is the part that has no obvious error to spot.

## What a prediction is not

A prediction is an estimate of the average outcome at that input, not a promise about any single case. The fitted line says what the model expects, and the spread of the data around it is the uncertainty that the prediction does not contain.

Reporting a prediction without that spread is the most common way a fitted model is over-trusted.

## Extrapolation

A model fitted on one range of inputs says nothing reliable outside it. A line through data from those inputs will happily extend past them, and the further it extends, the less the answer means, because nothing in the fitting constrained the behaviour there.

The check is to compare the new input with the range of the training data before using the prediction, and to say so when it falls outside.
