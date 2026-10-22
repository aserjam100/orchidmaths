# L'Hopital's Rule

**L'Hopital's rule:** for a quotient whose numerator and denominator both tend to zero or both to infinity, the limit of the quotient equals the limit of the quotient of derivatives, when that limit exists.

## The indeterminate case

When a quotient has numerator and denominator both tending to zero, the ratio is not determined by the two limits alone and is called an indeterminate form. L'Hopital's rule replaces the ratio of the functions by the ratio of their derivatives near the point, because the derivatives control how fast each side approaches zero. If the new ratio has a limit, that limit is the limit of the original quotient. The rule applies to the form infinity over infinity as well, by the same comparison of rates.

## Repeated application

When the ratio of the derivatives is again indeterminate, the rule may be applied again, as long as the functions stay differentiable and the new ratio has a limit. Applying it twice is what settles the limit of one minus the cosine over the square, because one application still leaves zero over zero and the second leaves a finite ratio. Each application must itself meet the hypotheses; applying the rule blindly when the ratio of derivatives has no limit proves nothing, even if the original limit exists.

## When it does not help

L'Hopital's rule is a sufficient tool, not a necessary one: the ratio of derivatives may have no limit while the original quotient does, and then the rule is silent. It also gives the wrong answer if the hypotheses are ignored, notably when the denominator has zeros accumulating at the point so that the derivative quotient oscillates. The safe practice is to confirm that the limit of the derivative ratio exists before quoting the rule, and to fall back on the Taylor expansion or a direct estimate when it does not.
