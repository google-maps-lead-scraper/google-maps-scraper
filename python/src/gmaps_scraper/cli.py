"""CLI: gmaps-scraper me | scrape "<keyword>" [--out path]."""

from __future__ import annotations

import argparse
import csv
import json
import sys
from pathlib import Path
from typing import Any

from . import __version__
from .client import Client
from .errors import ApiError


def _write_output(rows: list[dict[str, Any]], out: Path | None) -> None:
    if out is None:
        json.dump(rows, sys.stdout, ensure_ascii=False, indent=2)
        sys.stdout.write("\n")
        return

    out.parent.mkdir(parents=True, exist_ok=True)
    suffix = out.suffix.lower()
    if suffix == ".csv":
        fieldnames: list[str] = []
        seen: set[str] = set()
        for row in rows:
            for key in row:
                if key not in seen:
                    seen.add(key)
                    fieldnames.append(key)
        with out.open("w", encoding="utf-8", newline="") as fh:
            writer = csv.DictWriter(fh, fieldnames=fieldnames or ["Name"])
            writer.writeheader()
            for row in rows:
                writer.writerow({k: row.get(k, "") for k in fieldnames})
    else:
        out.write_text(
            json.dumps(rows, ensure_ascii=False, indent=2) + "\n",
            encoding="utf-8",
        )
    print(f"Wrote {len(rows)} rows to {out}", file=sys.stderr)


def cmd_me(_: argparse.Namespace) -> int:
    client = Client()
    me = client.me()
    json.dump(me, sys.stdout, ensure_ascii=False, indent=2)
    sys.stdout.write("\n")
    return 0


def cmd_scrape(args: argparse.Namespace) -> int:
    client = Client()
    rows = client.scrape(
        args.keyword,
        poll_interval_ms=args.poll_interval_ms,
        timeout_ms=args.timeout_ms,
    )
    out = Path(args.out) if args.out else None
    _write_output(rows, out)
    return 0


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        prog="gmaps-scraper",
        description="GMaps Lead Finder Agent API CLI",
    )
    parser.add_argument("--version", action="version", version=f"%(prog)s {__version__}")
    sub = parser.add_subparsers(dest="command", required=True)

    p_me = sub.add_parser("me", help="Show plan and credits")
    p_me.set_defaults(func=cmd_me)

    p_scrape = sub.add_parser("scrape", help="Scrape one keyword and print/save rows")
    p_scrape.add_argument("keyword", help="Exactly one Maps search keyword")
    p_scrape.add_argument(
        "--out",
        help="Output path (.json or .csv). Default: JSON to stdout",
    )
    p_scrape.add_argument(
        "--poll-interval-ms",
        type=int,
        default=2000,
        help="Job poll interval in milliseconds (default: 2000)",
    )
    p_scrape.add_argument(
        "--timeout-ms",
        type=int,
        default=600_000,
        help="Max wait for job completion in milliseconds (default: 600000)",
    )
    p_scrape.set_defaults(func=cmd_scrape)
    return parser


def main(argv: list[str] | None = None) -> None:
    parser = build_parser()
    args = parser.parse_args(argv)
    try:
        code = args.func(args)
    except ApiError as exc:
        print(f"error: {exc}", file=sys.stderr)
        if exc.status_code:
            print(f"status: {exc.status_code}", file=sys.stderr)
        raise SystemExit(1) from exc
    raise SystemExit(code)


if __name__ == "__main__":
    main()
