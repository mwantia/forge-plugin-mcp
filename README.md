# forge-plugin-mcp

Tools plugin that bridges the Forge agent to external [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) servers, aggregating their tools under a single plugin.

## Capabilities

| Capability | Supported |
|---|---|
| Tools | yes |
| Async execution | no |

## Configuration

Each entry under `servers` configures one MCP server. Two transports are supported: `sse` (HTTP/SSE) and `stdio` (local subprocess).

```hcl
plugin "mcp" {
  servers {
    # SSE transport — remote server
    my-remote-server {
      transport = "sse"
      url       = "https://mcp.example.com/sse"

      headers {
        Authorization = "Bearer <token>"
      }
    }

    # stdio transport — local subprocess
    my-local-server {
      transport = "stdio"
      command   = "/usr/local/bin/my-mcp-server"
      args      = ["--config", "/etc/mcp/config.json"]
    }
  }
}
```

| Field | Type | Description |
|---|---|---|
| `servers.<name>.transport` | string | `"sse"` or `"stdio"` |
| `servers.<name>.url` | string | Server endpoint (SSE only) |
| `servers.<name>.headers` | map | HTTP headers (SSE only) |
| `servers.<name>.command` | string | Executable path (stdio only) |
| `servers.<name>.args` | list | Command arguments (stdio only) |

## Tool naming

Tools from each MCP server are exposed as `<server-name>__<tool-name>` (double underscore separator). For example, a tool `read_file` from a server named `filesystem` is accessible as `filesystem__read_file`.

## Notes

- Servers that fail to connect at startup are logged and skipped — the plugin continues with any successfully connected servers.
- All tool discovery and execution is thread-safe.
