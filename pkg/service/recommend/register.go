package recommend

func Register(register func(...interface{}) error) error {
	return register(
		makeSaveUserPreferences,
	)
}
