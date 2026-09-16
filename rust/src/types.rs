//! Response types for the GMaps Lead Finder Agent HTTP API.

use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Debug, Clone, Deserialize, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct MeResponse {
    pub plan: String,
    pub credits_limit: f64,
    pub credits_used: f64,
    pub credits_remaining: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct CreateJobResponse {
    pub job_id: String,
    pub keyword_count: i32,
    pub credits_remaining: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct JobFile {
    pub id: Option<String>,
    pub keyword: Option<String>,
    pub position: Option<i32>,
    pub status: Option<String>,
    pub page_count: Option<i32>,
    pub row_count: Option<i32>,
    pub enrich_status: Option<String>,
    pub error: Option<String>,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct JobResponse {
    pub id: Option<String>,
    pub status: Option<String>,
    pub keywords: Option<Vec<String>>,
    pub keyword_count: Option<i32>,
    pub current_keyword_index: Option<i32>,
    pub row_count: Option<i32>,
    pub error: Option<String>,
    pub created_at: Option<String>,
    pub finished_at: Option<String>,
    pub files: Option<Vec<JobFile>>,
    pub credits_remaining: Option<f64>,
}

pub type PlaceRow = HashMap<String, String>;

#[derive(Debug, Clone, Deserialize, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct ResultsResponse {
    pub job_id: Option<String>,
    pub status: Option<String>,
    pub columns: Option<Vec<String>>,
    pub rows: Option<Vec<PlaceRow>>,
    pub count: Option<i32>,
    pub total: Option<i32>,
    pub cursor: Option<i32>,
    pub next_cursor: Option<String>,
}

#[derive(Debug, Clone, Default)]
pub struct ClientOptions {
    pub api_key: Option<String>,
    pub base_url: Option<String>,
    /// Per-request HTTP timeout in milliseconds (default 60_000).
    pub timeout_ms: Option<u64>,
}

#[derive(Debug, Clone, Default)]
pub struct ScrapeOptions {
    pub poll_interval_ms: Option<u64>,
    /// Max wait for job completion in milliseconds (default 600_000).
    pub timeout_ms: Option<u64>,
    pub result_limit: Option<u32>,
}

#[derive(Debug, Clone, Default)]
pub struct GetResultsOptions {
    pub limit: Option<u32>,
    pub cursor: Option<String>,
}

pub fn is_terminal_status(status: &str) -> bool {
    matches!(
        status.to_ascii_lowercase().as_str(),
        "completed" | "partial" | "failed"
    )
}
