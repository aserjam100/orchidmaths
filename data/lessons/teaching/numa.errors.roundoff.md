# Floating-Point Roundoff and Machine Epsilon

**Machine epsilon:** the gap between one and the next representable floating-point number, the natural unit of the rounding error that every machine arithmetic carries.

## Finite storage, infinite numbers

A machine stores a number as a sign, a fixed-length significand and an exponent, so only finitely many reals are representable. Every other real is rounded to the nearest one, and the relative error of that rounding is at most about half of machine epsilon. The significand length sets epsilon, which is roughly ten to the minus sixteen for the usual double precision.

## Worked: a number that will not terminate

One tenth is a finite decimal but its binary expansion never stops, so it cannot be stored exactly. The stored value is the nearest double, and the difference is a rounding error of the order of machine epsilon times one tenth. Adding one tenth ten times does not give exactly one, which is a direct consequence of this rounding.

## Rounding is not associative

Each floating-point operation rounds its result, so regrouping a sum changes which roundings happen and can change the answer. Adding a tiny number to a large one can even lose the tiny number entirely, because it falls below the last digit of the large one. This is why the order of a computation is part of the algorithm, not an afterthought.
