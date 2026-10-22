# Oscillation

**Oscillation:** the spread between the supremum and the infimum of a function on a set; a function is continuous at a point exactly when its oscillation on shrinking neighbourhoods tends to zero.

## Measuring spread on a set

The oscillation of a function on a set is the supremum of the function there minus its infimum there, so it measures how far apart the values get rather than how large they are. A constant function has oscillation zero on every set. A function that climbs from below zero to above one has oscillation at least one. This is a single number attached to a set, and it turns the question of continuity, which is about values being close at nearby points, into a question about a number that shrinks.

## Continuity through oscillation

A function is continuous at a point exactly when its oscillation on the intervals shrinking to that point tends to zero. Continuity says the values near the point stay close to the value at the point, and the oscillation of such a neighbourhood is bounded by twice the allowed error, so it can be made as small as wanted. Conversely, if the oscillation tends to zero then the values are forced close to the value at the point. Discontinuity is therefore a positive oscillation that survives every shrinking, which makes the notion quantitative rather than merely negative.

## The oscillating example

The function that takes the sine of one over its input, with any value at zero, has oscillation two on every interval containing zero, because the sine visits values arbitrarily close to both one and minus one near zero however small the interval. So the oscillation does not tend to zero and the function is discontinuous there, whatever value is assigned at the point. This is the standard witness that a function can be bounded and yet fail continuity, and the oscillation is the number that says so. Monotone functions behave oppositely: their oscillation on a small interval is exactly the jump across it.
