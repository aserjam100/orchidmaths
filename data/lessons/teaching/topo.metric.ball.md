# Open Balls

**Open ball:** In a metric space $(X, d)$, the open ball of radius $r > 0$ around $x$ is $B(x, r) = \{y \in X : d(x, y) < r\}$. The strict inequality is what makes the ball open rather than closed.

## Distance first, then the ball

The ball is defined by the metric, so two metrics on the same set produce different balls and different open sets. In $\mathbb{R}$ the ball $B(3, 2)$ is the interval $(1, 5)$, while in $\mathbb{R}^2$ the ball of radius 5 around the origin is the disk $x^2 + y^2 < 25$.

## Worked: balls in the line and the plane

1. In $\mathbb{R}$, $d(0, 3) = 3$, so 3 lies on the boundary of $B(0, 3)$ and is not in the open ball.
2. In $\mathbb{R}^2$, $d((0,0), (3,4)) = \sqrt{9 + 16} = 5$, so $(3,4)$ is on the boundary of $B((0,0), 5)$.
3. In $\mathbb{R}$ with the usual metric, the ball of radius 2 centred at 3 is $(1, 5)$: the lower endpoint is 1 and the upper endpoint is 5.

## Why balls generate the topology

A set $U$ is open exactly when every point of $U$ has a ball around it that stays inside $U$. That is the same as saying $U$ is a union of balls: take one ball at each point, small enough to fit. This is why the collection of all balls is a basis for the metric topology.
