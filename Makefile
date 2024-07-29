all: codegen

codegen: wire

wire:
	@echo "Generating wire files..."
	@wire . ./cmd

test:
	@go test -cover -coverpkg=./...

coverage:
	@go test -coverprofile=coverage.out -coverpkg=./...
	@go tool cover -html=coverage.out
