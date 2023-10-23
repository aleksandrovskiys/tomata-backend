package database

import "tomata-backend/interfaces"

var database = &RelationalDB{}

func GetDB() interfaces.Database {
	if !database.initialized {
		database.Init()
	}
	return database
}
