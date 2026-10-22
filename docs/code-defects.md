# Code defects found by a documentation audit

These are code defects, not documentation defects. A documentation sweep found both because
the prose and the code disagreed, and in each case the prose was closer to correct than a
reader would expect.

They are recorded separately from the documentation work on purpose. A fix here changes
product behavior, so it needs its own review and its own test.

## 1. The daily XP goal fallback is a literal, and the package above it exists to prevent that

**Location:** `internal/scoring/updater.go:73`

```go
dailyGoal := 10
if st, err := u.repo.GetStudent(studentID); err == nil && st != nil && st.DailyXPGoal > 0 {
    dailyGoal = st.DailyXPGoal
}
```

`xp.DefaultDailyGoal` is 30 (`internal/xp/policy.go:61`). This fallback applies to any learner
whose stored `DailyXPGoal` is absent, zero, or on a row `GetStudent` does not return. Such a
learner is then measured against 10 XP/day.

The package comment at `internal/xp/policy.go:53-56` records this exact failure:

> Daily XP goal policy. Owned here because the server, the estimator and the client all need the
> same numbers, and they had drifted: the shipped default was 10 while the planner's fallbacks
> said 30, so a learner with no stored goal planned against 3x the XP/day they were actually
> banking.

The same field on the struct just above (`updater.go:31-34`) carries the matching history for
`QuizGateXP`. The invariant was enforced everywhere except the one line that needed it most.

**Suggested fix.** Replace the literal with `xp.DefaultDailyGoal`. Add a test that calls the
updater for a student with no stored goal and asserts the fallback equals `xp.DefaultDailyGoal`,
so the two cannot drift apart again.

**Impact.** Wrong progress bars and wrong daily-goal completion percentages for learners with
no stored goal. No data loss, and no effect on XP actually awarded.

## 2. `-repair-grading` reconstructs progress with a fallback threshold

**Location:** `internal/repair/repair.go:144-155`

```go
reqStreak := 3
thresh := 10.0
if c := dag.Concept(conceptID); c != nil {
    reqStreak = c.MasteryThreshold.Streak
    thresh = c.MasteryThreshold.AvgTimeSeconds
}
if reqStreak <= 0 {
    reqStreak = 3
}
if thresh <= 0 {
    thresh = 10.0
}
```

**This is narrower than it first appears, and the correction matters.** The DAG is consulted
first, so a concept in the graph uses its own real `required_streak` and `time_threshold`. The
literals apply in two cases only: the concept is absent from the DAG, or its stored threshold
is zero or negative.

The residual concern is therefore real but small. For those two cases the replay uses
`required_streak = 3` where the live engine would have used a different value, so the rebuilt
`concept_progress` row can differ from the row the live path would have produced. `replay` is
already documented as best-effort at `repair.go:140-142`, and it states that it does not
reproduce halt/rush XP or per-topic learning-speed evolution exactly. The threshold fallback is
a further deviation that the comment does not name.

**Suggested fix.** Either log when the fallback fires, or extend the `replay` doc comment to
name this deviation alongside the two it already names. A test covering a concept with a zero
`AvgTimeSeconds` would pin the behavior either way.

**Impact.** Low. Affects only repair runs over concepts with a missing or zero threshold.

## 3. `complex.adv.polar` answers are rounded to whole degrees

**Location:** `internal/generator/complex/generators.go`, the `polarGen` generator.

**What it does.** For a rectangular number built from a Pythagorean triple, the generator computes
`arctan(b/a)` in degrees and rounds it to an integer before writing the polar answer:
`5(cos 53° + i sin 53°)`. For the `(3, 4)` triple the exact argument is about 53.13°, so the
answer states an approximate value as though it were exact.

**Why it has not broken.** The `complex` grader converts the polar form with a regex that accepts
integers only and then compares both sides in SymPy. Because expected and answer are both the rounded
string, the comparison agrees with itself. A learner who types the accurate angle, `53.13°`, either
fails the regex or is compared against the rounded value and is marked wrong.

**Why it was not fixed here.** The fix would change the answer format of an existing concept — a
generator and its grader — and that is outside the `complex_numbers` batch that found it. Choosing
between exact integer arguments (only multiples of 30° and 45° have exact sines and cosines), a
`±`-style tolerance in the answer, or dropping polar answers in favour of exponential form is a
decision about the concept, not a repair.

**Suggested fix.** Restrict `polarGen` to angles whose sine and cosine are exact — 30°, 45°, 60°,
90°, 120°, 135°, 150°, 180° and the reflex ones — and pick radii so the rectangular form stays
integral. `TestComplexAcceptsEquivalentForms` already covers the polar path with a whole-degree case,
so a change there would be caught by the grader test.

**Impact.** Low for grading (the generator and the grader agree), real for a learner: the concept
teaches that a polar form is exact when the number it came from is a triple, and a learner who works
out the true angle is marked wrong for being right.

## What this document is not

Neither defect is a documentation error, and neither is fixed by the documentation sweep. They
are listed here so that the sweep's findings are not read as a claim that the prose is the only
thing that can be wrong.
