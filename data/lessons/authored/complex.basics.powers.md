# Powers of complex numbers

Raising a complex number to a whole-number power is repeated multiplication, and repeated multiplication of \(i\) has a period: \(i, -1, -i, 1\) and then back to \(i\).

## The cycle of i

The powers of \(i\) repeat with period four, so \(i^n\) depends only on the remainder of \(n\) on division by four. With remainder 0 the value is 1, with remainder 1 it is \(i\), with remainder 2 it is \(-1\), and with remainder 3 it is \(-i\).

So \(i^{23} = i^{20} \cdot i^{3} = (i^4)^5 \cdot i^3 = -i\). The reduction is the whole method; no multiplication is needed beyond the remainder.

## Squaring an expression

Squaring \(a + bi\) means multiplying it by itself: \((a+bi)^2 = a^2 + 2abi + b^2 i^2 = a^2 - b^2 + 2abi\). The cross term doubles and the \(i^2\) contributes the negative sign that mixes the parts.

For \(2 + 3i\): \(a = 2\) and \(b = 3\), so the real part is \(4 - 9 = -5\) and the imaginary part is \(2 \cdot 2 \cdot 3 = 12\). The square is \(-5 + 12i\).

## Beyond squares

Higher powers follow the binomial theorem, and the same collapse happens: every \(i^2\) becomes \(-1\), so the result is again a real part plus an imaginary part. For large powers the polar form is faster, since \((r(\cos\theta + i\sin\theta))^n = r^n(\cos n\theta + i\sin n\theta)\).

A useful check is that powers of a number with modulus greater than one grow in modulus, and powers of a number on the unit circle stay on it. A square that leaves the unit circle when the original did not is an arithmetic error.
