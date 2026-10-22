# Closure and Interior

**Closure and interior:** the closure adds every limit point, the interior removes every boundary point, and the boundary is what lies between them.

## The closure

The closure of a set is the set together with all of its limit points, and it is the smallest closed set containing the original. A point lies in the closure exactly when every ball around it meets the set, and equivalently when some sequence drawn from the set converges to it. The closure of an open interval adds its two endpoints, and the closure of the rationals is the whole line, which is the density of the rationals restated. The closure operation is idempotent, since closing a closed set adds nothing.

## The interior

The interior of a set is the largest open set contained in it, consisting of the points that have some ball entirely inside the set. It is the complement, taken in the closure, of the closure of the complement, so the interior and closure are dual operations. A point of a set is interior when every nearby point is still in the set, and the interior can be empty, as it is for the rationals, which contain no open interval at all. The set is open exactly when it equals its interior.

## The boundary

The boundary of a set is the set of points every ball around which meets both the set and its complement. It equals the closure minus the interior, and it is always a closed set. The boundary of an open interval is its two endpoints, the boundary of a disc is its circle, and the boundary of the rationals is the whole line. A set is both open and closed, such as the empty set, exactly when its boundary is empty, and a set is closed precisely when it contains its boundary.
