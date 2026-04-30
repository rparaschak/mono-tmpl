//go:build integration

package handlers_test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/rparaschak/mono-tmpl/api/internal/testsupport/integration"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/testkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteSampleAPI(t *testing.T) {
	env := integration.NewAutotestEnv(t)
	factory := testkit.NewSamplesFactory(t, env.Deps)

	t.Run("DELETE /samples/{sampleId}", func(t *testing.T) {
		t.Run("deletes sample", func(t *testing.T) {
			sample := factory.CreateWarsaw()

			env.Expect.DELETE("/samples/{sampleId}", sample.Id.String()).
				Expect().
				Status(http.StatusNoContent)

			var count int64
			require.NoError(t, env.Deps.DB.Model(&persistence.Sample{}).Where("id = ?", sample.Id).Count(&count).Error, "deleted sample count query should succeed")
			assert.Equal(t, int64(0), count, "deleted sample should no longer exist")
		})

		t.Run("returns 404 for missing sample", func(t *testing.T) {
			env.Expect.DELETE("/samples/{sampleId}", uuid.NewString()).
				Expect().
				Status(http.StatusNotFound)
		})
	})
}
