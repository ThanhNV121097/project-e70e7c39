# Architecture Overview

## Stack

| Part | Choice |
|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind CSS v3, ESLint |
| Backend | Go 1.25 HTTP server with `net/http` and `pgx` |
| Database | PostgreSQL 16 |
| Runtime | `docker compose --profile local up --build` |

## Layout

- `code/backend/cmd/api`: single `main` package and HTTP entrypoint.
- `code/backend/migrations`: ordered SQL migration pairs, applied at startup.
- `code/frontend/app`: App Router shell and frozen shared CSS tokens.
- `code/frontend/components`: one default-exported component per story.
- `code/frontend/lib/mock`: UI-only fixtures, deleted when API lands.
- `docs/architecture`: shared architecture, ERD, API contracts.

## Contracts and conventions

Backend reads `DATABASE_URL`, applies migrations before listening, and exposes database-backed `/healthz`. Migration names sort lexically and records live in `schema_migrations`. API routes use `/v1/...`; deployment proxy owns `/api` prefix. JSON errors share documented envelope.

Frontend `app/page.tsx` is server composition root. Interactive story component starts with literal first line `"use client"`. Components default-export PascalCase function. CSS modules use only tokens from `app/globals.css`; globals covers color, spacing, typography, radius, shadow, motion.

PostgreSQL is required for restart-safe greeting. Rejected in-memory store: shorter scaffold, loses saved state. `net/http` plus `pgx` replaces heavier router/ORM: two endpoints and one table need neither.

## Environment

| Service | Keys |
|---|---|
| Backend | `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback |
| Frontend | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| Compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, port/resource overrides |

Examples list names only; never commit secrets.

## Run

Copy root `.env.example` to `.env` if overrides needed. Run `docker compose --profile local up --build`. Frontend: `http://localhost:3000`; backend health: `http://localhost:8080/healthz`. `docker compose --profile local down` stops stack; add `-v` only to remove persisted local data.

## Compatibility and rollout

Migrations run transactionally and only once per database. Initial migration seeds one row. Later schema changes require new ordered up/down pairs; never edit applied migration. No external services or auth.

## Requirement coverage

Greeting story reads current singleton through `GET /v1/greeting`, updates it through `PUT /v1/greeting`, trims text at API boundary, and renders status behavior in client component. Database remains source of truth after reload or restart. Validation rejects blank text; later successful write wins.

## Observability and failure handling

Server logs startup, migration failures, and request errors without credentials or greeting contents. It does not listen until database migration and probe succeed. Health endpoint returns non-200 when database unavailable. API returns generic JSON errors; frontend handles only approved validation and success copy in story scope.

## Product requirements retained

One public page only. Centered `main` holds one centered section, one `h1`, one labelled text input, one Save submit button, one polite status region. Default value is `Hello, World!`. Input and button remain 44px high; controls stack at 520px and below. No navigation, sign-in, external service, multiple greetings, history, animation, loading/error screen, or extra sections.

Saved value is trimmed. Blank or whitespace-only submission preserves stored value, focuses input, and shows `Enter a greeting before saving.`. Success updates heading and field without navigation and shows `Saved.`. Long unbroken greeting wraps with no horizontal page scroll from 320px. Public concurrent saves use last successful save.

## Accessibility

Input label is programmatic and visually hidden. Section references stable `h1` ID. Status has `role="status"` and `aria-live="polite"`. Input and button use visible 3px focus outline. Color contrast follows design-system audit. Heading explicitly uses heading weight and tight tracking because Tailwind preflight resets browser heading weight.

## Decisions

| Decision | Rejected alternative | Tradeoff |
|---|---|---|
| PostgreSQL singleton record | Browser or process memory | Requires migration; survives restart |
| Go `net/http` | Router framework | Manual routing; no extra dependency |
| `pgx` driver | ORM | SQL remains explicit; one table needs no mapping layer |
| Next App Router | Pages Router | Server/client boundary rules; matches stack |
| CSS custom-property tokens | Story-local visual literals | Initial token work; prevents visual drift |

## Risks and unknowns

| Item | Handling |
|---|---|
| Empty runtime database | Startup migration creates schema and default row |
| Database unreachable | Startup fails; health stays unavailable |
| Future multiple greetings | Add keyed entity and migration after scope approval |
| API error display | Not approved in design; keep generic contract until design changes |
| Deployment proxy path | Backend mounts `/v1`; proxy strips `/api` |
