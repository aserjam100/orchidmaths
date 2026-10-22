# Question Seed Script

Populates the orchid question bank with practice questions for each concept.

## Usage

### Generate questions using built-in generators

```bash
# From project root — seeds 8 questions per concept into orchid.db
go run scripts/seed-questions/main.go -db orchid.db -n 8
```

### Import questions from a JSON file

```bash
go run scripts/seed-questions/main.go -db orchid.db -import questions.json
```

## JSON Import Format

```json
[
  {
    "concept_id": "alg.quad.formula",
    "question": "Solve x^2 - 5x + 6 = 0",
    "answer": "x = 2, x = 3",
    "explanation": "Factor: (x-2)(x-3) = 0, so x = 2 or x = 3",
    "source": "math_dataset",
    "difficulty": 0.5
  }
]
```

Fields:
- `concept_id` (required) — must match a orchid concept ID (e.g. `alg.quad.formula`)
- `question` (required) — the problem statement (supports LaTeX)
- `answer` (required) — the final answer
- `explanation` (optional) — step-by-step solution
- `source` (optional) — name of the source dataset
- `difficulty` (optional, default 0.5) — 0.0 to 1.0

## Suggested External Datasets

| Dataset | License | Size | How to convert |
|---------|---------|------|----------------|
| [MATH (hendrycks)](https://github.com/hendrycks/math) | MIT | 12,500 | Use `scripts/convert-math.py` |
| [StackMathQA](https://huggingface.co/datasets/math-ai/StackMathQA) | CC BY-SA | 2M | Use HuggingFace `datasets` library |
| [OpenStax](https://openstax.org/) | CC BY-NC-SA 4.0 | 1000s per book | Scrape exercises from book JSON |

OpenStax exercises are non-commercial and share-alike, and each book's own panel forbids use "in
the training of large language models or otherwise ... in generative AI offerings" without
written permission. Read `THIRD_PARTY_NOTICES.md` before loading that row. It read CC BY 4.0 here
until 2026-10-17, which is wrong for every title this project has read.

Write a Python script that loads any of these, maps the subject tags to orchid concept IDs, and exports in the JSON format above.
