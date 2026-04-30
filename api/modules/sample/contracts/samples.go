package contracts

import (
	"time"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
)

type SampleDTO struct {
	Id        string    `json:"id"        format:"uuid" readOnly:"true"`
	Name      string    `json:"name"      minLength:"1"`
	CreatedAt time.Time `json:"createdAt" readOnly:"true"`
}

func NewSampleDTO(sample persistence.Sample) SampleDTO {
	return SampleDTO{
		Id:        sample.Id.String(),
		Name:      sample.Name,
		CreatedAt: sample.CreatedAt,
	}
}
