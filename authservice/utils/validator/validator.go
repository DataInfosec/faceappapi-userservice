package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/DataInfosec/faceappapi/authservice/entity"
)

var (
	ErrNoRow           = errors.New("no rows in result set")
	ErrPasswordInvalid = errors.New("Password is required")
	ErrEmailInvalid    = errors.New("Valid Email is required")
	ErrNilInvalid      = errors.New("Field is required")
)

// Validation for new user fields
func Validate(a entity.Login) error {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	switch {
	case len(strings.TrimSpace(a.Password)) == 0:
		return ErrPasswordInvalid
	case len(strings.TrimSpace(a.Email)) == 0 || !rxEmail.MatchString(a.Email):
		return ErrEmailInvalid
	default:
		return nil
	}
}
