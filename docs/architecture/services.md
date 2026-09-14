# Service Contracts

Base path is `/v1`; deployment proxy strips external `/api`. All JSON uses `Content-Type: application/json`.

## Error envelope

Every non-2xx response:

```json
{"error":{"code":"invalid_greeting","message":"Greeting must not be blank."}}
```

`code` is stable machine text; `message` is generic safe text. No database details or input echo.

## Get current greeting

`GET /v1/greeting`

Success `200`:

```json
{"greeting":"Hello, World!"}
```

Failures: `503 database_unavailable` when storage cannot serve request; `500 internal_error` otherwise.

## Replace current greeting

`PUT /v1/greeting`

Request:

```json
{"greeting":"Pipeline accepted"}
```

Server trims surrounding whitespace before validation and persistence. Success `200`:

```json
{"greeting":"Pipeline accepted"}
```

Failures: `400 invalid_json` for malformed body; `422 invalid_greeting` for missing, empty, or whitespace-only greeting; `503 database_unavailable`; `500 internal_error`. Later successful request wins.

## Health

`GET /healthz` returns `200` with `{"status":"ok"}` only after migrations completed and `SELECT 1` succeeds. It returns `503` with shared error envelope otherwise. Health endpoint is operational, not proxy-prefixed.

## Boundary rules

Only `GET` and `PUT` routes above exist. Requests use one JSON object, reject unknown shape only when contract grows. Backend caps request body at 64 KiB, accepts public text as UTF-8 JSON, and sends no-cache response headers for mutable greeting. CORS permits frontend origin in local compose; production proxy serves same public origin.
