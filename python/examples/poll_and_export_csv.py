#!/usr/bin/env python3
"""Create a job, poll manually, export CSV. Requires GMF_API_KEY."""

from __future__ import annotations

import csv
import time
from pathlib import Path

from gmaps_scraper import Client
from gmaps_scraper.models import TERMINAL_STATUSES

KEYWORD = "dentists in Austin TX"
OUT = Path("leads.csv")


def main() -> None:
    client = Client()
    created = client.create_job(KEYWORD)
    job_id = created["jobId"]
    print(f"jobId={job_id}")

    while True:
        job = client.get_job(job_id)
        status = (job.get("status") or "").lower()
        print(f"status={status} rowCount={job.get('rowCount')}")
        if status in TERMINAL_STATUSES:
            if status == "failed":
                raise SystemExit(job.get("error") or "job failed")
            break
        time.sleep(2)

    rows: list[dict] = []
    cursor: str | None = "0"
    while cursor is not None:
        page = client.get_results(job_id, limit=100, cursor=cursor)
        rows.extend(page.get("rows") or [])
        nxt = page.get("nextCursor")
        cursor = None if nxt is None else str(nxt)

    fieldnames: list[str] = []
    seen: set[str] = set()
    for row in rows:
        for key in row:
            if key not in seen:
                seen.add(key)
                fieldnames.append(key)

    with OUT.open("w", encoding="utf-8", newline="") as fh:
        writer = csv.DictWriter(fh, fieldnames=fieldnames or ["Name"])
        writer.writeheader()
        for row in rows:
            writer.writerow({k: row.get(k, "") for k in fieldnames})

    print(f"Wrote {len(rows)} rows to {OUT}")


if __name__ == "__main__":
    main()
