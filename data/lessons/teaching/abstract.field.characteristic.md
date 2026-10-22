# Characteristic of a Field

**Characteristic:** The characteristic of a field $F$ is the least positive integer $n$ with $n \cdot 1 = 0$, if such an $n$ exists, and otherwise it is 0. Here $n \cdot 1$ means $1 + 1 + \cdots + 1$ with $n$ terms.

## Repeatedly adding the identity

The decisive point is that the definition adds copies of 1, not arbitrary elements. If no positive number of copies of 1 sums to 0, the field has characteristic 0; that is the case for $Q$, $R$, $C$ and every field containing $Q$. If some number does sum to zero, the least such number must be prime: if $n = ab$ with both factors smaller, then $(a \cdot 1)(b \cdot 1) = 0$, which a field cannot have with both factors nonzero.

## Prime fields and their characteristic

A field of characteristic $p$ contains a copy of $Z_p$ as its prime subfield, obtained by sending $k$ to $k \cdot 1$. A field of characteristic 0 contains a copy of $Q$. The characteristic is therefore either 0 or a prime, never a composite number.

## Worked: F_9 has characteristic 3, not 9

The field $F_9$ has $9 = 3^2$ elements.

1. It contains $F_3$ as its prime subfield, and $1 + 1 + 1 = 0$ in $F_3$.
2. No smaller positive number of copies of 1 sums to zero.
3. So the characteristic is 3, while the order of the field is 9.

Order and characteristic are different invariants: every finite field has order $p^n$ and characteristic $p$, and the exponent $n$ does not appear in the characteristic.
