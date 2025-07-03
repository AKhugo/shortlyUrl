# Variables
BINARY_NAME=shortlyurl
MAIN_DIR=./cmd/server
UNIT_DIR=./test/unit
INTEGRATION_DIR=./test/integration

# Commands

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_DIR)

run:
	go run $(MAIN_DIR)/main.go

test-unit:
	go test -v $(UNIT_DIR)

test-integration:
	INTEGRATION_TESTS=true go test -v $(INTEGRATION_DIR)

test-all: test-unit test-integration

test: test-unit

clean:
	rm -f $(BINARY_NAME)

migrate:
	go run $(MAIN_DIR) migrate

fmt:
	go fmt ./...

vet:
	go vet ./...

# Targets

.PHONY: all build run test test-unit test-integration test-all clean migrate fmt vet