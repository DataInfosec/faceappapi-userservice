package auth

import (
	"github.com/DataInfosec/faceappapi/authservice/controller"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	route := router.Group("/auth")
	{
		route.POST("/login", controller.CreateAuthEndpoint)
		// route.POST("/verify", auth.VerifyByEmail)
		// route.GET("/test", auth.TestByEmail)
	}
}
