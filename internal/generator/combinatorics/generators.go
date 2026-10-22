// Package combinatorics generates the counting problems of the `discrete_math` combinatorics
// subdomain.
//
// Every concept here is `numeric`: a counting answer is an exact integer that a learner produces
// rather than an option they pick, which is the change section 4.2 of docs/content-expansion.md
// argues for in the domains that are entirely `multiple_choice`. The explanation is derived from
// the same integers as the answer, so the two cannot disagree.
package combinatorics

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/aserjam100/orchid/internal/generator"
)

// Register adds this subdomain's generators to reg.
func Register(reg *generator.Registry) {
	reg.Register("discrete.combinatorics.counting_principle", &countingPrincipleGen{})
	reg.Register("discrete.combinatorics.factorial", &factorialGen{})
	reg.Register("discrete.combinatorics.multinomial", &multinomialGen{})
	reg.Register("discrete.combinatorics.circular", &circularGen{})
	reg.Register("discrete.combinatorics.derangement", &derangementGen{})
	reg.Register("discrete.combinatorics.catalan", &catalanGen{})
	registerCountingTechniques(reg)
}

// pick returns a random int in [0, n) and never panics on n < 1, so a difficulty of 0 (which the
// easy-pool ratchet samples at 0.12 but a caller may pass as 0) yields a one-element range rather
// than a panic inside rand.Intn.
func pick(n int) int {
	if n < 1 {
		n = 1
	}
	return rand.Intn(n)
}

// span widens an integer range with difficulty: floor is the easy-pool width and the extra width
// grows to `extra` at difficulty 1.
func span(difficulty float64, floor, extra int) int {
	return floor + int(difficulty*float64(extra))
}

func fact(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// derangements returns !n, computed by the recurrence !n = (n-1)(!(n-1) + !(n-2)) rather than by
// the inclusion-exclusion sum, because the recurrence is exact in integers at every n here.
func derangements(n int) int {
	if n == 0 {
		return 1
	}
	prev, cur := 1, 0 // !0 = 1, !1 = 0
	for i := 2; i <= n; i++ {
		prev, cur = cur, (i-1)*(cur+prev)
	}
	return cur
}

// catalan returns C_n from C_0 = 1 by C_{n+1} = C_n * 2(2n+1)/(n+2), which stays exact because the
// division is exact at every step.
func catalan(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}

type countingPrincipleGen struct{}

func (g *countingPrincipleGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	a := 2 + pick(span(ctx.Difficulty, 2, 5))
	b := 2 + pick(span(ctx.Difficulty, 2, 5))
	// The concept is the two rules, so a third of the instances ask for the one that adds. Both
	// branches produce many distinct answers in the easy pool.
	if pick(3) == 0 {
		return generator.Problem{
			Question: fmt.Sprintf("A trip from here to the coast can be made on any of %d different trains or any of %d different buses. How many different ways are there to make the trip?", a, b),
			Answer:   fmt.Sprintf("%d", a+b),
			Explanation: fmt.Sprintf("The sum rule adds the choices when only one of them is taken: %d trains + %d buses = %d ways.",
				a, b, a+b),
			Facts: map[string]string{
				"trains": fmt.Sprintf("%d", a), "buses": fmt.Sprintf("%d", b),
				"ways": fmt.Sprintf("%d", a+b), "rule": "the sum rule",
			},
		}
	}
	return generator.Problem{
		Question: fmt.Sprintf("A deli offers %d kinds of bread and %d kinds of filling, and a sandwich is one bread with one filling. How many different sandwiches can be made?", a, b),
		Answer:   fmt.Sprintf("%d", a*b),
		Explanation: fmt.Sprintf("The product rule multiplies the choices at each step: %d breads x %d fillings = %d sandwiches.",
			a, b, a*b),
		Facts: map[string]string{
			"breads": fmt.Sprintf("%d", a), "fillings": fmt.Sprintf("%d", b),
			"sandwiches": fmt.Sprintf("%d", a*b), "rule": "the product rule",
		},
	}
}

type factorialGen struct{}

func (g *factorialGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	n := 3 + pick(span(ctx.Difficulty, 3, 5))
	return generator.Problem{
		Question: fmt.Sprintf("Compute \\(%d!\\) — the number of ways to arrange %d distinct books on a shelf.", n, n),
		Answer:   fmt.Sprintf("%d", fact(n)),
		Explanation: fmt.Sprintf("%d! is the product of every whole number from 1 to %d, so %d! = %s = %d.",
			n, n, n, productChain(n), fact(n)),
		Facts: map[string]string{"n": fmt.Sprintf("%d", n), "value": fmt.Sprintf("%d", fact(n))},
	}
}

// productChain renders "1 x 2 x ... x n" for an explanation, and "1" for n = 1.
func productChain(n int) string {
	if n < 2 {
		return "1"
	}
	parts := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		parts = append(parts, fmt.Sprintf("%d", i))
	}
	return strings.Join(parts, " x ")
}

type multinomialGen struct{}

func (g *multinomialGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	n := 4 + pick(span(ctx.Difficulty, 2, 4)) // 4..9
	// Two or three distinct letters, so the multinomial coefficient has two or three factorials.
	kinds := 2 + pick(2)
	if kinds > n {
		kinds = n
	}
	counts := composition(n, kinds)
	letters := "ABCD"
	var word strings.Builder
	for i, k := range counts {
		for j := 0; j < k; j++ {
			word.WriteByte(letters[i])
		}
	}
	// Shuffle the letters so the word is not "AAABBC" every time.
	runes := []rune(word.String())
	rand.Shuffle(len(runes), func(i, j int) { runes[i], runes[j] = runes[j], runes[i] })
	shuffled := string(runes)

	total := fact(n)
	denoms := make([]string, 0, kinds)
	denom := 1
	for _, k := range counts {
		denoms = append(denoms, fmt.Sprintf("%d!", k))
		denom *= fact(k)
	}
	return generator.Problem{
		Question: fmt.Sprintf("How many distinct arrangements are there of the letters in \\(%s\\)?", shuffled),
		Answer:   fmt.Sprintf("%d", total/denom),
		Explanation: fmt.Sprintf("The word has %d letters, so %d! orders them and dividing by %s removes the orders that only swap identical letters: %d / %d = %d.",
			n, n, strings.Join(denoms, " "), total, denom, total/denom),
		Facts: map[string]string{
			"letters": fmt.Sprintf("%d", n), "total_orders": fmt.Sprintf("%d", total),
			"divide_by": fmt.Sprintf("%d", denom), "arrangements": fmt.Sprintf("%d", total/denom),
		},
	}
}

// composition splits n into exactly k positive parts, in the same order every call for a given
// (n, k), so the word it builds is checkable.
func composition(n, k int) []int {
	if k < 1 {
		return nil
	}
	parts := make([]int, k)
	for i := range parts {
		parts[i] = 1
	}
	for rest := n - k; rest > 0; rest-- {
		parts[pick(k)]++
	}
	return parts
}

type circularGen struct{}

func (g *circularGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	n := 4 + pick(span(ctx.Difficulty, 2, 4)) // 4..9
	return generator.Problem{
		Question: fmt.Sprintf("How many ways can %d people sit around a round table, if two seatings that differ only by a rotation are the same seating?", n),
		Answer:   fmt.Sprintf("%d", fact(n-1)),
		Explanation: fmt.Sprintf("Fix one person to break the rotation, then arrange the remaining %d around them: (%d-1)! = %d! = %d seatings.",
			n-1, n, n-1, fact(n-1)),
		Facts: map[string]string{"n": fmt.Sprintf("%d", n), "seatings": fmt.Sprintf("%d", fact(n-1))},
	}
}

type derangementGen struct{}

func (g *derangementGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	n := 3 + pick(span(ctx.Difficulty, 3, 4)) // 3..9
	return generator.Problem{
		Question: fmt.Sprintf("A shop has %d hats and %d coats hanging on %d numbered pegs, one hat and one coat per peg. The hats are all put back on the wrong pegs, so no hat returns to its own peg. How many ways can the hats be placed?", n, n, n),
		Answer:   fmt.Sprintf("%d", derangements(n)),
		Explanation: fmt.Sprintf("This is the number of derangements of %d items, written !%d, and !%d = %d: count the %d! arrangements of the hats, then subtract those that leave at least one hat at home.",
			n, n, n, derangements(n), n),
		Facts: map[string]string{"n": fmt.Sprintf("%d", n), "derangements": fmt.Sprintf("%d", derangements(n))},
	}
}

type catalanGen struct{}

func (g *catalanGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	n := 2 + pick(span(ctx.Difficulty, 2, 4)) // 2..8
	return generator.Problem{
		Question: fmt.Sprintf("How many strings of %d pairs of parentheses are correctly matched, with every prefix never closing more than it has opened?", n),
		Answer:   fmt.Sprintf("%d", catalan(n)),
		Explanation: fmt.Sprintf("That count is the Catalan number C_%d, and C_%d = %d: of the %d ways to arrange %d '(' and %d ')', the ones that ever close early are excluded.",
			n, n, catalan(n), choose(2*n, n), n, n),
		Facts: map[string]string{"n": fmt.Sprintf("%d", n), "catalan": fmt.Sprintf("%d", catalan(n))},
	}
}

// choose is the binomial coefficient nCr, computed multiplicatively so it stays exact.
func choose(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	c := 1
	for i := 0; i < r; i++ {
		c = c * (n - i) / (i + 1)
	}
	return c
}
