package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranav93/Qlik_Assignment/controllers"
)

// Server Creates a gin instance and returns int
func Server() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	r.GET("/message/:id/", controllers.GetMessage)
	r.GET("/messages/", controllers.GetMessageList)
	r.POST("/messages/", controllers.CreateMessage)
	r.DELETE("/message/:id/", controllers.DeleteMessage)
	return r
}
