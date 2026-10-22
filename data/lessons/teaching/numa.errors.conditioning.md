# The Condition Number of a Problem

**Condition number:** the factor by which a small relative change in the data can be magnified into a relative change in the answer, a property of the problem rather than of any algorithm.

## Sensitivity of the answer

A problem is well conditioned when a small relative perturbation of its data changes the answer only a little, and ill conditioned when it can change the answer a lot. The condition number makes this quantitative. It depends on the problem alone, so no clever algorithm can remove it; a good algorithm can only avoid adding error on top.

## Worked: raising to a power

For the function that raises its argument to the power n, the relative condition number is the size of x times the derivative divided by the size of the function. The x cancels, leaving n. Squaring is therefore well conditioned near a value of order one, while raising to the twentieth power multiplies any relative error in the input by twenty.

## Conditioning is not stability

An ill-conditioned problem has sensitive answers even when solved exactly, while an unstable algorithm adds avoidable error to a problem that may be perfectly conditioned. The two ideas combine multiplicatively: the computed relative error is roughly the condition number times the backward error of the algorithm. Separating them tells you whether to change the problem or the method.
