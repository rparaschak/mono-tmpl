package bootstrap

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/modules/sample"
)

type Modules struct {
	Sample *sample.Module
}

func RegisterRoutes(api huma.API, deps modules.GlobalDependencies) Modules {
	registry := Modules{
		Sample: sample.New(deps),
	}

	registry.Sample.WithRestRouter(api)
	return registry
}
