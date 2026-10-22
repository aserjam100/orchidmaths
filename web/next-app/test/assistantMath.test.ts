import { describe, it, expect } from 'vitest'
import { splitMath } from '../lib/assistantMath'

// The assistant is asked to write mathematics between dollar signs, and a live
// session ignored that and used \( ... \) instead. These cases pin both halves:
// the delimiters we promise to parse, and the prose that must survive untouched.

describe('splitMath', () => {
	it('returns one text part when there is no mathematics', () => {
		expect(splitMath('Start with the diagnostic.')).toEqual([
			{ kind: 'text', text: 'Start with the diagnostic.' },
		])
	})

	it('reads dollar signs as inline mathematics', () => {
		expect(splitMath('The slope is $m = 2$ here.')).toEqual([
			{ kind: 'text', text: 'The slope is ' },
			{ kind: 'math', tex: 'm = 2', display: false },
			{ kind: 'text', text: ' here.' },
		])
	})

	it('normalises the LaTeX delimiters a model leaks', () => {
		expect(splitMath('a \\(x^2\\) b')).toEqual([
			{ kind: 'text', text: 'a ' },
			{ kind: 'math', tex: 'x^2', display: false },
			{ kind: 'text', text: ' b' },
		])
	})

	it('treats square brackets and double dollars as display mathematics', () => {
		expect(splitMath('\\[\\frac{dy}{dx}\\]')).toEqual([
			{ kind: 'math', tex: '\\frac{dy}{dx}', display: true },
		])
		expect(splitMath('$$\\int_0^1 x$$')).toEqual([
			{ kind: 'math', tex: '\\int_0^1 x', display: true },
		])
	})

	it('leaves an unmatched delimiter as text', () => {
		expect(splitMath('this \\( is unfinished')).toEqual([
			{ kind: 'text', text: 'this \\( is unfinished' },
		])
	})

	it('does not mistake money for mathematics', () => {
		expect(splitMath('It costs $5 and $6 a month.')).toEqual([
			{ kind: 'text', text: 'It costs $5 and $6 a month.' },
		])
	})

	it('keeps an escaped dollar literal', () => {
		expect(splitMath('costs \\$5')).toEqual([{ kind: 'text', text: 'costs $5' }])
	})

	it('leaves Markdown links as text, because links come from the server', () => {
		const text = 'see [x](https://evil.example)'
		expect(splitMath(text)).toEqual([{ kind: 'text', text }])
	})
})
