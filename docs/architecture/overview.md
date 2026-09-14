# Architecture Overview

## Stack

| Part | Choice | Reason |
|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind CSS v3, ESLint | Required UI stack; server composition root and client story components. |
| Backend | Go 1.22, `net/http`, `database/sql`, `pgx` stdlib adapter | Small JSON API; avoids framework overhead. |
| Database | PostgreSQL 16 | Required durable greeting storage. |
| Runtime | Docker Compose | Starts PostgreSQL, API, and UI together. |

## Layout

```text
code/backend/
  cmd/api/main.go       HTTP server, migration boot, health endpoint
  migrations/           Ordered SQL migration pairs
code/frontend/
  app/                  App Router shell and frozen shared CSS tokens
  components/           One default-export component per story
  lib/mock/             UI-stage mock data; removed when API lands
docs/architecture/      This overview, ERD, service contract
```

## Contracts and conventions

- API routes mount below `/v1`; deployment proxy owns `/api` prefix.
- API sends JSON and common error envelope from `services.md`.
- Backend reads `DATABASE_URL`, applies migrations before listening, then exposes database-backed `/healthz`.
- Migration names start with sortable timestamp and are recorded in `schema_migrations`; migration SQL is idempotent only through tracking.
- Go packages use lower-case names. Exported identifiers use PascalCase. JSON fields use lower camel case.
- Frontend components use PascalCase files and `export default function ComponentName()`.
- `app/page.tsx` only composes story components. Interactive components begin with literal first line `"use client"`.
- Shared visual values use tokens from `app/globals.css`; CSS modules must not hardcode design values or token fallbacks.
- Greeting text is public content. Server trims and rejects blank values; latest successful write wins.

## Environment

| Service | Variables |
|---|---|
| Backend | `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback |
| Frontend | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| Compose/PostgreSQL | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` |

Tracked `.env.example` files list every key and contain no secrets. Compose supplies local defaults; deployment injects runtime values.

## Run

1. Copy root `.env.example` to `.env` only when overriding local defaults.
2. Run `docker compose --profile local up --build` from repository root.
3. Open `http://localhost:3000`; API health is `http://localhost:8080/healthz`.

## Decisions

| Decision | Rejected alternative | Tradeoff |
|---|---|---|
| One-row greeting table | History or multi-record model | Cannot retain revisions; scope needs only current value. |
| Self-migrating API startup | Manual migration command | Startup owns empty runtime database; migration failure blocks readiness. |
| `net/http` routes | Web framework | Less routing abstraction; two routes do not justify dependency. |
| Client-side story editor | Server actions | Browser API contract remains explicit; one small client component required for form state. |

## Compatibility and rollout

First migration seeds `Hello, World!`. Existing empty databases upgrade on API boot. Later schema changes need ordered reversible migration pairs. No external integrations or credentials. Unknown: production `NEXT_PUBLIC_API_URL` must name deployed proxy URL, while `API_ORIGIN` must remain backend-reachable from frontend server.
