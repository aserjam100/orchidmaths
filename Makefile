run:
	go run ./cmd/orchid

build:
	go build -o bin/orchid ./cmd/orchid

build-all:
	mkdir -p dist
	GOOS=windows GOARCH=amd64 go build -o dist/orchid-windows-amd64.exe ./cmd/orchid
	GOOS=linux GOARCH=amd64 go build -o dist/orchid-linux-amd64 ./cmd/orchid
	GOOS=darwin GOARCH=amd64 go build -o dist/orchid-darwin-amd64 ./cmd/orchid
	GOOS=darwin GOARCH=arm64 go build -o dist/orchid-darwin-arm64 ./cmd/orchid

test:
	go test ./...

# The browser suite runs against the static export, so it needs a fresh `out/`
# first. CI shards this; locally it is one command. Nothing else runs it, which
# is how a renamed button and a moved page reached main green on the unit suite.
e2e:
	cd web/next-app && rm -rf out && npx next build
	cd web/next-app && npx playwright test

# Prove the browser suite does not depend on a real API. The bundle is normally
# built against the deployed API (web/next-app/.env.production), so an endpoint no
# stub claims is answered by production, and a page that waits for that answer can
# pass on a fast reply and hang on a slow one. This build points every request at
# 192.0.2.1, a documentation address that never answers, so a stub that misses its
# URL is a failure here rather than a flake in CI. See CONTRIBUTING.md.
e2e-hermetic:
	cd web/next-app && rm -rf out && NEXT_PUBLIC_API_URL=http://192.0.2.1:9099 npx next build
	cd web/next-app && npx playwright test

validate:
	go run scripts/validate_graph.go
	go run ./cmd/validate-mastery-units
	python3 scripts/audit_lessons.py
	python3 scripts/audit_solutions.py
	python3 scripts/counts.py --check
	python3 scripts/check_docs.py
	python3 scripts/curriculum_moe.py --check
	python3 scripts/check_ste.py
	node web/next-app/scripts/build-docs.mjs --check
	go test ./internal/generator/ -run TestLearnerDomainsHaveSchemas

# Regenerate the MOE Primary courses, destination and coverage report from
# data/curricula/sg-moe-primary-2021.json.
curriculum:
	python3 scripts/curriculum_moe.py

# Redraw Orchid's footer, login and icon artwork (needs rsvg-convert).
art:
	cd web/next-app && python3 art/orchid-art.py
	cd web/next-app && for f in art/out/footer-*.svg; do rsvg-convert "$$f" -o public/footer/$$(basename $${f%.svg}).png; done
	cd web/next-app && rsvg-convert art/out/login-orchid.svg -o public/hero/login-orchid.png
	cd web/next-app && sed 's|translate(256,246)">|translate(256,246) scale(0.72)">|' public/icons/icon.svg > art/out/icon-maskable.svg
	cd web/next-app && rsvg-convert -w 192 -h 192 public/icons/icon.svg -o public/icons/icon-192.png
	cd web/next-app && rsvg-convert -w 512 -h 512 public/icons/icon.svg -o public/icons/icon-512.png
	cd web/next-app && rsvg-convert -w 512 -h 512 art/out/icon-maskable.svg -o public/icons/icon-maskable-512.png
	cd web/next-app && rsvg-convert -w 64 -h 64 public/icons/icon.svg -o art/out/fav64.png && sips -s format ico art/out/fav64.png --out public/favicon.ico

# Strict LaTeX check: runs every math span in the lesson corpus through Compute
# Engine and reports the malformed ones. Needs node (for the CE dependency) and
# the Go toolchain (for latexdump). Exits 1 while defects remain, so it is not in
# `validate` yet -- see the CI step for why.
# The mastery-unit partition is a generated draft (ADR-046 D3). The band width is a parameter,
# not policy: 2, 3 and 4 are all candidates and the width decides how often a learner is
# assessed. Regenerating at a different width needs no code change.
mastery-units:
	go run ./cmd/masteryunits -band 3

latex-normalize:
	cd web/next-app && node scripts/normalize-latex.mjs

lint-go:
	go run github.com/curtbushko/go-ai-lint/cmd/go-ai-lint@v1.0.1-0.20260620203811-c6ce4ee5624f ./...

lint-slop:
	npx --yes aislop@0.16.1 scan

fuzz:
	go test ./internal/generator/... -run TestFuzz -count 1000

tidy:
	go mod tidy

serve:
	go run ./cmd/orchid --serve --port 8080

# One command for local development. The app is two programs: the Go API (port 8080) and the
# Next.js dev server (port 3000), which reloads the page on every save. This starts both and
# stops both on Ctrl-C. Open http://localhost:3000.
dev:
	go build -o bin/orchid ./cmd/orchid
	mkdir -p temp
	@trap 'kill 0' INT TERM EXIT; \
	PATH="$(CURDIR)/venv/bin:$$PATH" DATABASE_URL=sqlite:temp/orchid.db ./bin/orchid --serve --port 8080 & \
	cd web/next-app && NEXT_PUBLIC_API_URL=http://localhost:8080 npm run dev

# One program on one port: builds the site, then the Go server serves it with the API.
# No live reload; run it again after a change. Open http://localhost:8080.
start:
	go build -o bin/orchid ./cmd/orchid
	cd web/next-app && NEXT_PUBLIC_API_URL=http://localhost:8080 npx next build
	mkdir -p temp
	PATH="$(CURDIR)/venv/bin:$$PATH" DATABASE_URL=sqlite:temp/orchid.db ./bin/orchid --serve --port 8080
