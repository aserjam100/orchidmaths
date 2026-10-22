# Angular Velocity and Linear Speed

**Angular velocity:** the rate at which an angle turns, measured in degrees or radians per unit time; multiplying it by the radius gives the linear speed along the circle.

## Two rates for one motion

A point moving around a circle has two speeds worth distinguishing. The angular velocity says how fast the central angle sweeps, measured as an angle per unit time. The linear speed says how fast the point itself travels along the arc, measured as a distance per unit time. One full revolution is 360 degrees, or two pi radians, and the arc covered in that revolution is the circumference, so the two rates differ by the radius.

## Converting between them

If a wheel makes a whole number of revolutions in a given number of seconds, its angular velocity in degrees per second is 360 times the number of revolutions divided by the seconds. Multiplying that by the radius gives the linear speed of a point on the rim. In radians the arithmetic is cleaner, because one revolution is two pi radians and the arc length is the radius times the angle, so the linear speed is simply the radius times the angular velocity with no conversion factor.

## Why the distinction matters

The distinction explains why two points on the same rotating disc have the same angular velocity but different linear speeds: the outer point covers more arc in the same time. It is the reason a gear train changes speed at each mesh, since the teeth move at one linear speed while the angular velocities differ with the radii. Problems about wheels, pulleys, belts and satellites are all the same computation, differing only in which of the two rates is given.
