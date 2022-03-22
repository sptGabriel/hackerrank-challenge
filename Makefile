NAME=hackerrank-challenge
OS ?= linux

.PHONY: compile
compile:
	@echo "=> installing dependencies..."
	go mod tidy
	@echo "==> Compiling challenge..."
	go build -o build/server main.go

.PHONY: test
test:
	@echo "==> Running tests..."
	go test ./... --race -count=1 -v


.PHONY: lint
lint:
	@echo "Running golangci-lint"
	@golangci-lint run --fix