package httpserver

import (
	"sesi6-gin/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.HealthCheck)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.ReadUser)

	return router
}
