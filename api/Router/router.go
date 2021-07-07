package Router

import (
	"github.com/MT2022_PROJ03/Controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	bookRoute(r)
}

func bookRoute(r *gin.Engine) {
	r.GET("/books", Controllers.GetBooks)
	r.GET("/books/:id", Controllers.GetBook)
	r.POST("/books", Controllers.CreateBook)
	r.GET("/search/:value", Controllers.Search)
}
