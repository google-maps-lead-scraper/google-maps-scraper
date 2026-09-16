# HTTP API

Canonical live reference: [gmapsleadfinder.com/docs/api](https://gmapsleadfinder.com/docs/api)  
OpenAPI: [openapi-agent.yaml](https://gmapsleadfinder.com/openapi-agent.yaml) (repo copy: [`openapi/agent.yaml`](../openapi/agent.yaml))

## Base URL & auth

```http
Authorization: Bearer gmf_<your_key>
```

Base: `https://gmapsleadfinder.com`  
Optional override env: `GMF_BASE_URL`

## Endpoints

### `GET /api/v1/me`

Plan and credits snapshot.

```bash
curl -sS https://gmapsleadfinder.com/api/v1/me \
  -H "Authorization: Bearer $GMF_API_KEY"
```

### `POST /api/v1/jobs`

Queue a single-keyword scrape + enrich job.

```bash
curl -sS -X POST https://gmapsleadfinder.com/api/v1/jobs \
  -H "Authorization: Bearer $GMF_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"keyword":"dentists in Austin TX"}'
```

Response includes `jobId`, `keywordCount` (always `1`), `creditsRemaining`.

### `GET /api/v1/jobs/{id}`

Poll until `status` is `completed`, `partial`, or `failed` (also expect `queued` / `running` while in progress).

### `GET /api/v1/jobs/{id}/results`

Paginated rows. Query: `limit` (1–500, default 100), `cursor` (from previous `nextCursor`).

Follow `nextCursor` until it is `null`. Each row is `{ [columnHeader]: string }`.

## Typical workflow

1. `GET /api/v1/me` — confirm credits.
2. `POST /api/v1/jobs` — store `jobId`.
3. Poll `GET /api/v1/jobs/{id}`.
4. Fetch all pages of `…/results`.

SDKs expose the same flow via `Client.scrape()`.

## Errors & limits

| Code | Meaning |
|------|---------|
| 400 | Bad request (e.g. not exactly one keyword) |
| 401 | Missing or invalid Bearer key |
| 402 | No credits remaining |
| 403 | Plan cannot use Agent API |
| 404 | Job not found |
| 409 | Another job already running for this user |

Notes:

- One credit = one place row.
- Web UI, HTTP API, and MCP share one running-job lock.
- Empty contact fields mean nothing public was found.

## Prefer MCP?

If you use Claude, Cursor, or Codex, install Remote MCP instead of hand-rolling HTTP. See [mcp.md](mcp.md) and https://gmapsleadfinder.com/docs/agent.
