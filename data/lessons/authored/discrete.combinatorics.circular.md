# Circular arrangements

## Why rotations are the same

Arranging \(n\) people around a round table counts fewer orders than arranging them in a row, because a rotation of the whole table is not a new seating. The row has \(n!\) orders, and each seating at the table is written by exactly \(n\) of them, one for each person who could be counted as first.

## The (n-1)! rule

Dividing the \(n!\) row orders by the \(n\) rotations gives

\[(n-1)!\]

seatings around the table. Four people sit in \(3! = 6\) distinct ways, not \(4! = 24\).

The same count answers a different question about a loop: the number of distinct necklaces on \(n\) beads, all different, is \((n-1)!\), because a necklace is a circle of beads and turning the necklace is not a change.

## Reversals and reflections

If a seating and its mirror image are also the same, divide by 2 as well, which is allowed only when \(n \ge 3\), because with two people the mirror of a seating is a rotation of it and dividing by 2 would count the single seating as a half.

Six people can be seated around a table in \(5! = 120\) ways, and in \(120/2 = 60\) ways once reflections are counted as the same seating.
