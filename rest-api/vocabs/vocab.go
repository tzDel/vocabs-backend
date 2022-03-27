package vocabs

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VocabEntity struct {
	UserID     int    `json:"userId"`
	FirstTerm  string `json:"firstTerm"`
	SecondTerm string `json:"secondTerm"`
	Languages  string `json:"languages"`
}

var allVocabsOfUser = []VocabEntity{
	{UserID: 1, FirstTerm: "egg", SecondTerm: "Ei", Languages: "en-de"},
	{UserID: 1, FirstTerm: "kitchen", SecondTerm: "Küche", Languages: "en-de"},
	{UserID: 1, FirstTerm: "language", SecondTerm: "Sprache", Languages: "en-de"},
	{UserID: 1, FirstTerm: "entity", SecondTerm: "Entität", Languages: "en-de"},
	{UserID: 1, FirstTerm: "random", SecondTerm: "zufällig", Languages: "en-de"},
	{UserID: 1, FirstTerm: "example", SecondTerm: "Beispiel", Languages: "en-de"},
	{UserID: 1, FirstTerm: "test", SecondTerm: "Test", Languages: "en-de"},
	{UserID: 1, FirstTerm: "cat", SecondTerm: "Katze", Languages: "en-de"},
	{UserID: 1, FirstTerm: "", SecondTerm: "", Languages: "en-de"}}

func GetAllVocabs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, allVocabsOfUser)
}

func GetRandomVocabs(context *gin.Context) {
	outputCount := getOutputCount()
	randomVocabs := make([]VocabEntity, outputCount)
	uniqueRandomInts := make([]int, outputCount)
	uniqueRandomInts[0] = randomIntBetween(1, outputCount)

	for i := 0; i < outputCount; i++ {
		randomIndex := getUniqueRandom(1, len(allVocabsOfUser), uniqueRandomInts)
		randomVocabs[i] = allVocabsOfUser[randomIndex]
		uniqueRandomInts[i] = randomIndex
	}
	context.IndentedJSON(http.StatusOK, randomVocabs)
}

func getUniqueRandom(min int, max int, uniqueRandoms []int) int {
	random := randomIntBetween(min, max)
	for !isUnique(random, uniqueRandoms) {
		random = randomIntBetween(min, max)
	}
	return random
}

func isUnique(random int, randoms []int) bool {
	for _, r := range randoms {
		if r == random {
			return false
		}
	}
	return true
}

func getVocabByUserId(id int) (*VocabEntity, error) {
	for i, vocab := range allVocabsOfUser {
		if vocab.UserID == id {
			return &allVocabsOfUser[i], nil
		}
	}
	return nil, errors.New("")
}

func getOutputCount() int {
	totalCountOfVocabs := len(allVocabsOfUser)
	if totalCountOfVocabs < 5 {
		return totalCountOfVocabs
	}
	return 5
}

func randomIntBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}
