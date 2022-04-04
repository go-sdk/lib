GO 	    ?= GO111MODULE=on go
VERSION ?= $(shell git describe --exact-match --tags HEAD 2>/dev/null || echo "latest")
GITHASH ?= $(shell git rev-parse --short HEAD)

LDFLAGS := -s -w
LDFLAGS += -X "github.com/go-sdk/lib/app.VERSION=$(VERSION)"
LDFLAGS += -X "github.com/go-sdk/lib/app.GITHASH=$(GITHASH)"
LDFLAGS += -X "github.com/go-sdk/lib/app.BUILT=$(shell date +%FT%T%z)"

.PHONY: test tidy lint

test:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) test -race -ldflags '$(LDFLAGS)' -count=1 -cover -covermode=atomic -coverprofile=coverage.out -v $(shell $(GO) list ./... | grep -v lib/flag | grep -v lib/bot | grep -v lib/cron/locker)

tidy:
	$(GO) mod tidy

lint:
	golangci-lint run --skip-dirs-use-default --skip-dirs internal/dateparse
