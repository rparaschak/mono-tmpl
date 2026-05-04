package mcptest

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	mcpclient "github.com/mark3labs/mcp-go/client"
	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/rparaschak/mono-tmpl/api/internal/testsupport/integration"
	"github.com/stretchr/testify/require"
)

type AutotestEnv struct {
	*integration.AutotestEnv
	Server *httptest.Server
}

func NewAutotestEnv(t *testing.T) *AutotestEnv {
	t.Helper()

	env := integration.NewAutotestEnv(t)
	server := httptest.NewServer(env.App.Handler())
	t.Cleanup(server.Close)

	return &AutotestEnv{
		AutotestEnv: env,
		Server:      server,
	}
}

func (e *AutotestEnv) NewClient(t *testing.T) *mcpclient.Client {
	t.Helper()

	return NewClient(t, e.Server.URL)
}

func NewClient(t *testing.T, baseURL string) *mcpclient.Client {
	t.Helper()

	client, err := mcpclient.NewStreamableHttpClient(baseURL + "/mcp")
	require.NoError(t, err, "MCP client should initialize")
	t.Cleanup(func() {
		require.NoError(t, client.Close(), "MCP client should close")
	})

	ctx := t.Context()
	require.NoError(t, client.Start(ctx), "MCP client should start")
	_, err = client.Initialize(ctx, mcplib.InitializeRequest{
		Params: mcplib.InitializeParams{
			ProtocolVersion: mcplib.LATEST_PROTOCOL_VERSION,
			ClientInfo: mcplib.Implementation{
				Name:    "mono-tmpl-test",
				Version: "1.0.0",
			},
		},
	})
	require.NoError(t, err, "MCP client should complete protocol initialization")

	return client
}

func DecodeStructuredContent[T any](t *testing.T, content any) T {
	t.Helper()

	data, err := json.Marshal(content)
	require.NoError(t, err, "structured content should marshal")

	var output T
	require.NoError(t, json.Unmarshal(data, &output), "structured content should decode")
	return output
}
