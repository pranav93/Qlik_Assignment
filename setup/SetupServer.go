package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pranav93/Qlik_Assignment/controllers"
)

// Server Creates a gin instance and returns it
func Server() *gin.Engine {
	r := gin.Default()

	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	r.GET("/api/message/:id/", controllers.GetMessage)
	r.GET("/api/message/:id/palindrome/", controllers.CheckPalindrome)
	r.GET("/api/messages/", controllers.GetMessageList)
	r.POST("/api/messages/", controllers.CreateMessage)
	r.DELETE("/api/message/:id/", controllers.DeleteMessage)
	return r
}
