# phenotype-ops-mcp — MCP Server for NanoVM Instance Management

Model Context Protocol (MCP) server for the nanoVMs `ops` unikernel toolchain, extended with Phenotype-specific features.

**Core Mission**: Unified MCP interface for nanoVM operations within Phenotype's agent-driven infrastructure.

## Technology Stack

- **Language**: Go (1.20+)
- **MCP Protocol**: Model Context Protocol (SSE transport)
- **Upstream**: nanoVMs `ops` toolchain
- **Deployment**: Docker, systemd, or standalone binary

## Key Features

- **Instance Lifecycle**: Create, start, stop, delete nanoVM instances
- **Image Management**: List, build, pull nanoVM images
- **Multi-Tenant Isolation**: Namespace-aware instance grouping
- **Observability**: Structured logging + distributed tracing

## Related

- **nvms**: `/Users/kooshapari/CodeProjects/Phenotype/repos/nanovms` — microVM runtime
- **Pine**: `/Users/kooshapari/CodeProjects/Phenotype/repos/Pine` — Wine-equivalent compatibility layer
- **PhenoCompose**: `/Users/kooshapari/CodeProjects/Phenotype/repos/PhenoCompose` — unified NVMS interface
