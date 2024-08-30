package recommendations

func Register(register func(...interface{}) error) error {
	return register(
		makeSaveRecommendation,
	)
}
