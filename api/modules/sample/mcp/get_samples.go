package mcp

import (
	"context"

	mcplib "github.com/mark3labs/mcp-go/mcp"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
	"github.com/rparaschak/mono-tmpl/api/pkg/mcpapi"
)

const GetSamplesToolName = "get_samples"

type GetSamplesInput struct {
	Prefix    string `json:"prefix,omitempty"    jsonschema:"Optional sample name prefix filter"`
	SortField string `json:"sortField,omitempty" enum:"name,created_at" jsonschema:"Optional sort field"`
	SortOrder string `json:"sortOrder,omitempty" enum:"ASC,DESC"         jsonschema:"Optional sort order"`
}

type GetSamplesOutput struct {
	Samples []contracts.SampleMCPDTO `json:"samples" jsonschema:"Samples matching the requested filters"`
}

func GetSamplesTool() mcplib.Tool {
	return mcplib.NewTool(
		GetSamplesToolName,
		mcplib.WithDescription("List samples with optional prefix and sorting."),
		mcpapi.WithInputSchema[GetSamplesInput](),
		mcplib.WithOutputSchema[GetSamplesOutput](),
	)
}

func (h Handlers) GetSamplesHandler(ctx context.Context, _ mcplib.CallToolRequest, input GetSamplesInput) (*GetSamplesOutput, error) {
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

	output := &GetSamplesOutput{
		Samples: make([]contracts.SampleMCPDTO, len(samples)),
	}
	for i, sample := range samples {
		output.Samples[i] = contracts.NewSampleMCPDTO(sample)
	}

	return output, nil
}
