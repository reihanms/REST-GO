package app

import (
	"github.com/gin-gonic/gin"

	"gin-gorm/internal/handler"
)

func Run() {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.HelloBook)
	v1.GET("/books/:id", handler.BookDetail)
	v1.GET("/query", handler.BookQuery)
	v1.POST("/books", handler.SaveBook)

	router.Run()
}
