package errx

import (
	"testing"

	"github.com/go-sdk/lib/codec/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestError(t *testing.T) {
	t.Log(OK("OK"))
	t.Log(BadRequest("BadRequest"))
	t.Log(Unauthorized("Unauthorized"))
	t.Log(Forbidden("Forbidden"))
	t.Log(NotFound("NotFound"))
	t.Log(Conflict("Conflict"))
	t.Log(Internal("Internal"))
	t.Log(NotImplemented("NotImplemented"))

	e1 := OK("OK")
	t.Log(e1.WithMetadata(Metadata{"code": 0}))
	t.Log(e1)
	t.Log(json.MustMarshalToString(e1))

	e2 := BadRequest("BadRequest").WithMetadata(Metadata{"username": "6"})
	t.Log(e2)
	t.Log(e2.Status())
	t.Log(e2.Message())
	t.Log(e2.Metadata())
	t.Log(e2.GetMetadata("username"))

	e3 := status.New(codes.Unimplemented, "Unimplemented").Err()
	e4 := FromError(e3)
	t.Log(e4)
}
