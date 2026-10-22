package fractions

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/aserjam100/orchid/internal/generator"
)

// fracCompareUnitLikeGen compares the two kinds of fraction pair the Primary 2 syllabus
// allows, with denominators up to 12: two unit fractions (1/3 and 1/5), where the larger
// denominator is the smaller fraction, and two like fractions (3/8 and 5/8), where the
// larger numerator is the larger fraction. Neither needs a common denominator, which
// frac.ops.compare does and which the syllabus leaves to Primary 3.
type fracCompareUnitLikeGen struct{}

func (g *fracCompareUnitLikeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	maxDen := 6 + int(ctx.Difficulty*6)
	var aNum, aDen, bNum, bDen int
	var rule string
	if rand.Intn(2) == 0 {
		aNum, bNum = 1, 1
		aDen = rand.Intn(maxDen-1) + 2
		bDen = rand.Intn(maxDen-1) + 2
		for bDen == aDen {
			bDen = rand.Intn(maxDen-1) + 2
		}
		rule = "Both are unit fractions. The whole is cut into more parts for the larger denominator, so each part is smaller."
	} else {
		aDen = rand.Intn(maxDen-2) + 3
		bDen = aDen
		aNum = rand.Intn(aDen-1) + 1
		bNum = rand.Intn(aDen-1) + 1
		for bNum == aNum {
			bNum = rand.Intn(aDen-1) + 1
		}
		rule = "The denominators are the same, so the parts are the same size. More parts make the larger fraction."
	}
	ans := ">"
	if aNum*bDen < bNum*aDen {
		ans = "<"
	}
	return generator.Problem{
		Question:    fmt.Sprintf("Compare: \\(\\frac{%d}{%d} \\_\\_ \\frac{%d}{%d}\\) (enter > or <)", aNum, aDen, bNum, bDen),
		Answer:      ans,
		Explanation: fmt.Sprintf("%s So %d/%d %s %d/%d.", rule, aNum, aDen, ans, bNum, bDen),
		Facts:       map[string]string{"a_num": strconv.Itoa(aNum), "a_den": strconv.Itoa(aDen), "b_num": strconv.Itoa(bNum), "b_den": strconv.Itoa(bDen), "op": ans, "rule": rule},
	}
}
