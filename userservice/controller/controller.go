package controller

import (
	"net/http"
    // "fmt"
	"github.com/DataInfosec/faceappapi/userservice/entity"
	"github.com/DataInfosec/faceappapi/userservice/service"
	"github.com/DataInfosec/faceappapi/userservice/utils/shared/error"

	"github.com/gin-gonic/gin"
)

var repo service.UserServiceInterface = service.UserService()

// CreateUserEndpoint godoc
// @Summary Add a user
// @Description add by json user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.User true "User Data"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/users/ [post]
func CreateUserEndpoint(ctx *gin.Context) {
	var user entity.User
	payload, err := repo.Create(ctx, user)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var response = entity.Response{Data: payload, Message: "New user was created successfully"}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// UpdateUserEndpoint godoc
// @Summary Update a user
// @Description Update by json user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.UpdateUser true "User Data"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/users/ [put]
func UpdateUserEndpoint(ctx *gin.Context) {
	var user entity.UpdateUser
	payload, err := repo.Update(ctx, user)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var response = entity.ResponseUpdate{Data: payload, Message: "Update was successfully"}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// GetUsersEndpoint godoc
// @Summary get all users
// @Description get all user by json
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/users/ [get]
func GetUsersEndpoint(ctx *gin.Context) {
	payload, err := repo.Find()
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": payload})
}

// GetUserById godoc
// @Summary Show a user
// @Description get user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [get]
func FindUserById(ctx *gin.Context) {
    user, err := repo.FindById(ctx)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUserByEmailEndpoint godoc
// @Summary get user by email
// @Description getting the user details by supplying email
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.UserEmail true "User Data"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/users/email [put]
func GetUserByEmail(ctx *gin.Context) {
	user, err := repo.FindByEmail(ctx)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
