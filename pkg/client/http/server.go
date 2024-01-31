package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kkrajkumar1198/blog-grpc/grpc-test/docs"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	setPostRoutes(apiGroup) // Update to setPostRoutes instead of setPersonRoutes

	return router
}

func Start(port string) {
	router := setRouter()
	router.Run(":" + port)
}
