.PHONY: build test vendor

GO ?= go
BIN ?= ./securemyvps

LDFLAGS := -s -w -buildid=

build:
	CGO_ENABLED=0 $(GO) build -trimpath -buildvcs=false -mod=vendor \
		-ldflags="$(LDFLAGS)" -o $(BIN) ./cmd/securemyvps

test:
	$(GO) test -race ./...

vendor:
	$(GO) mod tidy
	$(GO) mod vendor
	$(GO) mod verify