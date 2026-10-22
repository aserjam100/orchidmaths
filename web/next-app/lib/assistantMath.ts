// Splitting a model's reply into prose and mathematics.
//
// The assistant is told to write mathematics between dollar signs, and nothing
// else in a reply is formatting. A model does not always comply: a live session
// answered "explain how to find a derivative" with the limit definition written
// as \(f'(x)=\lim_{h\to0}\frac{f(x+h)-f(x)}{h}\), because "name the method" beat
// "no LaTeX" whenever the method's name *is* a formula. Forbidding the delimiters
// again would only move the boundary; the reply is therefore parsed and the math
// is rendered, which is what the suggested libraries do too (assistant-ui's
// react-markdown preprocess rewrites \(...\) and \[...\] into remark-math's
// delimiters rather than deleting them).
//
// Parsing is deliberately not a Markdown parse. The panel renders prose as text,
// so `[x](https://evil)` stays the characters a person can read. Only the four
// math delimiters below are recognised, and only the server's own [[id]] table
// can produce a link.

export type MathPart =
	| { kind: 'text'; text: string }
	| { kind: 'math'; tex: string; display: boolean }

type Delimiter = { open: string; close: string; display: boolean }

// Longest openers first: "$$" must be tested before "$", and "\[" before a
// literal backslash is treated as text.
const DELIMITERS: Delimiter[] = [
	{ open: '$$', close: '$$', display: true },
	{ open: '\\[', close: '\\]', display: true },
	{ open: '\\(', close: '\\)', display: false },
	{ open: '$', close: '$', display: false },
]

// A dollar figure is money, not mathematics. "$5 and $6" would otherwise open a
// span at the first dollar and close it at the second, swallowing " and ".
function looksLikeMoney(tex: string): boolean {
	return /^\d[\d.,]*$/.test(tex.trim())
}

// Inline content may not begin or end with whitespace, which is what separates
// "$x$" from "$5 and $6" without guessing at the writer's intent.
function reasonableInline(tex: string): boolean {
	if (tex.trim() === '') return false
	if (tex !== tex.trim()) return false
	return !looksLikeMoney(tex)
}

// splitMath returns the reply as ordered parts, alternating prose and math. Text
// with no maths is returned as a single text part, so a caller can skip the math
// component entirely for the common case.
export function splitMath(input: string): MathPart[] {
	const parts: MathPart[] = []
	let text = ''
	let i = 0

	const flush = () => {
		if (text !== '') {
			parts.push({ kind: 'text', text })
			text = ''
		}
	}

	while (i < input.length) {
		// An escaped dollar is a literal dollar, not an opener.
		if (input[i] === '\\' && input[i + 1] === '$') {
			text += '$'
			i += 2
			continue
		}

		const opener = DELIMITERS.find(d => input.startsWith(d.open, i))
		if (!opener) {
			text += input[i]
			i += 1
			continue
		}

		const start = i + opener.open.length
		const end = input.indexOf(opener.close, start)
		// An unmatched delimiter is prose that happens to contain a bracket. It is
		// kept, because deleting it would hide the model's formatting mistake.
		if (end === -1) {
			text += input[i]
			i += 1
			continue
		}

		const tex = input.slice(start, end)
		if (!opener.display && !reasonableInline(tex)) {
			// Not maths after all: keep the dollar and continue past it, so the
			// rest of the sentence is still available to a later opener.
			text += opener.open
			i = start
			continue
		}

		flush()
		parts.push({ kind: 'math', tex, display: opener.display })
		i = end + opener.close.length
	}

	flush()
	return parts
}
