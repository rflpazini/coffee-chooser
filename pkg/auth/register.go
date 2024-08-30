package auth

func Register(register func(...interface{}) error) error {
	return register(
		makeCreateSessionToken,
		makeValidateSessionToken,
	)
}
