package coffee

func Register(register func(...interface{}) error) error {
	return register(
		makeSaveBrewingMethod,
		makeGetBrewingMethod,
		makeDeleteBrewingMethod,
	)
}
