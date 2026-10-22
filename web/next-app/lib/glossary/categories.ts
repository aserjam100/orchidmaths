import type { GlossaryCategoryId } from './types'

export interface GlossaryCategory {
  id: GlossaryCategoryId
  label: string
  description: string
}

/** The nine spec categories (A-I). Order is presentation order. */
export const GLOSSARY_CATEGORIES: GlossaryCategory[] = [
  {
    id: 'greek',
    label: 'Greek letters',
    description: 'The 24 Greek letters, their forms and common mathematical uses. Meanings vary by context.',
  },
  {
    id: 'arithmetic',
    label: 'Arithmetic',
    description: 'Elementary operations, comparison, fractions, powers, roots and basic notation.',
  },
  {
    id: 'algebra',
    label: 'Algebra and functions',
    description: 'Variables, equations, functions, exponentials, logarithms, sums and sequences.',
  },
  {
    id: 'logic-sets',
    label: 'Sets, logic and proof',
    description: 'Set notation, quantifiers, logical connectives and proof language.',
  },
  {
    id: 'geometry-trig',
    label: 'Geometry and trigonometry',
    description: 'Angles, shapes, congruence, coordinates and trigonometric functions.',
  },
  {
    id: 'calculus',
    label: 'Calculus and analysis',
    description: 'Limits, derivatives, integrals, series and epsilon-delta language.',
  },
  {
    id: 'linear-discrete',
    label: 'Linear algebra and discrete math',
    description: 'Vectors, matrices, eigenvalues, combinatorics, graphs and modular arithmetic.',
  },
  {
    id: 'probability',
    label: 'Probability and statistics',
    description: 'Events, random variables, expectation, distributions and inference.',
  },
  {
    id: 'advanced',
    label: 'Number theory and advanced math',
    description: 'Number sets, complex numbers, abstract algebra, topology and transforms.',
  },
]

export const categoryLabel = (id: GlossaryCategoryId): string =>
  GLOSSARY_CATEGORIES.find(c => c.id === id)?.label ?? id
