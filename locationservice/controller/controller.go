package controller

import (
	"net/http"
    // "fmt"
	"github.com/DataInfosec/faceappapi/locationservice/entity"
	"github.com/DataInfosec/faceappapi/locationservice/service"
	"github.com/DataInfosec/faceappapi/locationservice/utils/shared/error"

	"github.com/gin-gonic/gin"
)

var repo service.LocationServiceInterface = service.LocationService()

// CreateLocationEndpoint godoc
// @Summary Add a location
// @Description add by json location
// @Tags Locations
// @Accept  json
// @Produce  json
// @Param location body entity.Location true "Location Data"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/locations/ [post]
func CreateLocationEndpoint(ctx *gin.Context) {
	var location entity.Location
	payload, err := repo.Create(ctx, location)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var response = entity.Response{Data: payload, Message: "New location was created successfully"}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// GetLocationsEndpoint godoc
// @Summary get all locations
// @Description get all location by json
// @Tags Locations
// @Accept  json
// @Produce  json
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/locations/ [get]
func GetLocationsEndpoint(ctx *gin.Context) {
	payload, err := repo.Find()
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": payload})
}

// GetLocationById godoc
// @Summary Show a location
// @Description get location by ID
// @Tags Locations
// @Accept  json
// @Produce  json
// @Param id path string true "Location ID"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/locations/{id} [get]
func FindLocationById(ctx *gin.Context) {
    location, err := repo.FindById(ctx)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": location})
}

// GetLocationByUser godoc
// @Summary Show locations by a user
// @Description get location by User
// @Tags Locations
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} error.HTTPError200
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError404
// @Failure 500 {object} error.HTTPError500
// @Security ApiKeyAuth
// @Router /api/v1/locations/{id}/user [get]
func GetLocationByUser(ctx *gin.Context) {
    location, err := repo.FindByUser(ctx)
	if err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": location})
}