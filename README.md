# phenotype-ops-mcp

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
