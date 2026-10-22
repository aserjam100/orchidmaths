package numtheory

// The 2026-10-18 number-theory batch: ten concepts that take the domain from
// 32 to the target of 42.
//
// The domain already covers divisibility, gcd, congruences, Fermat, Euler phi,
// Wilson, Legendre, quadratic reciprocity, orders, primitive roots, Mobius, the
// prime number theorem, zeta, Chebyshev, elliptic curves, continued fractions,
// class numbers and RSA. Several obvious additions were already taught and were
// not repeated: modular inverses in the Bezout lesson, Euler's criterion in the
// Legendre lesson, perfect numbers and Mersenne primes in the divisor-sum
// lesson, twin primes and Brun's sieve in the sieve lesson, and Gaussian
// integers in the class-number lesson.
//
// The batch fills real gaps instead: the Jacobi symbol, discrete logarithms,
// square roots modulo a prime, sums of two squares, Pell's equation, Pisano
// periods, the Carmichael function, the Miller-Rabin test, prime gaps, and
// Bertrand's postulate.
//
// Eight concepts are numeric and two are multiple choice. Every numeric
// generator publishes the facts its answer follows from, so the tests recompute
// the answer from those facts. The wiring is scoped: only the new concepts plus
// the course's advertised RSA and Diophantine topics are targeted, so the
// introductory course does not become a catalogue of zeta functions, elliptic
// curves and class numbers.

import (
	"math/rand"
	"strconv"

	"github.com/aserjam100/orchid/internal/generator"
)

func registerBatch(reg *generator.Registry) {
	reg.Register("nt.adv.jacobi", &jacobiGen{})
	reg.Register("nt.adv.discrete_log", &discreteLogGen{})
	reg.Register("nt.adv.modular_sqrt", &modularSqrtGen{})
	reg.Register("nt.adv.two_squares", &twoSquaresGen{})
	reg.Register("nt.adv.pell", &pellGen{})
	reg.Register("nt.adv.pisano", &pisanoGen{})
	reg.Register("nt.adv.carmichael_function", &carmichaelFunctionGen{})
	reg.Register("nt.adv.miller_rabin", &millerRabinGen{})
	reg.Register("nt.adv.prime_gap", &primeGapGen{})
	reg.Register("nt.adv.bertrand", &bertrandGen{})
}

// ---------------------------------------------------------------------------
// 1. The Jacobi symbol (numeric)
// ---------------------------------------------------------------------------

type jacobiGen struct{}

func (g *jacobiGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, n   int
		answer string
		reason string
	}
	claims := []claim{
		{a: 2, n: 15, answer: "1", reason: "(2/15) = (2/3)(2/5) = (-1)(-1) = 1."},
		{a: 7, n: 15, answer: "-1", reason: "(7/15) = (7/3)(7/5) = (1/3)(2/5) = (1)(-1) = -1."},
		{a: 6, n: 35, answer: "-1", reason: "(6/35) = (2/35)(3/35) = (-1)(1) = -1."},
		{a: 5, n: 21, answer: "1", reason: "(5/21) = (5/3)(5/7) = (2/3)(5/7) = (-1)(-1) = 1."},
		{a: 10, n: 21, answer: "-1", reason: "(10/21) = (2/21)(5/21) = (-1)(1) = -1."},
		{a: -1, n: 15, answer: "-1", reason: "(-1/15) = (-1/3)(-1/5) = (-1)(1) = -1."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "Compute the Jacobi symbol (a/n) for a = " + strconv.Itoa(c.a) +
			" and n = " + strconv.Itoa(c.n) + ". (enter 1 or -1)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"a": strconv.Itoa(c.a), "n": strconv.Itoa(c.n), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 2. Discrete logarithms (numeric)
// ---------------------------------------------------------------------------

type discreteLogGen struct{}

func (g *discreteLogGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		p, g, h int
		answer  string
		reason  string
	}
	claims := []claim{
		{p: 7, g: 3, h: 6, answer: "3", reason: "3^1 = 3, 3^2 = 2, 3^3 = 6 mod 7, so x = 3."},
		{p: 11, g: 2, h: 9, answer: "6", reason: "The powers of 2 mod 11 are 2, 4, 8, 5, 10, 9, so 2^6 = 9 and x = 6."},
		{p: 13, g: 2, h: 11, answer: "7", reason: "The powers of 2 mod 13 reach 11 at the seventh step, so x = 7."},
		{p: 17, g: 3, h: 5, answer: "5", reason: "The powers of 3 mod 17 are 3, 9, 10, 13, 5, so 3^5 = 5 and x = 5."},
		{p: 7, g: 5, h: 4, answer: "2", reason: "5^1 = 5 and 5^2 = 25 = 4 mod 7, so x = 2."},
		{p: 11, g: 8, h: 5, answer: "8", reason: "The powers of 8 mod 11 reach 5 at the eighth step, so x = 8."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "Modulo " + strconv.Itoa(c.p) + ", the element " + strconv.Itoa(c.g) +
			" is a primitive root. Find x with " + strconv.Itoa(c.g) + "^x = " + strconv.Itoa(c.h) +
			" (mod " + strconv.Itoa(c.p) + "), where 1 <= x <= " + strconv.Itoa(c.p-1) +
			". (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts: map[string]string{
			"p": strconv.Itoa(c.p), "g": strconv.Itoa(c.g), "h": strconv.Itoa(c.h), "answer": c.answer,
		},
	}
}

// ---------------------------------------------------------------------------
// 3. Square roots modulo a prime (numeric)
// ---------------------------------------------------------------------------

type modularSqrtGen struct{}

func (g *modularSqrtGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		a, p   int
		answer string
		reason string
	}
	claims := []claim{
		{a: 4, p: 7, answer: "2", reason: "2^2 = 4 and 5^2 = 25 = 4 mod 7; the smaller root is 2."},
		{a: 2, p: 7, answer: "3", reason: "3^2 = 9 = 2 and 4^2 = 16 = 2 mod 7; the smaller root is 3."},
		{a: 9, p: 11, answer: "3", reason: "3^2 = 9 and 8^2 = 64 = 9 mod 11; the smaller root is 3."},
		{a: 5, p: 11, answer: "4", reason: "4^2 = 16 = 5 and 7^2 = 49 = 5 mod 11; the smaller root is 4."},
		{a: 3, p: 13, answer: "4", reason: "4^2 = 16 = 3 and 9^2 = 81 = 3 mod 13; the smaller root is 4."},
		{a: 10, p: 13, answer: "6", reason: "6^2 = 36 = 10 and 7^2 = 49 = 10 mod 13; the smaller root is 6."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "Solve x^2 = " + strconv.Itoa(c.a) + " (mod " + strconv.Itoa(c.p) +
			") for the smaller root x in 1.." + strconv.Itoa(c.p-1) + ". (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"a": strconv.Itoa(c.a), "p": strconv.Itoa(c.p), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 4. Sums of two squares (numeric)
// ---------------------------------------------------------------------------

type twoSquaresGen struct{}

func (g *twoSquaresGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		n      int
		answer string
		reason string
	}
	claims := []claim{
		{n: 5, answer: "8", reason: "The pairs are (±1, ±2) and (±2, ±1): eight in all."},
		{n: 10, answer: "8", reason: "The pairs are (±1, ±3) and (±3, ±1): eight in all."},
		{n: 25, answer: "12", reason: "The pairs are (±5, 0), (0, ±5), (±3, ±4) and (±4, ±3): twelve in all."},
		{n: 13, answer: "8", reason: "The pairs are (±2, ±3) and (±3, ±2): eight in all."},
		{n: 9, answer: "4", reason: "The pairs are (±3, 0) and (0, ±3): four in all."},
		{n: 3, answer: "0", reason: "No integer pair has x^2 + y^2 = 3, so there are no representations."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "How many integer pairs (x, y), counting signs and order, satisfy x^2 + y^2 = " +
			strconv.Itoa(c.n) + "? (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"n": strconv.Itoa(c.n), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 5. Pell's equation (numeric)
// ---------------------------------------------------------------------------

type pellGen struct{}

func (g *pellGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		d      int
		answer string
		reason string
	}
	claims := []claim{
		{d: 2, answer: "3", reason: "3^2 - 2*2^2 = 9 - 8 = 1, and no smaller positive x works."},
		{d: 3, answer: "2", reason: "2^2 - 3*1^2 = 4 - 3 = 1, so the smallest x is 2."},
		{d: 5, answer: "9", reason: "9^2 - 5*4^2 = 81 - 80 = 1, and no smaller positive x works."},
		{d: 6, answer: "5", reason: "5^2 - 6*2^2 = 25 - 24 = 1, so the smallest x is 5."},
		{d: 7, answer: "8", reason: "8^2 - 7*3^2 = 64 - 63 = 1, so the smallest x is 8."},
		{d: 8, answer: "3", reason: "3^2 - 8*1^2 = 9 - 8 = 1, so the smallest x is 3."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "Find the smallest positive integer x for which Pell's equation x^2 - " + strconv.Itoa(c.d) +
			"*y^2 = 1 has a positive integer solution y. (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"d": strconv.Itoa(c.d), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 6. Pisano periods (numeric)
// ---------------------------------------------------------------------------

type pisanoGen struct{}

func (g *pisanoGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		n      int
		answer string
		reason string
	}
	claims := []claim{
		{n: 2, answer: "3", reason: "Modulo 2 the Fibonacci pairs repeat as (0,1), (1,1), (1,0), (0,1), so the period is 3."},
		{n: 3, answer: "8", reason: "Modulo 3 the pair (0,1) first returns after 8 steps, so the period is 8."},
		{n: 4, answer: "6", reason: "Modulo 4 the pair (0,1) first returns after 6 steps."},
		{n: 5, answer: "20", reason: "Modulo 5 the pair (0,1) first returns after 20 steps."},
		{n: 6, answer: "24", reason: "Modulo 6 the period is the least common multiple of the periods modulo 2 and 3, which is 24."},
		{n: 7, answer: "16", reason: "Modulo 7 the pair (0,1) first returns after 16 steps."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "The Fibonacci sequence starts F0 = 0, F1 = 1 and satisfies F(n) = F(n-1) + F(n-2). " +
			"What is its Pisano period modulo " + strconv.Itoa(c.n) +
			", the smallest k > 0 with F(k) = 0 and F(k+1) = 1 mod " + strconv.Itoa(c.n) +
			"? (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"n": strconv.Itoa(c.n), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 7. The Carmichael function (numeric)
// ---------------------------------------------------------------------------

type carmichaelFunctionGen struct{}

func (g *carmichaelFunctionGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		n      int
		answer string
		reason string
	}
	claims := []claim{
		{n: 8, answer: "2", reason: "For 2^3 with exponent at least 3, lambda is 2^(3-2) = 2."},
		{n: 9, answer: "6", reason: "lambda(9) = phi(9) = 6, since 9 is an odd prime power."},
		{n: 10, answer: "4", reason: "lambda(10) = lcm(lambda(2), lambda(5)) = lcm(1, 4) = 4."},
		{n: 12, answer: "2", reason: "lambda(12) = lcm(lambda(4), lambda(3)) = lcm(2, 2) = 2."},
		{n: 15, answer: "4", reason: "lambda(15) = lcm(lambda(3), lambda(5)) = lcm(2, 4) = 4."},
		{n: 16, answer: "4", reason: "For 2^4 with exponent at least 3, lambda is 2^(4-2) = 4."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "What is the Carmichael function lambda(" + strconv.Itoa(c.n) +
			"), the exponent of the multiplicative group modulo " + strconv.Itoa(c.n) + "? (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"n": strconv.Itoa(c.n), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 8. The Miller-Rabin test
// ---------------------------------------------------------------------------

type millerRabinGen struct{}

func (g *millerRabinGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "The Miller-Rabin test is based on which congruence?",
			correct:     "a^(n-1) = 1 mod n together with the square-root-of-1 structure",
			distractors: []string{"n divides (n-1)! + 1", "a^2 = 1 mod n for every a", "n divides 2^n - 2"},
			reason:      "The test strengthens Fermat's congruence by examining the repeated square roots of 1 that appear on the way to a^(n-1).",
		},
		{
			question:    "A strong probable prime to base a...",
			correct:     "passes the Miller-Rabin test for that base",
			distractors: []string{"is necessarily prime", "fails Fermat's test", "is a Carmichael number"},
			reason:      "Passing one base is evidence of primality, not proof: a strong pseudoprime can pass a chosen base.",
		},
		{
			question:    "The Miller-Rabin error probability is reduced by...",
			correct:     "testing more independent bases",
			distractors: []string{"using a larger modulus only", "repeating the same base", "checking divisibility by 2"},
			reason:      "Each independent base cuts the false-positive probability, and a few dozen bases make the error astronomically small.",
		},
		{
			question:    "A composite that passes Miller-Rabin for a base is called...",
			correct:     "a strong pseudoprime to that base",
			distractors: []string{"a Carmichael prime", "a Mersenne prime", "a Fermat prime"},
			reason:      "Strong pseudoprimes are the composites that survive the strong test for a particular base.",
		},
		{
			question:    "Why does Miller-Rabin beat trial division for large n?",
			correct:     "it runs in polynomial time in log n",
			distractors: []string{"it factors n completely", "it proves primality with certainty", "it needs no modular arithmetic"},
			reason:      "Trial division is exponential in the number of digits; Miller-Rabin uses fast modular exponentiation and is polynomial in log n.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 9. Prime gaps and the next prime (numeric)
// ---------------------------------------------------------------------------

type primeGapGen struct{}

func (g *primeGapGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		p        int
		reason   string
	}
	claims := []claim{
		{
			question: "What is the gap between 23 and the next larger prime? (enter a number)",
			answer:   "6", kind: "gap", p: 23,
			reason: "The next prime after 23 is 29, so the gap is 29 - 23 = 6.",
		},
		{
			question: "What is the smallest prime greater than 100? (enter a number)",
			answer:   "101", kind: "next", p: 100,
			reason: "101 has no divisor other than 1 and itself, and every number from 100 down is composite or below the bound.",
		},
		{
			question: "What is the gap between 89 and the next larger prime? (enter a number)",
			answer:   "8", kind: "gap", p: 89,
			reason: "The next prime after 89 is 97, so the gap is 97 - 89 = 8.",
		},
		{
			question: "What is the smallest prime greater than 200? (enter a number)",
			answer:   "211", kind: "next", p: 200,
			reason: "211 is prime, and every number from 201 to 210 has a proper factor.",
		},
		{
			question: "What is the gap between 7 and the next larger prime? (enter a number)",
			answer:   "4", kind: "gap", p: 7,
			reason: "The next prime after 7 is 11, so the gap is 4.",
		},
		{
			question: "What is the smallest prime greater than 50? (enter a number)",
			answer:   "53", kind: "next", p: 50,
			reason: "53 is prime, and 51 = 3*17 and 52 = 4*13 are composite.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question:    c.question,
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"kind": c.kind, "p": strconv.Itoa(c.p), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 10. Bertrand's postulate
// ---------------------------------------------------------------------------

type bertrandGen struct{}

func (g *bertrandGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Bertrand's postulate says that for every integer n > 1 there is...",
			correct:     "a prime p with n < p < 2n",
			distractors: []string{"a prime p with 2n < p < 3n", "a perfect square between n and 2n", "a twin prime pair below n"},
			reason:      "The postulate guarantees at least one prime in every interval (n, 2n) with n > 1.",
		},
		{
			question:    "Which statement is equivalent to Bertrand's postulate?",
			correct:     "the prime counting function satisfies pi(2n) > pi(n) for n > 1",
			distractors: []string{"pi(n) grows like n/log n", "there are infinitely many twin primes", "every even number is a sum of two primes"},
			reason:      "A prime strictly between n and 2n is exactly what makes the count at 2n exceed the count at n.",
		},
		{
			question:    "Bertrand's postulate guarantees a prime between...",
			correct:     "100 and 200",
			distractors: []string{"200 and 300", "1000 and 1001", "1 and 2"},
			reason:      "Apply the postulate with n = 100: some prime lies strictly between 100 and 200.",
		},
		{
			question:    "Bertrand's postulate is weaker than which result?",
			correct:     "the prime number theorem",
			distractors: []string{"the fundamental theorem of arithmetic", "Euler's theorem", "the sieve of Eratosthenes"},
			reason:      "The prime number theorem gives the asymptotic density of the primes, from which the existence of a prime in (n, 2n) follows for large n.",
		},
		{
			question:    "Who first proved Bertrand's postulate?",
			correct:     "Chebyshev",
			distractors: []string{"Euler", "Fermat", "Gauss"},
			reason:      "Chebyshev proved the postulate in 1850; Erdős later gave an elementary proof.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}
