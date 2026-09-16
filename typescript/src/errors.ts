/** Exception types mapped from Agent HTTP API status codes. */

export class ApiError extends Error {
  readonly statusCode?: number;
  readonly body?: unknown;

  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message);
    this.name = "ApiError";
    this.statusCode = opts?.statusCode;
    this.body = opts?.body;
  }
}

export class AuthenticationError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "AuthenticationError";
  }
}

export class PlanNotAllowedError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "PlanNotAllowedError";
  }
}

export class InsufficientCreditsError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "InsufficientCreditsError";
  }
}

export class JobConflictError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "JobConflictError";
  }
}

export class NotFoundError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "NotFoundError";
  }
}

export class BadRequestError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "BadRequestError";
  }
}

export class TimeoutError extends ApiError {
  constructor(message: string, opts?: { statusCode?: number; body?: unknown }) {
    super(message, opts);
    this.name = "TimeoutError";
  }
}
