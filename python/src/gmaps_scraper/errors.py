"""Exception types mapped from Agent HTTP API status codes."""

from __future__ import annotations

from typing import Any


class ApiError(Exception):
    """Base error for GMaps Lead Finder API failures."""

    def __init__(
        self,
        message: str,
        *,
        status_code: int | None = None,
        body: Any = None,
    ) -> None:
        super().__init__(message)
        self.status_code = status_code
        self.body = body


class AuthenticationError(ApiError):
    """HTTP 401 — missing or invalid API key."""


class PlanNotAllowedError(ApiError):
    """HTTP 403 — key valid but plan cannot use the Agent API."""


class InsufficientCreditsError(ApiError):
    """HTTP 402 — no credits remaining."""


class JobConflictError(ApiError):
    """HTTP 409 — another job is already running for this user."""


class NotFoundError(ApiError):
    """HTTP 404 — job not found."""


class BadRequestError(ApiError):
    """HTTP 400 — invalid request (e.g. not exactly one keyword)."""


class TimeoutError(ApiError):
    """Local poll timeout waiting for a job to finish."""
