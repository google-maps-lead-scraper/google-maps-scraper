package gmaps

// Response types for the GMaps Lead Finder Agent HTTP API.

type MeResponse struct {
	Plan             string  `json:"plan"`
	CreditsLimit     float64 `json:"creditsLimit"`
	CreditsUsed      float64 `json:"creditsUsed"`
	CreditsRemaining float64 `json:"creditsRemaining"`
}

type CreateJobResponse struct {
	JobID            string  `json:"jobId"`
	KeywordCount     int     `json:"keywordCount"`
	CreditsRemaining float64 `json:"creditsRemaining"`
}

type JobFile struct {
	ID           string  `json:"id,omitempty"`
	Keyword      string  `json:"keyword,omitempty"`
	Position     int     `json:"position,omitempty"`
	Status       string  `json:"status,omitempty"`
	PageCount    int     `json:"pageCount,omitempty"`
	RowCount     int     `json:"rowCount,omitempty"`
	EnrichStatus string  `json:"enrichStatus,omitempty"`
	Error        *string `json:"error,omitempty"`
}

type JobResponse struct {
	ID                  string    `json:"id,omitempty"`
	Status              string    `json:"status,omitempty"`
	Keywords            []string  `json:"keywords,omitempty"`
	KeywordCount        int       `json:"keywordCount,omitempty"`
	CurrentKeywordIndex int       `json:"currentKeywordIndex,omitempty"`
	RowCount            int       `json:"rowCount,omitempty"`
	Error               *string   `json:"error,omitempty"`
	CreatedAt           string    `json:"createdAt,omitempty"`
	FinishedAt          *string   `json:"finishedAt,omitempty"`
	Files               []JobFile `json:"files,omitempty"`
	CreditsRemaining    float64   `json:"creditsRemaining,omitempty"`
}

// PlaceRow is one export row keyed by column header (Name, Phone, Website, Emails, …).
type PlaceRow map[string]string

type ResultsResponse struct {
	JobID      string     `json:"jobId,omitempty"`
	Status     string     `json:"status,omitempty"`
	Columns    []string   `json:"columns,omitempty"`
	Rows       []PlaceRow `json:"rows,omitempty"`
	Count      int        `json:"count,omitempty"`
	Total      int        `json:"total,omitempty"`
	Cursor     int        `json:"cursor,omitempty"`
	NextCursor *string    `json:"nextCursor,omitempty"`
}

type ClientOptions struct {
	APIKey    string
	BaseURL   string
	TimeoutMs int // per-request HTTP timeout; default 60000
}

type ScrapeOptions struct {
	PollIntervalMs int
	TimeoutMs      int // max wait for job completion; default 600000
	ResultLimit    int
}

type GetResultsOptions struct {
	Limit  int
	Cursor string
}

var TerminalStatuses = map[string]struct{}{
	"completed": {},
	"partial":   {},
	"failed":    {},
}
