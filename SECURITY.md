# Security Policy

## Reporting a vulnerability

Email **support@gmapsleadfinder.com** with a clear description and steps to reproduce.

Do **not** open a public GitHub issue for security-sensitive reports.

## API keys

- Keys start with `gmf_` and are personal credentials.
- Store them in environment variables (`GMF_API_KEY`) or a secret manager.
- Never commit keys, paste them into public chats, or embed them in front-end code.
- Rotate a compromised key at [Account → API key](https://gmapsleadfinder.com/account#api-key).
