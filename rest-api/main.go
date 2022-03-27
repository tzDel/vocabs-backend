package main

import "github.com/gin-gonic/gin"


var vocabEntities = []VocabEntity{
	{ID: 1, FirstTerm: "egg", SecondTerm: "Ei", Languages: "en-de"},
	{ID: 2, FirstTerm: "kitchen", SecondTerm: "Küche", Languages: "en-de"},
	{ID: 3, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"},
	{ID: 5, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"},
	{ID: 6, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"},
	{ID: 7, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"},
	{ID: 8, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"},
	{ID: 9, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"},
	{ID: 10, FirstTerm: "Search", SecondTerm: "Suche", Languages: "en-de"}}

func main() {
	runApi()
}

func runApi() {
	router := gin.Default()
	router.GET("/vocabs", getAllVocabs)
	router.GET("/random", getRandomVocabs)
	router.Run("localhost:8080")
}