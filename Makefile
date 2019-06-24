GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=./...
GOENTRYPOINT=./cmd
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid

build:
	@echo "Building $(GOENTRYPOINT) to ./bin"
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(GONAME) $(GOENTRYPOINT)

start: build
	@$(GOBIN)/$(GONAME)

get:
	@go mod download

run:
	@go run $(GOENTRYPOINT)

clean:
	@echo "Cleaning"
	@go clean $(GOFILES)
	@go mod tidy

test:
	@go test $(GOFILES)

shorttest:
	@go test -short $(GOFILES)

stest: shorttest

.PHONY: build start get run clean test shorttest stest