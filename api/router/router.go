package router

import (
	"vocabs-backend/api/controller"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	router := gin.Default()
	router.GET("/myvocabs", controller.GetVocabsByUserID)
	router.GET("/myvocabs/random", controller.GetRandomVocabsForUser)
	router.Run("localhost:8080")
}
