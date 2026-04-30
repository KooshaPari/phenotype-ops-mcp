# phenotype-ops-mcp

[![AI Slop Inside](https://sladge.net/badge.svg)](https://sladge.net)

> **Phenotype fork notice**
>
> This repository is a Phenotype-org fork of [nanovms/ops-mcp](https://github.com/nanovms/ops-mcp),
> the MCP server for the nanoVMs `ops` unikernel toolchain. It is maintained by
> [KooshaPari](https://github.com/KooshaPari) as part of the Phenotype MCP-first
> agent stack, where it acts as the bridge between agent tools and
> [`phenotype-nanovms-client`](https://github.com/KooshaPari/phenotype-shared)
> (promoted from `bare-cua`).
>
> - **Upstream:** https://github.com/nanovms/ops-mcp (tracked via `upstream` remote)
> - **Fork:** https://github.com/KooshaPari/phenotype-ops-mcp
> - **License:** Apache-2.0 (preserved from upstream — see `LICENSE`)
>
> Fork rationale: extend the upstream server with Phenotype-specific tooling
> (auth, multi-tenant instance isolation, observability hooks) without losing
> the ability to pull upstream updates. Contributions that are generic enough
> for upstream will be PR'd back to `nanovms/ops-mcp` first.

---

# ops-mcp (upstream README)

mcp server for ops

Build like normally.

Put this in your Claud config:

```
~/Library/Application Support/Claude/claude_desktop_config.json
```

Ensure your command is in the right path and more importantly the PATH
env is set to run.

```
{
"mcpServers": {
  "ops-mcp": {
      "command": "/Users/eyberg/go/src/github.com/nanovms/ops-mcp/ops-mcp",
      "args": [],
      "env": {
        "HOME":"/Users/eyberg",
        "LOGNAME":"eyberg",
        "PATH":"/bin:/Users/eyberg/.ops/bin",
        "SHELL":"/bin/zsh",
        "USER":"eyberg"
        }
    }
  }
}
```

Available tools:

```
List instances
```

```
List images
```

```
Instance create <image_name>
```

```
Instance create redis-server
```

Note: Very open to suggestions on how this all should work as this initial cut was done not having
ever used Claude or MCP.

## Tool Manifest

A canonical machine-readable manifest of every tool this server exposes lives at
[`tools.json`](./tools.json). It lists each tool's `name`, `description`, and
`inputSchema` (JSON Schema reflected from the handler argument struct) so
downstream consumers can discover the surface without reading `main.go`.

Regenerate after any tool change:

```
go run . --dump-tools
```

The [`Manifest check`](./.github/workflows/manifest-check.yml) workflow runs the
same command in CI and fails the build if `tools.json` drifts from the code, so
the manifest cannot go stale.

Note: `metoro-io/mcp-golang@v0.13.0` does not expose a public accessor for
registered tools (the internal `tools` map is unexported). The manifest is
built from the single `toolRegistrations()` source of truth in `main.go` and
re-reflects input schemas via `invopop/jsonschema` (the same library the SDK
uses internally). If upstream adds a `ListTools()` accessor, this plumbing can
be simplified.
