package handlers

import (
	"context"
	"net/http"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type CreateSampleRequest struct {
	Body contracts.SampleInputDTO
}

type CreateSampleResponse struct {
	Status int
	Body   contracts.SampleDTO
}

func (h Handlers) CreateSampleHandler(ctx context.Context, input *CreateSampleRequest) (*CreateSampleResponse, error) {
	sample, err := h.UseCases.CreateSample(ctx, usecases.CreateSampleInput{
		Sample: usecases.SampleInput{
			Name:      input.Body.Name,
			Latitude:  input.Body.Latitude,
			Longitude: input.Body.Longitude,
		},
	})
	if err != nil {
		return nil, err
	}

	return &CreateSampleResponse{
		Status: http.StatusCreated,
		Body:   contracts.NewSampleDTO(*sample),
	}, nil
}
