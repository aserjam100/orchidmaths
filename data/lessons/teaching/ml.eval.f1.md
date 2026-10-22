# The F1 Score

**F1 score:** $F_1 = 2PR/(P+R)$, the harmonic mean of precision $P$ and recall $R$. It is high only when both are high, so it summarises a classifier that must catch positives without flooding the positive predictions with false ones.

## The harmonic mean of precision and recall

The ordinary mean would let a large precision compensate for a tiny recall: precision 1 with recall 0 has mean 0.5. The harmonic mean punishes imbalance, giving $2 \cdot 1 \cdot 0/(1+0) = 0$. When $P$ and $R$ are equal, F1 equals their common value.

## Worked: two exact F1 values

1. $P = 0.6$ and $R = 0.6$: $F_1 = 2(0.36)/1.2 = 0.6$.
2. $P = 1$ and $R = 0.6$: $F_1 = 2(0.6)/1.6 = 0.75$.
3. $P = 0.9$ and $R = 0.9$: $F_1 = 0.9$.

The second case shows the price of perfect precision on six tenths recall: F1 lands between the two rather than near 1.

## When F1 is the wrong summary

F1 weights precision and recall equally, which is a choice about costs rather than a property of the classifier. If a false negative is far worse than a false positive, recall should carry more weight, and the F-beta score does that. On balanced data where both errors cost the same, accuracy is easier to interpret.
