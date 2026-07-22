# AGENTS.md

This file provides guidance to Claude Code, Codex, GitHub Copilot, and other AI coding agents
working in this repository.

## About This Project

`common.go` provides small, dependency-free utility packages (environment variable helpers,
generic slice mapping, structured logging setup, shared constants) used across the sweetrpg
platform's Go services and libraries. It has no internal sweetrpg dependencies itself, making it
the base of the platform's Go dependency graph.

## Consumers

Depended on directly or transitively by nearly every other sweetrpg Go module, including
`mongodb.go`, `api-core.go`, `model-core.go`, `catalog-objects.go`, `catalog-data.go`, and
`catalog-api`. Breaking changes here ripple across the whole platform - keep the public API
additive where possible, and bump appropriately (see Releases below) when it isn't.

## Committing Code

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

## Branches and Workflow

* `develop` - integration branch, default branch, target for all PRs.
* `master` - latest released state, nothing committed directly.
* `feature/*`, `fix/*` branched from `develop`; `hotfix/*` branched from `master`.

See `CONTRIBUTING.md` for the full workflow.

## Running Checks Locally

```bash
go build -v ./...
go vet ./...
go test -v -coverprofile coverage.out ./...
go tool cover -func coverage.out
```

## Releases

Merges to `develop` auto-tag a patch release via CI (`.github/workflows/go-ci.yml`). Use the
"Bump version" workflow (`.github/workflows/bump-version.yml`, manually dispatched) for a minor
or major bump instead.
