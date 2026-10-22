# Lattice paths

## Shortest paths on a grid

A shortest path from \((0,0)\) to \((m,n)\) that steps only right and up uses exactly \(m\) right steps and \(n\) up steps, in some order. Choosing which of the \(m+n\) positions hold the right steps determines the path, so the number is

\[\binom{m+n}{m}.\]

To \((3,2)\) there are \(\binom{5}{3} = 10\) paths.

## Words and paths

The same count is the number of words with \(m\) of one letter and \(n\) of another, which is the simplest case of the multinomial coefficient. The correspondence between paths and words is what makes lattice paths the usual model for a counting problem that looks geometric.

## Paths that stay on one side

If a path may not cross the diagonal, the count changes from a binomial coefficient to a Catalan number: the paths from \((0,0)\) to \((n,n)\) that never rise above the diagonal are the \(n\)th Catalan number, by the reflection principle. That is the same reflection that counts balanced strings of parentheses, seen as a picture rather than as a word.
