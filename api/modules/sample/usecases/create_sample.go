package usecases

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"github.com/rparaschak/mono-tmpl/api/pkg/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateSampleInput struct {
	Sample SampleInput
}

func (u UseCase) CreateSample(ctx context.Context, input CreateSampleInput) (persistence.Sample, error) {
	newSample := persistence.Sample{
		Name:        input.Sample.Name,
		Geolocation: database.NewGeolocation(input.Sample.Longitude, input.Sample.Latitude),
	}

	err := gorm.G[persistence.Sample](u.DB, clause.Returning{}).
		Select("name", "geolocation").
		Create(ctx, &newSample)
	if err != nil {
		return persistence.Sample{}, err
	}

	return newSample, nil
}
