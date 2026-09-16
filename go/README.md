# Google Maps Scraper

<p align="center">
  <img src="https://raw.githubusercontent.com/google-maps-lead-scraper/google-maps-scraper/main/google-maps-scraper.png" alt="Google Maps Lead Scraper" width="100%" />
</p>

**Google Maps Extractor · Google Maps Lead Scraper · Google Maps Lead Extractor**

Go SDK to scrape Google Maps places and export leads (name, phone, website, emails, and more) through a hosted Agent HTTP API. Powered by [GMaps Lead Finder](https://gmapsleadfinder.com). This is **not** a local browser crawler — jobs run in the cloud scrape-and-enrich pipeline.

- **Module:** [`github.com/google-maps-lead-scraper/google-maps-scraper/go`](https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go)
- **Import:** `gmaps "github.com/google-maps-lead-scraper/google-maps-scraper/go"`
- **Docs:** [pkg.go.dev](https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go)

## Get an API key

1. Sign in at [gmapsleadfinder.com](https://gmapsleadfinder.com).
2. Use a plan with Agent API (**Growth** or higher) — see [Pricing](https://gmapsleadfinder.com/pricing).
3. Open [Account → API key](https://gmapsleadfinder.com/account#api-key) and copy your `gmf_…` key.

```bash
export GMF_API_KEY=gmf_your_key_here
```

## Install

```bash
go get github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

Requires **Go 1.22+**. Standard library only.

## Quickstart

```go
package main

import (
	"fmt"
	"log"

	gmaps "github.com/google-maps-lead-scraper/google-maps-scraper/go"
)

func main() {
	client, err := gmaps.NewClient(nil) // reads GMF_API_KEY
	if err != nil {
		log.Fatal(err)
	}

	me, err := client.Me()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(me.Plan, me.CreditsRemaining)

	rows, err := client.Scrape("dentists in Austin TX", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(rows), rows[0])
}
```

CLI from this module:

```bash
cd go
go run ./cli me
go run ./cli scrape "dentists in Austin TX" --out leads.json
```

## Client API

| Method | Description |
|--------|-------------|
| `NewClient(opts)` | Build client from env / options |
| `Me()` | Plan & credits |
| `CreateJob(keyword)` | Queue one-keyword job |
| `GetJob(jobID)` | Poll status |
| `GetResults(jobID, opts)` | Paginated rows (`Limit`, `Cursor`) |
| `Scrape(keyword, opts)` | Create → poll → all rows |

`ScrapeOptions`: `PollIntervalMs` (default 2000), `TimeoutMs` (default 600000), `ResultLimit` (default 100).

## Errors

Typed errors: `AuthenticationError` (401), `PlanNotAllowedError` (403), `InsufficientCreditsError` (402), `JobConflictError` (409), `NotFoundError` (404), `BadRequestError` (400), `TimeoutError`, `APIError`.

## Limits

- Exactly one keyword per API job
- One running job per user (`409` if busy)
- 1 credit = 1 place row; Growth+ required

## Publish to pkg.go.dev

Go modules are published via **Git tags** (no separate upload):

```bash
git tag go/v0.1.1
git push origin go/v0.1.1
GOPROXY=https://proxy.golang.org go list -m github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

Then open: https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1

## Links

| Resource | URL |
|----------|-----|
| Website | https://gmapsleadfinder.com |
| HTTP API | https://gmapsleadfinder.com/docs/api |
| Agent & MCP | https://gmapsleadfinder.com/docs/agent |
| pkg.go.dev | https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go |
| Source | https://github.com/google-maps-lead-scraper/google-maps-scraper |

## License

MIT
