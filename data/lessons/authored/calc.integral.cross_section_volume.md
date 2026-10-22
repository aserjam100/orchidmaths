# Volumes from known cross-sections

## Slicing the solid

When every cross-section perpendicular to an axis is known, the volume is the integral of that area along the axis:

\[
V = \int_a^b A(x)\,dx.
\]

A solid whose sections are squares of side \(x\) has sections of area \(x^2\), so its volume over an interval starting at zero is a third of the cube of the endpoint.

## The integral of the area

The formula is the same reasoning as for a solid of revolution, with the area supplied directly instead of being computed from a radius. Solids of revolution are the special case where every section is a disc, and the cross-section method is the general one.

The step worth slowing down for is expressing the cross-section's area in terms of the axis variable: the side or radius usually comes from a curve the problem gives, and it is a function of \(x\) rather than a constant.

## Choosing the axis

The sections decide which variable integrates. Slicing perpendicular to the \(x\)-axis integrates in \(x\), and a solid better described by horizontal slices integrates in \(y\). Both give the same volume, but one is usually far easier.

For a solid with known sections the choice is not free, though: the sections are given at a particular orientation, and the integral must run along the direction they are stacked in.
