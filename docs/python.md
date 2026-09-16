# Python SDK

Package lives in [`python/`](../python/).

## Install

```bash
cd python
pip install -e .
```

Requires Python 3.10+. Uses the standard library only (no third-party HTTP deps).

## Configuration

```bash
export GMF_API_KEY=gmf_your_key_here
# optional:
# export GMF_BASE_URL=https://gmapsleadfinder.com
```

You can also pass `api_key=` / `base_url=` into `Client(...)`.

## Usage

```python
from gmaps_scraper import Client

client = Client()

me = client.me()
print(me["creditsRemaining"])

# High-level: create → poll → fetch all rows
rows = client.scrape("dentists in Austin TX")
for row in rows[:5]:
    print(row.get("Name"), row.get("Phone"), row.get("Website"))
```

### Low-level methods

```python
created = client.create_job("dentists in Austin TX")
job_id = created["jobId"]

job = client.get_job(job_id)
page = client.get_results(job_id, limit=100, cursor="0")
```

`scrape()` options: `poll_interval_ms` (default `2000`), `timeout_ms` (default `600_000`), `result_limit` (default `100`).

## CLI

```bash
gmaps-scraper me
gmaps-scraper scrape "dentists in Austin TX" --out leads.json
gmaps-scraper scrape "dentists in Austin TX" --out leads.csv
```

## Examples

- [`python/examples/quickstart.py`](../python/examples/quickstart.py)
- [`python/examples/poll_and_export_csv.py`](../python/examples/poll_and_export_csv.py)
- [`python/examples/async_batch_keywords.py`](../python/examples/async_batch_keywords.py) — sequential keywords (respects single running job)

## Errors

Typed exceptions in `gmaps_scraper.errors`: `AuthenticationError` (401), `PlanNotAllowedError` (403), `InsufficientCreditsError` (402), `JobConflictError` (409), `NotFoundError` (404), `ApiError` (other).

## Related

- [HTTP API](http-api.md)
- Live docs: https://gmapsleadfinder.com/docs/api
