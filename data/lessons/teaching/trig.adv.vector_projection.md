# Scalar Projection of a Vector

**Scalar projection:** the signed length of the shadow one vector casts on another, equal to the dot product divided by the length of the vector being projected onto.

## The shadow reading

The scalar projection of one vector onto another is the length of the component of the first along the direction of the second, with a sign that is negative when the two point more than ninety degrees apart. It is the dot product of the two vectors divided by the length of the second, which is the geometric meaning of the dot product: the product of the lengths times the cosine of the angle between them, divided by one length, leaves the other times the cosine.

## Computing it

With components, the dot product is the sum of the products of corresponding coordinates, and the length is the square root of the sum of the squares. For a target of three and four, the length is five, so the scalar projection is the combination three times the first coordinate plus four times the second, divided by five. The arithmetic is short and exact whenever the target is a Pythagorean pair and the coordinates are chosen to make the numerator divisible.

## Projection vector and uses

Multiplying the scalar projection by the unit vector in the direction of the target gives the vector projection, the actual shadow vector rather than its length. The difference between the original vector and its projection is perpendicular to the target, which is the decomposition into a component along the target and a component across it. This decomposition is the geometry behind work, where only the force along the displacement contributes, and behind the least-squares fit, where a residual is made perpendicular to a subspace.
