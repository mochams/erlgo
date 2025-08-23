# Target to run when no target is specified
.DEFAULT_GOAL := test

.PHONY: fmt vet test clean

TEST_FLAGS ?= -v

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test $(TEST_FLAGS) ./...

clean: 
	go clean
