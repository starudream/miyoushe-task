PROJECT ?= $(shell basename $(CURDIR))
MODULE  ?= $(shell go list -m)
VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)

BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "github.com/starudream/go-lib/core/v2/config/version.gitVersion=$(VERSION)"

.PHONY: init
init:
	git status -b -s
	go mod tidy

.PHONY: bin
bin: init
	CGO_ENABLED=0 go build -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o bin/$(PROJECT) $(MODULE)/cmd

.PHONY: run
run: bin
	DEBUG=true bin/$(PROJECT) $(ARGS)
