package router

import (
	"vocabs-backend/api/controller"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	router := gin.Default()
	router.POST("/vocabs", controller.AddVocab)
	router.DELETE("/myvocabs/:term", controller.DeleteVocabByTerm)
	router.GET("/myvocabs", controller.GetVocabsByUserID)
	router.GET("/myvocabs/random", controller.GetRandomVocabsForUser)
	router.Run("localhost:8080")
}
