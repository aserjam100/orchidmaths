# The Confusion Matrix

**Confusion matrix:** For a binary classifier the four counts are true positives TP, false positives FP, false negatives FN and true negatives TN. Every common classification metric is an arithmetic combination of these four cells.

## Four counts behind every metric

The rows of the matrix are the actual classes and the columns are the predicted classes. A case is on the diagonal when the prediction is correct. Accuracy is $(TP+TN)/\text{total}$, precision is $TP/(TP+FP)$, recall is $TP/(TP+FN)$, and the false positive rate is $FP/(FP+TN)$.

## Worked: reading the cells

Take $TP = 8$, $FP = 2$, $FN = 3$ and $TN = 7$.

1. Actual positives are $TP + FN = 11$, and actual negatives are $FP + TN = 9$.
2. Predicted positives are $TP + FP = 10$, and the total is $8 + 2 + 3 + 7 = 20$.
3. Correct predictions are $TP + TN = 15$, so accuracy is $15/20 = 0.75$.

## From cells to precision and recall

Precision uses the predicted-positive column, $8/(8+2) = 0.8$, while recall uses the actual-positive row, $8/(8+3) \approx 0.73$. Both are read from the same table, which is why quoting one without the other hides the trade the classifier made.
