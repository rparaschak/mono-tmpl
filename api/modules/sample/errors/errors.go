package errors

import (
	"net/http"

	"github.com/rparaschak/mono-tmpl/api/pkg/apperror"
)

var ErrSampleNotFound = apperror.New(http.StatusNotFound, "sample_not_found", "sample not found")
