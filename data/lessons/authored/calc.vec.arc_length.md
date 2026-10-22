# Arc length of a parametrised curve

## The speed integral

Moving along a parametrised curve, the distance covered in a short time is the speed times the time, so the length of the path is the integral of the speed:

\[L = \int_a^b \sqrt{\left(\frac{dx}{dt}\right)^2 + \left(\frac{dy}{dt}\right)^2}\,dt.\]

The integrand is the length of the velocity vector, which is why the formula holds for any parametrisation of the same path: reparametrising changes the speed and the interval together, and the two cancel.

## Straight parametrisations

For \(x = 3t\) and \(y = 4t\) with \(0 \le t \le 1\), the velocity is constant and its length is \(\sqrt{9 + 16} = 5\). The integral of a constant over an interval of length 1 is that constant, so the arc length is 5, which agrees with the distance between the endpoints.

A direction whose two legs form a right triangle with a whole hypotenuse therefore gives a whole length, and that is the easiest place to check the formula by hand.

## Why the formula works

The curve is approximated by short straight pieces. Each piece has length about \(\sqrt{(\Delta x)^2 + (\Delta y)^2}\), dividing by \(\Delta t\) and multiplying by it again turns the sum into a Riemann sum for the speed, and the integral is the limit.

The same reasoning gives the arc length of a graph, \(\int \sqrt{1 + (y')^2}\,dx\), by choosing \(x\) itself as the parameter.
