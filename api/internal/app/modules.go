package app

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/mark3labs/mcp-go/server"

	"github.com/rparaschak/mono-tmpl/api/internal/dependencies"
	"github.com/rparaschak/mono-tmpl/api/modules/core"
	"github.com/rparaschak/mono-tmpl/api/modules/sample"
)

type Modules struct {
	Core   *core.Module
	Sample *sample.Module
}

func NewModules(deps dependencies.Dependencies) *Modules {
	return &Modules{
		Core:   core.New(deps),
		Sample: sample.New(deps),
	}
}

func (m *Modules) RegisterHTTP(api huma.API) {
	m.Core.RegisterHTTP(api)
	m.Sample.RegisterHTTP(api)
}

func (m *Modules) RegisterMCP(server *server.MCPServer) {
	m.Sample.RegisterMCP(server)
}
