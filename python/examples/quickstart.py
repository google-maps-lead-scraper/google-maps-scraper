#!/usr/bin/env python3
"""Minimal scrape example. Requires GMF_API_KEY."""

from gmaps_scraper import Client

KEYWORD = "dentists in Austin TX"


def main() -> None:
    client = Client()
    me = client.me()
    print(f"plan={me['plan']} creditsRemaining={me['creditsRemaining']}")

    rows = client.scrape(KEYWORD)
    print(f"scraped {len(rows)} places for {KEYWORD!r}")
    if rows:
        sample = rows[0]
        print(
            sample.get("Name"),
            sample.get("Phone"),
            sample.get("Website"),
            sample.get("Emails"),
        )


if __name__ == "__main__":
    main()
