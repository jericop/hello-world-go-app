.PHONY: build test clean run install lint

# Build the application
build:
	mkdir -p bin
	go build -o bin/hello-world ./cmd/hello-world

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out
	rm -rf dist/

# Run the application
run:
	go run ./cmd/hello-world

# Install dependencies
install:
	go mod download
	go mod tidy

# Lint the code (requires golangci-lint)
lint:
	golangci-lint run ./...

# Build for all platforms (requires GoReleaser)
release-snapshot:
	goreleaser release --snapshot --clean
