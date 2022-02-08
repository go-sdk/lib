package errx

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/status"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/seq"
)

type Error struct {
	Status    int    `json:"status"`
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

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	sb.WriteString(e.Timestamp)
	sb.WriteString("] (")
	sb.WriteString(e.TraceId)
	sb.WriteString(") ")
	sb.WriteString(strconv.Itoa(e.Status))
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
	e.TraceId = seq.NewUUID().String()
	return e
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if se, ok := status.FromError(err); ok {
		return InternalError(se.Message()).WithCode(se.Code().String())
	}

	if e, ok := err.(*Error); ok {
		return e
	}

	return InternalError(err.Error())
}
