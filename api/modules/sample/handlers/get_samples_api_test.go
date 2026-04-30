//go:build integration

package handlers_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/rparaschak/mono-tmpl/api/internal/testsupport/integration"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/testkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSamplesAPI(t *testing.T) {
	env := integration.NewAutotestEnv(t)
	factory := testkit.NewSamplesFactory(t, env.Deps)

	t.Run("GET /samples", func(t *testing.T) {
		t.Run("returns default sort by created_at descending", func(t *testing.T) {
			prefix := factory.UniqueName("default-sort")
			older := factory.Create(testkit.InputNamed(prefix + "-older"))
			newer := factory.Create(testkit.InputNamed(prefix + "-newer"))
			factory.SetCreatedAt(older, time.Date(2026, 4, 29, 12, 0, 0, 0, time.UTC))
			factory.SetCreatedAt(newer, time.Date(2026, 4, 30, 12, 0, 0, 0, time.UTC))

			var body testkit.SampleListResponse
			env.Expect.GET("/samples").
				WithQuery("prefix", prefix).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Decode(&body)

			require.Len(t, body.Samples, 2, "prefix-filtered response should include only samples created for this test")
			assert.Equal(t, []string{newer.Name, older.Name}, testkit.DTONames(body.Samples), "samples should be sorted by created_at descending by default")
		})

		t.Run("filters by prefix and sorts by name", func(t *testing.T) {
			prefix := factory.UniqueName("alpha")
			second := factory.Create(testkit.InputNamed(prefix + "-2"))
			factory.Create(testkit.InputNamed(factory.UniqueName("beta")))
			first := factory.Create(testkit.InputNamed(prefix + "-1"))

			var body testkit.SampleListResponse
			env.Expect.GET("/samples").
				WithQuery("prefix", prefix).
				WithQuery("sortField", "name").
				WithQuery("sortOrder", "ASC").
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Decode(&body)

			require.Len(t, body.Samples, 2, "prefix-filtered response should exclude samples with other prefixes")
			assert.Equal(t, []string{first.Name, second.Name}, testkit.DTONames(body.Samples), "samples should be sorted by name ascending")
		})
	})
}
