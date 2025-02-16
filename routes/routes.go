// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ratnesh-maurya/api.ratn.tech/config"
	"github.com/ratnesh-maurya/api.ratn.tech/controllers"
)

// SetupRoutes initializes and returns the Gin router with defined routes

func RatnTechRoutes(router *gin.Engine) {

	blogview := config.GetRepoCollection("blogview")

	router.POST("/blogview/:slug", controllers.IncrementViews(blogview))
	router.GET("/blogview/:slug", controllers.GetViews(blogview))


	
}