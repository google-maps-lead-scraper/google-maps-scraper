# Getting an API key

You need a personal **GMaps Lead Finder** Agent API key (`gmf_…`) to use this repository’s SDK, CLI, or Remote MCP.

## Steps

1. Open [gmapsleadfinder.com](https://gmapsleadfinder.com) and sign in with Google.
2. Subscribe to a plan that includes **Agent API** — **Growth** or higher. See [Pricing](https://gmapsleadfinder.com/pricing).
3. Go to [Account → API key](https://gmapsleadfinder.com/account#api-key) (`/account#api-key`).
4. Copy the key (or regenerate if needed).
5. Export it locally:

```bash
export GMF_API_KEY=gmf_your_key_here
```

Or copy [`.env.example`](../.env.example) to `.env` and fill in the value (never commit `.env`).

## Verify

```bash
curl -sS https://gmapsleadfinder.com/api/v1/me \
  -H "Authorization: Bearer $GMF_API_KEY"
```

You should see JSON with `plan`, `creditsLimit`, `creditsUsed`, and `creditsRemaining`.

- **401** — missing/invalid key  
- **403** — key works but your plan cannot use the Agent API  

## Security

- Treat the key like a password.
- Use it only from servers, CLIs, agents, or secret stores — not public websites.
- Rotate at Account → API key if exposed.
- Report security issues to support@gmapsleadfinder.com (see [SECURITY.md](../SECURITY.md)).

## Related

- [HTTP API](http-api.md)
- [MCP setup](mcp.md)
- Live docs: https://gmapsleadfinder.com/docs/api
