# Factorials

## Definition

For a whole number \(n \ge 1\), the **factorial** \(n!\) is the product of every whole number from 1 to \(n\):

\[n! = 1 \times 2 \times \cdots \times n\]

So \(3! = 1 \times 2 \times 3 = 6\) and \(5! = 120\). The empty product is defined as \(0! = 1\), which is the value that makes the counting formulas below work for the empty case rather than a convention invented for its own sake.

## Counting arrangements

A factorial counts *orderings*. The number of ways to arrange \(n\) distinct objects in a row is \(n!\): pick which object goes first, \(n\) ways, then which goes second, \(n-1\) ways, and so on down to 1, and the product rule multiplies them.

Five distinct books can be shelved in \(5! = 120\) orders. Add one book and the count multiplies by 6, because the sixth book can be inserted into any of the six gaps in each existing order.

## Growth

Factorials grow faster than any fixed power or exponential of \(n\), which is why they appear as the dividing factor in every counting formula that treats some objects as identical. Comparing \(10! = 3{,}628{,}800\) with \(2^{10} = 1024\) shows the difference early; by \(20!\) the value has passed \(2.4 \times 10^{18}\).

Because the growth is so fast, count with the formula and cancel before multiplying, rather than computing a factorial and dividing afterwards.
