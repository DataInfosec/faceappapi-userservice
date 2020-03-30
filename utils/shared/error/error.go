package error

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// ErrNoRow example
	ErrNoRow = errors.New("no rows in result set")
)

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPError200 struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Item created successfully"`
}

type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Server Error"`
}

type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found"`
}

type HTTPError401 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Unauthorized user"`
}