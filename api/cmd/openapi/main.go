package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/pkg/openapi"
	"github.com/rparaschak/mono-tmpl/api/wiring"
)

func main() {
	out := flag.String("out", "", "file path to write the OpenAPI document to")
	flag.Parse()

	spec, err := openapi.JSON(wiring.RouteRegistrar(modules.GlobalDependencies{}))
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
