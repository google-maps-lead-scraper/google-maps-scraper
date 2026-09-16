package gmaps

import "fmt"

// APIError is the base error for Agent HTTP API failures.
type APIError struct {
	Message    string
	StatusCode int
	Body       any
}

func (e *APIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("%s (status %d)", e.Message, e.StatusCode)
	}
	return e.Message
}

type AuthenticationError struct{ APIError }
type PlanNotAllowedError struct{ APIError }
type InsufficientCreditsError struct{ APIError }
type JobConflictError struct{ APIError }
type NotFoundError struct{ APIError }
type BadRequestError struct{ APIError }
type TimeoutError struct{ APIError }

func raiseForStatus(status int, body any) error {
	msg := fmt.Sprintf("API request failed with status %d", status)
	if m := extractMessage(body); m != "" {
		msg = m
	}
	base := APIError{Message: msg, StatusCode: status, Body: body}
	switch status {
	case 400:
		return &BadRequestError{base}
	case 401:
		return &AuthenticationError{base}
	case 402:
		return &InsufficientCreditsError{base}
	case 403:
		return &PlanNotAllowedError{base}
	case 404:
		return &NotFoundError{base}
	case 409:
		return &JobConflictError{base}
	default:
		return &base
	}
}

func extractMessage(body any) string {
	m, ok := body.(map[string]any)
	if !ok {
		return ""
	}
	for _, key := range []string{"message", "error", "statusMessage"} {
		if v, ok := m[key].(string); ok && v != "" {
			return v
		}
	}
	return ""
}
