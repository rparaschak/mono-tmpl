//go:build integration

package sample_test

import "github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"

type sampleListResponse struct {
	Samples []contracts.SampleDTO `json:"samples"`
}
