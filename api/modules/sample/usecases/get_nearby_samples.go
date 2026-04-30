package usecases

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"github.com/rparaschak/mono-tmpl/api/pkg/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetNearbySamplesInput struct {
	Latitude  float64
	Longitude float64
	Limit     int
}

func (u UseCase) GetNearbySamples(ctx context.Context, input GetNearbySamplesInput) ([]persistence.Sample, error) {
	geolocation := database.NewGeolocation(input.Longitude, input.Latitude)

	return gorm.G[persistence.Sample](u.DB).
		Order(clause.Expr{
			SQL:  "geolocation <-> ?::geography",
			Vars: []any{geolocation},
		}).
		Limit(input.Limit).
		Find(ctx)
}
