package integration

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/rparaschak/mono-tmpl/api/internal/bootstrap"
	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/rparaschak/mono-tmpl/api/pkg/httpapi"
	"github.com/stretchr/testify/require"
)

type AutotestEnv struct {
	Deps   modules.GlobalDependencies
	Expect *httpexpect.Expect
}

func NewAutotestEnv(t *testing.T) *AutotestEnv {
	t.Helper()

	deps := NewAutotestDependencies(t)
	mux, api := httpapi.NewRouter()
	bootstrap.RegisterRoutes(api, deps)

	expect := httpexpect.WithConfig(httpexpect.Config{
		TestName: t.Name(),
		BaseURL:  "http://mono-tmpl.test",
		Client: &http.Client{
			Transport: httpexpect.NewBinder(mux),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewRequireReporter(t),
	})

	return &AutotestEnv{
		Deps:   deps,
		Expect: expect,
	}
}

func NewAutotestDependencies(t *testing.T) modules.GlobalDependencies {
	t.Helper()

	databaseURL := os.Getenv("DATABASE_URL")
	require.NotEmpty(t, databaseURL, "DATABASE_URL must be set for integration tests")

	cfg, err := config.Load()
	require.NoError(t, err, "autotest config should load")
	cfg.HTTPServer.Env = "autotest"
	cfg.Database.Env = "autotest"
	cfg.Database.URL = databaseURL

	deps, err := modules.NewDependencies(context.Background(), cfg)
	require.NoError(t, err, "autotest dependencies should initialize")

	sqlDB, err := deps.DB.DB()
	require.NoError(t, err, "autotest DB handle should be available")
	t.Cleanup(func() {
		require.NoError(t, sqlDB.Close(), "autotest DB handle should close cleanly")
	})

	return deps
}
