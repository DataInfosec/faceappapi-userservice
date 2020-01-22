package user

import (
	"github.com/DataInfosec/faceappapi/userservice/controller"
	"github.com/DataInfosec/faceappapi/userservice/middlewares/authentications"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	users := router.Group("/users")
	users.Use(authentications.AuthMiddleware())
	{
		users.POST("/", controller.CreateUserEndpoint)
		users.PUT("/", controller.UpdateUserEndpoint)
		users.GET("/", controller.GetUsersEndpoint)
		users.GET("/:id", controller.FindUserById)
		users.PUT("/email", controller.GetUserByEmail)
	}
}
