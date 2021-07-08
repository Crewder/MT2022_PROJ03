package main

import (
	"MT2022_PROJ03/api/Database"
	"MT2022_PROJ03/api/Router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	setup()
}

func setup() {
	Database.Init()
	r := gin.Default()
	Router.Setup(r)

	err := r.Run(":8080")
	log.Println(err)
}
