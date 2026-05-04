package sample

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/mark3labs/mcp-go/server"

	"github.com/rparaschak/mono-tmpl/api/internal/dependencies"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/handlers"
	samplemcp "github.com/rparaschak/mono-tmpl/api/modules/sample/mcp"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type Module struct {
	Deps        dependencies.Dependencies
	UseCases    *usecases.UseCase
	Handlers    *handlers.Handlers
	MCPHandlers *samplemcp.Handlers
}

func New(deps dependencies.Dependencies) *Module {
	useCases := &usecases.UseCase{Dependencies: deps}
	return &Module{
		Deps:        deps,
		UseCases:    useCases,
		Handlers:    handlers.NewHandlers(useCases),
		MCPHandlers: samplemcp.NewHandlers(useCases),
	}
}

func (m *Module) RegisterHTTP(api huma.API) {
	RegisterRoutes(api, m.Handlers)
}

func (m *Module) RegisterMCP(mcpServer *server.MCPServer) {
	RegisterTools(mcpServer, m.MCPHandlers)
}
