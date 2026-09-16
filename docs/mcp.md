# Remote MCP

Hosted MCP endpoint for AI clients. Live guide: [gmapsleadfinder.com/docs/agent](https://gmapsleadfinder.com/docs/agent) · Agents hub: [gmapsleadfinder.com/agents](https://gmapsleadfinder.com/agents)

## Endpoint

- URL: `https://gmapsleadfinder.com/mcp`
- Transport: Streamable HTTP (JSON-RPC)
- Auth: `Authorization: Bearer gmf_<key>`
- Server name (suggested): `gmaps-finder`

Same plan and credit rules as the HTTP Agent API (Growth+).

## Tools

| Tool | Purpose |
|------|---------|
| `gmaps_me` | Plan & credits |
| `gmaps_create_job` | Create one-keyword job |
| `gmaps_get_job` | Job status |
| `gmaps_get_results` | Paginated results |

## Install snippets

### Claude Code

```bash
claude mcp add --transport http gmaps-finder https://gmapsleadfinder.com/mcp \
  --header "Authorization: Bearer $GMF_API_KEY"
```

### Cursor

Add to MCP config (e.g. `~/.cursor/mcp.json`):

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

### Codex

```toml
[mcp_servers.gmaps-finder]
url = "https://gmapsleadfinder.com/mcp"
http_headers = { Authorization = "Bearer gmf_your_key_here" }
```

### Generic MCP client

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

## Agent tips

- Create one job, wait until complete, then fetch results (or ask the agent to poll).
- For multiple keywords, run jobs **one after another** — concurrent jobs return `409`.
- Never paste API keys into shared transcripts; use env vars or the product’s Agents UI when available.

## Related

- [Getting an API key](getting-api-key.md)
- [HTTP API](http-api.md)
- [AGENTS.md](../AGENTS.md)
