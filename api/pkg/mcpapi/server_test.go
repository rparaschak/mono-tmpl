package mcpapi

import (
	"encoding/json"
	"testing"

	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithInputSchemaAddsEnumTags(t *testing.T) {
	type input struct {
		SortField string `json:"sortField,omitempty" enum:"name,created_at" jsonschema:"Sort field"`
		SortOrder string `json:"sortOrder,omitempty" enum:"ASC,DESC"         jsonschema:"Sort order"`
	}

	tool := mcplib.NewTool("test_tool", WithInputSchema[input]())

	require.Empty(t, tool.InputSchema.Type, "raw input schema should replace the default input schema")
	require.NotEmpty(t, tool.RawInputSchema, "tool should expose a raw input schema")

	var schema map[string]any
	require.NoError(t, json.Unmarshal(tool.RawInputSchema, &schema), "input schema should decode")

	properties := schema["properties"].(map[string]any)
	sortField := properties["sortField"].(map[string]any)
	sortOrder := properties["sortOrder"].(map[string]any)

	assert.Equal(t, []any{"name", "created_at"}, sortField["enum"], "sort field should expose enum values")
	assert.Equal(t, []any{"ASC", "DESC"}, sortOrder["enum"], "sort order should expose enum values")
}
