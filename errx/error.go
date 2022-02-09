package errx

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/go-sdk/lib/codec/json"
)

type GRPCError interface {
	GRPCStatus() *status.Status
}

type Error struct {
	status   int
	message  string
	metadata Metadata
}

type Metadata map[string]interface{}

func New(status int, format string, a ...interface{}) *Error {
	return &Error{status: status, message: fmt.Sprintf(format, a...)}
}

func OK(format string, a ...interface{}) *Error {
	return New(http.StatusOK, format, a...)
}

func BadRequest(format string, a ...interface{}) *Error {
	return New(http.StatusBadRequest, format, a...)
}

func Unauthorized(format string, a ...interface{}) *Error {
	return New(http.StatusUnauthorized, format, a...)
}

func Forbidden(format string, a ...interface{}) *Error {
	return New(http.StatusForbidden, format, a...)
}

func NotFound(format string, a ...interface{}) *Error {
	return New(http.StatusNotFound, format, a...)
}

func Conflict(format string, a ...interface{}) *Error {
	return New(http.StatusConflict, format, a...)
}

func Internal(format string, a ...interface{}) *Error {
	return New(http.StatusInternalServerError, format, a...)
}

func NotImplemented(format string, a ...interface{}) *Error {
	return New(http.StatusNotImplemented, format, a...)
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se, ok := status.FromError(err); ok {
		e := New(HTTPStatusFromCode(se.Code()), se.Message())
		if len(se.Details()) > 0 {
			if v, y := se.Details()[0].(*structpb.Struct); y {
				return e.WithMetadata(v.AsMap())
			}
		}
		return e
	}
	return New(http.StatusInternalServerError, err.Error())
}

func (e *Error) Status() int {
	return e.status
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Metadata() Metadata {
	m := map[string]interface{}{}
	for k, v := range e.metadata {
		m[k] = v
	}
	return m
}

func (e *Error) GetMetadata(key string) interface{} {
	if len(e.metadata) > 0 {
		return e.metadata[key]
	}
	return nil
}

func (e *Error) Error() string {
	s := fmt.Sprintf("error: status: %d, message: %s", e.status, e.message)
	if len(e.metadata) > 0 {
		s += fmt.Sprintf(", metadata: %+v", e.metadata)
	}
	return s
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.ProtoMarshal(e.GRPCStatus().Proto())
}

func (e *Error) GRPCStatus() *status.Status {
	s := status.New(GRPCCodeFromStatus(e.status), e.message)
	if s.Code() == codes.OK || len(e.metadata) == 0 {
		return s
	}
	d, err := structpb.NewStruct(e.metadata)
	if err == nil {
		s, _ = s.WithDetails(d)
	}
	return s
}

func (e Error) WithMetadata(md Metadata) *Error {
	if e.status == http.StatusOK || len(md) == 0 {
		return &e
	}
	if len(e.metadata) > 0 {
		for k, v := range md {
			e.metadata[k] = v
		}
	} else {
		e.metadata = md
	}
	return &e
}
