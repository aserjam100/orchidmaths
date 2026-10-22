# Composite Rules

**Composite rules:** split the interval into many small pieces, apply a simple rule on each piece, and add the results, so the error falls with the width of the pieces.

## Many small pieces

A single trapezoid over a wide interval is inaccurate, but many narrow trapezoids together follow the curve closely. The composite rule applies the basic rule on each subinterval and sums. Interior sample points are shared by two neighbours, so in the composite trapezoid rule the two end values carry weight one half while every interior value carries weight one.

## Worked: a few subintervals

For the square function on the interval from nought to one with two subintervals, the width is one half and the sample values are nought, a quarter and one. The weighted sum is nought plus a quarter plus a half, times the width, giving 0.375. The exact value is one third, so the error has fallen from the single-trapezoid case as the pieces multiplied.

## The convergence rate

Because each piece contributes an error proportional to the cube of its width, the total error over a fixed interval is proportional to the square of the width, so halving the width quarters the error. For the composite Simpson rule the total error is proportional to the fourth power of the width, so halving it divides the error by sixteen. This is why composite rules are used in every serious quadrature routine.
