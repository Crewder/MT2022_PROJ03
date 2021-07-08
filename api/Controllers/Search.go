package Controllers

import (
	"bytes"
	"encoding/json"
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Helpers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Search(c *gin.Context) {
	var buf bytes.Buffer
	var value = c.Params.ByName("value")

	client, _ := Database.GetESClient()

	var fields []string

	//todo passé les element du models
	fieldResult := append(fields, "abstract", "title")

	var query = Helpers.MultiSearchQuery(fieldResult, value)

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
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
