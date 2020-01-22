package index

import (
	"github.com/DataInfosec/faceappapi/userservice/routes/user"
	"github.com/gin-gonic/gin"
)

func Index(router *gin.RouterGroup) {
	user.User(router)
}
