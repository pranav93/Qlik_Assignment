package setup

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/pranav93/Qlik_Assignment/controllers"
)

// Server Creates a gin instance and returns int
func Server() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	config.AllowOrigins = []string{allowedOrigin}

	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	r.GET("/message/:id/", controllers.GetMessage)
	r.GET("/message/:id/palindrome/", controllers.CheckPalindrome)
	r.GET("/messages/", controllers.GetMessageList)
	r.POST("/messages/", controllers.CreateMessage)
	r.DELETE("/message/:id/", controllers.DeleteMessage)
	return r
}
