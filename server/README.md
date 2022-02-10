# Server

```go
package service

import (
    "context"
    "fmt"
    "net/http"

    "github.com/go-sdk/lib/app"
    "github.com/go-sdk/lib/consts"
    "github.com/go-sdk/lib/errx"
    "github.com/go-sdk/lib/seq"
    "github.com/go-sdk/lib/server"

    "github.com/go-sdk/pb/admin"
    "github.com/go-sdk/pb/common"
    "github.com/go-sdk/pb/doc"
)

func Start(ctx context.Context) error {
    s := server.New(ctx)
    s.RegisterService(&admin.UserService_ServiceDesc, &UserService{})
    s.RegisterServiceHandler(admin.RegisterUserServiceHandlerFromEndpoint)
    g := s.Group("/v1")
    {
        g.HandlePath(http.MethodPost, "/ping", func(c *server.Context) (interface{}, error) {
            fmt.Println("ping")
            return "pong", nil
        })
        gg := g.Group("/user")
        {
            gg.HandlePath(http.MethodGet, "/create", func(c *server.Context) (interface{}, error) {
                return "world", errx.New(http.StatusBadRequest, "no").WithMetadata(errx.Metadata{"x": "y"})
            })
        }
    }
    s.HandlePath(http.MethodGet, "/v2", func(c *server.Context) (interface{}, error) {
        fmt.Println("v2")
        return "v2", nil
    })
    s.HandlePath(http.MethodGet, "/swagger.json", func(c *server.Context) (interface{}, error) {
        c.W.Header().Set(consts.ContentType, consts.ContentTypeJSON)
        return doc.SwaggerJSON, nil
    })
    return s.Start(":8999")
}

type UserService struct {
    admin.UnimplementedUserServiceServer
}

func (s *UserService) Health(_ context.Context, _ *common.Empty) (*common.Struct, error) {
    return common.NewStruct(app.VersionInfoMap())
}

func (s *UserService) CreateUser(ctx context.Context, _ *admin.User) (*common.Id, error) {
    return &common.Id{Id: seq.NewXID().String()}, nil
}
```
