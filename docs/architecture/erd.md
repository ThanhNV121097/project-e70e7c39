# ERD

## `greetings`

Singleton current public greeting.

| Column | PostgreSQL type | Constraints | Purpose |
|---|---|---|---|
| `id` | `boolean` | primary key, `CHECK (id)` | Enforces one row keyed as `TRUE` |
| `text` | `text` | `NOT NULL`, `CHECK (btrim(text) <> '')` | Current trimmed greeting |
| `updated_at` | `timestamptz` | `NOT NULL`, default `now()` | Last successful write time |

Initial migration inserts `(TRUE, 'Hello, World!')` with conflict ignored. No relationships: product owns one current greeting and no users, history, or child records.

## Migration tracking

`schema_migrations(version text primary key, applied_at timestamptz not null default now())` is infrastructure metadata. It tracks ordered migration filenames; it has no product relationship.

## Integrity rules

`id = TRUE` makes second product row impossible. API trims before write; database check prevents blank data from any alternate writer. `updated_at` is set by API on update. No foreign keys, indexes, soft deletion, or audit trail: one-row lookup needs none. Add history only after stakeholder scope change.
