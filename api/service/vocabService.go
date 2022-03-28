package service

import (
	"errors"
	"math/rand"
	"vocabs-backend/api/database"
	"vocabs-backend/api/model"

	"gorm.io/gorm"
)

var AllVocabsOfUser = []model.Vocab{
	{UserID: 1, FirstTerm: "egg", SecondTerm: "Ei", Languages: "en-de"},
	{UserID: 1, FirstTerm: "kitchen", SecondTerm: "Küche", Languages: "en-de"},
	{UserID: 1, FirstTerm: "language", SecondTerm: "Sprache", Languages: "en-de"},
	{UserID: 1, FirstTerm: "entity", SecondTerm: "Entität", Languages: "en-de"},
	{UserID: 1, FirstTerm: "random", SecondTerm: "zufällig", Languages: "en-de"},
	{UserID: 1, FirstTerm: "example", SecondTerm: "Beispiel", Languages: "en-de"},
	{UserID: 1, FirstTerm: "test", SecondTerm: "Test", Languages: "en-de"},
	{UserID: 1, FirstTerm: "cat", SecondTerm: "Katze", Languages: "en-de"},
	{UserID: 1, FirstTerm: "", SecondTerm: "", Languages: "en-de"}}

func AddVocab(toBeCreated model.Vocab) error {
	return database.AddVocab(toBeCreated)
}

func GetVocabsBy(userID string) (*[]model.Vocab, error) {
	return database.GetAllByUserId(userID)
}

func GetRandomVocabs(db *gorm.DB) *[]model.Vocab {
	outputCount := getOutputCount()
	randomVocabs := make([]model.Vocab, outputCount)
	uniqueRandomInts := make([]int, outputCount)
	uniqueRandomInts[0] = randomIntBetween(1, outputCount)

	for i := 0; i < outputCount; i++ {
		randomIndex := getUniqueRandom(1, len(AllVocabsOfUser), uniqueRandomInts)
		randomVocabs[i] = AllVocabsOfUser[randomIndex]
		uniqueRandomInts[i] = randomIndex
	}
	return &randomVocabs
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

func getVocabByUserId(id int) (*model.Vocab, error) {
	for i, vocab := range AllVocabsOfUser {
		if vocab.UserID == id {
			return &AllVocabsOfUser[i], nil
		}
	}
	return nil, errors.New("")
}

func getOutputCount() int {
	totalCountOfVocabs := len(AllVocabsOfUser)
	if totalCountOfVocabs < 5 {
		return totalCountOfVocabs
	}
	return 5
}

func randomIntBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}
