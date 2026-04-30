package errors

import "github.com/danielgtaylor/huma/v2"

var ErrSampleNotFound = huma.Error404NotFound("sample not found")
