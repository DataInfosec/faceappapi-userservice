package validator

import (
	"errors"
	"strings"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/DataInfosec/faceappapi/locationservice/entity"
)

//  example
var (
	ErrNoRow                      = errors.New("no rows in result set")
	ErrIDInvalid           = errors.New("Location's owner is required")
	ErrUserIdInvalid           = errors.New("_id is required")
	ErrLatitudeInvalid           = errors.New("Latitude Name is required")
	ErrLongitudeInvalid            = errors.New("Longitude Name is required")
)

// Validation for new location fields
func Validate(a entity.Location) error {
	id := a.User.Hex()
	// var id string
	//  if oid, ok := a.User.(primitive.ObjectID); ok {
	// 	 id = oid.Hex()
	// }
	switch {
	case len(strings.TrimSpace(a.Latitude)) == 0:
		return ErrLatitudeInvalid
	case len(strings.TrimSpace(a.Longitude)) == 0:
		return ErrLongitudeInvalid
	case len(strings.TrimSpace(id)) == 0:
		return ErrUserIdInvalid
	default:
		return nil
	}
}