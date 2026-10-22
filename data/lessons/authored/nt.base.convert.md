# Writing a number in another base

## Place value in any base

A numeral's value is the sum of its digits each multiplied by a power of the base, counted from the right starting at zero. Base ten is the familiar case, and nothing about the rule depends on the base being ten.

So reading a numeral in another base is the same expansion with the powers of ten replaced.

## Expanding the digits

Working from the left and repeatedly multiplying by the base and adding the next digit performs the expansion in one pass, which is why that method is the quick one. The equivalent from the right is to weight each digit by its power of the base and add.

Both give the same value, and the left-to-right pass needs no powers computed separately.

## Bases that are not decimal

Binary and hexadecimal appear wherever a machine's storage is described, because a group of bits is one hexadecimal digit. Base twelve survives in the division of the day and the year into twelfths, which is a reminder that the base is a convention rather than a fact about the numbers.

Converting between bases never changes the value, only the way it is written down.
