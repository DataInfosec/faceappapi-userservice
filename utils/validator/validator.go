package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/DataInfosec/faceappapi-userservice/entity"
)

//  example
var (
	ErrNoRow                      = errors.New("no rows in result set")
	ErrIDInvalid                  = errors.New("_id is required")
	ErrFirstNameInvalid           = errors.New("First Name is required")
	ErrLastNameInvalid            = errors.New("Last Name is required")
	ErrTypeInvalid                = errors.New("Type is required")
	ErrTypeIncorrect              = errors.New("Type is invalid")
	ErrPasswordInvalid            = errors.New("Password is required")
	ErrConfirmpasswordInvalid     = errors.New("Confirm password is required")
	ErrPasswordCombinationInvalid = errors.New("Password and Confirm password should be the same")
	ErrEmailInvalid               = errors.New("Valid Email is required")
	ErrOfficeIdInvalid            = errors.New("Valid Office Id is required")
	ErrSecretAnswerInvalid        = errors.New("Valid Secret Answer is required")
	ErrCompanyNotExist            = errors.New("Company does not exist")
	ErrCompanyInvalid             = errors.New("Company is required")
	ErrNilInvalid                 = errors.New("Field is required")
	ErrImageInvalid               = errors.New("Image is required")
	ADMIN                         = "admin"
	STAFF                         = "staff"
	SUPERADMIN                    = "super admin"
)

// Validation for new user fields
func Validate(a entity.User) error {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	switch {
	case len(strings.TrimSpace(a.Company)) == 0:
		return ErrCompanyInvalid
	case len(strings.TrimSpace(a.Firstname)) == 0:
		return ErrFirstNameInvalid
	case len(strings.TrimSpace(a.Lastname)) == 0:
		return ErrLastNameInvalid
	case len(strings.TrimSpace(a.Type)) == 0:
		return ErrTypeInvalid
	case strings.TrimSpace(a.Type) != ADMIN && strings.TrimSpace(a.Type) != SUPERADMIN && strings.TrimSpace(a.Type) != STAFF:
		return ErrTypeIncorrect
	case len(strings.TrimSpace(a.OfficeId)) == 0:
		return ErrOfficeIdInvalid
	case len(strings.TrimSpace(a.SecretAnswer)) == 0:
		return ErrSecretAnswerInvalid
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
	case len(strings.TrimSpace(a.Company)) == 0:
		return ErrCompanyInvalid
	case len(strings.TrimSpace(a.ID)) == 0:
		return ErrIDInvalid
	case len(strings.TrimSpace(a.Lastname)) == 0:
		return ErrLastNameInvalid
	case len(strings.TrimSpace(a.Type)) == 0:
		return ErrTypeInvalid
	case strings.TrimSpace(a.Type) != ADMIN && strings.TrimSpace(a.Type) != SUPERADMIN && strings.TrimSpace(a.Type) != STAFF:
		return ErrTypeIncorrect
	case len(strings.TrimSpace(a.OfficeId)) == 0:
		return ErrOfficeIdInvalid
	case len(strings.TrimSpace(a.SecretAnswer)) == 0:
		return ErrSecretAnswerInvalid
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

func ValidateImage(image entity.UpdateImage) error {
	switch {
	case len(strings.TrimSpace(image.Image)) == 0:
		return ErrImageInvalid
	default:
		return nil
	}
}
