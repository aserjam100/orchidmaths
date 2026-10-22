import { describe, it, expect } from 'vitest'
import {
  backTarget,
  learnConceptPath,
  normalizeReturnEntry,
  parseReturnStack,
  pushReturn,
  studyConceptPath,
  withReturnStack,
} from '../lib/learnReturn'

// The URL is the only state a navigation leaves standing, so these are the properties that keep
// a chain of prerequisite hops from stranding a learner or looping.

describe('parseReturnStack', () => {
  it('reads repeated return params, most-recent first', () => {
    const params = new URLSearchParams('concept=c&return=/study?concept=b&return=/learn?concept=a')
    expect(parseReturnStack(params)).toEqual(['/study?concept=b', '/learn?concept=a'])
  })

  it('reads a bare concept id as a learning-loop target', () => {
    const params = new URLSearchParams('return=frac.add.diff')
    expect(parseReturnStack(params)).toEqual(['/learn?concept=frac.add.diff'])
  })

  it('drops empty and duplicate entries', () => {
    const params = new URLSearchParams('return=&return=/learn?concept=a&return=/learn?concept=a')
    expect(parseReturnStack(params)).toEqual(['/learn?concept=a'])
  })

  it('is empty when there is no return at all', () => {
    expect(parseReturnStack(new URLSearchParams('concept=c'))).toEqual([])
  })
})

describe('normalizeReturnEntry', () => {
  it('passes a path through and expands a bare id', () => {
    expect(normalizeReturnEntry('/study?concept=x')).toBe('/study?concept=x')
    expect(normalizeReturnEntry('x.y')).toBe('/learn?concept=x.y')
    expect(normalizeReturnEntry('   ')).toBeNull()
  })
})

describe('pushReturn', () => {
  it('puts the current page on the front', () => {
    expect(pushReturn(['/learn?concept=a'], '/learn?concept=b')).toEqual([
      '/learn?concept=b',
      '/learn?concept=a',
    ])
  })

  it('drops an entry that points at the current concept, so a cycle cannot grow', () => {
    // A -> B -> A: the second hop back to A must collapse the stack, not append A again.
    expect(pushReturn(['/learn?concept=a'], '/study?concept=a')).toEqual(['/study?concept=a'])
  })

  it('does not push the current path twice', () => {
    expect(pushReturn(['/study?concept=b'], '/study?concept=b')).toEqual(['/study?concept=b'])
  })

  it('leaves the stack unchanged in spirit when empty', () => {
    expect(pushReturn([], '/learn?concept=b')).toEqual(['/learn?concept=b'])
  })
})

describe('backTarget', () => {
  it('returns null for a first visit', () => {
    expect(backTarget([])).toBeNull()
  })

  it('pops the most recent entry and carries the rest', () => {
    const back = backTarget(['/study?concept=b', '/learn?concept=a'])
    expect(back?.href).toBe('/study?concept=b&return=%2Flearn%3Fconcept%3Da')
    expect(back?.remaining).toEqual(['/learn?concept=a'])
  })

  it('gives a bare path for the last hop', () => {
    expect(backTarget(['/learn?concept=a'])?.href).toBe('/learn?concept=a')
  })
})

describe('withReturnStack', () => {
  it('uses & when the path already has a query', () => {
    expect(withReturnStack('/learn?concept=c', ['/learn?concept=a'])).toBe(
      '/learn?concept=c&return=%2Flearn%3Fconcept%3Da',
    )
  })

  it('uses ? when it does not', () => {
    expect(withReturnStack('/review', ['/learn?concept=a'])).toBe('/review?return=%2Flearn%3Fconcept%3Da')
  })

  it('appends nothing for an empty stack', () => {
    expect(withReturnStack('/learn?concept=c', [])).toBe('/learn?concept=c')
  })
})

describe('path builders', () => {
  it('escape ids', () => {
    expect(learnConceptPath('a b/c')).toBe('/learn?concept=a%20b%2Fc')
    expect(studyConceptPath('a b/c')).toBe('/study?concept=a%20b%2Fc')
  })
})
