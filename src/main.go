package main

import c "vocabs-backend/src/commons"

func main() {
	c.InitDatabaseConnection()
	defer c.CloseDatabase()
	c.MigrateDatabase()
	c.RunApi()
}
