package router

func Register(register func(...interface{}) error) error {
	return register(newEchoRouter)
}
