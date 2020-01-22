package jwt

import (
	"github.com/DataInfosec/faceappapi/middlewareservice/entity"
	"github.com/dgrijalva/jwt-go"
)

func DecodeJWT(tokenBearer string) (entity.User, error) {
	var tokenString = tokenBearer
	var user entity.User
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("elvisSecreyKey"), nil
	})
	if err != nil {
		return user, err
	}
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user = entity.User{
				Firstname: claims["firstname"].(string),
				Lastname:  claims["lastname"].(string),
				Email:     claims["email"].(string),
				// Authorized: claims["authorized"].(string),
				// Exp: claims["exp"].(string),
			}
			return user, nil
		} else {
			return user, err
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return user, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return user, err
		} else {
			return user, err
		}
	} else {
		return user, err
	}
}
