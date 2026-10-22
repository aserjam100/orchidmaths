# ADR-015: Stub lessons fail the gate

*2026-09-28*

9 committed lessons were SEO/placeholder stubs (5 Algebrica bodies that never landed, 4 fraction word-problem templates); the audit only scanned `teaching/*.md` non-recursively and never `algebrica/`, so it missed all of them. The gate now scans recursively and flags tiny word-count + boilerplate files, validated to zero false positives against the corpus. Remedy: refetch Algebrica bodies via `scripts/refetch_algebrica.py` (one-time network, then committed), author word lessons to match their KP sections, and repoint semantically wrong mappings (`alg.systems.elimination`/`substitution` pointed at the Rouché–Capelli theorem page) with rewritten KPs. One lesson serving many concepts stays legal; the gate checks sections resolve, not 1:1 files.
