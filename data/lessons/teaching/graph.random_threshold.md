# Random Graph Thresholds

**Random graph G(n,p):** The Erdős–Rényi model $G(n,p)$ has $n$ labelled vertices, and each of the $\binom{n}{2}$ possible edges is present independently with probability $p$.

## The model G(n,p)

In $G(n,p)$ the edge count is a binomial random variable with $\binom{n}{2}$ trials and success probability $p$, so on average there are $p \binom{n}{2}$ edges. The degree of a vertex is binomial with $n-1$ trials. A property is called monotone when adding edges cannot destroy it, and a threshold is a value of $p$ around which the probability of the property jumps from near zero to near one for large $n$.

## Worked: the giant component

At $p = c/n$ the behaviour depends on the constant $c$. When $c < 1$ every component has size $O(\log n)$, and when $c > 1$ a unique component of size linear in $n$ appears and coexists with small components. This phase transition at $c = 1$ is the Erdős–Rényi threshold for a giant component, and it is the sharpest early result about $G(n,p)$.

## Connectivity and isolated vertices

The threshold for $G(n,p)$ to be connected is $p = (\log n)/n$, later than the giant-component threshold. The reason is the last small obstruction: isolated vertices. Just below $(\log n)/n$ the expected number of isolated vertices is still large, so the graph is disconnected with high probability, and just above it the graph becomes connected. Many thresholds are explained by the disappearance of a specific small forbidden configuration.
