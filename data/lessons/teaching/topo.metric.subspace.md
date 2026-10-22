# The Subspace Topology

**Subspace topology:** For a subset $A$ of a space $X$, the open sets of $A$ are exactly the traces $A \cap U$ with $U$ open in $X$. This is the coarsest topology that makes the inclusion of $A$ into $X$ continuous.

## Restricting the opens

The subspace topology keeps the ambient opens but cuts them down to $A$. A set can be open in $A$ without being open in $X$, because the points outside $A$ are no longer available to witness the openness. The definition also makes the inclusion map continuous by construction.

## Worked: an open set inside [0,1]

1. Take $A = [0,1]$ as a subspace of $\mathbb{R}$.
2. The set $[0, 1/2)$ is the trace of $(-1, 1/2)$, which is open in $\mathbb{R}$.
3. So $[0, 1/2)$ is open in the subspace, while it is not open in $\mathbb{R}$: no ball around 0 stays inside it.

## Which properties are inherited

Hausdorffness and metrizability pass to subspaces: separating ambient opens can be intersected with $A$. Compactness and connectedness do not always pass, because the subspace may omit the points that made the ambient space behave. Openness and closedness are relative: a subspace can have sets that are both, or neither, even when the ambient space is connected.
