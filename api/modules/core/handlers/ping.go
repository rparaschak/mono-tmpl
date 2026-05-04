package handlers

import "context"

type PingInput struct{}

type PingOutput struct{}

func (h *Handlers) PingHandler(_ context.Context, _ *PingInput) (*PingOutput, error) {
	return &PingOutput{}, nil
}
