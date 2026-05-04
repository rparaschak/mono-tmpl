package integration

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/rparaschak/mono-tmpl/api/internal/app"
	"github.com/rparaschak/mono-tmpl/api/internal/dependencies"
	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/stretchr/testify/require"
)

type AutotestEnv struct {
	App    *app.App
	Deps   dependencies.Dependencies
	Expect *httpexpect.Expect
}

func NewAutotestEnv(t *testing.T) *AutotestEnv {
	t.Helper()

	cfg := NewAutotestConfig(t)
	application, err := app.New(context.Background(), cfg)
	require.NoError(t, err, "autotest app should initialize")
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		require.NoError(t, application.Shutdown(ctx), "autotest app should shut down cleanly")
	})

	expect := httpexpect.WithConfig(httpexpect.Config{
		TestName: t.Name(),
		BaseURL:  "http://mono-tmpl.test",
		Client: &http.Client{
			Transport: httpexpect.NewBinder(application.Handler()),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewRequireReporter(t),
	})

	return &AutotestEnv{
		App:    application,
		Deps:   application.Deps,
		Expect: expect,
	}
}

func NewAutotestDependencies(t *testing.T) dependencies.Dependencies {
	t.Helper()

	cfg := NewAutotestConfig(t)
	deps, err := dependencies.NewAutotest(context.Background(), cfg)
	require.NoError(t, err, "autotest dependencies should initialize")
	t.Cleanup(func() {
		require.NoError(t, deps.Close(), "autotest dependencies should close cleanly")
	})

	return deps
}

func NewAutotestConfig(t *testing.T) config.Config {
	t.Helper()

	databaseURL := os.Getenv("DATABASE_URL")
	require.NotEmpty(t, databaseURL, "DATABASE_URL must be set for integration tests")

	cfg, err := config.Load()
	require.NoError(t, err, "autotest config should load")
	cfg.HTTPServer.Env = "autotest"
	cfg.Database.Env = "autotest"
	cfg.Database.URL = databaseURL

	return cfg
}
