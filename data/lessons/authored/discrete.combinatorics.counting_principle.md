# The product and sum rules

Counting starts from two rules that turn a description of choices into a number.

## The product rule

If a task is done in two steps, and the first step has \(m\) outcomes while the second has \(n\) outcomes whatever the first was, then the task has \(m \times n\) outcomes. A deli with 3 kinds of bread and 4 kinds of filling makes \(3 \times 4 = 12\) sandwiches, because each bread pairs with each filling.

The rule extends to any number of steps: multiply the outcomes at every step. Four shirts, 5 pairs of trousers and 2 hats give \(4 \times 5 \times 2 = 40\) outfits.

## The sum rule

If a task is done in one of several *disjoint* ways, and the ways have \(m\) and \(n\) outcomes with none counted twice, then the task has \(m + n\) outcomes. A trip can be made on 5 trains or on 3 buses, and no bus is a train, so there are \(5 + 3 = 8\) ways to travel.

The word *disjoint* carries the work. If the two sets of outcomes overlapped, adding would count the shared ones twice, and the count would need the subtraction that inclusion and exclusion provides.

## Using both rules

Most problems need both, and the order is what makes the working readable. To count the 4-character codes that begin with a letter and end with a digit: 26 choices for each of the first three characters and 10 for the last, so the product rule gives \(26^3 \times 10 = 175{,}760\).

When a problem splits into cases that cannot both happen, count each case with the product rule and add the results with the sum rule.
