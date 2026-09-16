# Google Maps Scraper (GMaps Lead Finder Client)

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Python 3.10+](https://img.shields.io/badge/python-3.10+-3776AB.svg)](python/)
[![Node.js 18+](https://img.shields.io/badge/node-18+-339933.svg)](typescript/)

Official open-source **Python + TypeScript** client kit for [GMaps Lead Finder](https://gmapsleadfinder.com) — scrape Google Maps places (name, phone, website, emails, and more) through the hosted Agent HTTP API / Remote MCP.

This repo does **not** run a local browser crawler. It calls the same cloud scrape-and-enrich pipeline as the Online Lead Extractor.

## Get an API key

1. Sign in at [gmapsleadfinder.com](https://gmapsleadfinder.com) (Google OAuth).
2. Use a plan that includes Agent API (**Growth** or higher) — see [Pricing](https://gmapsleadfinder.com/pricing).
3. Open [Account → API key](https://gmapsleadfinder.com/account#api-key), copy your `gmf_…` key.

```bash
export GMF_API_KEY=gmf_your_key_here
```

Full walkthrough: [docs/getting-api-key.md](docs/getting-api-key.md).

## 60-second Quickstart

### Python

```bash
cd python
pip install -e .
gmaps-scraper me
gmaps-scraper scrape "dentists in Austin TX" --out leads.json
```

```python
from gmaps_scraper import Client

client = Client()  # reads GMF_API_KEY
rows = client.scrape("dentists in Austin TX")
print(len(rows), rows[0] if rows else None)
```

### TypeScript

```bash
cd typescript
npm install
npm run build
npx gmaps-scraper me
npx gmaps-scraper scrape "dentists in Austin TX" --out leads.json
```

```ts
import { Client } from "@gmapsleadfinder/google-maps-scraper";

const client = new Client(); // reads GMF_API_KEY
const rows = await client.scrape("dentists in Austin TX");
console.log(rows.length, rows[0]);
```

## Product & docs links

| Resource | URL |
|----------|-----|
| Website | https://gmapsleadfinder.com |
| Pricing | https://gmapsleadfinder.com/pricing |
| HTTP API docs | https://gmapsleadfinder.com/docs/api |
| Agent & MCP docs | https://gmapsleadfinder.com/docs/agent |
| Agents hub | https://gmapsleadfinder.com/agents |
| OpenAPI | https://gmapsleadfinder.com/openapi-agent.yaml |
| Account / API key | https://gmapsleadfinder.com/account#api-key |
| Support | support@gmapsleadfinder.com |

Repo docs: [HTTP](docs/http-api.md) · [MCP](docs/mcp.md) · [Python](docs/python.md) · [TypeScript](docs/typescript.md) · [AGENTS.md](AGENTS.md) · [llms.txt](llms.txt)

## Remote MCP (Claude / Cursor / Codex)

Prefer MCP when using an AI coding agent. Endpoint: `https://gmapsleadfinder.com/mcp`.

**Claude Code**

```bash
claude mcp add --transport http gmaps-finder https://gmapsleadfinder.com/mcp \
  --header "Authorization: Bearer $GMF_API_KEY"
```

**Cursor** (`~/.cursor/mcp.json` fragment)

```json
{
  "mcpServers": {
    "gmaps-finder": {
      "url": "https://gmapsleadfinder.com/mcp",
      "headers": {
        "Authorization": "Bearer gmf_your_key_here"
      }
    }
  }
}
```

MCP tools: `gmaps_me`, `gmaps_create_job`, `gmaps_get_job`, `gmaps_get_results`. Details: [docs/mcp.md](docs/mcp.md).

## Limits (read before batching)

- **Exactly one keyword per API job**
- **One running job per user** at a time (web + API + MCP share the lock) → HTTP `409` if busy
- **1 credit = 1 place row**; enrich is included
- Agent HTTP/MCP requires **Growth+**
- Results page size: `limit` 1–500 (default 100); follow `nextCursor`

## Repository layout

```text
python/          # SDK + CLI (gmaps-scraper)
typescript/      # SDK + CLI
docs/            # Human guides
openapi/         # OpenAPI snapshot
AGENTS.md        # Instructions for AI agents
llms.txt         # Machine-readable summary
```

## License

MIT — see [LICENSE](LICENSE).
