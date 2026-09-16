# AGENTS.md — GMaps Lead Finder client

Instructions for AI coding agents (Cursor, Claude Code, Codex, etc.) working in this repo or integrating the hosted API.

## What this project is

Official **client kit** for [GMaps Lead Finder](https://gmapsleadfinder.com). It talks to the **hosted** Agent HTTP API / Remote MCP. It is **not** a local Google Maps crawler.

- Base URL: `https://gmapsleadfinder.com`
- Auth: `Authorization: Bearer gmf_<key>`
- Env: `GMF_API_KEY` (optional `GMF_BASE_URL`)

## Preferred install (for end users / agents)

```bash
pip install google-maps-scraper-sdk
# import: from gmaps_scraper import Client

npm install @gmapsleadfinder/google-maps-scraper
# import: import { Client } from "@gmapsleadfinder/google-maps-scraper"

go get github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
# import: gmaps "github.com/google-maps-lead-scraper/google-maps-scraper/go"
```

## Prefer MCP vs HTTP

| Situation | Use |
|-----------|-----|
| User is in Claude / Cursor / Codex with MCP support | **Remote MCP** at `https://gmapsleadfinder.com/mcp` |
| Scripts, CI, custom apps | **HTTP** via Python/TS/Go SDK or raw `fetch`/`curl` |

MCP tools map 1:1 to HTTP: `gmaps_me`, `gmaps_create_job`, `gmaps_get_job`, `gmaps_get_results`.

## API key rules

1. Tell the user to get a key at https://gmapsleadfinder.com/account#api-key (Growth+ plan).
2. Store only in env / secret store — **never** hard-code, commit, or put in public frontends.
3. Keys look like `gmf_` + hex.

## HTTP workflow

1. `GET /api/v1/me` — confirm `creditsRemaining > 0` and plan allows API.
2. `POST /api/v1/jobs` with `{ "keyword": "<one query>" }` → `jobId`.
3. Poll `GET /api/v1/jobs/{id}` until `completed`, `partial`, or `failed`.
4. `GET /api/v1/jobs/{id}/results?limit=100` and follow `nextCursor` until `null`.

Or use SDK helpers: `scrape` / `Scrape` (Python / TypeScript / Go) which polls and paginates.

## Hard constraints

- Exactly **one** keyword per job (multi-keyword → `400`).
- Only **one** in-flight job per user (`409` if busy). For multiple keywords, run **sequentially** and wait for completion.
- `402` = no credits; `403` = plan cannot use Agent API; `401` = bad/missing key.
- Empty email/social fields mean nothing public was found — do not invent contacts.

## Result shape

Rows are objects keyed by export column headers, e.g. `Name`, `Phone`, `Website`, `Emails` (exact set follows the user’s Leads export preferences).

## Canonical links

- HTTP docs: https://gmapsleadfinder.com/docs/api
- Agent & MCP: https://gmapsleadfinder.com/docs/agent
- OpenAPI: https://gmapsleadfinder.com/openapi-agent.yaml (repo copy: `openapi/agent.yaml`)
- Pricing: https://gmapsleadfinder.com/pricing
- Repo guides: `docs/getting-api-key.md`, `docs/http-api.md`, `docs/mcp.md`, `docs/python.md`, `docs/typescript.md`, `go/README.md`

## When editing this repo

- Keep Python, TypeScript, and Go public APIs aligned.
- Do not add Playwright/Selenium scrapers or auth bypasses.
- Update examples and docs together when changing client surfaces.
- Go releases use Git tags `go/vX.Y.Z` (see `docs/publishing.md`).
