//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/Khan/genqlient"
	_ "github.com/atombender/go-jsonschema"
	_ "github.com/golang-migrate/migrate/v4"
	_ "golang.org/x/tools/cmd/godoc"
)
