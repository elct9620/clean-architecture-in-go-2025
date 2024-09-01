all: codegen

codegen: openapi grpc wire

openapi:
	@echo "Generating OpenAPI interface..."
	@go generate ./internal/api/rest
	@go generate ./internal/testability

grpc:
	@echo "Generating gRPC interface..."
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/orderspb/orders.proto

wire:
	@echo "Generating wire files..."
	@wire . ./cmd/...

test:
	@go test -cover -coverpkg=./...

coverage:
	@go test -coverprofile=coverage.out -coverpkg=./...
	@go tool cover -html=coverage.out
