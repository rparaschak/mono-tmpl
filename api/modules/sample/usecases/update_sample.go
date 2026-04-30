package usecases

import (
	"context"

	"github.com/google/uuid"
	sampleErrors "github.com/rparaschak/mono-tmpl/api/modules/sample/errors"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"github.com/rparaschak/mono-tmpl/api/pkg/database"
	"gorm.io/gorm/clause"
)

type UpdateSampleInput struct {
	Id     uuid.UUID
	Sample SampleInput
}

func (u UseCase) UpdateSample(ctx context.Context, input UpdateSampleInput) (persistence.Sample, error) {
	updatedSample := persistence.Sample{
		Id:          input.Id,
		Name:        input.Sample.Name,
		Geolocation: database.NewGeolocation(input.Sample.Longitude, input.Sample.Latitude),
	}

	result := u.DB.WithContext(ctx).
		Clauses(clause.Returning{}).
		Select("name", "geolocation").
		Where("id = ?", input.Id).
		Updates(&updatedSample)
	if result.Error != nil {
		return persistence.Sample{}, result.Error
	}
	if result.RowsAffected == 0 {
		return persistence.Sample{}, sampleErrors.ErrSampleNotFound
	}
	return updatedSample, nil
}
