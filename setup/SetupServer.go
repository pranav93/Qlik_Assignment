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
	// r.GET("/cart/:id/", controllers.GetCart)
	r.POST("/messages/", controllers.CreateMessage)
	// r.PATCH("/cart/save/:id/", controllers.UpdateCart)
	// r.DELETE("/cart/:id/", controllers.DeleteCart)
	// r.POST("/cart/checkout/:id/", controllers.CartCheckOut)
	return r
}
