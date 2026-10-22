# Ramsey Numbers

**Ramsey number:** $R(a,b)$ is the least $n$ such that every red-blue coloring of the edges of $K_n$ contains a red $K_a$ or a blue $K_b$. By symmetry $R(a,b) = R(b,a)$.

## Two colours

Colour every edge of a complete graph red or blue, with no restriction on the pattern. Ramsey's theorem guarantees that for large enough $n$ a large monochromatic clique must appear, no matter how the coloring is chosen. The number $R(a,b)$ is the exact threshold, and it is finite for every pair $a, b$. The definition counts a red $K_a$ and a blue $K_b$, so the two arguments trade roles under swapping the colours.

## Worked: R(3,3) equals 6

Five vertices are not enough: colour the edges of a 5-cycle red and its complement, another 5-cycle, blue, and there is no monochromatic triangle. Six vertices are enough. Pick any vertex; it has five incident edges, so at least three share a colour, say red, to neighbours $u, v, w$. If any pair among $u, v, w$ is joined by a red edge, that edge with the centre forms a red triangle. Otherwise all three pairs are blue, forming a blue triangle. So $R(3,3) = 6$.

## Known values and growth

Small values are known exactly: $R(2,k) = k$, $R(3,4) = 9$, $R(3,5) = 14$ and $R(4,4) = 18$. The diagonal growth is not settled: the best known bounds leave a large gap between a power of two and the exponential upper bound. Finding $R(5,5)$ is a famous open problem, and the known lower bounds come from explicit colorings while the upper bounds come from counting arguments.
