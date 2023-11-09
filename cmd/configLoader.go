package main

import (
	"os"

	"coffee-choose/pkg/config"
	"github.com/cristalhq/aconfig"
)

const (
	DEV  = "dev"
	STAG = "staging"
	PROD = "production"
)

func NewConfig() (*config.Config, error) {
	var cfg config.Config

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipEnv:   true,
		SkipFlags: true,
		Files:     []string{getConfigFile()},
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

func getConfigFile() string {
	path := "config/"
	fileExt := ".json"
	var fileName string

	env := os.Getenv("ENV")
	if env == PROD {
		fileName = PROD + fileExt
	} else if env == STAG {
		fileName = STAG + fileExt
	} else {
		fileName = DEV + fileExt
	}

	return path + fileName
}
