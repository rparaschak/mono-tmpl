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

func TestCreateSampleAPI(t *testing.T) {
	env := integration.NewAutotestEnv(t)
	factory := testkit.NewSamplesFactory(t, env.Deps)

	t.Run("POST /samples", func(t *testing.T) {
		t.Run("creates sample", func(t *testing.T) {
			input := testkit.WithName(testkit.WarsawInput(), factory.UniqueName("create-sample"))

			var body contracts.SampleDTO
			env.Expect.POST("/samples").
				WithJSON(input).
				Expect().
				Status(http.StatusCreated).
				JSON().
				Object().
				Decode(&body)

			parsedID, err := uuid.Parse(body.Id)
			require.NoError(t, err, "response should include a valid sample UUID")
			assert.Equal(t, input.Name, body.Name, "response should include created sample name")
			assert.Equal(t, input.Latitude, body.Latitude, "response should include created sample latitude")
			assert.Equal(t, input.Longitude, body.Longitude, "response should include created sample longitude")
			assert.False(t, body.CreatedAt.IsZero(), "response should include created timestamp")

			var persisted persistence.Sample
			require.NoError(t, env.Deps.DB.First(&persisted, "id = ?", parsedID).Error, "created sample should be persisted")
			assert.Equal(t, input.Name, persisted.Name, "persisted sample should have requested name")
			assert.Equal(t, input.Latitude, persisted.Geolocation.Lat, "persisted sample should have requested latitude")
			assert.Equal(t, input.Longitude, persisted.Geolocation.Lng, "persisted sample should have requested longitude")
		})

		t.Run("rejects invalid body", func(t *testing.T) {
			input := contracts.SampleInputDTO{
				Name:      "",
				Latitude:  91,
				Longitude: 21.0122,
			}

			env.Expect.POST("/samples").
				WithJSON(input).
				Expect().
				Status(http.StatusUnprocessableEntity)
		})
	})
}
