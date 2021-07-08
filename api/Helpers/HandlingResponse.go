package Helpers

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func isASuccessResponse(value int) bool {
	return 200 <= value && value < 300
}

func HandlingResponse(c *gin.Context, response *esapi.Response, err error) {

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else if !isASuccessResponse(response.StatusCode) {
		c.JSON(response.StatusCode, response)
	} else {
		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err == nil {
			c.JSON(response.StatusCode, result)
		}
	}
}
