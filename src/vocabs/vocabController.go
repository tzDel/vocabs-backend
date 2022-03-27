package vocabs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, AllVocabsOfUser)
}

func GetRandom(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, GetRandomVocabs)
}
