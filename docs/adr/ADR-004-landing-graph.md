# ADR-004: Landing graph

Landing graph: keep three.js/R3F, drop drei; fetch graph payload/positions instead of bundling the corpus; deterministic domain-sector spherical layout (`lib/layoutDAG3D.ts`, O(n log n), 5k nodes ≈24ms); demand frameloop with auto-rotate stopping after first interaction (2026-09-11) — mobile First Load JS and 1k–5k scale without a renderer migration
