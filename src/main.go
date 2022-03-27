package main

import (
	"vocabs-backend/src/vocabs"

	"github.com/gin-gonic/gin"
)

func main() {
	runApi()
}

func runApi() {
	router := gin.Default()
	router.GET("/vocabs", vocabs.GetAll)
	router.GET("/random", vocabs.GetRandom)
	router.Run("localhost:8080")
}
