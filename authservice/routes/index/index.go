package index

import (
	"github.com/DataInfosec/faceappapi/authservice/routes/auth"
	"github.com/gin-gonic/gin"
)

func Index(router *gin.RouterGroup) {
	auth.Auth(router)
}
