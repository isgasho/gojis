GOCMD=go
GOMODCMD=$(GOCMD) mod
GOTESTCMD=$(GOCMD) test
GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=./...
GOENTRYPOINT=./cmd
GONAME=$(shell basename "$(PWD)")
EXECUTABLE=$(GOBIN)/$(GONAME)
PID=/tmp/go-$(GONAME).pid

build:
	@echo "Building $(GOENTRYPOINT) to ./bin"
	@GOBIN=$(GOBIN) $(GOCMD) build -o $(EXECUTABLE) $(GOENTRYPOINT)

start: build
	@$(EXECUTABLE)

get:
	@$(GOMODCMD) download

run:
	@$(GOCMD) run $(GOENTRYPOINT)

clean:
	@echo "Cleaning"
	@rm $(EXECUTABLE)
	@$(GOCMD) clean $(GOFILES)
	@$(GOMODCMD) tidy

test:
	@$(GOTESTCMD) $(GOFILES)

shorttest:
	@$(GOTESTCMD) -short $(GOFILES)

stest: shorttest

.PHONY: build start get run clean test shorttest stest