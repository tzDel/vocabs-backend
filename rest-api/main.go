package main

import (
	v "vocabs-backend/rest-api/vocabs"

	"github.com/gin-gonic/gin"
)

func main() {
	runApi()
}

func runApi() {
	router := gin.Default()
	router.GET("/vocabs", v.GetAllVocabs)
	router.GET("/random", v.GetRandomVocabs)
	router.Run("localhost:8080")
}
