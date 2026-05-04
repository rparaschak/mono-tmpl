package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/rparaschak/mono-tmpl/api/internal/app"
	"github.com/rparaschak/mono-tmpl/api/internal/dependencies"
	"github.com/rparaschak/mono-tmpl/api/pkg/openapi"
)

func main() {
	out := flag.String("out", "", "file path to write the OpenAPI document to")
	flag.Parse()

	modules := app.NewModules(dependencies.Dependencies{})
	spec, err := openapi.JSON(modules.RegisterHTTP)
	if err != nil {
		slog.Error("failed to generate OpenAPI document", "error", err)
		os.Exit(1)
	}

	if *out != "" {
		if err := os.WriteFile(*out, spec, 0644); err != nil {
			slog.Error("failed to write OpenAPI document", "path", *out, "error", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println(string(spec))
}
