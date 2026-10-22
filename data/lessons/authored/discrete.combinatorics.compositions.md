# Compositions of an integer

## Ordered sums

A composition of \(n\) is a way of writing \(n\) as a sum of positive integers where the order matters, so \(1+2\) and \(2+1\) are different compositions of 3.

Fixing the number of parts gives a binomial coefficient: the number of compositions of \(n\) into exactly \(k\) positive parts is

\[\binom{n-1}{k-1}.\]

For \(n = 5\) and \(k = 2\) that is \(\binom{4}{1} = 4\), the four splits \(1+4\), \(2+3\), \(3+2\) and \(4+1\).

## Stars and bars with a twist

The formula is the stars-and-bars count with the bars placed between items rather than among them. Writing \(n\) ones in a row leaves \(n-1\) gaps, and choosing \(k-1\) gaps to cut gives the ordered parts directly.

Allowing zero parts is the ordinary stars and bars: non-negative solutions of \(x_1 + \cdots + x_k = n\) number \(\binom{n+k-1}{k-1}\), and insisting that every part be at least one is what removes the extra freedom.

## Counting all of them

Summing over \(k\) gives \(2^{n-1}\) compositions of \(n\) in total, since each of the \(n-1\) gaps is either cut or not. That is the fastest check on the formula: at \(n = 4\) the four values of \(k\) give \(1 + 3 + 3 + 1 = 8 = 2^3\).
