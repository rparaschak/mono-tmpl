package handlers

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type GetSamplesRequest struct {
	Prefix    string `query:"prefix"`
	SortField string `query:"sortField" enum:"name,createdAt"`
	SortOrder string `query:"sortOrder" enum:"ASC,DESC"`
}

type GetSamplesRequestBody struct{}

type GetSamplesResponse struct {
	Body GetSamplesResponseBody
}

type GetSamplesResponseBody struct {
	Samples []contracts.SampleDTO `json:"samples"`
}

func (h Handlers) GetSamplesHandler(ctx context.Context, input *GetSamplesRequest) (*GetSamplesResponse, error) {
	samples, err := h.UseCases.GetSamples(ctx, usecases.GetSamplesInput{
		Prefix: input.Prefix,
		Sort: usecases.GetSamplesSortInput{
			Field: usecases.SortField(input.SortField),
			Order: usecases.SortOrder(input.SortOrder),
		},
	})
	if err != nil {
		return nil, err
	}

	sampleDTOs := make([]contracts.SampleDTO, len(samples))
	for i, sample := range samples {
		sampleDTOs[i] = contracts.NewSampleDTO(sample)
	}

	return &GetSamplesResponse{
		Body: GetSamplesResponseBody{
			Samples: sampleDTOs,
		},
	}, nil
}
