package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type UpdateSampleRequest struct {
	SampleId uuid.UUID `path:"sampleId" format:"uuid"`
	Body     contracts.SampleInputDTO
}

type UpdateSampleResponse struct {
	Body contracts.SampleDTO
}

func (h Handlers) UpdateSampleHandler(ctx context.Context, input *UpdateSampleRequest) (*UpdateSampleResponse, error) {
	sample, err := h.UseCases.UpdateSample(ctx, usecases.UpdateSampleInput{
		Id: input.SampleId,
		Sample: usecases.SampleInput{
			Name:      input.Body.Name,
			Latitude:  input.Body.Latitude,
			Longitude: input.Body.Longitude,
		},
	})
	if err != nil {
		return nil, err
	}

	return &UpdateSampleResponse{
		Body: contracts.NewSampleDTO(sample),
	}, nil
}
