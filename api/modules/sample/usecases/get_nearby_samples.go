package usecases

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetNearbySamplesInput struct {
	Latitude  float64
	Longitude float64
	Limit     int
}

func (u UseCase) GetNearbySamples(ctx context.Context, input GetNearbySamplesInput) ([]persistence.Sample, error) {
	return gorm.G[persistence.Sample](u.DB).
		Order(clause.Expr{
			SQL:  "geolocation <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography",
			Vars: []any{input.Longitude, input.Latitude},
		}).
		Limit(input.Limit).
		Find(ctx)
}
