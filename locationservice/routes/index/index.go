package index

import (
	"github.com/DataInfosec/faceappapi/locationservice/routes/location"
	"github.com/gin-gonic/gin"
)

func Index(router *gin.RouterGroup) {
	location.Location(router)
}
