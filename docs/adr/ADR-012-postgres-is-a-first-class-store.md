# ADR-012: Postgres is a first-class store

*2026-09-27*

`internal/storage/parity_test.go` runs one identical scenario against SQLite and Postgres and requires the same observable result; CI runs it against a `postgres:16` service and `deploy` depends on it (`TEST_POSTGRES_DSN` gates the test). The parity test caught `GetDailyActivity` comparing a TEXT `timestamp` to `timestamptz` — now cast with `::timestamptz`. Known remaining divergence: Postgres `attempts`/`active_sessions` omit the SQLite foreign keys (align in a later pass).
