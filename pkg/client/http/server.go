package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kkrajkumar1198/blog-grpc/docs"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	setPostRoutes(apiGroup)

	return router
}

func Start(port string) {
	router := setRouter()
	router.Run(":" + port)
}
