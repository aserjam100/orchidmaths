# The dot product

## Multiplying two vectors to get a number

The product of two vectors could be another vector, and sometimes it is. The dot product is the other
choice: multiply the matching components and add, so
\(\langle a, b \rangle \cdot \langle c, d \rangle = ac + bd\), and the result is a single number.

For \(\langle 3, -2 \rangle\) and \(\langle 4, 5 \rangle\) it is \(12 - 10 = 2\). The number is the
whole output.

## What the sign says

The sign of the dot product reports how the two directions relate. Two vectors pointing broadly the
same way give a positive result, because both products are positive or they outweigh the negative
terms. Two pointing broadly apart give a negative one.

A dot product of zero is the sharp case: neither vector is zero, yet the number is, and the two are
perpendicular. The zero comes from the two products cancelling exactly, which is what a right angle
between the directions means in components.

## The angle behind the number

The dot product is also the product of the magnitudes times the cosine of the angle between the
vectors: \(\mathbf{u} \cdot \mathbf{v} = |\mathbf{u}||\mathbf{v}|\cos\theta\). That identity is why
it is useful, and the component formula is what makes it computable.

Reading it backwards gives the angle, and the perpendicular case falls out of it: if the cosine is
zero, the angle is \(90^{\circ}\). Nothing here needs the vectors to be drawn.
