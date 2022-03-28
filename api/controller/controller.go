package controller

import (
	"vocabs-backend/api/service"

	"github.com/gin-gonic/gin"
)

func GetVocabsByUserID(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	vocabs, err := service.GetVocabsBy(userID)
	if err != nil {
		c.AbortWithError(404, err)
	} else {
		c.JSON(200, vocabs)
	}
}

func GetRandomVocabsForUser(c *gin.Context) {
	// userID := c.Params.ByName("user_id")
	// service.GetRandomVocabs(userID)
}
