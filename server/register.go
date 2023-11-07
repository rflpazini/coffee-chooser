package server

func Register(register func(...interface{}) error) error {
	return register(makeServer, makeStart)
}
