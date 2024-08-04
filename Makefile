all: codegen

codegen: openapi wire

openapi:
	@echo "Generating OpenAPI interface..."
	@go generate ./internal/api/rest
	@go generate ./internal/testability

wire:
	@echo "Generating wire files..."
	@wire . ./cmd

test:
	@go test -cover -coverpkg=./...

coverage:
	@go test -coverprofile=coverage.out -coverpkg=./...
	@go tool cover -html=coverage.out
