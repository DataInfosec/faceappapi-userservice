package authentications

import (
	"net/http"
	"strings"
	// "fmt"

	"github.com/DataInfosec/faceappapi/proto/middleware"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header["Authorization"]
		if len(authorization) == 0 || len(strings.Fields(authorization[0])) < 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorised to perform this action"})
			ctx.Abort()
			return
		}
		token := strings.Fields(authorization[0])[1]

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorised to perform this action"})
			ctx.Abort()
			return
		}

		conn, err := grpc.Dial(":50052", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client := middleware.NewJwtServiceClient(conn)

		req := &middleware.JwtRequest{Token: token}
		if _, err := client.JwtService(ctx, req); err == nil {
			// fmt.Println("result from response :: ", response)
		} else {
			// fmt.Println("result from response tok :: ", resp)
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorised to perform this action"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
