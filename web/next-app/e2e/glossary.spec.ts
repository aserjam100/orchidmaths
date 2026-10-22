import { test, expect } from './fixtures'

// The glossary is a static reference: no API routes are stubbed here on
// purpose. If the page ever needs the network to render, these tests fail,
// which is exactly the regression this feature must not have.
test('glossary loads with categories, search and a result count', async ({ page }) => {
  const errors: string[] = []
  page.on('console', msg => {
    if (msg.type() === 'error') errors.push(msg.text())
  })
  await page.goto('/glossary')
  await expect(page.getByRole('heading', { name: 'Mathematical glossary' })).toBeVisible({ timeout: 30_000 })
  // Category browsing is visible with no query.
  await expect(page.getByRole('heading', { name: /Greek letters/ })).toBeVisible()
  await expect(page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })).toBeVisible()
  await expect(page.getByText(/\d+ entries/)).toBeVisible()
  expect(errors).toEqual([])
})

test('glossary searches by symbol, name and alias', async ({ page }) => {
  await page.goto('/glossary')
  const box = page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })
  await expect(box).toBeVisible({ timeout: 30_000 })

  await box.fill('∀')
  await expect(page.getByRole('heading', { name: 'Universal quantifier' })).toBeVisible()

  await box.fill('integral')
  await expect(page.getByRole('heading', { name: 'Integral', exact: true })).toBeVisible()

  await box.fill('n choose k')
  await expect(page.getByRole('heading', { name: 'Binomial coefficient' })).toBeVisible()
})

test('glossary filters by category, clears, and handles empty results', async ({ page }) => {
  await page.goto('/glossary')
  await expect(page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })).toBeVisible({ timeout: 30_000 })

  await page.getByRole('button', { name: 'Tags' }).click()
  await page.getByRole('button', { name: 'Geometry and trigonometry' }).click()
  await expect(page.getByRole('heading', { name: 'Sine', exact: true })).toBeVisible()
  await expect(page.getByRole('heading', { name: 'Integral', exact: true })).toHaveCount(0)

  await page.getByRole('button', { name: 'Clear search and filters' }).click()
  await expect(page.getByRole('heading', { name: /Greek letters/ })).toBeVisible()

  await page.getByRole('searchbox', { name: 'Search symbols, names or meanings' }).fill('zzzqqq')
  await expect(page.getByText('No entries match')).toBeVisible()
  // Two clear controls are visible in the empty state (search bar + empty
  // panel); either resets the page.
  await page.getByRole('button', { name: 'Clear search and filters' }).first().click()
  await expect(page.getByText('No entries match')).toHaveCount(0)
})

test('glossary shows contextual meanings for an ambiguous symbol', async ({ page }) => {
  await page.goto('/glossary')
  const box = page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })
  await expect(box).toBeVisible({ timeout: 30_000 })
  await box.fill('π')
  await expect(page.getByRole('heading', { name: 'Pi', exact: true })).toBeVisible()
  await expect(page.getByRole('heading', { name: 'Pi (circle constant)' })).toBeVisible()
})

test('glossary paginates twenty per page with URL state', async ({ page }) => {
  await page.goto('/glossary')
  await expect(page.getByRole('heading', { name: 'Mathematical glossary' })).toBeVisible({ timeout: 30_000 })
  await expect(page.getByText(/Showing 1–20 of \d+ entries/)).toBeVisible()
  // Page 1 holds Greek letters; page 2 must show different cards.
  await expect(page.getByRole('heading', { name: 'Alpha', exact: true })).toBeVisible()
  await page.getByRole('button', { name: 'Next page' }).click()
  await expect(page).toHaveURL(/page=2/)
  await expect(page.getByText(/Showing 21–40 of \d+ entries/)).toBeVisible()
  await expect(page.getByRole('heading', { name: 'Alpha', exact: true })).toHaveCount(0)
  // Back takes us to page 1 through history.
  await page.goBack()
  await expect(page.getByRole('heading', { name: 'Alpha', exact: true })).toBeVisible()
})

test('glossary clamps bad page params and resets on new queries', async ({ page }) => {
  await page.goto('/glossary?page=999')
  await expect(page.getByRole('heading', { name: 'Mathematical glossary' })).toBeVisible({ timeout: 30_000 })
  // Clamped to the last page: no Next page available.
  await expect(page.getByRole('button', { name: 'Next page' })).toBeDisabled()
  const box = page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })
  await box.fill('integral')
  await expect(page.getByRole('heading', { name: 'Integral', exact: true })).toBeVisible()
  await expect(page).not.toHaveURL(/page=/)
})

test('glossary related links jump across pages', async ({ page }) => {
  await page.goto('/glossary')
  await expect(page.getByRole('heading', { name: 'Mathematical glossary' })).toBeVisible({ timeout: 30_000 })
  // Summation lives far from the letter sigma; follow the cross-link.
  const box = page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })
  await box.fill('sigma')
  await expect(page.getByRole('heading', { name: 'Sigma', exact: true })).toBeVisible()
  await page.getByRole('button', { name: /Show Summation/ }).click()
  await expect(page.getByRole('heading', { name: 'Summation', exact: true })).toBeVisible()
})

test('glossary math mode searches by typed notation', async ({ page }) => {
  await page.goto('/glossary')
  await expect(page.getByRole('heading', { name: 'Mathematical glossary' })).toBeVisible({ timeout: 30_000 })
  await page.getByRole('button', { name: 'Math: f(x)' }).click()
  const field = page.locator('math-field')
  await expect(field).toBeVisible({ timeout: 30_000 })
  // Focus first: under parallel CI load the click alone can land before the
  // field is ready and keystrokes go nowhere (the same class of flake the
  // answer-input helpers guard against with explicit focus asserts).
  await field.click()
  await expect(field).toBeFocused({ timeout: 10_000 })
  // A pasted/tapped symbol arrives as the glyph itself (the mobile keypad
  // inserts commands that serialize the same way); typing a raw backslash
  // command char-by-char does not, which is MathLive input behavior, not search.
  const fieldLatex = (): Promise<string> =>
    field.evaluate(el => {
      const withValue = el as unknown as { getValue?: (f: string) => string; value?: string }
      try {
        return withValue.getValue ? withValue.getValue('latex') : (withValue.value ?? '')
      } catch {
        return withValue.value ?? ''
      }
    })
  await page.keyboard.type('∀', { delay: 40 })
  // Poll for arrival instead of sleeping once: dropped keystrokes under load
  // leave an empty field, and no fixed wait can distinguish that from parsing.
  try {
    await expect.poll(fieldLatex, { timeout: 15_000 }).not.toBe('')
  } catch {
    await field.click()
    await page.keyboard.type('∀', { delay: 80 })
    await expect.poll(fieldLatex, { timeout: 15_000 }).not.toBe('')
  }
  await expect(page.getByRole('heading', { name: 'Universal quantifier' })).toBeVisible({ timeout: 15_000 })
  // Switching back to text restores the text box and clears the math query.
  await page.getByRole('button', { name: 'Text', exact: true }).click()
  await expect(page.getByRole('searchbox', { name: 'Search symbols, names or meanings' })).toBeVisible({ timeout: 15_000 })
  await expect(page.getByRole('heading', { name: /Greek letters/ })).toBeVisible()
})

test('glossary tags toggle reveals category filters', async ({ page }) => {
  await page.goto('/glossary')
  await expect(page.getByRole('heading', { name: 'Mathematical glossary' })).toBeVisible({ timeout: 30_000 })
  const toggle = page.getByRole('button', { name: 'Tags' })
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(page.getByRole('button', { name: 'Geometry and trigonometry' })).toHaveCount(0)
  await toggle.click()
  await expect(toggle).toHaveAttribute('aria-expanded', 'true')
  await expect(page.getByRole('button', { name: 'Geometry and trigonometry' })).toBeVisible()
  await toggle.click()
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(page.getByRole('button', { name: 'Geometry and trigonometry' })).toHaveCount(0)
})
