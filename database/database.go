package database

import "tomata-backend/interfaces"

var database = &InMemoryDatabase{}

func GetDB() interfaces.Database {
	return database
}
