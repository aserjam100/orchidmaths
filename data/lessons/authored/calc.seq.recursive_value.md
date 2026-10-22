# Recursive sequences

## Defined by what each term does

A recursive sequence states a first term and a rule that turns each term into the next. That is a complete definition: given the start and the rule, every later term is forced, though reaching one means walking through the terms before it.

The rule is usually written with the next index on the left, since it describes a step rather than a value.

## Iterating

To find a term far along, apply the rule that many times. The arithmetic is repetitive and exact, and a table of the intermediate terms is the natural record of it.

For an affine rule, where each term is a multiple of the previous one plus a constant, the growth is geometric after a shift, which is why the terms can become large quickly.

## Closed forms and recursion

Some recursive sequences have a closed form that gives a term directly, and finding it is a separate problem. The recursion is the definition; the closed form is a convenience, and not every recursion has one.

When the closed form is unknown, iterating is still a complete answer, which is why recursive definitions are useful in computing even when the mathematics prefers a formula.
