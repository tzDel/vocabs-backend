package controller

import (
	"vocabs-backend/api/model"
	"vocabs-backend/api/service"

	"github.com/gin-gonic/gin"
)

func AddVocab(c *gin.Context) {
	var requestBody model.Vocab
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(400, err)
	}
	if err := service.AddVocab(requestBody); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(201, requestBody)
	}
}

func DeleteVocabByTerm(c *gin.Context) {
	term := c.Params.ByName("term")
	if err := service.DeleteVocabBy(term); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, gin.H{"data": true})
	}
}

func GetVocabsByUserID(c *gin.Context) {
	// dummy implementation
	// userID := c.Params.ByName("user_id")
	if vocabs, err := service.GetVocabsBy("1"); err != nil {
		c.AbortWithError(404, err)
	} else {
		c.JSON(200, vocabs)
	}
}

func GetRandomVocabsForUser(c *gin.Context) {
	// dummy implementation
	if vocabs, err := service.GetRandomVocabsFor(); err != nil {
		c.AbortWithError(404, err)
	} else {
		c.JSON(200, vocabs)
	}
}
