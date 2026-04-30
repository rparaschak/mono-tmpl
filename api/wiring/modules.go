package wiring

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/modules/sample"
	"github.com/rparaschak/mono-tmpl/api/pkg/httpapi"
)

type ModulesRepo struct {
	Sample *sample.Module
}

func (r *ModulesRepo) Init(deps modules.GlobalDependencies) {
	r.Sample = sample.New(deps)
}

func (r *ModulesRepo) Start(api huma.API) {
	r.Sample.WithRestRouter(api)
}

func RegisterModules(api huma.API, deps modules.GlobalDependencies) ModulesRepo {
	registry := ModulesRepo{}
	registry.Init(deps)
	registry.Start(api)
	return registry
}

func RouteRegistrar(deps modules.GlobalDependencies) httpapi.RouteRegistrar {
	return func(api huma.API) {
		RegisterModules(api, deps)
	}
}
