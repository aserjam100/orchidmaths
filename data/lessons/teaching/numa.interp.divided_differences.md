# Newton's Divided Differences

**Divided differences:** the successive slopes that build the Newton form of the interpolant, which can be extended to a new node without rebuilding the earlier terms.

## Successive slopes

The first divided difference of two points is their slope. Higher divided differences divide the difference of two lower ones by the distance between the outer nodes. They generalise the idea of a derivative to data that are spread out, and the leading divided difference is the leading coefficient of the interpolant.

## Worked: a first difference

For the points one, two and four, eleven, the first divided difference is eleven minus two divided by four minus one, which is three. That slope is the coefficient of the first power in the Newton form of the interpolant. Further points add further terms without changing the ones already computed.

## The Newton form

The interpolant is the first value plus the first difference times the first distance plus the second difference times two distances, and so on. Because each term only involves the new node, a new point appends one term and leaves the rest untouched. This is the practical advantage of divided differences over the Lagrange construction.
