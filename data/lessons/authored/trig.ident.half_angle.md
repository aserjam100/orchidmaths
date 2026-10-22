# The half-angle formulas

## Halving an angle

The double-angle formulas describe \(2\theta\) from \(\theta\), and they can be read backwards.
Solving for the ratios at half an angle turns \(\cos(2\theta) = \cos^{2}\theta - \sin^{2}\theta\)
into a statement about \(\theta/2\), which is what the half-angle formulas are.

The reason to want them is that some angles have no exact values of their own but are twice an angle
that does. Half of \(45^{\circ}\) is \(22.5^{\circ}\), whose sine is on no table, yet the formula
reaches it from \(\cos 45^{\circ}\).

## The two formulas and their roots

With \(\cos(2t) = \cos^{2}t - \sin^{2}t\) and \(\cos^{2}t + \sin^{2}t = 1\), writing \(\theta\) for
\(2t\) leaves \(\cos\theta = 2\cos^{2}(\theta/2) - 1\). Solving that for each ratio gives
\(\sin(\theta/2) = \pm\sqrt{(1 - \cos\theta)/2}\) and
\(\cos(\theta/2) = \pm\sqrt{(1 + \cos\theta)/2}\).

Both are square roots, so both carry a sign that the formula alone cannot decide. The sign is settled
by the quadrant \(\theta/2\) falls in: if \(\theta\) is between \(0^{\circ}\) and \(90^{\circ}\) then
so is \(\theta/2\), and both roots are positive.

## When the root comes out exact

A square root is usually where an exact answer would be lost, but a cosine taken from a right
triangle makes the radicand a perfect square. Writing the cosine as \((q^{2} - 2p^{2})/q^{2}\), with
\(p\), \(b\) and \(q\) the legs and hypotenuse of a triple, puts \((1 - \cos\theta)/2\) equal to
\(p^{2}/q^{2}\), whose root is \(p/q\); the other radicand is \(b^{2}/q^{2}\).

With \(\cos\theta = 7/25\) from the \(3, 4, 5\) triple, \((1 - 7/25)/2 = 9/25\), so
\(\sin(\theta/2) = 3/5\), and the cosine's radicand is \(16/25\), so \(\cos(\theta/2) = 4/5\). Both
legs of the triple are the two answers.
