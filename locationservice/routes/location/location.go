package location

import (
	"github.com/DataInfosec/faceappapi/locationservice/controller"
	"github.com/DataInfosec/faceappapi/locationservice/middlewares/authentications"
	"github.com/gin-gonic/gin"
)

func Location(router *gin.RouterGroup) {
	locations := router.Group("/locations")
	locations.Use(authentications.AuthMiddleware())
	{
		locations.POST("/", controller.CreateLocationEndpoint)
		locations.GET("/", controller.GetLocationsEndpoint)
		locations.GET("/:id", controller.FindLocationById)
		locations.GET("/:id/user", controller.GetLocationByUser)
	}
}
