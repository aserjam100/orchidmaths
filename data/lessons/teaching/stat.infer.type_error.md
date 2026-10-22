# Type I and II Errors and Power

**Two errors:** A Type I error rejects a true null (a false alarm, rate $\alpha$); a Type II error fails to reject a false null (a missed detection, rate $\beta$). Power is $1 - \beta$, the chance of detecting a real effect.

## Two ways to be wrong

The four outcomes form a table: a true null is either kept correctly or rejected wrongly, and a false null is either detected or missed. Lowering alpha shrinks the false-alarm region, but with a fixed sample it also enlarges the missed-detection region, because the rejection threshold moves.

## Worked: the trade in one table

1. Test at alpha = 0.05: 5 percent of true nulls are rejected, and a real effect is detected with whatever power the design provides.
2. Tighten to alpha = 0.01: false alarms fall to 1 percent, but the critical value moves outward and power drops.
3. Increase the sample size: the standard error shrinks, so power rises without raising alpha.

## Power and sample size

Power depends on the effect size, the variance, the sample size and alpha. The usual way to raise it is to collect more data, which shrinks the standard error and separates the true effect from noise. Power is worth computing before an experiment: an underpowered study can miss a real effect and still look perfectly valid.
