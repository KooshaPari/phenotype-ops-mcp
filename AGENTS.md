# AGENTS.md — phenotype-ops-mcp

## Agent Onboarding

This repository contains Go code for: Go MCP server for Phenotype operations

### Key Paths
- `src/` — Source code
- `crates/` — Workspace crates (if applicable)
- `tests/` — Test suites
- `Cargo.toml` / `package.json` — Dependencies

### Before Starting Work

1. Install dependencies: `cargo fetch` / `npm install` / `go mod download`
2. Verify baseline: `go test ./...`
3. Check lints: `go vet ./...`

### Committing

Use conventional commits with scope:
```
feat(module): add new feature
fix(module): resolve bug
docs: update README
test: add test coverage
```

### CI/CD

Local verification required before push:
- `go test ./...`
- `go vet ./...`
- `go build ./...`
