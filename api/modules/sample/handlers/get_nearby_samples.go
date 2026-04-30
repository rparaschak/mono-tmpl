package handlers

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type GetNearbySamplesRequest struct {
	Latitude  float64 `query:"lat"   required:"true" minimum:"-90"  maximum:"90"`
	Longitude float64 `query:"lng"   required:"true" minimum:"-180" maximum:"180"`
	Limit     int     `query:"limit" required:"true" minimum:"1"    maximum:"100"`
}

type GetNearbySamplesRequestBody struct{}

type GetNearbySamplesResponse struct {
	Body GetNearbySamplesResponseBody
}

type GetNearbySamplesResponseBody struct {
	Samples []contracts.SampleDTO `json:"samples"`
}

func (h Handlers) GetNearbySamplesHandler(ctx context.Context, input *GetNearbySamplesRequest) (*GetNearbySamplesResponse, error) {
	samples, err := h.UseCases.GetNearbySamples(ctx, usecases.GetNearbySamplesInput{
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Limit:     input.Limit,
	})
	if err != nil {
		return nil, err
	}

	sampleDTOs := make([]contracts.SampleDTO, len(samples))
	for i, sample := range samples {
		sampleDTOs[i] = contracts.NewSampleDTO(sample)
	}

	return &GetNearbySamplesResponse{
		Body: GetNearbySamplesResponseBody{
			Samples: sampleDTOs,
		},
	}, nil
}
