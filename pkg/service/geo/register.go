package geo

func Register(register func(...interface{}) error) error {
	return register(
		makeGeoIPService,
	)
}
