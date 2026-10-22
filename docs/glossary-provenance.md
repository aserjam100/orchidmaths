# Glossary provenance and audit log

This is the evidence behind the glossary's source attributions. [`THIRD_PARTY_NOTICES.md`](../THIRD_PARTY_NOTICES.md) states
each license obligation; this file records the work that established it — which page was opened,
what was verified there, and what was corrected as a result. Keep it: an obligation whose audit
trail is gone cannot be defended or re-verified.

Rounds are newest first. Each round names its batch, the sources opened, and the findings. A
term counts as source-verified only when a claim was checked against a page that was opened, not
against a search result.


## Specialist expansion (282 → 342; six fields, 60 entries)

### Method
Gap inventory per field confirmed near-zero dedicated coverage (only
manifold, isomorphism/homomorphism, tensor-product, metric-space,
tangent-line as related terms). 8/10/10/11/8/14 entries across category
theory, operator theory, homological algebra, algebraic geometry,
representation theory, differential geometry. Deliberately excluded and
recorded: monads/comonads and cartesian-closed categories (no verified
open source this run), trace-class/Hilbert–Schmidt operators
(too specialized for verified sourcing), spectral sequences
(mission-optional), divisors (needs dedicated sourcing).

### Sources opened this round
- SEP Category Theory (functor laws, naturality, adjoints F⊣G,
  equivalence-vs-isomorphism as the notion that matters, limits/colimits
  as adjoints, products/coproducts/pullbacks/pushouts, zero object).
- Wikipedia Chain complex (d²=0, cycles/boundaries, H=ker/im, exact and
  short exact sequences, chain maps/homotopies; CC BY-SA 4.0 footer
  confirmed from the page).
- Wikipedia Self-adjoint operator (adjoint, symmetric-vs-self-adjoint
  domain distinction, bounded case, spectrum) and Spectral theorem
  (normal/unitary diagonalization, compact case, projections).
- Wikipedia Algebraic variety (affine/projective, Zariski, coordinate
  rings, Nullstellensatz, dimension, singularities, scheme
  generalization) and Sheaf (presheaf, locality+gluing, stalks).
- Wikipedia Group representation (ρ:G→GL(V), invariant/irreducible,
  Maschke hypothesis, equivalence) and Schur's lemma (full statement
  incl. scalar self-maps over algebraically closed fields).
- Wikipedia Differentiable manifold (charts/atlas/transitions, smooth
  structure, tangent/cotangent/tensor bundles, Lie/exterior calculus,
  Riemannian structures; all 14 entries' topics confirmed present).
- UT Austin / Hamburg / Columbia homological-algebra notes used only as
  search excerpts for Tor/Ext/resolution structure (recorded as such);
  PlanetMath chaincomplex URL did NOT resolve (returned homepage) and
  was never cited.

### Findings
- Data gate caught two real collisions pre-merge (compact-operator vs
  compact-set symbol; cohomology→differential-form forward reference),
  both fixed (symbol widened; link retargeted then restored).
- All examples hand-checked (rotation eigenvalues ±i; Tor₁(Z/m,Z/n);
  logistic-free elementary verifications throughout).

## Final two deferred terms (280 → 282)

### Sigma-algebra
- Opened https://planetmath.org/sigmaalgebra: full definition (X∈M,
  complements, countable unions), MSC 28A60. License opened separately
  (https://planetmath.org/?op=license): entries under CC BY-SA.
- Entry states all three closure properties and the finite-vs-countable
  distinction in commonMistakes, per the brief. Example is the power set
  on two points (verifiable by hand from the stated axioms).

### Hilbert space
- Opened the LibreTexts Fowler function-spaces page: confirms the
  mathematician's definition (inner-product spaces of normalizable
  functions where convergent sequences have limits). Page license is
  NOT declared, so the entry records exactly that — consultation only,
  no reuse claim.
- Completeness framing cross-checked against UBC metric-completeness
  excerpts and a CalPoly Hilbert-space thesis (search excerpts, recorded
  as such in the entry note, not as opened sources).
- Entry keeps the required triple (inner product, induced norm,
  completeness) in two sentences plus a finite-dimensional example.

## Coverage expansion (229 → 280; batches A–D plus trig-identities)

### Method
Taxonomy of 21 areas (MSC-coded) in `web/next-app/scripts/glossary-coverage.json` with
per-area inventories, entry mapping, and honest statuses (20
verified_complete relative to their documented inventories; 1
gaps_remaining). Existing 229 audited automatically (no dup names, no
same-category dup symbols, all readings present and prefix-free,
definitions substantive, categories valid) — shared-symbol pairs are all
legitimate cross-category distinctions. 51 new entries in 4 batches, every
example hand-recomputed, every exact claim checked against an opened source
or a documented derivation.

### Sources opened this round
- NIST DADS theta.html / omegaCapital.html / bigOnotation.html (formal
  asymptotic definitions; Ω uppercase confirmed; little-omega page
  correctly NOT cited for Big-Omega).
- Britannica aleph-null (body confirmed: cardinality of naturals,
  denumerable, continuum larger). Proprietary; consultation only.
- DLMF 27.2 (π(x) sum, PNT, Möbius, totient table confirms φ(9) = 6).
- MathWorld PiContinuedFraction (π = [3; 7, 15, 1, 292, …], convergents).
- LibreTexts topology 1.1 (axioms; page states Public Domain) and IVC
  differential-equation terminology (ODE/PDE/order/linear/homogeneous;
  CC BY-NC-SA 4.0).
- UBC Rechnitzer/Martens metric-completeness notes excerpts (epsilon-N
  Cauchy, completeness); UChicago May REU topology notes excerpts
  (compact/connected/homeomorphism/Heine-Borel); UCSD Bender ch.10
  excerpts (Fibonacci GF x/(1−x−x²) for F₀=0,F₁=1).
- SEP Turing Machines (full formal definition, Church–Turing thesis).
- Scholarpedia Entropy, eq.5 (H = −Σp log₂p; CC BY-NC-SA 3.0 footer),
  cited via DOI (direct https fails; DOI resolves 200).
- MacTutor Fermat's Last Theorem (statement, Wiles 1993–95, Diophantus
  margin context; © St Andrews, consultation only).
- Encyclopedia of Mathematics rejected (persistent 502s); PlanetMath
  license confirmed (CC BY-SA) but no specific articles opened, so NOT
  cited anywhere.

### Deferred with reasons
- sigma-algebra, hilbert-space: no openly-licensed reference verifiable in
  this run; graduate texts not openly fetchable. Ledger gaps_remaining.
- Big-Theta/Omega note: DADS Ω definition is stronger (for-all-large-n)
  than some textbooks' infinitely-often version — recorded in the entry.

## Coverage batch 2 (26 new entries, 203 → 229)

### Scope
Genuine gaps only; the three still-excluded v1 candidates (standalone nabla,
slash-vs-fraction, per-letter Greek variants, standalone Stirling,
sheaf/scheme) remain excluded as duplicates, by design, or out of band.
Resolved two deferred exclusions with dedicated sources: aleph-null
(Britannica) and big-Theta/big-Omega (NIST DADS formal pages).

### Research log (all pages opened, not snippets)
- https://xlinux.nist.gov/dads/HTML/theta.html (NIST DADS, opened): formal
  Theta definition (c1, c2, k; "upper-case Greek letter Theta"). Supports
  `big-theta` definition and example.
- https://xlinux.nist.gov/dads/HTML/omegaCapital.html (NIST DADS, opened):
  formal Big-Omega definition (uppercase confirmed) plus the note that it is
  stronger than the traditional math definition. Supports `big-omega`
  including its qualification. The lowercase `omega.html` page is
  little-omega and was NOT cited for the entry.
- https://xlinux.nist.gov/dads/HTML/bigOnotation.html (NIST DADS, opened):
  formal O, "equal to expressed by Θ(n)", "capital O, not omicron".
- https://www.britannica.com/science/aleph-null (opened; body confirmed):
  aleph-null = cardinality of the naturals; denumerable/countably infinite;
  continuum c larger. Supports `aleph-null`. Britannica is proprietary, so
  the license string records consultation with no reuse.
- https://dlmf.nist.gov/27.2 (DLMF, opened): 27.2.2 π(x) = sum over p≤x;
  27.2.3 PNT π(x) ~ x/ln x; 27.2.12 Möbius definition; totient table
  confirms φ(9) = 6. Supports `prime-counting` and `mobius-function`.
- https://mathworld.wolfram.com/PiContinuedFraction.html (opened): π =
  [3; 7, 15, 1, 292, …], convergents incl. 22/7. Supports
  `continued-fraction` example. © Wolfram; consultation only.
- DLMF 5.2.3 (previously verified): Euler constant limit. Supports
  `euler-mascheroni`.
- Hand-recomputed: law-of-sines b = 5√2 ≈ 7.07; law-of-cosines c² = 39;
  Bayes ≈ 0.17 (0.0099/0.0594); IBP derivative check; kernel (−2,1);
  SE = 1.5; CI (48.04, 51.96); rotation QᵀQ = I.

### Findings and changes
- `bayes-theorem` first written with symbol P(A|B), colliding with
  `conditional-probability`; the data gate caught it. Symbol changed to the
  full theorem statement. Status: verified.
- All 26 new entries: verified at authoring. No other corrections.

## How a round is recorded

Each batch records: claims checked, sources actually opened (URL + what was
verified there), and findings. "Searched" is not "verified": an entry counts
as source-verified only when a claim was checked against an opened source.
Render-testing is covered separately by the KaTeX gate
(`test/glossaryData.test.ts` "renders every example through KaTeX") and the
Playwright spec; the ledger distinguishes the two.

## Batch: advanced (27 entries)

### Research log (all pages opened, not snippets)
- https://dlmf.nist.gov/5.2 (DLMF gamma definitions): 5.2.1 Euler integral
  for Gamma; analytic-continuation paragraph. Supports `gamma-function`
  (integral definition; integer values follow the standard recurrence).
- https://dlmf.nist.gov/5.12 (DLMF beta function): 5.12.1
  B(a,b) = Gamma(a)Gamma(b)/Gamma(a+b). Exact confirmation of the
  `beta-function` example.
- https://dlmf.nist.gov/7.2 (DLMF error functions): 7.2.1
  erf z = (2/sqrt(pi)) integral_0^z e^{-t^2} dt. Exact confirmation of the
  `error-function` example.
- https://dlmf.nist.gov/10.4 and https://dlmf.nist.gov/10.8 (DLMF Bessel):
  connection formulas and power-series framework (10.8 references the J
  series 10.2.2). J_0(0) = 1 follows from the defining series (k=0 term is
  1/(0! Gamma(1)) = 1); recorded as derivation, not lookup.
- https://mathworld.wolfram.com/BetaFunction.html (opened): on-topic page
  citing A&S section 6.2; equation glyphs did not render in text fetch, so
  the relation claim rests on DLMF 5.12.1 above.
- Number-set, group, isomorphism, congruence, gcd/lcm, totient, tensor,
  direct-sum, argmin, Fourier, Laplace and big-O examples all recomputed by
  hand: phi(9) = 6 (1,2,4,5,7,8); (R,+) iso (R>0,x) via exp; L{f'} formula;
  all correct.

### Findings and changes
- No corrections required. Status: verified.

## Batch: linear-discrete + probability (37 entries)

### Research log (all pages opened, not snippets)
- https://openstax.org/books/introductory-statistics/pages/6-1-the-standard-normal-distribution
  (OpenStax Introductory Statistics, section 6.1): z = (x-mu)/sigma;
  Z ~ N(0,1); Empirical Rule 68/95/99.7%. Confirms `z-score`, `normal-distribution`
  (modulo the convention note below) and `standard-deviation` entries.
- Same page footer (licensing): "Except where otherwise noted, textbooks on
  this site are licensed under CC BY-NC-SA 4.0", plus an LLM-training
  restriction (not applicable: no text reused anywhere in the glossary).
  This corrects the previously recorded "CC BY 4.0 (OpenStax)" on 105
  source lines across the dataset and in THIRD_PARTY_NOTICES.md.
- Linear-algebra examples (transpose, 2x2 determinant, trace 5, rank 1,
  eigenvalue equation, eigenvector check, inner/dot/cross products,
  Kronecker sift, P(5,2) = 20, C(5,2) = 10, mod remainders, floor/ceiling)
  and probability examples (die mean 3.5, Var identity, z = 2, Binomial pmf,
  p-hat 0.6, y-hat 14, p = 0.03 decision) recomputed by hand: all correct.

### Findings and changes
- `normal-distribution`: added a convention note — probability theory
  writes N(mu, sigma-squared) with the variance, while OpenStax and some
  texts write N(mu, sigma) with the standard deviation. The entry previously
  presented one convention as universal.
- Licensing: 105 source license strings corrected to CC BY-NC-SA 4.0 with
  "consulted as reference, original wording, no text reused"; ledger entry
  updated. No text was ever copied, so no ShareAlike trigger; the record is
  now accurate.
- No other corrections. Status: verified (`normal-distribution`:
  corrected and verified).

## Batch: geometry-trig + calculus (42 entries)

### Research log (all pages opened, not snippets)
- https://openstax.org/books/calculus-volume-3/pages/6-5-divergence-and-curl
  (OpenStax Calculus Vol 3, section 6.5): div F = Px+Qy+Rz (scalar);
  curl F = (Ry-Qz)i + (Pz-Rx)j + (Qx-Py)k (vector); curl-as-rotation with
  paddlewheel interpretation; div(curl F) = 0. Confirms the `divergence`
  example (1+1 = 2), the `curl` example (1-(-1) = 2k), and both definitions.
- https://tutorial.math.lamar.edu/classes/calciii/curldivergence.aspx
  (Lamar University, Paul's Online Notes): nabla operator, curl as cross
  product, divergence as dot product. Second confirmation of the same
  formulas and of the `nabla-gradient` framing.
- Geometry/trig examples (3-4-5 distance, midpoint (2,3), circle center and
  radius, sin30/cos60/tan45, arcsin principal value, vector length 5) and
  all calculus examples (limits, derivatives, partials, integrals,
  epsilon-delta statement, sup/inf of (0,1), Euler limit, Riemann limit)
  recomputed by hand during review: all correct.

### Findings and changes
- No corrections required. Every definition, example, explanation and
  cross-link reviewed field by field. Status: verified.

## Batch: algebra + logic-sets (41 entries)

### Research log (all pages opened, not snippets)
- Metamath Proof Explorer `qnnen` (formal proof, `ℚ ≈ ℕ`), Math SE #1311
  (rationals equicardinal with integers/naturals), and a cardinality
  exposition with explicit bijections: all confirm |N| = |Q|. Supports the
  `cardinality` example `|N| = |Q|`.
- Algebra examples (quadratic formula roots x = -1, -2 for x^2+3x+2;
  summation 1+4+9+16 = 30; log_2(8) = 3; arcsin(1/2) = 30 degrees principal
  value; (a+b)^3 expansion) verified by direct computation during review.

### Findings and changes
- `disjunction`: example explanation suggested "both conditions" could hold
  for (x=0) v (x=1), which is impossible there. Reworded: inclusive-or allows
  both in general, though not in this instance.
- Remaining algebra entries (function notation, inverse, composition,
  mapping arrow, exp/log/ln, sequences, summation, product, binomial,
  proportion, delta, ellipsis, slope-intercept, quadratic formula) and logic
  entries (membership, subsets, union/intersection/difference/complement,
  empty and power sets, set-builder, quantifiers, connectives, implication,
  iff, therefore/because, cardinality, QED, predicates): every field
  reviewed; examples recomputed; cross-links checked. Status: verified
  (`disjunction`: corrected and verified).

## Batch: greek + arithmetic (56 entries)

### Research log (all pages opened, not snippets)
- https://dlmf.nist.gov/25.6 (DLMF, Riemann zeta integer values): confirms
  25.6.1 zeta(2) = pi^2/6. Supports `zeta` example and `zeta-function` example.
- https://mathworld.wolfram.com/MangoldtFunction.html (MathWorld, specialist
  reference): von Mangoldt function "sometimes also called the lambda
  function", denoted by capital lambda. Supports `lambda` correction.
- https://leanprover-community.github.io/mathlib4_docs/Mathlib/NumberTheory/ArithmeticFunction/VonMangoldt.html
  (Lean Mathlib docs): "We use the standard notation Λ." Second confirmation.
- https://www.tauday.com/tau-manifesto (Hartl, primary source of the proposal)
  and Wikipedia "Tau (mathematics)": tau-as-circle-constant is a 2010
  proposal; "π is used almost exclusively in mainstream mathematical
  education and practice". Supports `tau` correction.
- Wikipedia "Correlation ratio" + https://encyclopediaofmath.org/wiki/Correlation_ratio
  (specialist encyclopedia): correlation ratio denoted eta. Supports `eta`
  correction (was "partial regression coefficient").

### Findings and changes
- `beta`: claimed lowercase beta denotes "the beta function" — wrong, that is
  capital B. Reworded to regression coefficients / beta-distribution shape.
  Related link `beta-function` replaced with `regression-note`.
- `eta`: "partial regression coefficient" unverified — replaced with
  correlation ratio (eta), verified above.
- `lambda`: "uppercase Lambda is used in logic and set theory" unverified —
  replaced with von Mangoldt function (number theory), verified above.
- `tau`: presented the tau proposal as standard usage — qualified as a
  proposal; mainstream uses pi. Example explanation qualified.
- `upsilon`: overstated mathematical use — corrected to rarely used.
- `alpha`: "root of an equation" is unusual — replaced with regression
  intercept / significance level.
- `power`: x^{-1} = 1/x now states nonzero x.
- `nth-root`: x^{1/n} now states the domain qualification.
- Remaining Greek letters (gamma, delta, epsilon, theta, iota, kappa, mu, nu,
  xi, omicron, pi, rho, sigma, phi, chi, psi, omega) and all other arithmetic
  entries: symbol, definition, examples and cross-links reviewed field by
  field against the entries' cited sources (Unicode charts for glyph identity,
  OpenStax Prealgebra/Elementary/Intermediate Algebra for meaning); no issues.
  Status: verified.
