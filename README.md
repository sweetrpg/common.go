# common.go

[![CI](https://github.com/sweetrpg/common.go/actions/workflows/ci.yaml/badge.svg)](https://github.com/sweetrpg/common.go/actions/workflows/ci.yaml)
[![License](https://img.shields.io/github/license/sweetrpg/common.go.svg)](https://img.shields.io/github/license/sweetrpg/common.go.svg)
[![Issues](https://img.shields.io/github/issues/sweetrpg/common.go.svg)](https://img.shields.io/github/issues/sweetrpg/common.go.svg)
[![PRs](https://img.shields.io/github/issues-pr/sweetrpg/common.go.svg)](https://img.shields.io/github/issues-pr/sweetrpg/common.go.svg)
[![Dependabot](https://badgen.net/github/dependabot/sweetrpg/common.go)](https://badgen.net/github/dependabot/sweetrpg/common.go)

Dependency-free utility packages shared across sweetrpg's Go services and libraries: generic
slice mapping, environment variable helpers, structured logging setup, and common constants.
It sits at the base of the platform's Go dependency graph - nearly every other sweetrpg Go
module depends on it directly or transitively.

## Install

```bash
go get github.com/sweetrpg/common.go
```

## Packages

- `util` - `GetEnv`/`GetEnvInt` (environment variable helpers with defaults), `Map`/`NullMap`
  (generic slice transforms)
- `logging` - `Init()` configures a `logf`-backed logger from the `LOG_LEVEL` environment
  variable; `Logger` is the shared instance
- `constants` - environment variable names and log-level values used by the packages above

## Documentation

Package documentation: [pkg.go.dev/github.com/sweetrpg/common.go](https://pkg.go.dev/github.com/sweetrpg/common.go).
Test coverage reports are published to [sweetrpg.github.io/common.go](https://sweetrpg.github.io/common.go)
on every merge to `develop`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the development workflow and
[RELEASE.md](RELEASE.md) for how versions get cut.
