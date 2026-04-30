//go:build integration

package sample_test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/rparaschak/mono-tmpl/api/internal/testsupport/integration"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/testkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateSampleAPI(t *testing.T) {
	env := integration.NewAutotestEnv(t)
	factory := testkit.NewSamplesFactory(t, env.Deps)

	t.Run("PUT /samples/{sampleId}", func(t *testing.T) {
		t.Run("updates sample", func(t *testing.T) {
			sample := factory.CreateWarsaw()
			input := testkit.WithName(testkit.PoznanInput(), factory.UniqueName("updated-sample"))

			var body contracts.SampleDTO
			env.Expect.PUT("/samples/{sampleId}", sample.Id.String()).
				WithJSON(input).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Decode(&body)

			assert.Equal(t, sample.Id.String(), body.Id, "response should keep the updated sample id")
			assert.Equal(t, input.Name, body.Name, "response should include updated sample name")
			assert.Equal(t, input.Latitude, body.Latitude, "response should include updated sample latitude")
			assert.Equal(t, input.Longitude, body.Longitude, "response should include updated sample longitude")

			var persisted persistence.Sample
			require.NoError(t, env.Deps.DB.First(&persisted, "id = ?", sample.Id).Error, "updated sample should still exist")
			assert.Equal(t, input.Name, persisted.Name, "persisted sample should have updated name")
			assert.Equal(t, input.Latitude, persisted.Geolocation.Lat, "persisted sample should have updated latitude")
			assert.Equal(t, input.Longitude, persisted.Geolocation.Lng, "persisted sample should have updated longitude")
		})

		t.Run("returns 404 for missing sample", func(t *testing.T) {
			env.Expect.PUT("/samples/{sampleId}", uuid.NewString()).
				WithJSON(testkit.PoznanInput()).
				Expect().
				Status(http.StatusNotFound)
		})
	})
}
