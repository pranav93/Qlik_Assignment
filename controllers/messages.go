package controllers

import (
	"log"
	"net/http"
	"strconv"

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
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"messageID": messageID}})
}

// GetMessage GetMessage
func GetMessage(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}

	message, err := models.GetMessage(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": message}})
}

// DeleteMessage DeleteMessage
func DeleteMessage(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}

	err = models.DeleteMessage(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"deleted": true}})
}

// GetMessageList GetMessageList
func GetMessageList(c *gin.Context) {
	message, err := models.GetMessageList()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"message_list": message}})
}
