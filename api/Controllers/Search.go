package Controllers

import (
	"encoding/json"
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Helpers"
	"github.com/MT2022_PROJ03/Models"
	"github.com/gin-gonic/gin"
	"log"
)

func Search(c *gin.Context) {
	client := Database.GetESClient()

	var value = c.Params.ByName("value")

	fields := Models.GetFields()
	var query = Helpers.MultiSearchQuery(fields, value)

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	response, err := client.Search(
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	Helpers.HandlingResponse(c, response, err)
}
