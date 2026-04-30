package usecases

import (
	"cmp"
	"context"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"gorm.io/gorm"
)

type SortField string
type SortOrder string

const (
	SortFieldName      SortField = "name"
	SortFieldCreatedAt SortField = "created_at"

	SortOrderASC  SortOrder = "ASC"
	SortOrderDESC SortOrder = "DESC"

	DefaultSortField SortField = SortFieldCreatedAt
	DefaultSortOrder SortOrder = SortOrderDESC
)

type GetSamplesInput struct {
	Prefix string
	Sort   GetSamplesSortInput
}

type GetSamplesSortInput struct {
	Field SortField
	Order SortOrder
}

func (u UseCase) GetSamples(ctx context.Context, input GetSamplesInput) ([]persistence.Sample, error) {
	input.Sort.Field = cmp.Or(input.Sort.Field, DefaultSortField)
	input.Sort.Order = cmp.Or(input.Sort.Order, DefaultSortOrder)
	sortStr := string(input.Sort.Field) + " " + string(input.Sort.Order)

	q := gorm.G[persistence.Sample](u.DB).
		Order(sortStr)

	if input.Prefix != "" {
		q = q.Where("name ILIKE ?", input.Prefix+"%")
	}

	return q.Find(ctx)
}
