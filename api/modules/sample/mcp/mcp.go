package mcp

import (
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type Handlers struct {
	UseCases *usecases.UseCase
}

func NewHandlers(useCases *usecases.UseCase) *Handlers {
	return &Handlers{UseCases: useCases}
}
