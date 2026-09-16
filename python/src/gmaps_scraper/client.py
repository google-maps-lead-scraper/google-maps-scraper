"""HTTP client for GMaps Lead Finder Agent API."""

from __future__ import annotations

import json
import os
import time
import urllib.error
import urllib.parse
import urllib.request
from typing import Any

from .errors import (
    ApiError,
    AuthenticationError,
    BadRequestError,
    InsufficientCreditsError,
    JobConflictError,
    NotFoundError,
    PlanNotAllowedError,
    TimeoutError,
)
from .models import (
    TERMINAL_STATUSES,
    CreateJobResponse,
    JobResponse,
    MeResponse,
    ResultsResponse,
)

DEFAULT_BASE_URL = "https://gmapsleadfinder.com"
DEFAULT_POLL_INTERVAL_MS = 2000
DEFAULT_TIMEOUT_MS = 600_000
DEFAULT_RESULT_LIMIT = 100
USER_AGENT = "gmaps-scraper-python/0.1.0"


def _raise_for_status(status: int, body: Any) -> None:
    message = None
    if isinstance(body, dict):
        message = body.get("message") or body.get("error") or body.get("statusMessage")
    if not message:
        message = f"API request failed with status {status}"

    if status == 400:
        raise BadRequestError(str(message), status_code=status, body=body)
    if status == 401:
        raise AuthenticationError(str(message), status_code=status, body=body)
    if status == 402:
        raise InsufficientCreditsError(str(message), status_code=status, body=body)
    if status == 403:
        raise PlanNotAllowedError(str(message), status_code=status, body=body)
    if status == 404:
        raise NotFoundError(str(message), status_code=status, body=body)
    if status == 409:
        raise JobConflictError(str(message), status_code=status, body=body)
    raise ApiError(str(message), status_code=status, body=body)


class Client:
    """Thin client around `/api/v1` plus a high-level `scrape()` helper."""

    def __init__(
        self,
        api_key: str | None = None,
        *,
        base_url: str | None = None,
        timeout_s: float = 60.0,
    ) -> None:
        key = api_key if api_key is not None else os.environ.get("GMF_API_KEY")
        if not key:
            raise AuthenticationError(
                "Missing API key. Set GMF_API_KEY or pass api_key=... "
                "(https://gmapsleadfinder.com/account#api-key)"
            )
        self.api_key = key.strip()
        self.base_url = (
            base_url
            or os.environ.get("GMF_BASE_URL")
            or DEFAULT_BASE_URL
        ).rstrip("/")
        self.timeout_s = timeout_s

    def me(self) -> MeResponse:
        return self._request("GET", "/api/v1/me")  # type: ignore[return-value]

    def create_job(self, keyword: str) -> CreateJobResponse:
        if not keyword or not str(keyword).strip():
            raise BadRequestError("keyword is required")
        return self._request(
            "POST",
            "/api/v1/jobs",
            json_body={"keyword": str(keyword).strip()},
        )  # type: ignore[return-value]

    def get_job(self, job_id: str) -> JobResponse:
        return self._request("GET", f"/api/v1/jobs/{job_id}")  # type: ignore[return-value]

    def get_results(
        self,
        job_id: str,
        *,
        limit: int = DEFAULT_RESULT_LIMIT,
        cursor: str | int | None = None,
    ) -> ResultsResponse:
        query: dict[str, str] = {"limit": str(limit)}
        if cursor is not None:
            query["cursor"] = str(cursor)
        return self._request(
            "GET",
            f"/api/v1/jobs/{job_id}/results",
            query=query,
        )  # type: ignore[return-value]

    def scrape(
        self,
        keyword: str,
        *,
        poll_interval_ms: int = DEFAULT_POLL_INTERVAL_MS,
        timeout_ms: int = DEFAULT_TIMEOUT_MS,
        result_limit: int = DEFAULT_RESULT_LIMIT,
    ) -> list[dict[str, Any]]:
        """Create a job, poll until terminal, return all result rows."""
        created = self.create_job(keyword)
        job_id = created["jobId"]
        deadline = time.monotonic() + (timeout_ms / 1000.0)

        while True:
            job = self.get_job(job_id)
            status = (job.get("status") or "").lower()
            if status in TERMINAL_STATUSES:
                if status == "failed":
                    raise ApiError(
                        job.get("error") or f"Job {job_id} failed",
                        status_code=None,
                        body=job,
                    )
                break
            if time.monotonic() >= deadline:
                raise TimeoutError(
                    f"Timed out after {timeout_ms}ms waiting for job {job_id} "
                    f"(last status={job.get('status')!r})"
                )
            time.sleep(max(poll_interval_ms, 100) / 1000.0)

        return self._fetch_all_rows(job_id, result_limit=result_limit)

    def _fetch_all_rows(
        self,
        job_id: str,
        *,
        result_limit: int,
    ) -> list[dict[str, Any]]:
        rows: list[dict[str, Any]] = []
        cursor: str | None = "0"
        while cursor is not None:
            page = self.get_results(job_id, limit=result_limit, cursor=cursor)
            rows.extend(page.get("rows") or [])
            next_cursor = page.get("nextCursor")
            cursor = None if next_cursor is None else str(next_cursor)
        return rows

    def _request(
        self,
        method: str,
        path: str,
        *,
        json_body: dict[str, Any] | None = None,
        query: dict[str, str] | None = None,
    ) -> Any:
        url = f"{self.base_url}{path}"
        if query:
            url = f"{url}?{urllib.parse.urlencode(query)}"

        data = None
        headers = {
            "Authorization": f"Bearer {self.api_key}",
            "Accept": "application/json",
            "User-Agent": USER_AGENT,
        }
        if json_body is not None:
            data = json.dumps(json_body).encode("utf-8")
            headers["Content-Type"] = "application/json"

        req = urllib.request.Request(url, data=data, headers=headers, method=method)
        try:
            with urllib.request.urlopen(req, timeout=self.timeout_s) as resp:
                raw = resp.read().decode("utf-8")
                if not raw:
                    return {}
                return json.loads(raw)
        except urllib.error.HTTPError as exc:
            raw = exc.read().decode("utf-8", errors="replace")
            body: Any
            try:
                body = json.loads(raw) if raw else {"message": exc.reason}
            except json.JSONDecodeError:
                body = {"message": raw or str(exc.reason)}
            _raise_for_status(exc.code, body)
            raise  # pragma: no cover
        except urllib.error.URLError as exc:
            raise ApiError(f"Network error: {exc.reason}") from exc
