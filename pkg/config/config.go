//go:generate easyjson -lower_camel_case $GOFILE
package config

import "time"

// easyjson:json
type Config struct {
	Server *ServerConfig `json:"server"`
	Mongo  *MongoConfig  `json:"mongo"`
	OpenAI *OpenAIConfig `json:"openai"`
	JWT    *JwtConfig    `json:"jwt"`
}

type ServerConfig struct {
	CookieName        string        `json:"cookieName"`
	AppName           string        `json:"appName"`
	Port              string        `json:"port"`
	PprofPort         string        `json:"pprofPort"`
	Mode              string        `json:"mode"`
	AppVersion        string        `json:"appVersion"`
	Repository        string        `json:"repository"`
	CtxDefaultTimeout time.Duration `json:"ctxDefaultTimeout"`
	WriteTimeout      time.Duration `json:"writeTimeout"`
	ReadTimeout       time.Duration `json:"readTimeout"`
	SSL               bool          `json:"ssl"`
	CSRF              bool          `json:"csrf"`
	Debug             bool          `json:"debug"`
}

type MongoConfig struct {
	URI     string `json:"uri"`
	Name    string `json:"name"`
	Timeout int    `json:"timeout"`
}

type OpenAIConfig struct {
	Key string `json:"key"`
}

type JwtConfig struct {
	Secret   string `json:"secret"`
	Audience string `json:"audience"`
	Issuer   string `json:"issuer"`
}
