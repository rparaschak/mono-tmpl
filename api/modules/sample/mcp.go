package sample

import (
	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	samplemcp "github.com/rparaschak/mono-tmpl/api/modules/sample/mcp"
)

func RegisterTools(mcpServer *server.MCPServer, h *samplemcp.Handlers) {
	mcpServer.AddTool(samplemcp.GetSamplesTool(), mcplib.NewStructuredToolHandler(h.GetSamplesHandler))
}
