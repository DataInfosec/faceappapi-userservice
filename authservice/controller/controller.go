package controller

import (
	"fmt"
	"net/http"

	"github.com/DataInfosec/faceappapi/authservice/entity"
	"github.com/DataInfosec/faceappapi/authservice/service"
	"github.com/DataInfosec/faceappapi/authservice/utils/shared/error"

	"github.com/gin-gonic/gin"
)

var auth service.AuthService = service.Authentication()

// CreateAuthEndpoint godoc
// @Summary authenticate user
// @Description authenticate user by providing username and password
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param user body entity.Login true "User Data"
// @Success 200 {object} entity.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/auth/login [post]
func CreateAuthEndpoint(ctx *gin.Context) {
	var loginBody entity.Login
	fmt.Println("application")
	e := ctx.ShouldBindJSON(&loginBody)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}
	user, err := auth.Authenticate(loginBody)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// VerifyByEmail godoc
// @Summary verify grpc
// @Description Verify verify grpc
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param user body entity.UserEmail true "User Data"
// @Success 200 {object} entity.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/auth/verify [post]
// func VerifyByEmail(ctx *gin.Context) {
// 	var user userm.UserEmail
// 	err := ctx.ShouldBindJSON(&user)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}
// 	client := proto.NewUserServiceClient(conn)

// 	req := &proto.Request{Email: user.Email}
// 	if response, err := client.UserDetails(ctx, req); err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"result details": fmt.Sprint(response),
// 		})
// 	} else {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}
// }

// VerifyByEmail godoc
// @Summary testing grpc
// @Description testing grpc
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/auth/test [get]
// func TestByEmail(ctx *gin.Context) {
// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}
// 	client := proto.NewUserServiceClient(conn)

// 	req := &proto.Request{Email: "elvis@lendsqr.com"}
// 	if response, err := client.UserDetails(ctx, req); err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"result details docker 2334": fmt.Sprint(response),
// 		})
// 	} else {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}
// }
