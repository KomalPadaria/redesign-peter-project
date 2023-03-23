package cache

import (
	"go.uber.org/fx"
)

// NewModule returns new module for uber fx
//
//nolint:gocritic
func NewModule() Client {
	return New(1, 1)
}

// Module for uber fx
var Module = fx.Options(fx.Provide(NewModule))
