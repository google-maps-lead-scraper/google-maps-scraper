# Google Maps Scraper

<p align="center">
  <img src="https://raw.githubusercontent.com/google-maps-lead-scraper/google-maps-scraper/main/google-maps-scraper.png" alt="Google Maps Lead Scraper" width="100%" />
</p>

**Google Maps Extractor · Google Maps Lead Scraper · Google Maps Lead Extractor**

Rust SDK to scrape Google Maps places and export leads (name, phone, website, emails, and more) through a hosted Agent HTTP API. Powered by [GMaps Lead Finder](https://gmapsleadfinder.com). This is **not** a local browser crawler — jobs run in the cloud scrape-and-enrich pipeline.

- **crates.io:** [`google-maps-scraper-sdk`](https://crates.io/crates/google-maps-scraper-sdk)
- **docs.rs:** https://docs.rs/google-maps-scraper-sdk
- **CLI:** `gmaps-scraper`

## Get an API key

1. Sign in at [gmapsleadfinder.com](https://gmapsleadfinder.com).
2. Use a plan with Agent API (**Growth** or higher) — see [Pricing](https://gmapsleadfinder.com/pricing).
3. Open [Account → API key](https://gmapsleadfinder.com/account#api-key) and copy your `gmf_…` key.

```bash
export GMF_API_KEY=gmf_your_key_here
```

## Install

```bash
cargo add google-maps-scraper-sdk
# or: cargo install google-maps-scraper-sdk --bin gmaps-scraper
```

Requires **Rust 1.70+**.

## Quickstart

```rust
use google_maps_scraper_sdk::{Client, ClientOptions};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::new(ClientOptions::default())?; // reads GMF_API_KEY
    let me = client.me()?;
    println!("{} {}", me.plan, me.credits_remaining);

    let rows = client.scrape("dentists in Austin TX", Default::default())?;
    println!("{} {:?}", rows.len(), rows.first());
    Ok(())
}
```

CLI:

```bash
cargo install google-maps-scraper-sdk --bin gmaps-scraper
gmaps-scraper me
gmaps-scraper scrape "dentists in Austin TX" --out leads.json
```

From this repo:

```bash
cd rust
cargo run --example quickstart
cargo run --bin gmaps-scraper -- me
```

## Client API

| Method | Description |
|--------|-------------|
| `Client::new(opts)` | From env / options |
| `me()` | Plan & credits |
| `create_job(keyword)` | Queue one-keyword job |
| `get_job(job_id)` | Poll status |
| `get_results(job_id, opts)` | Paginated rows |
| `scrape(keyword, opts)` | Create → poll → all rows |

## Errors

`Error` variants: `Authentication` (401), `PlanNotAllowed` (403), `InsufficientCredits` (402), `JobConflict` (409), `NotFound` (404), `BadRequest` (400), `Timeout`, `Api`, `Http`, `Json`.

## Limits

- Exactly one keyword per API job
- One running job per user (`409` if busy)
- 1 credit = 1 place row; Growth+ required

## Publish to crates.io

```bash
cargo login
cd rust
cargo publish --dry-run
cargo publish
```

## Links

| Resource | URL |
|----------|-----|
| Website | https://gmapsleadfinder.com |
| HTTP API | https://gmapsleadfinder.com/docs/api |
| crates.io | https://crates.io/crates/google-maps-scraper-sdk |
| Source | https://github.com/google-maps-lead-scraper/google-maps-scraper |

## License

MIT
