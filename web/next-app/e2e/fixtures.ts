import { test as base } from '@playwright/test'

// Every spec imports `test` and `expect` from here rather than from
// `@playwright/test`.
//
// The assistant dock is chrome: it is mounted once in `app/layout.tsx`, so every
// page asks `/api/config` on load to find out whether the assistant is on. That
// is one request the suite never stubbed, and it is not a request to the page's
// own server — a normal `next build` bakes `NEXT_PUBLIC_API_URL` from
// `.env.production`, which is the deployed API. So `/glossary`, a purely static
// reference, was reaching `https://orchid.fly.dev` from `localhost:4173`, and
// the browser logged the CORS refusal. `glossary.spec.ts` asserts that the
// console holds no errors, and it was right to fail: the page should not need
// the network, and now it did.
//
// Answering the request locally is the honest fix. The stub returns
// `ai_enabled: false`, which is the state a visitor with no assistant sees — the
// trigger still renders (that is deliberate, see ADR-055) and no further request
// is made, so a spec that asserts "this page renders without a network" keeps
// meaning what it says. A spec that wants the assistant on can register its own
// route for `/api/config` after this one; the last registration wins.
export const test = base.extend({
  page: async ({ page }, use) => {
    await page.route('**/api/config', (route) => route.fulfill({ json: { ai_enabled: false } }))
    await use(page)
  },
})

export * from '@playwright/test'
