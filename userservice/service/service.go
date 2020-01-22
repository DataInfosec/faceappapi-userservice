package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/DataInfosec/faceappapi/userservice/entity"
	"github.com/DataInfosec/faceappapi/userservice/utils/validator"
	"github.com/DataInfosec/faceappapi/userservice/utils/hash"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// user "github.com/DataInfosec/faceappapi/authservice/models/user"
	"github.com/DataInfosec/faceappapi/userservice/utils/connection"
	"github.com/gin-gonic/gin"
)

var collection *mongo.Collection = connection.Connection()

type UserServiceInterface interface {
	Create(ctx *gin.Context, body entity.User) (entity.UserDB, error)
	Update(ctx *gin.Context, body entity.UpdateUser) (entity.UpdateUser, error)
	Find() ([]entity.UserDB, error)
	FindById(ctx *gin.Context) (entity.UserDB, error)
	FindByEmail(ctx *gin.Context) (entity.UserDB, error)
}

type Service struct{}

func UserService() UserServiceInterface {
	return &Service{}
}

func (s *Service) Create(ctx *gin.Context, user entity.User) (entity.UserDB, error) {
	var payload entity.UserDB

	body, err_validator := CreateValidation(ctx, user)
	if err_validator != nil {
		return payload, err_validator
	}
	
	// check if email already exist
	_, errUserExist := FindUserByEmail(body.Email);
	if errUserExist == nil {
		return payload, errors.New("User already exist")
	}
	//password encryption 
	body.Password = string(hash.Hash(body.Password))
	body.Confirmpassword = ""
	
	//authenticate the user with email and password
	res_u, err_u := collection.InsertOne(context.TODO(), body)
	if err_u != nil {
		return payload, err_u
	}
	 var id string
	 if oid, ok := res_u.InsertedID.(primitive.ObjectID); ok {
		 id = oid.Hex()
	}
	payload = entity.UserDB{
	ID: id,
	Firstname: body.Firstname,
	Lastname: body.Lastname,
	Email: body.Email,
	Type: body.Type,
	}
    
	return payload, nil
}

func CreateValidation(ctx *gin.Context, user entity.User)(entity.User, error) {
	e := ctx.ShouldBindJSON(&user)
	if e != nil {
		return user, e
	}
	//validate if all the form are field filled with the correct input
	err := validator.Validate(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *Service) Update(ctx *gin.Context, user entity.UpdateUser) (entity.UpdateUser, error) {
	body, err_validator := UpdateValidation(ctx, user)
	if err_validator != nil {
		return user, err_validator
	}
	
	userId, errInvalidId := primitive.ObjectIDFromHex(body.ID)
	if errInvalidId != nil {
		return user, errors.New("Invalid _id supplied")
	}

	_, errIdNotFound := FindUserByID(userId)
	if errIdNotFound != nil {
		return user, errors.New("User does not exist")
	}
	filter := bson.D{{"_id", userId}}
	payload := bson.D{
		{"$set", bson.D{
		{"firstname", body.Firstname},
		{"lastname", body.Lastname},
		{"type", body.Type},
		}},
		}
	_, err := collection.UpdateOne(context.TODO(), filter, payload)
	if err != nil {
		return body, err
	}
    
	return body, nil
}

func UpdateValidation(ctx *gin.Context, user entity.UpdateUser) (entity.UpdateUser, error) {
	e := ctx.ShouldBindJSON(&user)
	if e != nil {
		return user, e
	}
	if err := validator.ValidateUpdate(user); err != nil {
		return user, err
	}
	return user, nil
}

func FindUserByEmail(email string) (entity.UserDB, error) {
	var user entity.UserDB
	err := collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func FindUserByID(id primitive.ObjectID) (entity.UserDB, error) {
	var user entity.UserDB
	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *Service) Find() ([]entity.UserDB, error) {
	var users []entity.UserDB
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil { return  users, err}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var user entity.UserDB
		err := cur.Decode(&user)
		if err != nil { return  users, err }
		users = append(users,  user)

		// To get the raw bson bytes use cursor.Current
		// raw := cur.Current
		// fmt.Println("current raw :: ", raw)
	}
	if err := cur.Err(); err != nil {
		return  users, err
	}
	return users, nil
}

func (s *Service) FindById(ctx *gin.Context) (entity.UserDB, error)  {
	id := ctx.Param("id")
	var user entity.UserDB
	userId, errInvalidId := primitive.ObjectIDFromHex(id)
	if errInvalidId != nil {
		return user, errors.New("Invalid _id supplied")
	}
	user, errIdNotFound := FindUserByID(userId)
	if errIdNotFound != nil {
		return user, errors.New("User does not exist")
	}
    return user, nil
}

func (s *Service) FindByEmail(ctx *gin.Context) (entity.UserDB, error)  {
	var user entity.UserDB
	var userEmail entity.UserEmail
	userEmail, errInvalidEmail := ValidateEmail(ctx)
	if errInvalidEmail != nil {
		return user, errInvalidEmail
	}
	user, errIdNotFound := FindUserByEmail(userEmail.Email)
	if errIdNotFound != nil {
		return user, errors.New("User does not exist")
	}
    return user, nil
}

func ValidateEmail(ctx *gin.Context) (entity.UserEmail, error) {
	var user entity.UserEmail
	e := ctx.ShouldBindJSON(&user)
	if e != nil {
		return user, e
	}
	if err := validator.ValidateEmail(user); err != nil {
		return user, err
	}
	return user, nil
}