package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gin-gorm/internal/dto"
)

func HelloBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "How to become famous",
		"subtitle": "Think yourself",
	})
}

func BookDetail(c *gin.Context) {
	// tangkap id yang dikirim
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"idSend": id})
}

func BookQuery(c *gin.Context) {
	// tangkap id yang dikirim
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{"title": title})
}

func SaveBook(c *gin.Context) {
	// struct diatas ditampung pada sebuah variable
	var bookInput dto.BookInput

	// validator
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, msg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
