package main

import (
	"coffee-choose/pkg/config"
	"github.com/cristalhq/aconfig"
)

func NewConfig() (*config.Config, error) {
	var cfg config.Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipEnv:   true,
		SkipFlags: true,
		Files:     []string{"config/dev.json"},
	})
	if err := loader.Load(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
