//go:build integration

package handlers_test

import (
	"net/http"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/internal/testsupport/integration"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/testkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNearbySamplesAPI(t *testing.T) {
	env := integration.NewAutotestEnv(t)
	factory := testkit.NewSamplesFactory(t, env.Deps)

	t.Run("GET /samples/nearby", func(t *testing.T) {
		t.Run("returns nearest samples within limit", func(t *testing.T) {
			nearest := factory.Create(contracts.SampleInputDTO{
				Name:      "new-york",
				Latitude:  40.7128,
				Longitude: -74.0060,
			})
			secondNearest := factory.Create(contracts.SampleInputDTO{
				Name:      "newark",
				Latitude:  40.7357,
				Longitude: -74.1724,
			})
			factory.Create(contracts.SampleInputDTO{
				Name:      "los-angeles",
				Latitude:  34.0522,
				Longitude: -118.2437,
			})

			var body testkit.SampleListResponse
			env.Expect.GET("/samples/nearby").
				WithQuery("lat", 40.7128).
				WithQuery("lng", -74.0060).
				WithQuery("limit", 2).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Decode(&body)

			require.Len(t, body.Samples, 2, "nearby response should respect requested limit")
			assert.Equal(t, []string{nearest.Name, secondNearest.Name}, testkit.DTONames(body.Samples), "nearby response should order samples by distance")
		})

		t.Run("validates query params", func(t *testing.T) {
			env.Expect.GET("/samples/nearby").
				WithQuery("lat", 91).
				WithQuery("lng", 21.0122).
				WithQuery("limit", 2).
				Expect().
				Status(http.StatusUnprocessableEntity)
		})
	})
}
