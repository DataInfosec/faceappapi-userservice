package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/DataInfosec/faceappapi/userservice/entity"
)

//  example
var (
	ErrNoRow                      = errors.New("no rows in result set")
	ErrIDInvalid           = errors.New("_id is required")
	ErrFirstNameInvalid           = errors.New("First Name is required")
	ErrLastNameInvalid            = errors.New("Last Name is required")
	ErrTypeInvalid                = errors.New("Type is required")
	ErrPasswordInvalid            = errors.New("Password is required")
	ErrConfirmpasswordInvalid     = errors.New("Confirm password is required")
	ErrPasswordCombinationInvalid = errors.New("Password and Confirm password should be the same")
	ErrEmailInvalid               = errors.New("Valid Email is required")
	ErrNilInvalid                 = errors.New("Field is required")
)

// Validation for new user fields
func Validate(a entity.User) error {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	switch {
	case len(strings.TrimSpace(a.Firstname)) == 0:
		return ErrFirstNameInvalid
	case len(strings.TrimSpace(a.Lastname)) == 0:
		return ErrLastNameInvalid
	case len(strings.TrimSpace(a.Type)) == 0:
		return ErrTypeInvalid
	case len(strings.TrimSpace(a.Password)) == 0:
		return ErrPasswordInvalid
	case len(strings.TrimSpace(a.Confirmpassword)) == 0:
		return ErrConfirmpasswordInvalid
	case strings.TrimSpace(a.Confirmpassword) != strings.TrimSpace(a.Password):
		return ErrPasswordCombinationInvalid
	case len(strings.TrimSpace(a.Email)) == 0 || !rxEmail.MatchString(a.Email):
		return ErrEmailInvalid
	default:
		return nil
	}
}

// Validation for update fields
func ValidateUpdate(a entity.UpdateUser) error {
	// var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	switch {
	case len(strings.TrimSpace(a.ID)) == 0:
		return ErrIDInvalid
	case len(strings.TrimSpace(a.Lastname)) == 0:
		return ErrLastNameInvalid
	case len(strings.TrimSpace(a.Type)) == 0:
		return ErrTypeInvalid
	// case len(strings.TrimSpace(a.Email)) == 0 || !rxEmail.MatchString(a.Email):
	// 	return ErrEmailInvalid
	default:
		return nil
	}
}

func ValidateEmail(a entity.UserEmail) error {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	switch {
	case len(strings.TrimSpace(a.Email)) == 0 || !rxEmail.MatchString(a.Email):
		return ErrEmailInvalid
	default:
		return nil
	}
}
