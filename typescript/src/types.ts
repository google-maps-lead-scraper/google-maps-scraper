/** Typed shapes for Agent HTTP API responses. */

export type MeResponse = {
  plan: string;
  creditsLimit: number;
  creditsUsed: number;
  creditsRemaining: number;
};

export type CreateJobResponse = {
  jobId: string;
  keywordCount: number;
  creditsRemaining: number;
};

export type JobFile = {
  id?: string;
  keyword?: string;
  position?: number;
  status?: string;
  pageCount?: number;
  rowCount?: number;
  enrichStatus?: string;
  error?: string | null;
};

export type JobResponse = {
  id?: string;
  status?: string;
  keywords?: string[];
  keywordCount?: number;
  currentKeywordIndex?: number;
  rowCount?: number;
  error?: string | null;
  createdAt?: string;
  finishedAt?: string | null;
  files?: JobFile[];
  creditsRemaining?: number;
};

export type PlaceRow = Record<string, string>;

export type ResultsResponse = {
  jobId?: string;
  status?: string;
  columns?: string[];
  rows?: PlaceRow[];
  count?: number;
  total?: number;
  cursor?: number;
  nextCursor?: string | null;
};

export type ClientOptions = {
  apiKey?: string;
  baseUrl?: string;
  /** Per-request fetch timeout hint (ms). Default 60_000. */
  timeoutMs?: number;
};

export type ScrapeOptions = {
  pollIntervalMs?: number;
  timeoutMs?: number;
  resultLimit?: number;
};

export type GetResultsOptions = {
  limit?: number;
  cursor?: string | number;
};

export const TERMINAL_STATUSES = new Set(["completed", "partial", "failed"]);
