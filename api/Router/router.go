package Router

import (
	"MT2022_PROJ03/api/Controllers"
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
	r.POST("/books/:id", Controllers.UpdateBook)
	r.DELETE("/books/:id", Controllers.DeleteBook)
}
