package config

func (v *Config) Register(register func(...interface{}) error) error {
	return register(
		func() *ServerConfig { return v.Server },
		func() *MongoConfig { return v.Mongo },
	)
}
