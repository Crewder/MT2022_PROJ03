package Controllers

import (
	"bytes"
	"encoding/json"
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Helpers"
	"github.com/MT2022_PROJ03/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func CreateBook(c *gin.Context) {
	var input Models.Book
	client, _ := Database.GetESClient()

	err := c.ShouldBindJSON(&input)
	if err != nil {
		return
	}

	book, _ := hydrateBook(input)

	res, err := client.Index(
		"books",
		strings.NewReader(string(book)),
	)

	var r map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, r)
}

func GetBooks(c *gin.Context) {
	client, _ := Database.GetESClient()

	var buf bytes.Buffer
	const field = "_index"
	const value = "books"
	var query = Helpers.SearchQuery(field, value)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := client.Search(
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	var r map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, r)
}

func GetBook(c *gin.Context) {

	var id = c.Params.ByName("id")
	client, _ := Database.GetESClient()
	res, err := client.Get(
		"books",
		id,
	)
	var r map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, r)
}

func DeleteBook(c *gin.Context) {

}

func UpdateBook(c *gin.Context) {

}

func hydrateBook(input Models.Book) ([]byte, error) {
	model := &Models.Book{
		Author:   input.Author,
		Abstract: input.Abstract,
		Title:    input.Title,
	}
	return json.Marshal(model)
}
