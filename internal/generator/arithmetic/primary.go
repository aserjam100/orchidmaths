package arithmetic

// Primary 1 and 2 concepts that the MOE syllabus scopes more tightly than the
// corpus's general concepts: ordering numbers to 100 and to 1000, dividing within
// the 2, 3, 4, 5 and 10 tables, and word problems that stay inside those tables.
// The general concepts (arith.compare.whole, arith.div.basic, arith.mult.word,
// arith.div.word) are left alone because Primary 3 and later still use them.
//
// Word problems use names a Singapore classroom would recognise and never refer
// to the character with a pronoun.

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerPrimary(reg *generator.Registry) {
	reg.Register("arith.compare.within100", &compareRangeGen{lo: 1, hi: 100})
	reg.Register("arith.compare.within1000", &compareRangeGen{lo: 100, hi: 999})
	reg.Register("arith.div.tables", &divTablesGen{})
	reg.Register("arith.mult.word_tables", &multWordTablesGen{})
	reg.Register("arith.div.word_tables", &divWordTablesGen{})
}

// primaryTables are the multiplication tables the syllabus sets for Primary 2.
var primaryTables = []int{2, 3, 4, 5, 10}

var primaryNames = []string{"Mei Ling", "Ahmad", "Priya", "Wei Jie", "Siti", "Ravi", "Hui Min", "Farid"}

// compareRangeGen orders three different numbers in [lo, hi]. At higher difficulty the
// numbers share their leading digit, so the learner has to look past the first place.
type compareRangeGen struct{ lo, hi int }

func (g *compareRangeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	lead := 10 // the leading place: tens for numbers to 100, hundreds for numbers to 1000
	if g.hi > 100 {
		lead = 100
	}
	pick := func() int { return rand.Intn(g.hi-g.lo+1) + g.lo }
	shareLead := ctx.Difficulty >= 0.5
	first := pick()
	if shareLead {
		// Keep the first number out of the top block, which can be too small to hold three
		// numbers (100 is alone in its tens block).
		first = rand.Intn(g.hi-lead-g.lo+1) + g.lo
	}
	nums := []int{first}
	for len(nums) < 3 {
		n := pick()
		if shareLead {
			n = first/lead*lead + rand.Intn(lead)
			if n < g.lo {
				continue
			}
		}
		dup := false
		for _, m := range nums {
			dup = dup || m == n
		}
		if !dup {
			nums = append(nums, n)
		}
	}
	if nums[0] < nums[1] && nums[1] < nums[2] {
		nums[0], nums[2] = nums[2], nums[0]
	}
	ordered := append([]int(nil), nums...)
	sort.Ints(ordered)
	parts := make([]string, 3)
	for i, n := range ordered {
		parts[i] = strconv.Itoa(n)
	}
	answer := strings.Join(parts, ",")
	display := fmt.Sprintf("%d; %d; %d", nums[0], nums[1], nums[2])
	return generator.Problem{
		Question:    fmt.Sprintf("Order these numbers from smallest to greatest: %s. Enter them separated by commas.", display),
		Answer:      answer,
		Explanation: fmt.Sprintf("Compare the digits from the left; the order is %s.", strings.ReplaceAll(answer, ",", ", ")),
		Facts:       map[string]string{"numbers": display, "ordered": answer},
	}
}

// tableFact draws a factor from the Primary 2 tables and a multiplier from 1 to 10.
// Easier questions keep the multiplier small.
func tableFact(ctx generator.GeneratorContext) (factor, n int) {
	factor = primaryTables[rand.Intn(len(primaryTables))]
	n = rand.Intn(min(10, 4+int(ctx.Difficulty*7))) + 1
	return factor, n
}

type divTablesGen struct{}

func (g *divTablesGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	divisor, quotient := tableFact(ctx)
	dividend := divisor * quotient
	return generator.Problem{
		Question:    fmt.Sprintf("\\(%d \\div %d = ?\\)", dividend, divisor),
		Answer:      strconv.Itoa(quotient),
		Explanation: fmt.Sprintf("%d × %d = %d, so %d ÷ %d = %d", divisor, quotient, dividend, dividend, divisor, quotient),
		Facts:       map[string]string{"dividend": strconv.Itoa(dividend), "divisor": strconv.Itoa(divisor), "quotient": strconv.Itoa(quotient)},
	}
}

type multWordTablesGen struct{}

func (g *multWordTablesGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	perGroup, groups := tableFact(ctx)
	name := primaryNames[rand.Intn(len(primaryNames))]
	set := [][2]string{{"boxes", "erasers"}, {"bags", "marbles"}, {"plates", "curry puffs"}, {"packets", "stickers"}, {"cups", "pencils"}}[rand.Intn(5)]
	total := groups * perGroup
	return generator.Problem{
		Question:    fmt.Sprintf("%s has %d %s. There are %d %s in each of the %s. How many %s are there altogether?", name, groups, set[0], perGroup, set[1], set[0], set[1]),
		Answer:      strconv.Itoa(total),
		Explanation: fmt.Sprintf("%d groups of %d: %d × %d = %d %s", groups, perGroup, groups, perGroup, total, set[1]),
		Facts:       map[string]string{"groups": strconv.Itoa(groups), "per_group": strconv.Itoa(perGroup), "total": strconv.Itoa(total), "item": set[1], "name": name},
	}
}

// divWordTablesGen asks both kinds of division the syllabus teaches: sharing (the number
// of groups is known) and grouping (the size of a group is known). Both branches publish
// the same facts, so one solution schema covers them.
type divWordTablesGen struct{}

func (g *divWordTablesGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	divisor, quotient := tableFact(ctx)
	total := divisor * quotient
	name := primaryNames[rand.Intn(len(primaryNames))]
	item := []string{"sweets", "stickers", "marbles", "pencils", "cards"}[rand.Intn(5)]
	var question string
	if rand.Intn(2) == 0 {
		question = fmt.Sprintf("%s shares %d %s equally among %d friends. How many %s does each friend get?", name, total, item, divisor, item)
	} else {
		question = fmt.Sprintf("%s puts %d %s into bags of %d. How many bags are filled?", name, total, item, divisor)
	}
	return generator.Problem{
		Question:    question,
		Answer:      strconv.Itoa(quotient),
		Explanation: fmt.Sprintf("%d ÷ %d = %d, because %d × %d = %d", total, divisor, quotient, divisor, quotient, total),
		Facts:       map[string]string{"total": strconv.Itoa(total), "divisor": strconv.Itoa(divisor), "quotient": strconv.Itoa(quotient), "item": item, "name": name},
	}
}
