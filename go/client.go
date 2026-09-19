package gmaps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	DefaultBaseURL        = "https://gmapsleadfinder.com"
	DefaultPollIntervalMs = 2000
	DefaultTimeoutMs      = 600_000
	DefaultResultLimit    = 100
	DefaultHTTPTimeoutMs  = 60_000
	userAgent             = "gmaps-scraper-go/0.1.1"
)

// Client talks to the GMaps Lead Finder Agent HTTP API.
type Client struct {
	APIKey    string
	BaseURL   string
	HTTP      *http.Client
	TimeoutMs int
}

// NewClient reads GMF_API_KEY (and optional GMF_BASE_URL) unless overridden in opts.
func NewClient(opts *ClientOptions) (*Client, error) {
	if opts == nil {
		opts = &ClientOptions{}
	}
	key := strings.TrimSpace(opts.APIKey)
	if key == "" {
		key = strings.TrimSpace(os.Getenv("GMF_API_KEY"))
	}
	if key == "" {
		return nil, &AuthenticationError{APIError{
			Message: "Missing API key. Set GMF_API_KEY or pass APIKey (https://gmapsleadfinder.com/account#api-key)",
		}}
	}
	base := strings.TrimRight(opts.BaseURL, "/")
	if base == "" {
		base = strings.TrimRight(os.Getenv("GMF_BASE_URL"), "/")
	}
	if base == "" {
		base = DefaultBaseURL
	}
	timeoutMs := opts.TimeoutMs
	if timeoutMs <= 0 {
		timeoutMs = DefaultHTTPTimeoutMs
	}
	return &Client{
		APIKey:    key,
		BaseURL:   base,
		TimeoutMs: timeoutMs,
		HTTP: &http.Client{
			Timeout: time.Duration(timeoutMs) * time.Millisecond,
		},
	}, nil
}

func (c *Client) Me() (*MeResponse, error) {
	var out MeResponse
	if err := c.request("GET", "/api/v1/me", nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreateJob(keyword string) (*CreateJobResponse, error) {
	trimmed := strings.TrimSpace(keyword)
	if trimmed == "" {
		return nil, &BadRequestError{APIError{Message: "keyword is required"}}
	}
	var out CreateJobResponse
	if err := c.request("POST", "/api/v1/jobs", map[string]string{"keyword": trimmed}, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetJob(jobID string) (*JobResponse, error) {
	var out JobResponse
	if err := c.request("GET", "/api/v1/jobs/"+url.PathEscape(jobID), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetResults(jobID string, opts *GetResultsOptions) (*ResultsResponse, error) {
	if opts == nil {
		opts = &GetResultsOptions{}
	}
	limit := opts.Limit
	if limit <= 0 {
		limit = DefaultResultLimit
	}
	q := url.Values{}
	q.Set("limit", fmt.Sprintf("%d", limit))
	if opts.Cursor != "" {
		q.Set("cursor", opts.Cursor)
	}
	var out ResultsResponse
	if err := c.request("GET", "/api/v1/jobs/"+url.PathEscape(jobID)+"/results", nil, q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreateReviewsJob(place string) (*CreateJobResponse, error) {
	trimmed := strings.TrimSpace(place)
	if trimmed == "" {
		return nil, &BadRequestError{APIError{Message: "place is required"}}
	}
	var out CreateJobResponse
	if err := c.request("POST", "/api/v1/review-jobs", map[string]string{"place": trimmed}, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetReviewsJob(jobID string) (*JobResponse, error) {
	var out JobResponse
	if err := c.request("GET", "/api/v1/review-jobs/"+url.PathEscape(jobID), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetReviewsResults(jobID string, opts *GetResultsOptions) (*ResultsResponse, error) {
	if opts == nil {
		opts = &GetResultsOptions{}
	}
	limit := opts.Limit
	if limit <= 0 {
		limit = DefaultResultLimit
	}
	q := url.Values{}
	q.Set("limit", fmt.Sprintf("%d", limit))
	if opts.Cursor != "" {
		q.Set("cursor", opts.Cursor)
	}
	var out ResultsResponse
	if err := c.request("GET", "/api/v1/review-jobs/"+url.PathEscape(jobID)+"/results", nil, q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreatePhotosJob(place string) (*CreateJobResponse, error) {
	trimmed := strings.TrimSpace(place)
	if trimmed == "" {
		return nil, &BadRequestError{APIError{Message: "place is required"}}
	}
	var out CreateJobResponse
	if err := c.request("POST", "/api/v1/photo-jobs", map[string]string{"place": trimmed}, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPhotosJob(jobID string) (*JobResponse, error) {
	var out JobResponse
	if err := c.request("GET", "/api/v1/photo-jobs/"+url.PathEscape(jobID), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPhotosResults(jobID string, opts *GetResultsOptions) (*ResultsResponse, error) {
	if opts == nil {
		opts = &GetResultsOptions{}
	}
	limit := opts.Limit
	if limit <= 0 {
		limit = DefaultResultLimit
	}
	q := url.Values{}
	q.Set("limit", fmt.Sprintf("%d", limit))
	if opts.Cursor != "" {
		q.Set("cursor", opts.Cursor)
	}
	var out ResultsResponse
	if err := c.request("GET", "/api/v1/photo-jobs/"+url.PathEscape(jobID)+"/results", nil, q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Scrape creates a job, polls until terminal, and returns all result rows.
func (c *Client) Scrape(keyword string, opts *ScrapeOptions) ([]PlaceRow, error) {
	if opts == nil {
		opts = &ScrapeOptions{}
	}
	created, err := c.CreateJob(keyword)
	if err != nil {
		return nil, err
	}
	if err := c.waitForJob(created.JobID, opts, "maps"); err != nil {
		return nil, err
	}
	resultLimit := opts.ResultLimit
	if resultLimit <= 0 {
		resultLimit = DefaultResultLimit
	}
	return c.fetchAllRows(created.JobID, resultLimit, "maps")
}

// ScrapeReviews creates a reviews job, polls until terminal, and returns all review rows.
func (c *Client) ScrapeReviews(place string, opts *ScrapeOptions) ([]PlaceRow, error) {
	if opts == nil {
		opts = &ScrapeOptions{}
	}
	created, err := c.CreateReviewsJob(place)
	if err != nil {
		return nil, err
	}
	if err := c.waitForJob(created.JobID, opts, "reviews"); err != nil {
		return nil, err
	}
	resultLimit := opts.ResultLimit
	if resultLimit <= 0 {
		resultLimit = DefaultResultLimit
	}
	return c.fetchAllRows(created.JobID, resultLimit, "reviews")
}

// ScrapePhotos creates a photos job, polls until terminal, and returns all photo rows.
func (c *Client) ScrapePhotos(place string, opts *ScrapeOptions) ([]PlaceRow, error) {
	if opts == nil {
		opts = &ScrapeOptions{}
	}
	created, err := c.CreatePhotosJob(place)
	if err != nil {
		return nil, err
	}
	if err := c.waitForJob(created.JobID, opts, "photos"); err != nil {
		return nil, err
	}
	resultLimit := opts.ResultLimit
	if resultLimit <= 0 {
		resultLimit = DefaultResultLimit
	}
	return c.fetchAllRows(created.JobID, resultLimit, "photos")
}

func (c *Client) waitForJob(jobID string, opts *ScrapeOptions, kind string) error {
	pollMs := opts.PollIntervalMs
	if pollMs <= 0 {
		pollMs = DefaultPollIntervalMs
	}
	timeoutMs := opts.TimeoutMs
	if timeoutMs <= 0 {
		timeoutMs = DefaultTimeoutMs
	}
	deadline := time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)

	for {
		var job *JobResponse
		var err error
		switch kind {
		case "reviews":
			job, err = c.GetReviewsJob(jobID)
		case "photos":
			job, err = c.GetPhotosJob(jobID)
		default:
			job, err = c.GetJob(jobID)
		}
		if err != nil {
			return err
		}
		status := strings.ToLower(job.Status)
		if _, ok := TerminalStatuses[status]; ok {
			if status == "failed" {
				msg := fmt.Sprintf("Job %s failed", jobID)
				if job.Error != nil && *job.Error != "" {
					msg = *job.Error
				}
				return &APIError{Message: msg, Body: job}
			}
			return nil
		}
		if time.Now().After(deadline) {
			return &TimeoutError{APIError{
				Message: fmt.Sprintf("Timed out after %dms waiting for job %s (last status=%q)", timeoutMs, jobID, job.Status),
			}}
		}
		time.Sleep(time.Duration(max(pollMs, 100)) * time.Millisecond)
	}
}

func (c *Client) fetchAllRows(jobID string, resultLimit int, kind string) ([]PlaceRow, error) {
	var rows []PlaceRow
	cursor := "0"
	for cursor != "" {
		var page *ResultsResponse
		var err error
		opts := &GetResultsOptions{Limit: resultLimit, Cursor: cursor}
		switch kind {
		case "reviews":
			page, err = c.GetReviewsResults(jobID, opts)
		case "photos":
			page, err = c.GetPhotosResults(jobID, opts)
		default:
			page, err = c.GetResults(jobID, opts)
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, page.Rows...)
		if page.NextCursor == nil || *page.NextCursor == "" {
			break
		}
		cursor = *page.NextCursor
	}
	return rows, nil
}

func (c *Client) request(method, path string, jsonBody any, query url.Values, out any) error {
	u := c.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var body io.Reader
	if jsonBody != nil {
		b, err := json.Marshal(jsonBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	if jsonBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return &APIError{Message: "Network error: " + err.Error()}
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var parsed any
	if len(raw) > 0 {
		_ = json.Unmarshal(raw, &parsed)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if parsed == nil {
			parsed = map[string]any{"message": string(raw)}
		}
		return raiseForStatus(res.StatusCode, parsed)
	}
	if out != nil && len(raw) > 0 {
		if err := json.Unmarshal(raw, out); err != nil {
			return err
		}
	}
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
