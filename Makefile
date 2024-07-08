all: codegen

codegen: wire

wire:
	@echo "Generating wire files..."
	@wire . ./cmd
