package coffeeTypes

func Register(register func(...interface{}) error) error {
	return register(
		makeGetAllCoffeeVarieties,
	)
}
