package routing

import (
	"context"
	"net/http"
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
	register(grp, http.MethodGet, path, summary, handler)
}

func POST[I, O any](grp huma.API, path, summary string, handler func(context.Context, *I) (*O, error), modifiers ...func(*huma.Operation)) {
	register(grp, http.MethodPost, path, summary, handler, modifiers...)
}

func PUT[I, O any](grp huma.API, path, summary string, handler func(context.Context, *I) (*O, error), modifiers ...func(*huma.Operation)) {
	register(grp, http.MethodPut, path, summary, handler, modifiers...)
}

func DELETE[I, O any](grp huma.API, path, summary string, handler func(context.Context, *I) (*O, error), modifiers ...func(*huma.Operation)) {
	register(grp, http.MethodDelete, path, summary, handler, modifiers...)
}

func WithErrors(statusCodes ...int) func(*huma.Operation) {
	return func(o *huma.Operation) {
		o.Errors = append(o.Errors, statusCodes...)
	}
}

func WithDefaultStatus(statusCode int) func(*huma.Operation) {
	return func(o *huma.Operation) {
		o.DefaultStatus = statusCode
	}
}

func register[I, O any](grp huma.API, method, path, summary string, handler func(context.Context, *I) (*O, error), modifiers ...func(*huma.Operation)) {
	operation := huma.Operation{
		OperationID: toOperationID(summary),
		Method:      method,
		Path:        path,
		Summary:     summary,
	}
	for _, modifier := range modifiers {
		modifier(&operation)
	}

	huma.Register(grp, operation, handler)
}
