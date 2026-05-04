package usecases

import "github.com/rparaschak/mono-tmpl/api/internal/dependencies"

type UseCase struct {
	dependencies.Dependencies
}

type SampleInput struct {
	Name      string
	Latitude  float64
	Longitude float64
}
