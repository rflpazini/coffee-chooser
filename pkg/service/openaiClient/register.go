package openaiClient

func Register(register func(...interface{}) error) error {
	return register(makeOpenAIService)
}
