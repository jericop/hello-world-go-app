# Hello World Go App

A simple Go command-line tool that prints "hello world".

## Installation

### Homebrew (macOS)

Install via Homebrew from our custom tap:

```bash
brew tap jericop/hello-world-go-app
brew install hello-world
```

Or install directly:

```bash
brew install jericop/hello-world-go-app/hello-world
```

### Download Binaries

Download the latest release for your platform from the [releases page](https://github.com/jericop/hello-world-go-app/releases).

**NOTE** The binaries are not notarized by apple, so you will need to allow them in security settings. 😢

### Build from Source

See the [Building](#building) section below.

## Project Structure

```
.
├── cmd/
│   └── hello-world/       # Main application
│       ├── main.go
│       └── main_test.go
├── pkg/
│   └── greeting/          # Public package with greeting constant
│       ├── greeting.go
│       └── greeting_test.go
├── internal/              # Private application code (currently empty)
├── .github/
│   └── workflows/
│       ├── test.yml       # CI workflow for testing
│       └── release.yml    # Release workflow with GoReleaser
├── .goreleaser.yml        # GoReleaser configuration
└── go.mod
```

## Building

Build the application:

```bash
mkdir -p bin
go build -o bin/hello-world ./cmd/hello-world
```

## Running

Run the application:

```bash
./bin/hello-world
```

Or run directly with Go:

```bash
go run ./cmd/hello-world
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

## CI/CD

### GitHub Actions Testing

The project uses GitHub Actions to automatically run tests on pull requests and pushes to the main branch. The workflow is defined in [.github/workflows/test.yml](.github/workflows/test.yml).

### Releases with GoReleaser

The project uses GoReleaser to build and publish binaries for multiple platforms:

- **Operating Systems**: Linux, macOS
- **Architectures**: amd64, arm64

To create a release:

1. Create and push a new tag:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. GitHub Actions will automatically trigger the release workflow, which will:
   - Build binaries for all supported platforms
   - Create release archives
   - Generate checksums
   - Publish the release to GitHub
   - Update the Homebrew tap with the new version

The release workflow is defined in [.github/workflows/release.yml](.github/workflows/release.yml).

## License

See [LICENSE](LICENSE) file for details.
