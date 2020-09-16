package controllers

import (
	"log"
	"net/http"

	"github.com/pranav93/Qlik_Assignment/models"

	"github.com/gin-gonic/gin"
)

// CreateMessageInput CreateMessageInput
type CreateMessageInput struct {
	Message string `json:"message" binding:"required"`
}

// CreateMessage CreateMessage
func CreateMessage(c *gin.Context) {
	var input CreateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)
	messageID, err := models.CreateMessage(input.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"messageID": messageID}})
}
