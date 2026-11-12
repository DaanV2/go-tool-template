# go-tool-template

A starting point for Go CLI tools.

- [cobra](https://github.com/spf13/cobra) for commands and flags
- [fang](https://github.com/charmbracelet/fang) for styled help, errors, `--version`, completions and man pages
- [log](https://github.com/charmbracelet/log) for structured logging (`--log-level`, `--log-format text|json|logfmt`)
- [just](https://github.com/casey/just) as the task runner
- [golangci-lint](https://golangci-lint.run) v2 for linting and formatting (gofumpt + goimports)
- [GoReleaser](https://goreleaser.com) for releases
- GitHub Actions for CI, releases, plus Dependabot

## Using this template

1. Create a new repo from this template.
2. Replace the module path and binary name everywhere:
   - `github.com/DaanV2/go-tool-template` → your module path (`go.mod`, imports, `.justfile`, `.golangci.yml`, `.goreleaser.yaml`)
   - `go-tool-template` → your binary name (`cmd/root.go`, `.justfile`, `.goreleaser.yaml`)
3. Remove `cmd/hello.go` and add your own commands in `cmd/`.

## Development

```sh
just          # list recipes
just onboard  # go install golangci-lint and goreleaser
just run hello gopher
just test     # go test ./...
just test-race
just lint     # golangci-lint run
just fmt      # golangci-lint fmt
just build    # ./bin/go-tool-template
just ci       # tidy + lint + test + build
```

## Layout

```
main.go                  entrypoint, calls cmd.Execute
cmd/                     cobra commands (root.go wires fang + logging flags)
infrastructure/logging/        charm log setup
infrastructure/version/        version info injected via -ldflags
```

## CI/CD

| Workflow               | Trigger             | What it does                                                                                                     |
| ---------------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ci.yml`               | push to `main`, PRs | `go mod tidy` check, golangci-lint, tests (+ race on Linux) and build on Linux/macOS/Windows, `goreleaser check` |
| `auto-label-pr.yaml`   | PRs                 | Labels PRs from changed files and branch name (`.github/labeler.yml`)                                            |
| `release-drafter.yaml` | push to `main`      | Keeps a draft release with categorized notes up to date (`.github/release-drafter.yml`)                          |
| `release-assets.yaml`  | release published   | GoReleaser builds linux/darwin/windows × amd64/arm64 and uploads the archives to that release                    |
| `codeql.yml`           | push, PRs, weekly   | CodeQL security analysis                                                                                         |

Dependabot updates Go modules (charm packages grouped separately) and GitHub Actions weekly.

## Release version

To release: open the draft release on GitHub, check the notes and version, and publish it.
The version bump defaults to patch; add a `minor` or `major` label to a PR to bump further.
Branches starting with `feature` or `bug` are labeled automatically.
