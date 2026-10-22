# Class Imbalance

**Class imbalance:** The classes occur at very different frequencies, so a model can score well by always predicting the majority. Accuracy hides this; precision, recall and the confusion matrix expose it.

## Why accuracy hides the minority

With 990 negatives and 10 positives, always predicting negative is 99 percent accurate and catches none of the positives. The error that matters is concentrated in the 10 cases the average ignores, so a single number cannot describe the classifier.

## Worked: a 99 percent majority classifier

1. The always-negative model is right on 990 cases and wrong on 10.
2. Accuracy is $990/1000 = 0.99$.
3. Recall on the positive class is $0/10 = 0$, and precision is undefined because the model never predicts positive.

## Resampling and class weights

Oversampling the minority, undersampling the majority, or weighting the loss by inverse class frequency all change what the model optimises. Resampling changes the data and can overfit duplicated minority points; class weights keep the data and change the cost. Either way, evaluate on the original distribution, not the resampled one.
