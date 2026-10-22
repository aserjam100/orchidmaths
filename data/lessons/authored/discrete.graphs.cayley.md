# Labelled trees

A **labelled tree** is a tree on a fixed vertex set, so two trees count as different when they use different edges even if they have the same shape. **Cayley's formula** counts them: there are \(n^{n-2}\) labelled trees on \(n\) given vertices.

## Counting labelled trees

For three vertices the formula gives \(3^1 = 3\), which is right: the three trees are the three choices of which vertex sits in the middle. For four vertices it gives \(4^2 = 16\), and the count includes the four path shapes and the twelve stars-with-a-tail arrangements.

The exponent \(n - 2\) is what makes the formula memorable, and the usual proof is a bijection rather than a count.

## Prüfer sequences

A **Prüfer sequence** is a string of \(n - 2\) labels from the \(n\) vertices. It is built from a tree by repeatedly deleting the smallest leaf and recording its neighbour; it is turned back into a tree by reversing the record.

The correspondence is a bijection between labelled trees and sequences of length \(n-2\) over an alphabet of size \(n\), so the count is \(n^{n-2}\) by counting sequences. The smallest-leaf rule is what makes the correspondence canonical: without a fixed deletion rule the map would not be well defined.

## Using the formula

For five vertices, \(5^3 = 125\) labelled trees; for six, \(6^4 = 1296\). The numbers grow fast enough that enumerating by hand is hopeless beyond five vertices, which is the point of a closed formula.

The formula also settles questions about random trees: if a tree is chosen uniformly, the Prüfer sequence is a uniform random string, so the degree of a vertex is one plus the number of times its label appears — a geometric distribution with mean approaching two for large \(n\).
