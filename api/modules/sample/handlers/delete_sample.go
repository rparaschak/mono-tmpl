package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type DeleteSampleRequest struct {
	SampleId uuid.UUID `path:"sampleId" format:"uuid"`
}

type DeleteSampleResponse struct {
	Status int
}

func (h Handlers) DeleteSampleHandler(ctx context.Context, input *DeleteSampleRequest) (*DeleteSampleResponse, error) {
	if err := h.UseCases.DeleteSample(ctx, input.SampleId); err != nil {
		return nil, err
	}

	return &DeleteSampleResponse{Status: http.StatusNoContent}, nil
}
