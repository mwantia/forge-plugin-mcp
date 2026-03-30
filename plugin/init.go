package plugin

import (
	"github.com/mwantia/forge-plugin-mcp/internal/mcp"
	"github.com/mwantia/forge-sdk/pkg/plugins"
)

func init() {
	plugins.Register(mcp.PluginName, mcp.PluginDescription, mcp.NewMCPDriver)
}
