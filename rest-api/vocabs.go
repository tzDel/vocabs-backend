package main

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VocabEntity struct {
	UserID     int
	ID         int    `json:"id"`
	FirstTerm  string `json:"firstTerm"`
	SecondTerm string `json:"secondTerm"`
	Languages  string `json:"languages"`
	Score      int    `json:"score"`
}

func getAllVocabs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, vocabEntities)
}

func getRandomVocabs(context *gin.Context) {
	outputCount := getOutputCount()
	randomVocabs := make([]VocabEntity, outputCount)
	uniqueRandomInts := make([]int, outputCount)
	uniqueRandomInts[0] = randomIntBetween(1, outputCount)

	for i := 0; i < outputCount; i++ {
		uniqueRandomInts[i] = getUniqueRandom(1, len(vocabEntities), uniqueRandomInts)
		vocabEntity, error := getVocabById(uniqueRandomInts[i])
		if error == nil {
			randomVocabs[i] = *vocabEntity
		}
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

func getVocabById(id int) (*VocabEntity, error) {
	for i, vocab := range vocabEntities {
		if vocab.ID == id {
			return &vocabEntities[i], nil
		}
	}
	return nil, errors.New("")
}

func getOutputCount() int {
	totalCountOfVocabs := len(vocabEntities)
	if totalCountOfVocabs < 5 {
		return totalCountOfVocabs
	}
	return 5
}

func randomIntBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}