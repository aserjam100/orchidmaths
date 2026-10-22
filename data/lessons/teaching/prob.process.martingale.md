# Martingales

**Martingale:** a sequence whose next value, given the whole past, has conditional expectation equal to the present value, the mathematical model of a fair game.

## The conditional-mean property

For a martingale the best forecast of tomorrow given everything known today is today's value. Taking expectations of the defining identity shows that the mean is constant over time, so a martingale has no drift. This is the precise sense in which a fair game offers no advantage.

## Examples

A symmetric random walk is a martingale, because each step has mean zero. The square of a symmetric walk is not, but subtracting the number of steps restores the property, since the second moment grows exactly linearly. Building a martingale out of a process is often the first step in proving a limit theorem about it.

## Optional stopping

The optional stopping theorem says that stopping a martingale at a suitable random time does not change its mean. The conditions matter: without boundedness or integrability the conclusion can fail, which is why the theorem is a common source of errors. Gambling systems that look profitable usually violate exactly these conditions.
