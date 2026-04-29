package routing

import (
	"context"
	"regexp"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)

func toOperationID(summary string) string {
	lower := strings.ToLower(summary)
	id := nonAlphanumeric.ReplaceAllString(lower, "-")
	return strings.Trim(id, "-")
}

func GET[I, O any](grp huma.API, path, summary string, handler func(context.Context, *I) (*O, error)) {
	huma.Register(grp, huma.Operation{
		OperationID: toOperationID(summary),
		Method:      "GET",
		Path:        path,
		Summary:     summary,
	}, handler)
}
