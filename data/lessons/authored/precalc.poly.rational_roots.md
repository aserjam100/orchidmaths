# The rational-root theorem

## Where the candidates come from

If a polynomial with whole-number coefficients has a rational root, then in lowest terms that root has the form

\[
\pm\frac{\text{a divisor of the constant term}}{\text{a divisor of the leading coefficient}}.
\]

Both the numerator and the denominator are constrained, which is what makes the list finite. A polynomial whose leading coefficient is 1 has only the divisors of its constant term to test.

## Counting the candidates

The number of candidates is the number of divisors of the constant term times the number of divisors of the leading coefficient, and then doubled for the two signs.

For a leading coefficient of 2 and a constant term of 12 the divisor counts are 2 and 6, so there are \(2 \times 6 = 12\) candidates before the signs are added.

## What it is and is not

The theorem gives candidates, not roots: a candidate may fail, and a polynomial with no rational root still has a full list of candidates that all fail. It also says nothing about irrational or complex roots, so a cubic can have three real roots and no rational one.

The practical use is to reduce an unbounded search to a short list, and then test that list by substitution, which is where the remainder theorem does the work.
