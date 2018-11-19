package logging

import (
	"net/http"
)

// ContextKey is
type ContextKey string

const (
	// ContextKeyRequestID is the context key for RequestID
	ContextKeyRequestID ContextKey = "requestID"

	// LogFieldKeyRequestID is the logfield key for RequestID
	LogFieldKeyRequestID = "requestId"
)

// GetRequestID will get reqID from a http request and return it as a string
func GetRequestID(req *http.Request) string {
	ctx := req.Context()
	reqID := ctx.Value(ContextKeyRequestID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}
