package controller

import (
	"errors"
	"net/http"

	// "fmt"
	"github.com/DataInfosec/faceappapi-userservice/entity"
	"github.com/DataInfosec/faceappapi-userservice/service"
	errorm "github.com/DataInfosec/faceappapi-userservice/utils/shared/error"

	"github.com/gin-gonic/gin"
)

var repo service.UserServiceInterface = service.UserService()

// CreateUserEndpoint godoc
// @Summary Add a user testing
// @Description add by json user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.User true "User Data"
// @Success 200 {object} errorm.HTTPError200
// @Failure 400 {object} errorm.HTTPError
// @Failure 404 {object} errorm.HTTPError404
// @Failure 500 {object} errorm.HTTPError500
// @Security ApiKeyAuth
// @Router /user/api/v1/users/ [post]
func CreateUserEndpoint(ctx *gin.Context) {
	var user entity.User
	err_privilege := CheckPrivilege(ctx)
	if err_privilege != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err_privilege)
		return
	}
	ctx.ShouldBindJSON(&user)
	payload, err := repo.Create(user)
	if err != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err)
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
// @Success 200 {object} errorm.HTTPError200
// @Failure 400 {object} errorm.HTTPError
// @Failure 404 {object} errorm.HTTPError404
// @Failure 500 {object} errorm.HTTPError500
// @Security ApiKeyAuth
// @Router /user/api/v1/users/ [put]
func UpdateUserEndpoint(ctx *gin.Context) {
	var user entity.UpdateUser
	err_privilege := CheckPrivilege(ctx)
	if err_privilege != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err_privilege)
		return
	}
	ctx.ShouldBindJSON(&user)
	payload, err := repo.Update(user)
	if err != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var response = entity.ResponseUpdate{Data: payload, Message: "Update was successfully"}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// GetUsersEndpoint godoc
// @Summary get all the users
// @Description get all user by json
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} errorm.HTTPError200
// @Failure 400 {object} errorm.HTTPError
// @Failure 404 {object} errorm.HTTPError404
// @Failure 500 {object} errorm.HTTPError500
// @Security ApiKeyAuth
// @Router /user/api/v1/users/ [get]
func GetUsersEndpoint(ctx *gin.Context) {
	err_privilege := CheckPrivilege(ctx)
	if err_privilege != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err_privilege)
		return
	}
	payload, err := repo.Find()
	if err != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err)
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
// @Success 200 {object} errorm.HTTPError200
// @Failure 400 {object} errorm.HTTPError
// @Failure 404 {object} errorm.HTTPError404
// @Failure 500 {object} errorm.HTTPError500
// @Security ApiKeyAuth
// @Router /user/api/v1/users/{id} [get]
func FindUserById(ctx *gin.Context) {
	err_privilege := CheckPrivilege(ctx)
	if err_privilege != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err_privilege)
		return
	}
	id := ctx.Param("id")
	user, err := repo.FindById(id)
	if err != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err)
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
// @Success 200 {object} errorm.HTTPError200
// @Failure 400 {object} errorm.HTTPError
// @Failure 404 {object} errorm.HTTPError404
// @Failure 500 {object} errorm.HTTPError500
// @Security ApiKeyAuth
// @Router /user/api/v1/users/email [put]
func GetUserByEmail(ctx *gin.Context) {
	err_privilege := CheckPrivilege(ctx)
	if err_privilege != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err_privilege)
		return
	}
	var userEmail entity.UserEmail
	ctx.ShouldBindJSON(&userEmail)
	user, err := repo.FindByEmail(userEmail)
	if err != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateCompanyEndpoint godoc
// @Summary Update a user's image
// @Description Update user's image
// @Tags Users
// @Accept  json
// @Produce  json
// @Param company body entity.UpdateImage true "User Image"
// @Success 200 {object} errorm.HTTPError200
// @Failure 400 {object} errorm.HTTPError
// @Failure 404 {object} errorm.HTTPError404
// @Failure 500 {object} errorm.HTTPError500
// @Security ApiKeyAuth
// @Router /user/api/v1/users/imageupload [put]
func UpdateImage(ctx *gin.Context) {
	var imageUrl entity.UpdateImage
	userId := ctx.GetHeader("user")
	ctx.ShouldBindJSON(&imageUrl)
	user, err := repo.UpdateImage(imageUrl, userId)
	if err != nil {
		errorm.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

//checking whether the user have the privilege
// to perform a particular action
func CheckPrivilege(ctx *gin.Context) error {
	userType := ctx.GetHeader("type")
	if userType == "staff" {
		return errors.New("You don't permission to perform this action")
	}
	return nil
}
