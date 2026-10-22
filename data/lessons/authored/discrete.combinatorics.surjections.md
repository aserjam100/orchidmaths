# Counting onto functions

## Choosing the image set

A function from an \(m\)-element set to an \(n\)-element set is onto when every element of the codomain is hit at least once. Counting them means removing the cases that miss an image.

The removal is the hard part, because the cases overlap. To count by image set instead, fix which elements are hit: if the image has exactly \(n\) elements there is nothing to choose, but a surjection is equivalently a partition of the \(m\)-element domain into \(n\) nonempty labelled blocks, each block being the preimage of one codomain element. Splitting the count that way gives

\[\operatorname{onto}(m,n) = n!\,S(m,n),\]

where \(S(m,n)\) is the Stirling number of the second kind, the number of ways to split \(m\) labelled objects into \(n\) unlabelled nonempty blocks. The \(n!\) puts the labels back. This is the version to use when the blocks are what you are counting; the inclusion-exclusion formula below is the version to use when you only want the number.

## Inclusion and exclusion

Start from all \(n^m\) functions and subtract those that miss at least one image. Missing a specific image leaves \((n-1)^m\) functions, there are \(n\) images to miss, and adding back the ones that miss two keeps the accounting right:

\[\sum_{i=0}^{n} (-1)^i \binom{n}{i} (n-i)^m.\]

For \(m = 4\) and \(n = 2\) the value is \(2^4 - 2 \cdot 1^4 = 14\). The sum alternates because a function missing \(k\) images has been subtracted \(k\) times and added back \(\binom{k}{2}\) times.

The alternating sum is the inclusion-exclusion principle applied to the events \(A_j\) = "image \(j\) is missed". A function that misses exactly \(k\) images lies in \(\binom{k}{i}\) of the \(i\)-fold intersections, so its total contribution is \(\sum_i (-1)^i\binom{k}{i}\), which is 0 unless \(k=0\), and 1 when \(k=0\). Every non-surjection cancels and every surjection counts once.

Two more values are worth computing to see the formula work. For \(m=5\), \(n=3\):

\[3^5 - 3\cdot2^5 + 3\cdot1^5 = 243 - 96 + 3 = 150,\]

and the bookkeeping agrees: \(S(5,3)=25\), so \(3!\cdot25=150\). For \(m=4\), \(n=4\) the sum is \(4^4 - 4\cdot3^4 + 6\cdot2^4 - 4\cdot1^4 = 256 - 324 + 96 - 4 = 24\), which is \(4!\), as it must be.

## The case that matters

The count is zero when \(n > m\), which the formula produces on its own rather than needing a separate rule, and it is \(n!\) when \(m = n\), since an onto function between equal finite sets is a bijection.

The two boundaries have a common explanation. A surjection needs at least one domain element per image, so \(n>m\) is impossible and the alternating sum cancels to 0; and when \(m=n\) the blocks are all singletons, so the only freedom is which domain element goes to which image, which is \(n!\).

The first case past the diagonal is worth having, because it is where the formula stops being an alternating cancellation and starts being combinatorial. With \(m=n+1\), exactly one image is hit twice and the rest once, so

\[\operatorname{onto}(n+1,n) = n!\binom{n+1}{2},\]

choosing the pair that shares an image and then labelling the \(n\) images. For \(n=3\) that is \(6\cdot6=36\). The alternating sum gives the same number for \(m=4\), \(n=3\):

\[3^4 - 3\cdot2^4 + 3\cdot1^4 = 81 - 48 + 3 = 36.\]

The two routes are the same count, which is the point of keeping both: the Stirling form says what a surjection *is*, and the alternating sum computes how many there are.

A common error is to count the onto functions as \(n!\cdot S(m,n)\) with the blocks filled in the wrong order, or to use \(n^m - n(n-1)^m\) and stop there. The second undercounts precisely when \(m\) is large enough for a function to miss two images, which the \(m=5, n=3\) value shows: the truncated expression gives \(243-96=147\), and 150 is right.
