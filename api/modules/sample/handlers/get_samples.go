package handlers

import "context"

type GetSamplesRequest struct{}

type GetSamplesResponse struct {
	Body GetSamplesResponseBody
}

type GetSamplesResponseBody struct {
	Samples []string `json:"samples"`
}

func (h Handlers) GetSamplesHandler(ctx context.Context, input *GetSamplesRequest) (*GetSamplesResponse, error) {
	return &GetSamplesResponse{
		Body: GetSamplesResponseBody{
			Samples: h.UseCases.GetSamples(),
		},
	}, nil
}
