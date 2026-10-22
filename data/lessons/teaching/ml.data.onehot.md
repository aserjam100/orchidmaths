# One-Hot Encoding

**One-hot encoding:** Represent a categorical feature with $C$ categories by $C$ binary columns, putting a 1 in the column of the observed category and 0 in every other. It is the standard way to feed unordered labels to a model that multiplies features by weights.

## One column per category

Integer codes would impose an order and a distance that the categories do not have: coding red, green and blue as 1, 2 and 3 tells the model that blue is twice green. One-hot removes that false arithmetic by making every category equidistant from every other.

## Worked: encoding three categories

1. Categories red, green and blue get columns [1, 0, 0], [0, 1, 0] and [0, 0, 1].
2. A row with category green becomes [0, 1, 0].
3. A linear model learns one weight per column, so each category gets its own effect.

## The dummy-variable trap

With an intercept, the $C$ one-hot columns sum to the all-ones column, so the design matrix is rank deficient. Dropping one column gives $C-1$ columns and removes the collinearity; the dropped category becomes the reference level. Tree models do not need the drop, and some implementations handle categories natively.
