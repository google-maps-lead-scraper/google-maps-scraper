//! Google Maps Scraper SDK for the GMaps Lead Finder Agent HTTP API.

pub mod client;
pub mod errors;
pub mod types;

pub use client::Client;
pub use errors::{Error, Result};
pub use types::{
    ClientOptions, CreateJobResponse, GetResultsOptions, JobFile, JobResponse, MeResponse,
    PlaceRow, ResultsResponse, ScrapeOptions,
};
