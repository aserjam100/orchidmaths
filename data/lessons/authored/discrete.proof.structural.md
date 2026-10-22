# Structural induction

Objects defined recursively are proved about recursively. **Structural induction** has one case per rule of the definition: the base case proves the property for the objects the definition builds from nothing, and the step proves it for an object built by a rule, assuming it for the objects that rule used.

## Induction over a definition

Take the definition "a binary tree is a leaf, or a node with a left and a right subtree". A property \(P\) holds for every such tree if it holds for every leaf, and if whenever it holds for two trees it holds for the node built from them.

That is exactly the shape of ordinary induction, with the definition taking the place of "\(n\) follows \(n-1\)". The difference is that a rule may consume several smaller objects, which is why the hypothesis covers all of them rather than only the immediate predecessor.

## The two cases

For the tree definition, the base case is a single leaf and the step is a node with two subtrees. Proving that every full binary tree has an odd number of nodes goes: a leaf has one node; a node with subtrees of \(2a+1\) and \(2b+1\) nodes has \(2a+2b+3\), which is odd.

For a grammar such as \(S \to a \mid aS\), the base case is the string \(a\) and the step prepends an \(a\) to a string already known to have the property. Strings built this way have length equal to their number of \(a\)'s, and the two-case proof writes itself.

## Choosing the principle

Use structural induction when the object is defined by rules rather than by a number, and use strong induction when the object is a natural number but the step needs several earlier values. The two are the same idea — well-founded recursion — applied to different orders.

What does not work is ordinary induction on "the number of vertices" when the step removes an arbitrary vertex: the remainder need not be the same kind of object, and the hypothesis does not apply. The cases have to follow the definition.
