package logging

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
func GetRequestID(ctx context.Context) string {

	reqID := ctx.Value(ContextKeyRequestID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}

// AttachRequestID will attach a brand new request ID to a http request
func AttachRequestID(ctx context.Context) context.Context {

	reqID := uuid.New()

	return context.WithValue(ctx, ContextKeyRequestID, reqID.String())
}

// NewLoggerWithRequestID creates a *logrus.Entry that has requestID as a field
func NewLoggerWithRequestID(ctx context.Context, log logrus.FieldLogger) logrus.FieldLogger {

	reqID := GetRequestID(ctx)

	return log.WithField(LogFieldKeyRequestID, reqID)
}
