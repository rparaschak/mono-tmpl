package usecases

import (
	"context"

	"github.com/google/uuid"
	sampleErrors "github.com/rparaschak/mono-tmpl/api/modules/sample/errors"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"gorm.io/gorm"
)

func (u UseCase) DeleteSample(ctx context.Context, id uuid.UUID) error {
	rowsAffected, err := gorm.G[persistence.Sample](u.DB).
		Where("id = ?", id).
		Delete(ctx)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sampleErrors.ErrSampleNotFound
	}
	return nil
}
