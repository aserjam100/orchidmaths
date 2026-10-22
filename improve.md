# Orchid vs MathAcademy.com — Deep Research & Gap Analysis

> Saved conversation: 2026-08-28 — plan-mode research + user-approved gap-closure plan.
>
> **This is a historical research record, not current documentation.** Every figure in it is the
> state of the project on 2026-08-28: the corpus then held ~437 concepts, XP was a flat 10/5 per
> lesson and review, and the mastery check was gated at 150 XP. None of that is true now (657
> concepts, an effort-derived XP base, a 50 XP gate). The decisions this research produced are
> recorded as they landed in `docs/adr/`; the code comments that cite this file cite it as the
> provenance of an idea, not as a description of the present.

---

## 1. How Research Was Done

- **Online:** `mathacademy.com` (`/`, `/how-it-works`, `/pedagogy`, `/how-our-ai-works`, `/courses`, `/faq`), *The Math Academy Way* (Justin Skycak PDF via `justinmath.com` + ResearchGate draft), 3rd-party reviews (HomeschoolTools parent guide).
- **Training data:** Cognitive-science pillars cited by Math Academy — Bloom 2-sigma, mastery learning, distributed practice, interleaving, layering, non-interference, deliberate practice, automaticity, testing effect, cognitive load theory, worked-example effect, dual-coding, subgoal labeling.
- **Local codebase:** `README.md:24`, `README.md:114`, `DESIGN.md:1`, `data/concepts/*.json` (17 domains), `data/courses.json:1`, `data/lessons/lessons.json:1`, `internal/concepts/concepts.go:8`, `internal/concepts/loader.go:13`, `internal/concepts/dag.go`, `internal/diagnostic/cat.go:1`, `internal/scheduler/scheduler.go:1`, `internal/scheduler/sm2.go:1`, `internal/mastery/machine.go:1`, `internal/engine/engine.go:1`, `internal/planning/planner.go:1`, `internal/lessons/lesson.go:1`, `internal/grader/grader.go:1`, `internal/generator/registry.go:1`, `internal/scoring/updater.go:1`, `internal/server/server.go:1`, `web/next-app/package.json:1`.

---

## 2. Executive Summary

Orchid correctly implements Math Academy's core philosophy: DAG-guarded, speed-and-accuracy mastery gating with generated problems and SM-2. The gap is not philosophy — it is **depth, scale, and adaptive intelligence**. Math Academy is a 10-year, thousands-of-topics expert system with 4 interlocking algorithms (knowledge graph, student model, diagnostic, task-selection); Orchid is a solid ~437-concept MVP with 1 scheduler and 1 simple CAT diagnostic. Roughly: **Orchid ≈ 30% of Math Academy's pedagogy surface; 60% of its engineering surface; 15% of its content/ops surface.**

---

## 3. Math Academy — What It Is

### 3.1 Knowledge Graph + Student Model + Diagnostic + Task Selection (the 4 algorithms)
- **Knowledge graph** stores: all topics, easiest↔hardest variations per topic, background knowledge required for each topic, prerequisites, encompassings, key prerequisites, interference links. Answers: "what are the topics?", "what are easiest/hardest variations?", "what background is needed?", "if student struggles on X, what specific background is most relevant?"
- **Student model** overlays answers on KG → knowledge profile, using spaced repetition to determine when previously learned material needs review. Tracks student's spaced repetition profile (darker blue oval = more stable).
- **Diagnostic algorithm:** Compresses KG into smallest covering set of topics at sufficient granularity, then repeatedly picks the topic whose assessment provides most information about knowledge profile. Correct → positive evidence for topic + prerequisites + related; incorrect → negative evidence for topic + postrequisites + related. Weights conflicting evidence. Finds "knowledge frontier" (boundary between known/unknown) in ~30-45 min adaptive exam. Assesses both course content AND lower-grade foundations. Reports: placement recommendation, estimated completion dates at varied daily XP, foundational gaps, mastery levels per topic, per-topic performance.
- **Task-selection algorithm:** Maximizes learning per unit time. Uses mastery learning, layering, spaced repetition, interleaving, minimizing associative interference. Adapts to pace: student-topic learning speed (how easy a topic is for that student → how many spaced reps a review is worth), forces explicit review on challenging topics, remedial reviews on halt/re-attempt, immediate remedial review after quiz miss. Quiz every ~150 XP, 80% difficulty target (adaptive), retake available.

### 3.2 Pedagogy — 8 Strategies + Worked Examples
1. **Automaticity & Computational Fluency** — Working memory holds 3-5 novel elements; long-term memory is unlimited. Practice until effortless/automatic.
2. **Distributed Practice / Spaced Repetition** — Reviews spaced at progressively greater intervals; massed practice is inefficient.
3. **Mixed Review / Interleaving** — Present problems out of context, spread over many assignments covering broad mixes.
4. **Active Learning** — Minimal lecturing, maximize problem solving with immediate feedback.
5. **Layering** — Learn advanced material that exercises prerequisites (retroactive facilitation); increases connections to existing knowledge.
6. **Cognitive Noninterference** — Teach dissimilar concepts simultaneously; space similar concepts over time.
7. **Deliberate Practice** — Structured, intentional, at-current-level, with expert feedback + repetition.
8. **Mastery Learning** — Bloom's mastery learning at fully granular level, only via 1-on-1 tutoring / automated platform; establishes performance threshold before advancing.

**Plus:** Lesson scaffolding — 10× granularity vs textbook (calculus: 100 textbook steps vs ~1000 steps = 300 topics ×3 knowledge points). Each knowledge point starts with fully worked example (worked-example effect, subgoal labeling, dual-coding visualizations/diagrams), then up to 5 practice problems — 2-in-a-row correct to advance.

### 3.3 Gamification & Ops
- XP ≈ 1 min focused effort; awarded per task type (lesson/review/multistep/quiz differ), bonus for perfect, negative XP for rushing/guessing.
- Daily XP goal 20-40 adjustable (3-5 days/week), leagues weekly promotion/demotion (fresh start), leaderboard, oversight dashboard (parent/teacher share), pause 30-90 days, accommodations (timing adjustable), 30-day refund.
- Courses: 4th grade → university (Calc I/II, Linear Algebra, Multivariable, Methods of Proof, ML, + Prob/Discrete/DiffEq/Real Analysis/Abstract Algebra), WASC accredited, UC a-g approved, transcripts, AP/SAT test-prep tracks.

---

## 4. Orchid — What It Is (Current)

- **Tagline:** "An open-source adaptive math learning engine — arithmetic to calculus, mastery-gated, locally-first." (`README.md:14`)
- **Graph:** 437 concepts across 17 domains (`README.md:114`), single Go binary + Next.js web, 5-layer architecture (UI → API → Core Engine → Grading → Storage, `README.md:123`).
- **Mastery gating:** 5 states `UNSEEN → LEARNING → PRACTICING → MASTERED → DECAYING` (`README.md:71`). Each concept has `streak` + `avg_time_seconds` thresholds (`internal/concepts/concepts.go:8`). `DECAYING` after 14 days (`internal/scheduler/scheduler.go:95`). `internal/mastery/machine.go:22` advances only when both thresholds met.
- **Scheduler:** `priority = 0.5*days + 0.2*(1-mastery) + 0.3*weakness +5 if decaying` (`internal/scheduler/scheduler.go:121`). Hard rules: prerequisites not mastered → never surface, never same concept twice in row, ~70% new / 30% review. `selectWithBalance:175` enforces ratio.
- **Spaced repetition:** Simplified SM-2 (`internal/scheduler/sm2.go:15`, `internal/mastery/machine.go:55`), interval `1→6→*EFactor`, `EFactor` starts 2.5. Reviews woven into normal sessions; separate `/api/reviews/due|session|answer` (`internal/server/server.go:1106`).
- **Generators:** Parameterized per-concept, `GeneratorContext{Difficulty: 0.3-1.0}`, `internal/generator/registry.go:62` self-check (grades own answer, rejects empty), `BatchGenerate` deduping. Covers arithmetic, fractions, algebra, geometry, trig, precalc, calc, stats, linalg, discrete, complex, numtheory, ode, abstract, ml, topology. Thin domains: ML 1, topology 3, abstract 7, diffEq 8.
- **Grader:** `internal/grader/grader.go:1` Supports numeric/polynomial/expression/symbolic/multiple_choice/comparison/ordering/tuple/complex + SymPy fallback — broader than MA's multiple-choice-only critique.
- **Lessons:** `internal/lessons/lesson.go:28` loads `data/lessons/lessons.json:1` (2110 entries) mapping concept→markdown files, mostly `algebrica/*.md` (CC BY-NC 4.0, `README.md:177`) plus `teaching/*.md`. One lesson serves multiple concepts. Only 8 diagram mappings (`internal/engine/engine.go:319`). D2 diagrams are architecture-only (`DESIGN.md:175`).
- **Courses:** `data/courses.json:1` — 10 courses terminating at Precalculus (`pc`). No Calculus I/II despite 71 calc concepts; no university course wiring for linalg/ml/topology-abstract-diffEq.
- **XP:** `internal/engine/engine.go:525` `computeXP: base 10 (5 review) * timeMultiplier (0.5-1.5) * streakMultiplier (1+0.1*streak)`, `internal/server/server.go:1039` daily XP goal (`150` default, `internal/scoring/updater.go:59`), weekly leaderboard (`internal/server/server.go:361`), streak calendar (`internal/scoring/updater.go:88`).
- **Planning:** `internal/planning/planner.go:57` `PrerequisitesOf` / `collectChain` (transitive closure of targets), `WeaknessMap` (`internal/engine/engine.go:666`) + propagation via `dependentsOf` (`internal/engine/engine.go:824`), `AdjustPlan` decay cleanup.
- **Design:** `DESIGN.md:1` Gold/teal taxonomy, IBM Plex Serif + Mono, dark-first `#0b0f1a` / parchment `#fefcf4`, no validation gaps in this session.

---

## 5. Gap Matrix

| Dimension | Math Academy | Orchid | Severity | File References |
|---|---|---|---|---|
| **Knowledge graph edges** | prerequisites + encompassings + key prerequisites + interference groups + difficulty variants | prerequisites only; single edge type; single time threshold | High — blocks layering & non-interference | `internal/concepts/concepts.go:8`, `loader.go:13`, `concepts/dag.go`, `scripts/validate_graph.go` |
| **Knowledge graph scale** | 3k-5k topics; ~1000 steps for calculus alone | 437 concepts; calculus 71 | High — 10× granularity missing | `README.md:114`, `data/concepts/*.json` |
| **Student model** | Per-topic learning speed, spaced-rep profile (blue oval), implicit vs explicit rep distinction | Vanilla SM-2, generic `EFactor`, `WeaknessScore` proxy | High — no individualized spacing | `internal/scheduler/sm2.go:1`, `internal/mastery/machine.go:55`, `internal/storage/store.go` |
| **Diagnostic** | Compressed covering set, info-gain selection, evidence propagation ±, confidence weighting, mastery+automaticity, supplemental diag, report (placement, gaps, completion dates) | Binary search on linear order, 2 probes/concept, 12 min Qs, ratio thresholds 0.75/0.25, no confidence/report | Critical — frontier estimate biased | `internal/diagnostic/cat.go:1`, `internal/engine/engine.go:578`, `internal/server/server.go:374` |
| **Task selection** | Maximize learning/time, layering turbo-boost, non-interference choice of 3-5 dissimilar, interleaving window, re-attempt+remedial batch | Single top priority, ratio-only, no layering, no dissimilarity, no interleaving window | High — no max-efficiency | `internal/scheduler/scheduler.go:121`, `175`, `internal/engine/engine.go:164` |
| **Quiz subsystem** | Every ~150 XP timed closed-book, diverse recent topics, 80% difficulty target, immediate review, optional retake | No quiz mode (reviews only), not timed-closed-book, no retake | Medium — missing testing effect | `internal/server/server.go:1106` vs `internal/quiz/*` (missing) |
| **Spaced repetition** | Per-topic speed-adapted intervals + retroactive facilitation | Flat SM-2, 14-day decay | Medium | `internal/scheduler/scheduler.go:95` |
| **Lessons** | 3 KPs/topic, worked example → subgoal-labeled + dual-coded diagrams → 5 problems 2-in-a-row | One markdown per concept(s), shared sources, 8 diagrams, no KP scaffolding | High — cognitive load not minimized | `internal/lessons/lesson.go:28`, `internal/engine/engine.go:319`, `DESIGN.md:175` |
| **Task taxonomy & XP** | lesson/review/multistep/quiz with distinct XP, bonus perfect, negative guessing | Single problem type, flat XP (10/5), no negative, no task-type diff | Medium | `internal/engine/engine.go:525` |
| **Gamification** | Daily XP goal 20-40, leagues weekly (promotion/demotion), leaderboard, accommodations | Weekly leaderboard, daily goal stored but not enforced, no leagues, no accommodations, no pause | Medium | `internal/server/server.go:1039`, `internal/scoring/updater.go:59` |
| **Courses** | 4th → university (Calc I/II, LinAlg, Multivariable, Proof, ML, + planned Prob/Discrete/DiffEq/Real Analysis/Abstract Algebra), integrated/standard | 10 courses → Precalc only; calc/uni concepts orphaned | High — blocks parity | `data/courses.json:1`, `internal/planning/planner.go:57` |
| **Accreditation & transcript** | WASC, UC a-g, transcripts | None, no diagnostic report, no completion estimator | High — accreditation-track blocked | `internal/diagnostic/report.go` (missing), `internal/planning/planner.go` |
| **Oversight** | Parent/teacher panel, sharing, analytics dashboard | Auth single student, `/api/activity`, `/api/weaknesses` only, no role | Medium | `internal/server/server.go:444`, `794` |
| **Efficacy instrumentation** | 4× acceleration claim, 93% first-pass, 98% second-pass metrics, AP BC 8th graders | No instrumentation beyond `Attempt` log | Low — next phase | `internal/storage/store.go`, `internal/engine/engine.go:477` |

---

## 6. What Orchid Does Well — Preserve

- DAG integrity: `scripts/validate_graph.go`, `loader.go:50` validate (no cycles/orphans/dup IDs), Kahn's topo-sort, strong `*_test.go` suite.
- Grading breadth (above), local-first single-binary deploy (`go.mod:1`, `Dockerfile`, `fly.toml`), design system (`DESIGN.md:1`).
- Speed+accuracy gating faithful to MA philosophy (`README.md:30`).

---

## 7. Approved Plan — Full Parity Push, Session-Split, Incremental PRs

### User choices (2026-08-28):
- **Scope:** Full parity push but divided into sessions; this session = core pedagogy only.
- **Lessons:** Augment Algebrica but source maths lessons online to fill gaps.
- **Courses:** Accreditation-track (WASC/UC a-g).
- **Delivery:** Incremental PRs (one pillar per PR, each shippable & tested).

### Sessions
- **Session 1 — THIS SESSION: Core pedagogy** (diagnostic, student model, task selection, quiz, spaced repetition)
- Session 2 — Lesson scaffolding + content expansion
- Session 3 — Gamification & oversight
- Session 4 — Catalog scale & accreditation
- Session 5 — Scale, efficacy & polish

### Session 1 — Detailed PRs

**PR 1.1 — Knowledge-graph extension** `internal/concepts/concepts.go:8`, `internal/concepts/dag.go`, `internal/concepts/loader.go:13`, `data/concepts/*.json`, `scripts/validate_graph.go`
- Extend `Concept`: `Encompasses []string`, `KeyPrerequisites []string`, `InterferenceGroup string`, `Variants []Variant{Difficulty, TimeThresh, Label}`.
- Index `encompassedBy`, `interferersOf` in `DAG`; validate no dangling refs, non-trivial groups, cycles via encompasses included in Kahn's check.
- Add `data/concepts/enrichment.json` migration (heuristic: `calc.*` encompasses `alg.*`, `trig.*` interferes among siblings).
- Verify: `go test ./internal/concepts`, `go run scripts/validate_graph.go`, new `TestEncompassings`/`TestInterferenceGroups`.

**PR 1.2 — Student model + spaced repetition** `internal/scheduler/sm2.go:1`, `internal/mastery/machine.go:1`, `internal/storage/store.go`, `internal/storage/sqlite.go`, `internal/scoring/updater.go:1`
- New table `student_topic_speed (student_id, concept_id, efactor, interval, repetitions, learning_speed)` in `storage/migrate.go`.
- `ComputeSM2` takes `learningSpeed` (0.5-2.0), scales interval `*= EFactor * learningSpeed`; update speed from `timeRatio` + streak.
- Expose `spaced_reps` profile via `GET /api/scores` (blue-oval darkness).
- Tests: `scheduler/sm2_test.go` speed 0.7 vs 1.5, `storage/storage_test.go` round-trip.

**PR 1.3 — Diagnostic rewrite** `internal/diagnostic/cat.go:1`, `internal/engine/engine.go:578`, `internal/server/server.go:374`, `internal/planning/planner.go:57`
- Replace binary search: `compressedCover()` minimal covering set, `infoGain(topic)` max entropy reduction, evidence propagation (+0.3 prerequisites on correct, -0.3 dependents on incorrect), Bayes conflict resolution.
- Per-concept `KnowledgeConfidence` 0-1, `Frontier` as highest drop; measure mastery+automaticity (`fast` already plumbed in `cat.go:51`).
- New `internal/diagnostic/report.go` → `DiagnosticReport{PlacementCourseID, FrontierIdx, GapsByDomain, MasteryLevels, CompletionEstimates map[int]string}` returned on done-branch.
- Add `supplementalDiagnostic()` when confidence <0.7; 25-45 adaptive Qs.
- Tests: simulate strong/weak students, frontier ±2, e2e `server/server_test.go`.

**PR 1.4 — Task selection + quiz** `internal/scheduler/scheduler.go:1`, `internal/engine/engine.go:164`, `internal/engine/detail.go`, `internal/server/server.go:232`
- Enhance `computePriority:121`: layering bonus (+2 if encompasses weak prereq), non-interference penalty (-3 if same `InterferenceGroup`), interleaving window (exclude same `Subdomain` in last 2).
- Choice: return top 3 dissimilar candidates for client pick; keep 70/30 via `selectWithBalance:175`.
- New `internal/quiz/quiz.go` → `NextQuizQuestion` every 150 XP, timed closed-book, diverse recents, 80% difficulty target (`weakness→difficulty` map via `engine.computeDifficulty:146`), immediate remedial enqueue, optional retake.
- New `POST /api/quiz/session|answer`, update `handleDueReviews:1106`.
- Tests: dissimilarity + layering, quiz flow.

**PR 1.5 — Halt / re-attempt / accommodations** `internal/engine/engine.go:356`, `internal/mastery/machine.go:22`
- Halt = 2 consec incorrect → cut short (XP=0), re-queue easier `Variant` + `KeyPrerequisites` remedial batch.
- `Student.Settings.timeMultiplier` applied to `timeThreshold` in `SubmitAnswer`; via `GET|PUT /api/settings:1065` `accommodations: {extra_time: 1.25}`.
- Negative XP: `elapsed <2s && incorrect` twice → `-5 XP`.

**Files touched Session 1:** `internal/concepts/*`, `internal/diagnostic/*`, `internal/scheduler/*`, `internal/mastery/*`, `internal/quiz/*` (new), `internal/engine/*`, `internal/storage/*`, `internal/server/*`, `web/next-app/app/**` (thin quiz/report UI), `data/concepts/*.json`, `scripts/*`. **Out of scope Session 1:** lesson KP scaffolding, diagram 8→100, leagues, oversight role, new courses beyond `data/courses.json:1` (ships Session 4).

**Effort:** ~3 engineer-weeks, 5 PRs, each `go test ./...`, `npm run test`, `go run scripts/validate_graph.go`, `playwright` gated.

**Risks & mitigations:** Encompass backfill wrong → transitive-closure verifier; info-gain complexity → entropy proxy + synthetic-student corpus + fallback flag; SM2 migration → lazy migration on first write; quiz disturbing ratio → quiz counted as review.

**Accreditation-track:** Report schema (`DiagnosticReport`, `MasteryLevels`, `CompletionEstimates`) designed for WASC evidence on adaptive placement + mastery transparency; source WASC/UC a-g specs online in Session 4, schema ships Session 1.

---

## 8. Next Actions

- [ ] PR 1.1 graph extension (first PR of Session 1).
- [ ] Then PR 1.2 → 1.5 sequentially per incremental-PR rule.
- Sessions 2-5 queued per plan above.

---

*Provenance note, added 2026-10-12. This file was gitignored while `.gitignore` called it "local
research — do not commit", but six Go files and several documents cite it by line (`engine.go:744`,
`scheduler.go:166`, `leaderboard/leagues.go:11`, `lessons/lesson.go:30`, `docs/efficacy.md:30`,
`docs/wasc-uc-ag-research.md`), so every reference dangled for anyone who cloned the repo. It is
committed now for the same reason the ADRs are: a decision's provenance is part of the decision.
It is deliberately not held to the current-fact checks in `scripts/check_docs.py`, because its
numbers are the numbers of the day it was written.*
