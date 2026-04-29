package sample

import (
	"testing"

	"github.com/rparaschak/mono-tmpl/api/modules"
)

func TestNewPassesDependenciesToUseCases(t *testing.T) {
	deps := modules.GlobalDependencies{}

	module := New(deps)

	if module.Deps != deps {
		t.Fatalf("Deps = %#v, want %#v", module.Deps, deps)
	}

	if module.UseCases == nil {
		t.Fatal("UseCases is nil")
	}

	if module.UseCases.GlobalDependencies != deps {
		t.Fatalf("UseCases.GlobalDependencies = %#v, want %#v", module.UseCases.GlobalDependencies, deps)
	}

	if module.Handlers == nil {
		t.Fatal("Handlers is nil")
	}

	if module.Handlers.UseCases != module.UseCases {
		t.Fatal("Handlers.UseCases does not point to module UseCases")
	}
}
