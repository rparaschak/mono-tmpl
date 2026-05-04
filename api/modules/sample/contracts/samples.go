package contracts

import (
	"time"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
)

type SampleDTO struct {
	Id        string    `json:"id"        format:"uuid" readOnly:"true"`
	Name      string    `json:"name"      minLength:"1"`
	Latitude  float64   `json:"latitude"  minimum:"-90"  maximum:"90"`
	Longitude float64   `json:"longitude" minimum:"-180" maximum:"180"`
	CreatedAt time.Time `json:"createdAt" readOnly:"true"`
}

type SampleMCPDTO struct {
	Id        string    `json:"id"        jsonschema:"Sample ID"`
	Name      string    `json:"name"      jsonschema:"Sample name"`
	Latitude  float64   `json:"latitude"  jsonschema:"Sample latitude"`
	Longitude float64   `json:"longitude" jsonschema:"Sample longitude"`
	CreatedAt time.Time `json:"createdAt" jsonschema:"Sample creation time"`
}

type SampleInputDTO struct {
	Name      string  `json:"name"      minLength:"1"`
	Latitude  float64 `json:"latitude"  minimum:"-90"  maximum:"90"`
	Longitude float64 `json:"longitude" minimum:"-180" maximum:"180"`
}

func NewSampleDTO(sample persistence.Sample) SampleDTO {
	return SampleDTO{
		Id:        sample.Id.String(),
		Name:      sample.Name,
		Latitude:  sample.Geolocation.Lat,
		Longitude: sample.Geolocation.Lng,
		CreatedAt: sample.CreatedAt,
	}
}

func NewSampleMCPDTO(sample persistence.Sample) SampleMCPDTO {
	return SampleMCPDTO{
		Id:        sample.Id.String(),
		Name:      sample.Name,
		Latitude:  sample.Geolocation.Lat,
		Longitude: sample.Geolocation.Lng,
		CreatedAt: sample.CreatedAt,
	}
}
