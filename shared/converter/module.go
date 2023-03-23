// Package converted contains function which helps to work with Pandoc file converter
package converter

import (
	"go.uber.org/fx"
)

// NewModule returns new module for uber fx
//
//nolint:gocritic
func NewModule() *File {
	return &File{}
}

// Module for uber fx
var Module = fx.Options(
	fx.Provide(
		NewModule,
	),
)
