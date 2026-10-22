# Limit Superior and Limit Inferior

**Limit superior:** the least number that the suprema of the tails of a sequence approach, with the limit inferior the matching infimum; the two agree exactly when the sequence converges.

## The eventual envelope

The limit superior of a bounded sequence is the infimum, over all starting points, of the supremum of the tail from that point onward. As the starting point moves to the right the tail loses terms and its supremum can only fall or stay level, so these suprema form a decreasing sequence bounded below, and its infimum is the limit superior. The limit inferior is the same construction with supremum and infimum swapped. Together they describe the eventual envelope of the sequence: the highest and lowest levels the terms keep returning near.

## Convergence and the two limits

A bounded sequence converges exactly when its limit superior and limit inferior are equal. When the two numbers are the same, the whole tail is squeezed between values that are equal in the limit, so the sequence has that common value as a limit. When the two differ, the sequence keeps visiting two separated levels and cannot converge, as the alternating sequence between minus one and one does, with limit superior one and limit inferior minus one. The pair therefore detects oscillation that a single limit cannot.

## Uses and comparison

Limit superior and limit inferior are the right tool when a limit may fail to exist but a decision is still needed. The ratio test and the root test for series are stated with limit superior precisely because the ratios need not converge. The oscillation of a function, and hence continuity, can be expressed through the same shrinking-envelope idea. And a sequence has a convergent subsequence exactly when its limit inferior is at most its limit superior and both are finite, which is the Bolzano-Weierstrass theorem in this language.
