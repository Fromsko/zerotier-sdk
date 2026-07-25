.PHONY: help build build-mcp test lint clean

help:
	@echo "ZeroTier SDK - Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  build          - Build all packages"
	@echo "  build-mcp      - Build MCP binaries for all platforms"
	@echo "  build-mcp-linux   - Build MCP for Linux"
	@echo "  build-mcp-darwin  - Build MCP for macOS"
	@echo "  build-mcp-windows - Build MCP for Windows"
	@echo "  test           - Run tests"
	@echo "  lint           - Run linter"
	@echo "  clean          - Clean build artifacts"

build:
	go build -o /dev/null ./...
	go build -o /dev/null ./example/...
	go build -o /dev/null ./cmd/mcp/...

build-mcp: build-mcp-linux build-mcp-darwin build-mcp-windows
	@echo "✅ All MCP binaries built successfully"
	@ls -lh dist/zerotier-mcp-*

build-mcp-linux:
	@mkdir -p dist
	@echo "Building Linux x86_64..."
	GOOS=linux GOARCH=amd64 go build -o dist/zerotier-mcp-linux-amd64 ./cmd/mcp
	@echo "Building Linux ARM64..."
	GOOS=linux GOARCH=arm64 go build -o dist/zerotier-mcp-linux-arm64 ./cmd/mcp

build-mcp-darwin:
	@mkdir -p dist
	@echo "Building macOS Intel..."
	GOOS=darwin GOARCH=amd64 go build -o dist/zerotier-mcp-darwin-amd64 ./cmd/mcp
	@echo "Building macOS Apple Silicon..."
	GOOS=darwin GOARCH=arm64 go build -o dist/zerotier-mcp-darwin-arm64 ./cmd/mcp

build-mcp-windows:
	@mkdir -p dist
	@echo "Building Windows x86_64..."
	GOOS=windows GOARCH=amd64 go build -o dist/zerotier-mcp-windows-amd64.exe ./cmd/mcp

test:
	go test -v -race ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf dist/
	go clean ./...
