# Open and closed sets in the plane

A set is **open** when every point of it has a small disk around it that stays inside the set, and **closed** when its complement is open. The boundary is the set of points that every disk around them meets both inside and outside.

## Disks and interior points

The disk \(\{z : |z| < r\}\) is open: around any point of it there is room up to the boundary, so a smaller disk fits. The point on the boundary is not in the set, which is what makes the whole interior safe.

The set \(\{z : |z| \le r\}\) is not open, because a point with \(|z| = r\) has no disk that stays inside: every disk around it leaves through the circle.

## Boundaries and closure

The boundary of both sets is the circle \(|z| = r\). The closure of the open disk is the closed disk, and the interior of the closed disk is the open one, so the boundary is exactly the difference.

A half-plane such as \(\{z : \operatorname{Re} z > 0\}\) is open, and its boundary is the imaginary axis. Included (closed) and excluded (open) boundaries are the whole distinction.

## Classifying a set

Ask two questions: does the set contain its boundary, and does its complement? A set that contains no boundary point is open, one that contains all of them is closed, and one that contains some is neither.

"Both open and closed" is possible, but only for the empty set and the whole plane in this setting — not for a disk or a half-plane, which is why the option is a distractor rather than a classification.
