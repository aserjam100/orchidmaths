# Newton's Law of Cooling

**Newton's law of cooling:** The rate at which an object cools is proportional to the difference between its temperature and the ambient temperature: $T' = -k(T - T_{env})$ with $k > 0$. The solution is $T(t) = T_{env} + (T_0 - T_{env})e^{-kt}$.

## The rate follows the temperature gap

When the object is far from room temperature it cools quickly; as the gap shrinks, cooling slows. The model therefore predicts an exponential approach rather than a linear one, and the constant $k$ depends on the object and its surroundings.

## Worked: a cooling body

1. A body at 90 C is placed in a 20 C room, and the model is $T(t) = 20 + 70e^{-0.1t}$.
2. The temperature excess is $70e^{-0.1t}$.
3. After 10 minutes the excess is $70e^{-1} \approx 25.75$, so the body is about 45.75 C.

## The ambient limit

As $t \to \infty$ the exponential term vanishes and $T \to T_{env}$: the object approaches room temperature but never crosses it in the model. The ambient temperature is the constant term, and the initial temperature is $T_{env}$ plus the coefficient of the exponential. A negative $k$ would model heating away from the ambient value, which is unphysical for passive cooling.
