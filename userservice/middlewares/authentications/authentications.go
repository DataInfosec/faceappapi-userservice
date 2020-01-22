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

	// func AuthMiddleware(ctx *gin.Context) {
	// 	token := ctx.Request.Header["token"]
	// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		client := middleware.NewJwtServiceClient(conn)

	// 		req := &middleware.JwtRequest{Token: token}
	// 		if response, err := client.JwtService(ctx, req); err == nil {
	// 			fmt.Println("token details :: ", req)
	// 			ctx.JSON(http.StatusOK, gin.H{
	// 				"result details": fmt.Sprint(response),
	// 			})
	// 		} else {
	// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		}

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 	log.SetOutput(os.Stdout) // logs go to Stderr by default
	// 	log.Println(r.Method, r.URL)
	// 	h.ServeHTTP(w, r) // call ServeHTTP on the original handler
	// })
}
