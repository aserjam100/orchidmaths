# Valid argument forms and fallacies

An argument is a list of premises and a conclusion, and it is **valid** when every assignment that makes the premises true also makes the conclusion true. Validity says nothing about whether the premises happen to be true: it is a claim about the shape of the inference.

## Valid forms

Two forms carry most of the work. **Modus ponens** has premises \(p \to q\) and \(p\), and concludes \(q\). **Modus tollens** has premises \(p \to q\) and \(\lnot q\), and concludes \(\lnot p\). Both are valid for the same reason: the implication rules out exactly one case, \(p\) true with \(q\) false, so ruling out either half of that case forces the other.

Disjunctive syllogism, \(p \lor q\) with \(\lnot p\) concluding \(q\), is the third form worth naming. Each is a truth-table fact before it is a rule of reasoning.

## The two fallacies

**Affirming the consequent** takes \(p \to q\) and \(q\) and concludes \(p\). It is invalid: \(q\) may be true because of something else. Rain makes the ground wet, and a sprinkler does too, so "the ground is wet" does not establish rain.

**Denying the antecedent** takes \(p \to q\) and \(\lnot p\) and concludes \(\lnot q\). It is invalid for the same reason in the other direction: the implication says nothing about what happens when \(p\) is false.

## Reading an argument

To classify an argument, write it in the standard form and match it against the forms you know. Premises "if the machine is on then the light is green" and "the light is green", with conclusion "the machine is on", match affirming the consequent: it is invalid, and one counterexample assignment — machine off, light green — satisfies both premises and falsifies the conclusion.

The same two premises with the conclusion "if the light is not green then the machine is off" is the contrapositive of the first premise, and that argument is valid because it restates the premise.
