package errx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-sdk/lib/seq"
)

const TraceId = "x-request-id"

type Error struct {
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Code      string `json:"code"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"ts"`
	TraceId   string `json:"tid"`
}

func (e *Error) WithContext(ctx context.Context) *Error {
	if req, ok := ctx.Value(0).(*http.Request); ok {
		e.TraceId = req.Header.Get(TraceId)
	}
	return e
}

func (e *Error) String() string {
	return fmt.Sprintf("[%s] (%s) %d %s, code: %s, message: %s", e.Timestamp, e.TraceId, e.Status, e.Error, e.Code, e.Message)
}

func OK(code, message string) *Error {
	return New(http.StatusOK, code, message)
}

func BadRequest(code, message string) *Error {
	return New(http.StatusBadRequest, code, message)
}

func Unauthorized(code, message string) *Error {
	return New(http.StatusUnauthorized, code, message)
}

func Forbidden(code, message string) *Error {
	return New(http.StatusForbidden, code, message)
}

func NotFound(code, message string) *Error {
	return New(http.StatusNotFound, code, message)
}

func Conflict(code, message string) *Error {
	return New(http.StatusConflict, code, message)
}

func InternalError(code, message string) *Error {
	return New(http.StatusInternalServerError, code, message)
}

func New(status int, code, message string) *Error {
	e := &Error{
		Status:    status,
		Code:      code,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}
	e.Error = http.StatusText(e.Status)
	e.TraceId = seq.NewUUID().String()
	return e
}
