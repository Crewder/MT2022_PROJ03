package main

import (
	"github.com/MT2022_PROJ03/Database"
	"github.com/MT2022_PROJ03/Router"
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
