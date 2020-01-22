package router

import (
	// "errors"
	// "fmt"
	// "net/http"

	"github.com/gin-gonic/gin"

	// "github.com/DataInfosec/faceappapi/shareservice/httputil"
	_ "github.com/DataInfosec/faceappapi/locationservice/docs"
	"github.com/DataInfosec/faceappapi/locationservice/routes/index"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-contrib/cors"
)

func Router() {
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")
	{
		index.Index(v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("0.0.0.0:8884")
}