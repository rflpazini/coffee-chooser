package main

import (
	"os"

	"coffee-choose/pkg/config"
	"github.com/cristalhq/aconfig"
)

const (
	DEV  = "dev"
	PROD = "production"
)

func NewConfig() (*config.Config, error) {
	var cfg config.Config
	var cfgFile string

	env := os.Getenv("ENV")
	if env == PROD {
		cfgFile = PROD + ".json"
	} else {
		cfgFile = DEV + ".json"

	}

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipEnv:   true,
		SkipFlags: true,
		Files:     []string{"config/" + cfgFile},
	})
	if err := loader.Load(); err != nil {
		return nil, err
	}

	mongo := os.Getenv("MONGODB_URL")
	if mongo != "" {
		cfg.Mongo.URI = mongo
	}

	return &cfg, nil
}
