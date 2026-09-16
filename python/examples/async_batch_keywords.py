#!/usr/bin/env python3
"""
Run multiple keywords sequentially.

The Agent API allows only one running job per user. Do not parallelize.
Requires GMF_API_KEY.
"""

from __future__ import annotations

from gmaps_scraper import Client, JobConflictError

KEYWORDS = [
    "dentists in Austin TX",
    "coffee shops in Austin TX",
]


def main() -> None:
    client = Client()
    for keyword in KEYWORDS:
        print(f"=== scraping {keyword!r} ===")
        try:
            rows = client.scrape(keyword)
        except JobConflictError:
            print("Another job is running (HTTP 409). Wait and retry.")
            raise
        print(f"got {len(rows)} rows")


if __name__ == "__main__":
    main()
