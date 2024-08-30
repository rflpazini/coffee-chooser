package utils

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

	opKey := os.Getenv("OPENAI_KEY")
	if opKey != "" {
		cfg.OpenAI.Key = opKey
	}

	version := os.Getenv("APP_VERSION")
	if version != "" {
		cfg.Server.AppVersion = version
	}

	jwt := os.Getenv("JWT_KEY")
	if jwt != "" {
		cfg.Server.JwtSecretKey = jwt
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
