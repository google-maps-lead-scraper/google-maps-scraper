//! Error types mapped from Agent HTTP API status codes.

use thiserror::Error;

#[derive(Debug, Error)]
pub enum Error {
    #[error("{message}")]
    Api {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    Authentication {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    PlanNotAllowed {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    InsufficientCredits {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    JobConflict {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    NotFound {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    BadRequest {
        message: String,
        status_code: Option<u16>,
        body: Option<serde_json::Value>,
    },

    #[error("{message}")]
    Timeout { message: String },

    #[error(transparent)]
    Http(#[from] reqwest::Error),

    #[error(transparent)]
    Json(#[from] serde_json::Error),
}

pub type Result<T> = std::result::Result<T, Error>;

pub(crate) fn raise_for_status(status: u16, body: Option<serde_json::Value>) -> Error {
    let message = body
        .as_ref()
        .and_then(|b| {
            b.get("message")
                .or_else(|| b.get("error"))
                .or_else(|| b.get("statusMessage"))
                .and_then(|v| v.as_str())
                .map(|s| s.to_string())
        })
        .unwrap_or_else(|| format!("API request failed with status {status}"));

    match status {
        400 => Error::BadRequest {
            message,
            status_code: Some(status),
            body,
        },
        401 => Error::Authentication {
            message,
            status_code: Some(status),
            body,
        },
        402 => Error::InsufficientCredits {
            message,
            status_code: Some(status),
            body,
        },
        403 => Error::PlanNotAllowed {
            message,
            status_code: Some(status),
            body,
        },
        404 => Error::NotFound {
            message,
            status_code: Some(status),
            body,
        },
        409 => Error::JobConflict {
            message,
            status_code: Some(status),
            body,
        },
        _ => Error::Api {
            message,
            status_code: Some(status),
            body,
        },
    }
}
