package main

import (
	"github.com/mwantia/forge-plugin-mcp/internal/mcp"
	"github.com/mwantia/forge-sdk/pkg/plugins/grpc"
)

func main() {
	grpc.Serve(mcp.NewMCPDriver)
}
