package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setPostRoutes(router *gin.RouterGroup) {
	router.POST("/post", postPost)
	router.GET("/post/:id", getPost)
	router.DELETE("/post/:id", deletePost)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
