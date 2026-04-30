package testkit

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/persistence"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
	"github.com/stretchr/testify/require"
)

type SamplesFactory struct {
	t      *testing.T
	ctx    context.Context
	deps   modules.GlobalDependencies
	prefix string
	seq    int
}

func NewSamplesFactory(t *testing.T, deps modules.GlobalDependencies) *SamplesFactory {
	t.Helper()

	return &SamplesFactory{
		t:      t,
		ctx:    context.Background(),
		deps:   deps,
		prefix: sanitizeName(t.Name()) + "-" + uuid.NewString()[:8],
	}
}

func (f *SamplesFactory) UniqueName(base string) string {
	f.t.Helper()

	f.seq++
	return fmt.Sprintf("%s-%s-%03d", sanitizeName(base), f.prefix, f.seq)
}

func (f *SamplesFactory) Create(input contracts.SampleInputDTO) persistence.Sample {
	f.t.Helper()

	input.Name = f.UniqueName(input.Name)
	sampleUseCases := &usecases.UseCase{GlobalDependencies: f.deps}
	sample, err := sampleUseCases.CreateSample(f.ctx, usecases.CreateSampleInput{
		Sample: usecases.SampleInput{
			Name:      input.Name,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
		},
	})
	require.NoError(f.t, err, "factory should create sample through usecase")
	require.NotNil(f.t, sample, "created sample should not be nil")

	return *sample
}

func (f *SamplesFactory) CreateMany(inputs ...contracts.SampleInputDTO) []persistence.Sample {
	f.t.Helper()

	samples := make([]persistence.Sample, len(inputs))
	for i, input := range inputs {
		samples[i] = f.Create(input)
	}
	return samples
}

func (f *SamplesFactory) CreateWarsaw() persistence.Sample {
	f.t.Helper()

	return f.Create(WarsawInput())
}

func (f *SamplesFactory) CreateKrakow() persistence.Sample {
	f.t.Helper()

	return f.Create(KrakowInput())
}

func (f *SamplesFactory) CreateWroclaw() persistence.Sample {
	f.t.Helper()

	return f.Create(WroclawInput())
}

func (f *SamplesFactory) CreateGdansk() persistence.Sample {
	f.t.Helper()

	return f.Create(GdanskInput())
}

func (f *SamplesFactory) SetCreatedAt(sample persistence.Sample, createdAt time.Time) {
	f.t.Helper()

	result := f.deps.DB.Model(&persistence.Sample{}).
		Where("id = ?", sample.Id).
		Update("created_at", createdAt)
	require.NoError(f.t, result.Error, "factory should update sample created_at")
	require.Equal(f.t, int64(1), result.RowsAffected, "factory should update exactly one sample")
}

func Names(samples []persistence.Sample) []string {
	names := make([]string, len(samples))
	for i, sample := range samples {
		names[i] = sample.Name
	}
	return names
}

var unsafeNameChars = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func sanitizeName(value string) string {
	sanitized := strings.Trim(unsafeNameChars.ReplaceAllString(strings.ToLower(value), "-"), "-")
	if sanitized == "" {
		return "sample"
	}
	return sanitized
}

func DTONames(samples []contracts.SampleDTO) []string {
	names := make([]string, len(samples))
	for i, sample := range samples {
		names[i] = sample.Name
	}
	return names
}
