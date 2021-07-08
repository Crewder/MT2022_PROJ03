package Controllers

import (
	"bytes"
	"encoding/json"
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Helpers"
	"github.com/MT2022_PROJ03/Models"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

var input Models.Book
var buf bytes.Buffer

func CreateBook(c *gin.Context) {
	client := Database.GetESClient()

	err := c.ShouldBindJSON(&input)
	if err != nil {
		return
	}

	book, _ := Models.HydrateBook(input)

	response, err := client.Index(
		"books",
		strings.NewReader(string(book)),
	)

	Helpers.HandlingResponse(c, response, err)
}

func GetBooks(c *gin.Context) {
	client := Database.GetESClient()

	const field = "_index"
	const value = "books"
	var query = Helpers.SearchQuery(field, value)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	response, err := client.Search(
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	Helpers.HandlingResponse(c, response, err)
}

func GetBook(c *gin.Context) {
	var id = c.Params.ByName("id")
	client := Database.GetESClient()

	response, err := client.Get(
		"books",
		id,
	)

	Helpers.HandlingResponse(c, response, err)
}

func DeleteBook(c *gin.Context) {
	var id = c.Params.ByName("id")
	client := Database.GetESClient()

	response, err := client.Delete(
		"books",
		id,
	)

	Helpers.HandlingResponse(c, response, err)
}

func UpdateBook(c *gin.Context) {
	client := Database.GetESClient()

	var id = c.Params.ByName("id")

	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	var query = Helpers.UpdateQuery(input)

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	response, err := client.Update(
		"books",
		id,
		&buf,
		client.Update.WithPretty(),
	)

	Helpers.HandlingResponse(c, response, err)
}
