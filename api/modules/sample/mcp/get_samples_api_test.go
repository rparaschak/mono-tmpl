//go:build integration

package mcp_test

import (
	"testing"

	mcplib "github.com/mark3labs/mcp-go/mcp"
	samplemcp "github.com/rparaschak/mono-tmpl/api/modules/sample/mcp"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/testkit"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
	"github.com/rparaschak/mono-tmpl/api/pkg/mcpapi/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSamplesMCP(t *testing.T) {
	env := mcptest.NewAutotestEnv(t)

	t.Run("POST /mcp", func(t *testing.T) {
		t.Run("initializes and lists sample tool", func(t *testing.T) {
			client := env.NewClient(t)

			tools, err := client.ListTools(t.Context(), mcplib.ListToolsRequest{})
			require.NoError(t, err, "client should list MCP tools")
			require.Len(t, tools.Tools, 1, "MCP should expose exactly one tool")

			tool := tools.Tools[0]
			assert.Equal(t, samplemcp.GetSamplesToolName, tool.Name, "tool should expose sample list")
			assert.Contains(t, tool.InputSchema.Properties, "prefix", "tool should expose prefix input")
			assert.Contains(t, tool.InputSchema.Properties, "sortField", "tool should expose sort field input")
			assert.Contains(t, tool.InputSchema.Properties, "sortOrder", "tool should expose sort order input")
			assert.Contains(t, tool.OutputSchema.Properties, "samples", "tool should expose samples output")
		})

		t.Run("returns filtered samples from usecase", func(t *testing.T) {
			client := env.NewClient(t)
			factory := testkit.NewSamplesFactory(t, env.Deps)
			alphaZ := factory.Create(testkit.InputNamed("mcp-alpha-z"))
			alphaA := factory.Create(testkit.InputNamed("mcp-alpha-a"))
			factory.Create(testkit.InputNamed("mcp-beta"))

			result, err := client.CallTool(t.Context(), mcplib.CallToolRequest{
				Params: mcplib.CallToolParams{
					Name: samplemcp.GetSamplesToolName,
					Arguments: samplemcp.GetSamplesInput{
						Prefix:    "mcp-alpha",
						SortField: string(usecases.SortFieldName),
						SortOrder: string(usecases.SortOrderASC),
					},
				},
			})
			require.NoError(t, err, "sample list tool should execute")
			require.False(t, result.IsError, "sample list tool should not return a tool error")

			output := mcptest.DecodeStructuredContent[samplemcp.GetSamplesOutput](t, result.StructuredContent)
			require.Len(t, output.Samples, 2, "sample list tool should return matching samples only")
			assert.Equal(t, alphaA.Name, output.Samples[0].Name, "sample list tool should sort by name ascending")
			assert.Equal(t, alphaZ.Name, output.Samples[1].Name, "sample list tool should include the second matching sample")
		})

		t.Run("returns tool error for invalid sort", func(t *testing.T) {
			client := env.NewClient(t)

			result, err := client.CallTool(t.Context(), mcplib.CallToolRequest{
				Params: mcplib.CallToolParams{
					Name: samplemcp.GetSamplesToolName,
					Arguments: samplemcp.GetSamplesInput{
						SortField: "id",
					},
				},
			})
			require.NoError(t, err, "invalid sort should be returned as a tool result error")
			require.True(t, result.IsError, "invalid sort should mark the tool result as an error")
		})
	})
}
