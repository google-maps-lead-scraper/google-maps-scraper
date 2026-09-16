"""Typed shapes for Agent HTTP API responses (dicts at runtime)."""

from __future__ import annotations

from typing import Any, TypedDict


class MeResponse(TypedDict):
    plan: str
    creditsLimit: float
    creditsUsed: float
    creditsRemaining: float


class CreateJobResponse(TypedDict):
    jobId: str
    keywordCount: int
    creditsRemaining: float


class JobFile(TypedDict, total=False):
    id: str
    keyword: str
    position: int
    status: str
    pageCount: int
    rowCount: int
    enrichStatus: str
    error: str | None


class JobResponse(TypedDict, total=False):
    id: str
    status: str
    keywords: list[str]
    keywordCount: int
    currentKeywordIndex: int
    rowCount: int
    error: str | None
    createdAt: str
    finishedAt: str | None
    files: list[JobFile]
    creditsRemaining: float


class ResultsResponse(TypedDict, total=False):
    jobId: str
    status: str
    columns: list[str]
    rows: list[dict[str, Any]]
    count: int
    total: int
    cursor: int
    nextCursor: str | None


TERMINAL_STATUSES = frozenset({"completed", "partial", "failed"})
