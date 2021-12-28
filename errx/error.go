package errx

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/seq"
)

type Error struct {
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"ts"`
	TraceId   string `json:"tid"`
}

func (e *Error) WithContext(ctx context.Context) *Error {
	if req, ok := ctx.Value(0).(*http.Request); ok {
		e.TraceId = req.Header.Get(consts.TraceId)
	}
	return e
}

func (e *Error) WithCode(code string) *Error {
	e.Code = code
	return e
}

func (e *Error) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	sb.WriteString(e.Timestamp)
	sb.WriteString("] (")
	sb.WriteString(e.TraceId)
	sb.WriteString(") ")
	sb.WriteString(strconv.Itoa(e.Status))
	sb.WriteString(" ")
	sb.WriteString(e.Error)
	if e.Code != "" {
		sb.WriteString(", ")
		sb.WriteString(e.Code)
	}
	if e.Message != "" {
		sb.WriteString(", ")
		sb.WriteString(e.Message)
	}
	return sb.String()
}

func OK(message string) *Error {
	return New(http.StatusOK, message)
}

func BadRequest(message string) *Error {
	return New(http.StatusBadRequest, message)
}

func Unauthorized(message string) *Error {
	return New(http.StatusUnauthorized, message)
}

func Forbidden(message string) *Error {
	return New(http.StatusForbidden, message)
}

func NotFound(message string) *Error {
	return New(http.StatusNotFound, message)
}

func NotAllowed(message string) *Error {
	return New(http.StatusMethodNotAllowed, message)
}

func Conflict(message string) *Error {
	return New(http.StatusConflict, message)
}

func InternalError(message string) *Error {
	return New(http.StatusInternalServerError, message)
}

func New(status int, message string) *Error {
	e := &Error{
		Status:    status,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}
	e.Error = http.StatusText(e.Status)
	e.TraceId = seq.NewUUID().String()
	return e
}
