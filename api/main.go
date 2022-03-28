package main

import (
	"vocabs-backend/api/database"
	"vocabs-backend/api/router"
)

func main() {
	database.InitConnection()
	defer database.CloseConnection()
	database.Migrate()
	router.RunApi()
}
