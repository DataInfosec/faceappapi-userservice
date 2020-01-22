package service

import (
	"context"

	"github.com/DataInfosec/faceappapi/authservice/entity"
	"github.com/DataInfosec/faceappapi/authservice/utils/hash"
	"github.com/DataInfosec/faceappapi/authservice/utils/jwt"
	"github.com/DataInfosec/faceappapi/authservice/utils/validator"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	// user "github.com/DataInfosec/faceappapi/authservice/models/user"
	"github.com/DataInfosec/faceappapi/authservice/utils/connection"
)

var collection *mongo.Collection = connection.Connection()

type AuthService interface {
	Authenticate(loginBody entity.Login) (entity.Payload, error)
}

type Service struct{}

func Authentication() AuthService {
	return &Service{}
}

func (s *Service) Authenticate(loginBody entity.Login) (entity.Payload, error) {
	var user entity.User
	var payload entity.Payload

	//validate if all the form are field filled with the correct input
	err := validator.Validate(loginBody)
	if err != nil {
		return payload, err
	}

	//authenticate the user with email and password
	errm := collection.FindOne(context.TODO(), bson.D{{"email", loginBody.Email}}).Decode(&user)
	if errm != nil {
		return payload, errm
	}
	// fmt.Println("result user details :: ", user)
	// ctx.JSON(http.StatusOK, gin.H{"data": user})

	// verify if the password from the user is same as the one in the db
	if errp := VerifyPassword(loginBody, user); errp != nil {
		return payload, errp
	}

	//Generate jwt token for the user
	data, errorToken := GenerateJwtToken(user)
	return data, errorToken
}

func VerifyPassword(loginBody entity.Login, user entity.User) error {
	return hash.CompareHashValue(loginBody.Password, user.Password)
}

func GenerateJwtToken(user entity.User) (entity.Payload, error) {
	var payload entity.Payload
	token, errorToken := jwt.GenerateJWT(user)
	if errorToken != nil {
		return payload, errorToken
	} else {
        payload = entity.Payload{
          Firstname: user.Firstname,
          Lastname: user.Lastname,
          ID: user.ID,
          Email: user.Email,
		  Type: user.Type,
		  Token: token,
		} 
		return payload, nil
	}
}
