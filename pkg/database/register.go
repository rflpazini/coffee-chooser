package database

func Register(register func(...interface{}) error) error {
	return register(
		makeMongoClient,
		makeMongoDB,
		makeMongoPing,
		makeMongoDisconnect,
		makeBrewingCollection,
	)
}
