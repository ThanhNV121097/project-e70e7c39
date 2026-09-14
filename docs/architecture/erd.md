# Entity Relationship Diagram

## Tables

### `greetings`

| Column | PostgreSQL type | Constraints | Purpose |
|---|---|---|---|
| `id` | `smallint` | primary key; must equal `1` | Enforces single current greeting row. |
| `text` | `text` | not null; trimmed non-empty enforced by API | Current public greeting. |
| `updated_at` | `timestamptz` | not null; default `now()` | Last successful save time. |

Seed migration inserts `(1, 'Hello, World!')`. Reads address `id = 1`; updates target same row. No foreign keys or relationships: scope has one independent persisted record.

## Migration rules

- Migration pair format: `YYYYMMDDHHMMSS_name.up.sql` and matching `.down.sql`.
- API applies `.up.sql` files in lexical order and records names in `schema_migrations`.
- Down migration removes table only for local rollback before dependent data exists.
