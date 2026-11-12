set windows-shell := ["pwsh", "-NoLogo", "-NoProfile", "-Command"]

binary := "go-tool-template"
version := if os_family() == "windows" { `git describe --tags --always --dirty 2>$null || echo dev` } else { `git describe --tags --always --dirty 2>/dev/null || echo dev` }
commit := if os_family() == "windows" { `git rev-parse --short HEAD 2>$null || echo ""` } else { `git rev-parse --short HEAD 2>/dev/null || echo ""` }
ldflags := "-s -w -X github.com/DaanV2/go-tool-template/infrastructure/version.Version=" + version + " -X github.com/DaanV2/go-tool-template/infrastructure/version.Commit=" + commit
bin := "bin/" + binary + if os_family() == "windows" { ".exe" } else { "" }

# List available recipes
default:
    @just --list

# Install the development tools (keep golangci-lint in sync with .github/workflows/ci.yml)
[group("setup")]
onboard:
    go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0
    go install github.com/goreleaser/goreleaser/v2@latest

# Build the binary into ./bin
[group("build")]
build:
    go build -trimpath -ldflags "{{ ldflags }}" -o {{ bin }} .

# Run the tool with the given arguments
[group("build")]
run *args:
    go run . {{ args }}

# Remove build artifacts
[group("build")]
clean:
    go clean
    {{ if os_family() == "windows" { "'bin', 'dist', 'coverage.out' | Where-Object { Test-Path $_ } | Remove-Item -Recurse -Force" } else { "rm -rf bin dist coverage.out" } }}

# Run all tests
[group("test")]
test:
    go test ./...

# Run all tests with the race detector (requires cgo / a C toolchain)
[group("test")]
[env("CGO_ENABLED", "1")]
test-race:
    go test -race ./...

# Run tests with coverage report
[group("test")]
cover:
    go test -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

# Run golangci-lint
[group("lint")]
lint:
    golangci-lint run

# Run golangci-lint and apply fixes
[group("lint")]
lint-fix:
    golangci-lint run --fix

# Format code
[group("lint")]
fmt:
    golangci-lint fmt

# Tidy and verify go modules
[group("lint")]
tidy:
    go mod tidy
    go mod verify

# Run everything CI runs
[group("ci")]
ci: tidy fmt lint test build

# Build a local snapshot release with goreleaser
[group("release")]
snapshot:
    goreleaser release --snapshot --clean
