package Controllers

import (
	"bytes"
	"encoding/json"
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Helpers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func Search(c *gin.Context) {
	var buf bytes.Buffer

	//var value = c.Params.ByName("value")

	client, _ := Database.GetESClient()

	const value = "lire"
	var query = Helpers.MultiSearchQuery(value)

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	log.Println(&buf)
	var test io.Reader
	res, err := client.Msearch(
		test,
		client.Msearch.WithPretty(),
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
