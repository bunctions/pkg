
TEST_TARGET ?= $(shell go list ./... | grep -v examples)

.PHONY: all
all: dep test

.PHONY: dep
dep:
	@go mod tidy

.PHONY: test
test:
	@go test -v -race $(TEST_TARGET)
