package graphtheory

// The first graph-theory batch: twenty concepts that open the domain with
// isomorphism, basic constructions, connectivity, matching and coloring.
//
// The discrete-math domain already has graphs, paths, trees, Eulerian and
// Hamiltonian graphs, coloring, bipartite matchings, planarity, Cayley and
// degree sequences. This domain starts after those and goes deeper: it treats
// isomorphism and homomorphism as structure-preserving maps, counts distances
// and cuts, proves the matching/cover min-max pair, and develops chromatic and
// list coloring. The second batch will add planarity, spectral graph theory,
// extremal and random graphs, and flows.
//
// Ten concepts are numeric and ten are multiple choice. Every numeric
// generator publishes the graph parameters its answer follows from, so the
// tests recompute the answer from those facts.

import (
	"math/rand"
	"strconv"

	"github.com/aserjam100/orchid/internal/generator"
)

// ---------------------------------------------------------------------------
// 1. Graph isomorphism
// ---------------------------------------------------------------------------

type isomorphismGen struct{}

func (g *isomorphismGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Which quantity is an isomorphism invariant of a graph?",
			correct:     "the multiset of vertex degrees",
			distractors: []string{"the drawing of the graph", "the names of the vertices", "the order in which the edges are listed"},
			reason:      "An isomorphism permutes vertices, so each vertex keeps its degree and the degree multiset is unchanged.",
		},
		{
			question:    "Two graphs are isomorphic when...",
			correct:     "a bijection of the vertices preserves adjacency and non-adjacency",
			distractors: []string{"they have the same number of vertices", "they have the same number of edges", "they are drawn the same way"},
			reason:      "Matching vertex and edge counts are necessary but not sufficient; the bijection must preserve the whole adjacency relation.",
		},
		{
			question:    "Graphs with degree sequences (2,2,2,2) and (3,3,1,1) are...",
			correct:     "not isomorphic",
			distractors: []string{"isomorphic", "both trees", "both disconnected"},
			reason:      "The degree multisets differ, and the degree multiset is an isomorphism invariant.",
		},
		{
			question:    "A graph and its complement can be...",
			correct:     "isomorphic or not, depending on the graph",
			distractors: []string{"never isomorphic", "always isomorphic", "isomorphic only when complete"},
			reason:      "Self-complementary graphs exist (P_4 is the smallest), and many graphs are not self-complementary.",
		},
		{
			question:    "Isomorphism is an equivalence relation because...",
			correct:     "it is reflexive, symmetric and transitive",
			distractors: []string{"it preserves degrees", "it is a bijection", "it is a graph property"},
			reason:      "Identity, inverse and composition of adjacency-preserving bijections give the three relation laws.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 2. The complement of a graph (numeric)
// ---------------------------------------------------------------------------

type complementGen struct{}

func (g *complementGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		n, m   int
		answer string
		reason string
	}
	claims := []claim{
		{n: 5, m: 4, answer: "6", reason: "A graph on 5 vertices has 10 possible edges, so the complement has 10 - 4 = 6."},
		{n: 6, m: 7, answer: "8", reason: "A graph on 6 vertices has 15 possible edges, so the complement has 15 - 7 = 8."},
		{n: 4, m: 3, answer: "3", reason: "A graph on 4 vertices has 6 possible edges, so the complement has 6 - 3 = 3."},
		{n: 5, m: 8, answer: "2", reason: "A graph on 5 vertices has 10 possible edges, so the complement has 10 - 8 = 2."},
		{n: 6, m: 10, answer: "5", reason: "A graph on 6 vertices has 15 possible edges, so the complement has 15 - 10 = 5."},
		{n: 7, m: 12, answer: "9", reason: "A graph on 7 vertices has 21 possible edges, so the complement has 21 - 12 = 9."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "A simple graph has " + strconv.Itoa(c.n) + " vertices and " + strconv.Itoa(c.m) +
			" edges. How many edges does its complement have? (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"n": strconv.Itoa(c.n), "m": strconv.Itoa(c.m), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 3. Subgraphs and induced subgraphs
// ---------------------------------------------------------------------------

type subgraphGen struct{}

func (g *subgraphGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "An induced subgraph on a vertex set S contains...",
			correct:     "all edges of G with both ends in S",
			distractors: []string{"all edges of G", "only edges chosen by S", "a spanning tree of G"},
			reason:      "Induction keeps every edge whose two endpoints survive, and removes every edge that touches a deleted vertex.",
		},
		{
			question:    "A spanning subgraph has...",
			correct:     "the same vertex set as G",
			distractors: []string{"the same edge set", "one vertex removed", "no cycles"},
			reason:      "Spanning means every vertex is kept; the edge set is allowed to shrink.",
		},
		{
			question:    "Deleting a vertex from a graph...",
			correct:     "deletes the vertex and every incident edge",
			distractors: []string{"deletes only the vertex", "deletes only the incident edges", "keeps all edges"},
			reason:      "The remaining object must still be a graph, so edges with a missing endpoint cannot stay.",
		},
		{
			question:    "Which graph is a subgraph of every graph?",
			correct:     "the graph on the same vertices with no edges",
			distractors: []string{"a spanning tree", "a Hamiltonian cycle", "the complement"},
			reason:      "The empty-edge graph is contained in every simple graph; the others need not exist.",
		},
		{
			question:    "An induced subgraph is determined by...",
			correct:     "its vertex set",
			distractors: []string{"its edge set", "a spanning tree", "its complement"},
			reason:      "Once the vertex set is fixed, the induced subgraph keeps exactly the edges inside it.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 4. Graph homomorphisms
// ---------------------------------------------------------------------------

type homomorphismGen struct{}

func (g *homomorphismGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "A graph homomorphism f: G -> H maps...",
			correct:     "every edge of G to an edge of H",
			distractors: []string{"every edge of H to an edge of G", "vertices to non-adjacent vertices", "only the degrees"},
			reason:      "The defining condition is that adjacency in G implies adjacency of the images in H.",
		},
		{
			question:    "A proper k-coloring is a homomorphism to...",
			correct:     "K_k",
			distractors: []string{"the path P_k", "the cycle C_k", "the complement of K_k"},
			reason:      "The color classes are the preimages of the k vertices of K_k, and adjacent vertices must get different colors.",
		},
		{
			question:    "A graph isomorphism is a homomorphism that is...",
			correct:     "bijective with a homomorphism inverse",
			distractors: []string{"injective only", "surjective only", "degree-preserving only"},
			reason:      "An isomorphism must preserve adjacency in both directions, which is exactly a bijective homomorphism whose inverse is a homomorphism.",
		},
		{
			question:    "A homomorphism to K_2 is...",
			correct:     "a proper 2-coloring",
			distractors: []string{"a spanning tree", "an Eulerian trail", "a matching"},
			reason:      "Two colors, and every edge forces different colors: the definition of a proper 2-coloring.",
		},
		{
			question:    "Homomorphisms are more flexible than isomorphisms because...",
			correct:     "they need not be injective",
			distractors: []string{"they preserve non-adjacency", "they preserve degrees", "they are always bijections"},
			reason:      "Several vertices may map to the same target, which is what allows a coloring to collapse a large graph onto a small one.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 5. Graph products (numeric)
// ---------------------------------------------------------------------------

type productGen struct{}

func (g *productGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		n1, m1 int
		n2, m2 int
		answer string
		reason string
	}
	claims := []claim{
		{n1: 3, m1: 2, n2: 3, m2: 2, answer: "12", reason: "The Cartesian product has m1 n2 + n1 m2 = 2*3 + 3*2 = 12 edges."},
		{n1: 4, m1: 4, n2: 2, m2: 1, answer: "12", reason: "m1 n2 + n1 m2 = 4*2 + 4*1 = 12."},
		{n1: 3, m1: 3, n2: 2, m2: 1, answer: "9", reason: "m1 n2 + n1 m2 = 3*2 + 3*1 = 9."},
		{n1: 4, m1: 3, n2: 2, m2: 1, answer: "10", reason: "m1 n2 + n1 m2 = 3*2 + 4*1 = 10."},
		{n1: 3, m1: 3, n2: 3, m2: 3, answer: "18", reason: "m1 n2 + n1 m2 = 3*3 + 3*3 = 18."},
		{n1: 4, m1: 6, n2: 2, m2: 1, answer: "16", reason: "m1 n2 + n1 m2 = 6*2 + 4*1 = 16."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question: "Graphs G and H have (" + strconv.Itoa(c.n1) + " vertices, " + strconv.Itoa(c.m1) +
			" edges) and (" + strconv.Itoa(c.n2) + " vertices, " + strconv.Itoa(c.m2) +
			" edges). How many edges does their Cartesian product have? (enter a number)",
		Answer:      c.answer,
		Explanation: c.reason,
		Facts: map[string]string{
			"n1": strconv.Itoa(c.n1), "m1": strconv.Itoa(c.m1),
			"n2": strconv.Itoa(c.n2), "m2": strconv.Itoa(c.m2), "answer": c.answer,
		},
	}
}

// ---------------------------------------------------------------------------
// 6. Distance and diameter (numeric)
// ---------------------------------------------------------------------------

type distanceGen struct{}

func (g *distanceGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n        int
		rows     int
		cols     int
		reason   string
	}
	claims := []claim{
		{question: "What is the diameter of the path P_5? (enter a number)", answer: "4", kind: "path", n: 5, reason: "The endpoints are four edges apart, and no pair is farther."},
		{question: "What is the diameter of the cycle C_6? (enter a number)", answer: "3", kind: "cycle", n: 6, reason: "The farthest pair is three steps apart around the cycle."},
		{question: "What is the diameter of the complete graph K_5? (enter a number)", answer: "1", kind: "complete", n: 5, reason: "Every pair of distinct vertices is adjacent, so every distance is 1."},
		{question: "What is the diameter of the path P_7? (enter a number)", answer: "6", kind: "path", n: 7, reason: "The endpoints are six edges apart."},
		{question: "What is the diameter of the cycle C_8? (enter a number)", answer: "4", kind: "cycle", n: 8, reason: "The farthest pair is half the cycle apart."},
		{question: "What is the diameter of the 3-by-4 grid graph? (enter a number)", answer: "5", kind: "grid", rows: 3, cols: 4, reason: "Opposite corners are (3-1) + (4-1) = 5 steps apart."},
	}
	c := claims[rand.Intn(len(claims))]
	facts := map[string]string{"kind": c.kind, "answer": c.answer}
	if c.kind == "grid" {
		facts["rows"] = strconv.Itoa(c.rows)
		facts["cols"] = strconv.Itoa(c.cols)
	} else {
		facts["n"] = strconv.Itoa(c.n)
	}
	return generator.Problem{Question: c.question, Answer: c.answer, Explanation: c.reason, Facts: facts}
}

// ---------------------------------------------------------------------------
// 7. Cut vertices (numeric)
// ---------------------------------------------------------------------------

type cutVertexGen struct{}

func (g *cutVertexGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n        int
		reason   string
	}
	claims := []claim{
		{question: "How many cut vertices does the path P_4 have? (enter a number)", answer: "2", kind: "path", n: 4, reason: "The two internal vertices disconnect the path when removed."},
		{question: "How many cut vertices does the path P_5 have? (enter a number)", answer: "3", kind: "path", n: 5, reason: "All three internal vertices are cut vertices."},
		{question: "How many cut vertices does the star K_{1,4} have? (enter a number)", answer: "1", kind: "star", n: 4, reason: "Removing the centre isolates all four leaves; no leaf is a cut vertex."},
		{question: "How many cut vertices does the cycle C_5 have? (enter a number)", answer: "0", kind: "cycle", n: 5, reason: "A cycle stays connected after any single vertex is removed."},
		{question: "How many cut vertices does a path on 6 vertices have? (enter a number)", answer: "4", kind: "path", n: 6, reason: "Every internal vertex is a cut vertex, and there are n - 2 = 4."},
		{question: "How many cut vertices does the bowtie graph (two triangles sharing one vertex) have? (enter a number)", answer: "1", kind: "bowtie", reason: "The shared vertex disconnects the two triangles; no other vertex does."},
	}
	c := claims[rand.Intn(len(claims))]
	facts := map[string]string{"kind": c.kind, "answer": c.answer}
	if c.n > 0 {
		facts["n"] = strconv.Itoa(c.n)
	}
	return generator.Problem{Question: c.question, Answer: c.answer, Explanation: c.reason, Facts: facts}
}

// ---------------------------------------------------------------------------
// 8. Bridges (numeric)
// ---------------------------------------------------------------------------

type bridgeGen struct{}

func (g *bridgeGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n        int
		reason   string
	}
	claims := []claim{
		{question: "How many bridges does the path P_4 have? (enter a number)", answer: "3", kind: "path", n: 4, reason: "Every edge of a path is a bridge."},
		{question: "How many bridges does the cycle C_5 have? (enter a number)", answer: "0", kind: "cycle", n: 5, reason: "Every edge of a cycle lies on a cycle, so none is a bridge."},
		{question: "How many bridges does the star K_{1,4} have? (enter a number)", answer: "4", kind: "star", n: 4, reason: "Each leaf edge is a bridge."},
		{question: "How many bridges does a tree on 6 vertices have? (enter a number)", answer: "5", kind: "tree", n: 6, reason: "Every edge of a tree is a bridge, and a tree on n vertices has n - 1 edges."},
		{question: "How many bridges does a graph made of two triangles joined by a single edge have? (enter a number)", answer: "1", kind: "join", reason: "Only the joining edge disconnects the graph when removed."},
		{question: "How many bridges does a square with one diagonal have? (enter a number)", answer: "0", kind: "chord", reason: "Every edge lies on a cycle, so none is a bridge."},
	}
	c := claims[rand.Intn(len(claims))]
	facts := map[string]string{"kind": c.kind, "answer": c.answer}
	if c.n > 0 {
		facts["n"] = strconv.Itoa(c.n)
	}
	return generator.Problem{Question: c.question, Answer: c.answer, Explanation: c.reason, Facts: facts}
}

// ---------------------------------------------------------------------------
// 9. Vertex and edge connectivity (numeric)
// ---------------------------------------------------------------------------

type connectivityGen struct{}

func (g *connectivityGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		a, b     int
		reason   string
	}
	claims := []claim{
		{question: "What is the vertex connectivity kappa(K_5)? (enter a number)", answer: "4", kind: "complete", a: 5, reason: "Removing any four vertices leaves one vertex, and removing fewer cannot disconnect K_5."},
		{question: "What is the vertex connectivity kappa(C_6)? (enter a number)", answer: "2", kind: "cycle", a: 6, reason: "Two vertices can disconnect a cycle, and no single vertex can."},
		{question: "What is the vertex connectivity kappa(P_5)? (enter a number)", answer: "1", kind: "path", a: 5, reason: "Any internal vertex disconnects a path."},
		{question: "What is the vertex connectivity kappa(K_{3,3})? (enter a number)", answer: "3", kind: "bipartite", a: 3, b: 3, reason: "Removing one side disconnects the graph, and no smaller set does."},
		{question: "What is the vertex connectivity kappa(K_{2,4})? (enter a number)", answer: "2", kind: "bipartite", a: 2, b: 4, reason: "Removing the two vertices of the smaller side disconnects the graph."},
		{question: "What is the vertex connectivity kappa of two disjoint cycles? (enter a number)", answer: "0", kind: "disjoint", reason: "A disconnected graph has vertex connectivity 0."},
	}
	c := claims[rand.Intn(len(claims))]
	facts := map[string]string{"kind": c.kind, "answer": c.answer}
	if c.a > 0 {
		facts["a"] = strconv.Itoa(c.a)
	}
	if c.b > 0 {
		facts["b"] = strconv.Itoa(c.b)
	}
	return generator.Problem{Question: c.question, Answer: c.answer, Explanation: c.reason, Facts: facts}
}

// ---------------------------------------------------------------------------
// 10. Menger's theorem
// ---------------------------------------------------------------------------

type mengerGen struct{}

func (g *mengerGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Menger's theorem equates the maximum number of internally vertex-disjoint u-v paths with...",
			correct:     "the minimum size of a vertex cut separating u and v",
			distractors: []string{"the maximum degree", "the number of edges", "the diameter"},
			reason:      "The theorem is a min-max duality: the largest family of disjoint paths is limited by, and equals, the smallest separating cut.",
		},
		{
			question:    "The edge version of Menger's theorem uses...",
			correct:     "edge-disjoint paths and edge cuts",
			distractors: []string{"vertex cuts only", "matchings", "spanning trees"},
			reason:      "Replacing internally disjoint paths by edge-disjoint paths and vertex cuts by edge cuts gives the edge form.",
		},
		{
			question:    "Menger's theorem implies which inequality?",
			correct:     "kappa(G) <= lambda(G) <= delta(G)",
			distractors: []string{"lambda(G) <= kappa(G) <= delta(G)", "delta(G) <= kappa(G) <= lambda(G)", "kappa(G) = delta(G)"},
			reason:      "A vertex cut can be made from the edges at a vertex, and an edge cut can be made from the edges at a minimum-degree vertex.",
		},
		{
			question:    "A graph is k-connected exactly when...",
			correct:     "every pair of vertices has k internally disjoint paths",
			distractors: []string{"every vertex has degree at least k", "it has k edges", "it has no cut vertices"},
			reason:      "That is the path form of k-connectivity, equivalent to the minimum vertex cut being at least k.",
		},
		{
			question:    "Menger's theorem is called a min-max theorem because...",
			correct:     "a maximum packing equals a minimum cut",
			distractors: []string{"it counts vertices", "it uses Euler's formula", "it applies only to trees"},
			reason:      "Packing paths and cutting them are dual optimisation problems with the same optimum.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 11. Matchings (numeric)
// ---------------------------------------------------------------------------

type matchingGen struct{}

func (g *matchingGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n        int
		reason   string
	}
	claims := []claim{
		{question: "What is the maximum matching size in the path P_4? (enter a number)", answer: "2", kind: "path", n: 4, reason: "Two disjoint edges cover all four vertices."},
		{question: "What is the maximum matching size in the path P_5? (enter a number)", answer: "2", kind: "path", n: 5, reason: "Two disjoint edges is the best possible; one vertex must be left out."},
		{question: "What is the maximum matching size in the cycle C_6? (enter a number)", answer: "3", kind: "cycle", n: 6, reason: "Alternate edges around the cycle give three disjoint edges."},
		{question: "What is the maximum matching size in the star K_{1,5}? (enter a number)", answer: "1", kind: "star", n: 5, reason: "Every edge shares the centre, so no two are disjoint."},
		{question: "What is the maximum matching size in the cycle C_5? (enter a number)", answer: "2", kind: "cycle", n: 5, reason: "Two disjoint edges is the maximum; an odd cycle cannot be perfectly matched."},
		{question: "What is the maximum matching size in the path P_6? (enter a number)", answer: "3", kind: "path", n: 6, reason: "Three disjoint edges cover all six vertices."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question:    c.question,
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"kind": c.kind, "n": strconv.Itoa(c.n), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 12. Perfect matchings
// ---------------------------------------------------------------------------

type perfectMatchingGen struct{}

func (g *perfectMatchingGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "A perfect matching covers...",
			correct:     "every vertex exactly once",
			distractors: []string{"every edge exactly once", "every vertex at least twice", "only the leaves"},
			reason:      "A matching uses disjoint edges; perfect means the disjoint edges touch every vertex.",
		},
		{
			question:    "Which graph cannot have a perfect matching?",
			correct:     "a path with an odd number of vertices",
			distractors: []string{"a cycle with an even number of vertices", "a complete graph with an even number of vertices", "a star with an even number of leaves"},
			reason:      "Every perfect matching covers vertices in pairs, so an odd vertex count is impossible.",
		},
		{
			question:    "Every regular bipartite graph...",
			correct:     "has a perfect matching",
			distractors: []string{"has no perfect matching", "is a tree", "is complete"},
			reason:      "Hall's condition follows from regularity, so a perfect matching exists.",
		},
		{
			question:    "A graph with a perfect matching must have...",
			correct:     "an even number of vertices",
			distractors: []string{"an odd number of vertices", "a cut vertex", "a bridge"},
			reason:      "The matching pairs the vertices, so their number is twice the number of matched edges.",
		},
		{
			question:    "A perfect matching is also...",
			correct:     "a maximum matching",
			distractors: []string{"a minimum vertex cover", "a spanning tree", "an Eulerian circuit"},
			reason:      "No matching can use more edges than one that already covers every vertex.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 13. Tutte's theorem
// ---------------------------------------------------------------------------

type tutteGen struct{}

func (g *tutteGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Tutte's theorem characterizes...",
			correct:     "graphs with a perfect matching",
			distractors: []string{"graphs with an Eulerian circuit", "planar graphs", "bipartite graphs"},
			reason:      "Tutte gives a necessary and sufficient condition for the existence of a perfect matching.",
		},
		{
			question:    "Tutte's condition counts...",
			correct:     "odd components of G - S for every vertex set S",
			distractors: []string{"even components", "edges", "cycles"},
			reason:      "The condition is that every S leaves at most |S| odd components.",
		},
		{
			question:    "If |V| is odd, Tutte fails at S = empty because...",
			correct:     "the single component is odd and larger than |S|",
			distractors: []string{"there are no components", "the graph is disconnected", "the matching is perfect"},
			reason:      "An odd number of vertices cannot be perfectly matched, and the condition detects it with the empty set.",
		},
		{
			question:    "Tutte's theorem generalizes...",
			correct:     "Hall's marriage theorem",
			distractors: []string{"Euler's formula", "Brooks' theorem", "Menger's theorem"},
			reason:      "Hall handles bipartite graphs; Tutte removes the bipartite restriction.",
		},
		{
			question:    "A bridge in a graph with a perfect matching...",
			correct:     "can exist, but each side must contain an odd number of vertices",
			distractors: []string{"cannot exist", "forces the graph to be disconnected", "is a perfect matching"},
			reason:      "The matching must use the bridge itself, and then each side has one unmatched vertex.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 14. Vertex covers (numeric)
// ---------------------------------------------------------------------------

type vertexCoverGen struct{}

func (g *vertexCoverGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		a, b     int
		reason   string
	}
	claims := []claim{
		{question: "What is the minimum vertex cover size of the path P_4? (enter a number)", answer: "2", kind: "path", a: 4, reason: "The two internal vertices cover every edge."},
		{question: "What is the minimum vertex cover size of the cycle C_6? (enter a number)", answer: "3", kind: "cycle", a: 6, reason: "Alternate vertices cover all six edges, and two cannot."},
		{question: "What is the minimum vertex cover size of the star K_{1,5}? (enter a number)", answer: "1", kind: "star", a: 5, reason: "The centre covers every edge."},
		{question: "What is the minimum vertex cover size of K_{2,3}? (enter a number)", answer: "2", kind: "bipartite", a: 2, b: 3, reason: "Koenig's theorem gives the smaller side, 2."},
		{question: "What is the minimum vertex cover size of the path P_5? (enter a number)", answer: "2", kind: "path", a: 5, reason: "Two well-placed internal vertices cover all four edges."},
		{question: "What is the minimum vertex cover size of K_{3,3}? (enter a number)", answer: "3", kind: "bipartite", a: 3, b: 3, reason: "Koenig's theorem gives 3, one entire side."},
	}
	c := claims[rand.Intn(len(claims))]
	facts := map[string]string{"kind": c.kind, "answer": c.answer}
	if c.a > 0 {
		facts["a"] = strconv.Itoa(c.a)
	}
	if c.b > 0 {
		facts["b"] = strconv.Itoa(c.b)
	}
	return generator.Problem{Question: c.question, Answer: c.answer, Explanation: c.reason, Facts: facts}
}

// ---------------------------------------------------------------------------
// 15. Independent sets (numeric)
// ---------------------------------------------------------------------------

type independentSetGen struct{}

func (g *independentSetGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		a, b     int
		reason   string
	}
	claims := []claim{
		{question: "What is the maximum independent set size of the path P_4? (enter a number)", answer: "2", kind: "path", a: 4, reason: "The two endpoints and the two internal vertices split into two independent pairs."},
		{question: "What is the maximum independent set size of the cycle C_6? (enter a number)", answer: "3", kind: "cycle", a: 6, reason: "Alternate vertices are independent and there are three."},
		{question: "What is the maximum independent set size of the star K_{1,5}? (enter a number)", answer: "5", kind: "star", a: 5, reason: "All five leaves are pairwise non-adjacent."},
		{question: "What is the maximum independent set size of K_{2,3}? (enter a number)", answer: "3", kind: "bipartite", a: 2, b: 3, reason: "The larger side of a complete bipartite graph is independent."},
		{question: "What is the maximum independent set size of the path P_5? (enter a number)", answer: "3", kind: "path", a: 5, reason: "Vertices 1, 3 and 5 are pairwise non-adjacent."},
		{question: "What is the maximum independent set size of the cycle C_5? (enter a number)", answer: "2", kind: "cycle", a: 5, reason: "Two non-adjacent vertices is the maximum in an odd cycle."},
	}
	c := claims[rand.Intn(len(claims))]
	facts := map[string]string{"kind": c.kind, "answer": c.answer}
	if c.a > 0 {
		facts["a"] = strconv.Itoa(c.a)
	}
	if c.b > 0 {
		facts["b"] = strconv.Itoa(c.b)
	}
	return generator.Problem{Question: c.question, Answer: c.answer, Explanation: c.reason, Facts: facts}
}

// ---------------------------------------------------------------------------
// 16. The chromatic polynomial (numeric)
// ---------------------------------------------------------------------------

type chromaticPolynomialGen struct{}

func (g *chromaticPolynomialGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n, k     int
		reason   string
	}
	claims := []claim{
		{question: "What is P(P_4, 3), the chromatic polynomial of the path P_4 evaluated at 3? (enter a number)", answer: "24", kind: "path", n: 4, k: 3, reason: "P(P_n,k) = k(k-1)^(n-1), so 3*2^3 = 24."},
		{question: "What is P(C_4, 3), the chromatic polynomial of the cycle C_4 evaluated at 3? (enter a number)", answer: "18", kind: "cycle", n: 4, k: 3, reason: "P(C_n,k) = (k-1)^n + (-1)^n(k-1), so 2^4 + 2 = 18."},
		{question: "What is P(K_3, 4), the chromatic polynomial of K_3 evaluated at 4? (enter a number)", answer: "24", kind: "complete", n: 3, k: 4, reason: "The colors must be distinct, so 4*3*2 = 24."},
		{question: "What is P(T_5, 2), the chromatic polynomial of a tree on 5 vertices evaluated at 2? (enter a number)", answer: "2", kind: "tree", n: 5, k: 2, reason: "A tree is bipartite, so with two colors the first vertex has 2 choices and each later vertex has 1."},
		{question: "What is P(C_3, 3), the chromatic polynomial of the triangle evaluated at 3? (enter a number)", answer: "6", kind: "cycle", n: 3, k: 3, reason: "P(C_3,k) = (k-1)^3 - (k-1), so 8 - 2 = 6."},
		{question: "What is P(P_3, 2), the chromatic polynomial of the path P_3 evaluated at 2? (enter a number)", answer: "2", kind: "path", n: 3, k: 2, reason: "P(P_3,k) = k(k-1)^2, so 2*1 = 2."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question:    c.question,
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"kind": c.kind, "n": strconv.Itoa(c.n), "k": strconv.Itoa(c.k), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 17. Brooks' theorem
// ---------------------------------------------------------------------------

type brooksGen struct{}

func (g *brooksGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "Brooks' theorem bounds the chromatic number by...",
			correct:     "the maximum degree, except for complete graphs and odd cycles",
			distractors: []string{"the minimum degree", "the number of edges", "the clique number"},
			reason:      "The greedy bound is Delta + 1; Brooks removes the extra color except for the two exceptional families.",
		},
		{
			question:    "The exceptions in Brooks' theorem are...",
			correct:     "complete graphs and odd cycles",
			distractors: []string{"trees and paths", "bipartite graphs", "planar graphs"},
			reason:      "K_n needs n = Delta + 1 colors and odd cycles need 3 = Delta + 1.",
		},
		{
			question:    "A connected graph with maximum degree 3 has chromatic number at most...",
			correct:     "3 unless it is K_4",
			distractors: []string{"2", "4 always", "5"},
			reason:      "Brooks gives at most Delta = 3 except when the graph is K_4, which needs 4.",
		},
		{
			question:    "Brooks strengthens which bound?",
			correct:     "the greedy bound chi <= Delta + 1",
			distractors: []string{"chi <= omega", "chi <= n", "chi <= alpha"},
			reason:      "Greedy coloring already gives Delta + 1; Brooks lowers it to Delta off the exceptions.",
		},
		{
			question:    "A connected graph with maximum degree 2 is...",
			correct:     "a path or a cycle",
			distractors: []string{"complete", "a tree with a branch", "bipartite but not a path"},
			reason:      "Every vertex has at most two neighbours, so the component is a path or a cycle.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 18. Edge coloring (numeric)
// ---------------------------------------------------------------------------

type edgeColoringGen struct{}

func (g *edgeColoringGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n        int
		reason   string
	}
	claims := []claim{
		{question: "What is the edge chromatic number of K_4? (enter a number)", answer: "3", kind: "complete", n: 4, reason: "K_n with n even needs n - 1 = 3 colors."},
		{question: "What is the edge chromatic number of K_5? (enter a number)", answer: "5", kind: "complete", n: 5, reason: "K_n with n odd needs n = 5 colors."},
		{question: "What is the edge chromatic number of C_4? (enter a number)", answer: "2", kind: "cycle", n: 4, reason: "An even cycle alternates two colors."},
		{question: "What is the edge chromatic number of C_5? (enter a number)", answer: "3", kind: "cycle", n: 5, reason: "An odd cycle needs three colors."},
		{question: "What is the edge chromatic number of the path P_4? (enter a number)", answer: "2", kind: "path", n: 4, reason: "Adjacent edges must differ, so two colors suffice and one does not."},
		{question: "What is the edge chromatic number of the star K_{1,4}? (enter a number)", answer: "4", kind: "star", n: 4, reason: "All four edges meet at the centre, so all four colors are needed."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question:    c.question,
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"kind": c.kind, "n": strconv.Itoa(c.n), "answer": c.answer},
	}
}

// ---------------------------------------------------------------------------
// 19. Perfect graphs
// ---------------------------------------------------------------------------

type perfectGraphGen struct{}

func (g *perfectGraphGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question    string
		correct     string
		distractors []string
		reason      string
	}
	claims := []claim{
		{
			question:    "A graph is perfect when...",
			correct:     "omega(H) = chi(H) for every induced subgraph H",
			distractors: []string{"chi(H) <= 2", "it is planar", "it is regular"},
			reason:      "The clique number must equal the chromatic number on the graph and on every induced subgraph.",
		},
		{
			question:    "Which graphs are perfect?",
			correct:     "bipartite graphs",
			distractors: []string{"odd cycles of length at least 5", "all graphs", "only complete graphs"},
			reason:      "Bipartite graphs have clique number at most 2 and chromatic number at most 2, and the same holds for induced subgraphs.",
		},
		{
			question:    "An odd cycle of length at least 5 is...",
			correct:     "not perfect",
			distractors: []string{"perfect", "bipartite", "a tree"},
			reason:      "Its clique number is 2 but its chromatic number is 3, so the equality fails.",
		},
		{
			question:    "The strong perfect graph theorem forbids...",
			correct:     "odd holes and odd antiholes",
			distractors: []string{"even cycles", "trees", "complete graphs"},
			reason:      "A graph is perfect exactly when neither it nor its complement contains an induced odd cycle of length at least 5.",
		},
		{
			question:    "Chordal graphs are...",
			correct:     "perfect",
			distractors: []string{"never perfect", "always bipartite", "always planar"},
			reason:      "Chordal graphs have a perfect elimination ordering, which makes clique number and chromatic number agree on every induced subgraph.",
		},
	}
	c := claims[rand.Intn(len(claims))]
	p := generator.MultipleChoice(c.question, c.correct, c.distractors...)
	p.Explanation = c.reason
	return p
}

// ---------------------------------------------------------------------------
// 20. List coloring (numeric)
// ---------------------------------------------------------------------------

type listColoringGen struct{}

func (g *listColoringGen) Generate(ctx generator.GeneratorContext) generator.Problem {
	type claim struct {
		question string
		answer   string
		kind     string
		n        int
		reason   string
	}
	claims := []claim{
		{question: "What is the list chromatic number of C_4? (enter a number)", answer: "2", kind: "cycle", n: 4, reason: "Even cycles are 2-choosable."},
		{question: "What is the list chromatic number of C_5? (enter a number)", answer: "3", kind: "cycle", n: 5, reason: "Odd cycles are not 2-choosable and need three colors."},
		{question: "What is the list chromatic number of the path P_5? (enter a number)", answer: "2", kind: "path", n: 5, reason: "A path is 2-choosable: color from one end, avoiding the previous color."},
		{question: "What is the list chromatic number of C_6? (enter a number)", answer: "2", kind: "cycle", n: 6, reason: "Even cycles are 2-choosable."},
		{question: "What is the list chromatic number of C_7? (enter a number)", answer: "3", kind: "cycle", n: 7, reason: "Odd cycles need three colors even with arbitrary lists."},
		{question: "What is the list chromatic number of the path P_3? (enter a number)", answer: "2", kind: "path", n: 3, reason: "A path with at least one edge is 2-choosable."},
	}
	c := claims[rand.Intn(len(claims))]
	return generator.Problem{
		Question:    c.question,
		Answer:      c.answer,
		Explanation: c.reason,
		Facts:       map[string]string{"kind": c.kind, "n": strconv.Itoa(c.n), "answer": c.answer},
	}
}
