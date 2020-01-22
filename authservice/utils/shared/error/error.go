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
