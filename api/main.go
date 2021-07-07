package main

import (
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Router"
	_ "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func main() {
	setup()
}

func setup() *gin.Engine {
	_, err := Database.GetESClient()

	if err != nil {
		return nil
	}

	r := gin.Default()
	Router.Setup(r)

	err = r.Run(":8080")

	if err != nil {
		return nil
	}
	return r
}
