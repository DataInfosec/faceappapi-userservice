package router

import (
	// "errors"
	"fmt"
	"os"
	// "net/http"

	"github.com/gin-gonic/gin"

	// "github.com/DataInfosec/faceappapi/shareservice/httputil"
	_ "github.com/DataInfosec/faceappapi-userservice/docs"
	"github.com/DataInfosec/faceappapi-userservice/routes/index"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	var rootUrl string
	if os.Getenv("LOCAL") != "" {
		rootUrl = "user/api/v1"
	} else {
		rootUrl = "/api/v1"
	}
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group(rootUrl)
	{
		index.Index(v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("0.0.0.0:8883")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authorization testing")
		if len(c.GetHeader("Authorization")) == 0 {
			// httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
