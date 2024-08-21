package recommendation

func Register(register func(...interface{}) error) error {
	return register(makeRecommendationService)
}
