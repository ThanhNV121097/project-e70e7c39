# Service Contracts

Base path reaches backend after deployment proxy strips `/api`. Paths below deliberately omit that prefix.

## Shared rules

- Content type is `application/json; charset=utf-8` for JSON responses.
- Greeting text is trimmed by API. Empty or whitespace-only values are invalid.
- No authentication. Last successful update wins.
- Errors use one envelope:

```json
{"error":{"code":"invalid_request","message":"Greeting must not be blank."}}
```

`code` is stable machine text. `message` is safe plain text. Never return database or internal error details.

## Endpoints

### `GET /v1/greeting`

Returns current stored greeting.

**200 response**

```json
{"greeting":"Hello, World!"}
```

**Errors:** `500 internal_error` when persistence cannot be read.

### `PUT /v1/greeting`

Replaces current greeting.

**Request**

```json
{"greeting":"Pipeline accepted"}
```

**200 response**

```json
{"greeting":"Pipeline accepted"}
```

**Errors:** `400 invalid_request` for malformed JSON, missing greeting, or blank trimmed greeting; `405 method_not_allowed` for unsupported methods; `500 internal_error` when persistence cannot be updated.

### `GET /healthz`

Operational endpoint. Returns `200` only after migrations complete and database `SELECT 1` succeeds. Response body is plain `ok\n`. Returns `503` otherwise.
