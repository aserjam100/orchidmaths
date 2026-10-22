# Logarithmic Scales

**Logarithmic scale:** A scale that reports the logarithm of a ratio rather than the ratio itself. Equal steps mean equal multiplicative factors, so a large range of physical values compresses into a small range of numbers.

## Equal steps, multiplicative changes

On a base-10 log scale, one unit is a factor of 10, and two units are a factor of 100. This is why an earthquake of magnitude 6 releases far more energy than one of magnitude 4: the difference of 2 is two factors of 10 in amplitude.

## Worked: pH, Richter and decibels

1. pH is $-\log_{10}[H^+]$, so $[H^+] = 10^{-3}$ gives pH 3.
2. Richter magnitude is $\log_{10}(A/A_0)$, so an amplitude ratio of $10^4$ gives magnitude 4.
3. Sound level is $10\log_{10}(I/I_0)$ decibels, so an intensity ratio of $10^6$ gives 60 dB.

## Reading a scale backwards

To recover the physical quantity, invert the logarithm: $[H^+] = 10^{-\mathrm{pH}}$, $A/A_0 = 10^{M}$, and $I/I_0 = 10^{L/10}$. A change of 1 on the pH scale is a factor of 10 in concentration; a change of 10 dB is a factor of 10 in intensity. The scales are convenient precisely because the underlying quantities span many orders of magnitude.
