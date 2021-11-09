GO 	    ?= GO111MODULE=on go
VERSION ?= test
GITHASH ?= -

LDFLAGS := -s -w
LDFLAGS += -X "github.com/go-sdk/lib/app.VERSION=$(VERSION)"
LDFLAGS += -X "github.com/go-sdk/lib/app.GITHASH=$(GITHASH)"

test:
	CGO_ENABLED=1 $(GO) test -race -ldflags '$(LDFLAGS)' -count=1 -cover -v ./...
