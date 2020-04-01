package service

import (
	"context"
	"errors"

	// "fmt"

	"github.com/DataInfosec/faceappapi-userservice/entity"
	"github.com/DataInfosec/faceappapi-userservice/utils/validator"
	"github.com/DataInfosec/faceappapi/proto/finduserbyid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	// user "github.com/DataInfosec/faceappapi/authservice/models/user"
	"github.com/DataInfosec/faceappapi-userservice/utils/connection"
)

var collection *mongo.Collection = connection.Connection()

type UserServiceInterface interface {
	Create(body entity.User) (entity.UserDB, error)
	Update(body entity.UpdateUser) (entity.UpdateUser, error)
	Find() ([]entity.UserDB, error)
	FindById(id string) (entity.UserDB, error)
	UpdateImage(image entity.UpdateImage, userId string) (string, error)
	FindByEmail(userEmail entity.UserEmail) (entity.UserDB, error)
}

type Service struct{}

func UserService() UserServiceInterface {
	return &Service{}
}

func (s *Service) Create(user entity.User) (entity.UserDB, error) {
	var payload entity.UserDB
	body, err_validator := CreateValidation(user)
	if err_validator != nil {
		return payload, err_validator
	}

	// check if email already exist
	_, errUserExist := FindUserByEmail(body.Email)
	if errUserExist == nil {
		return payload, errors.New("User already exist")
	}
	//password encryption
	// body.Password = string(hash.Hash(body.Password))
	// body.Confirmpassword = ""

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
		ID:        id,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Email:     body.Email,
		Type:      body.Type,
		Company:   body.Company,
		OfficeId:  body.OfficeId,
	}

	return payload, nil
}

func CreateValidation(user entity.User) (entity.User, error) {
	//validate if all the form are field filled with the correct input
	err := validator.Validate(user)
	if err != nil {
		return user, err
	}
	_, companyError := FindCompanyById(user.Company)
	if companyError != nil {
		return user, errors.New("Company does not exist")
	}
	return user, nil
}

func (s *Service) Update(user entity.UpdateUser) (entity.UpdateUser, error) {
	body, err_validator := UpdateValidation(user)
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

func (s *Service) UpdateImage(image entity.UpdateImage, userId string) (string, error) {
	imageResponse := ""
	_, err_validator := UpdateImageValidation(image)
	if err_validator != nil {
		return imageResponse, err_validator
	}

	userObj, errInvalidId := primitive.ObjectIDFromHex(userId)
	if errInvalidId != nil {
		return imageResponse, errors.New("Invalid _id supplied")
	}

	_, errIdNotFound := FindUserByID(userObj)
	if errIdNotFound != nil {
		return imageResponse, errors.New("User does not exist")
	}
	filter := bson.D{{"_id", userObj}}
	payload := bson.D{
		{"$set", bson.D{
			{"image", image.Image},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, payload)
	if err != nil {
		return imageResponse, err
	}

	return "Image upload was successful", nil
}

func UpdateValidation(user entity.UpdateUser) (entity.UpdateUser, error) {
	if err := validator.ValidateUpdate(user); err != nil {
		return user, err
	}
	_, companyError := FindCompanyById(user.Company)
	if companyError != nil {
		return user, errors.New("Company does not exist")
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
	if err != nil {
		return users, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var user entity.UserDB
		err := cur.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (s *Service) FindById(id string) (entity.UserDB, error) {
	var user entity.UserDB
	userId, errInvalidId := primitive.ObjectIDFromHex(id)
	if errInvalidId != nil {
		return user, errors.New("Invalid id supplied")
	}
	user, errIdNotFound := FindUserByID(userId)
	if errIdNotFound != nil {
		return user, errors.New("User does not exist")
	}
	return user, nil
}

func (s *Service) FindByEmail(userEmail entity.UserEmail) (entity.UserDB, error) {
	var user entity.UserDB
	userEmail, errInvalidEmail := ValidateEmail(userEmail)
	if errInvalidEmail != nil {
		return user, errInvalidEmail
	}
	user, errIdNotFound := FindUserByEmail(userEmail.Email)
	if errIdNotFound != nil {
		return user, errors.New("User does not exist")
	}
	return user, nil
}

func ValidateEmail(user entity.UserEmail) (entity.UserEmail, error) {
	if err := validator.ValidateEmail(user); err != nil {
		return user, err
	}
	return user, nil
}

func UpdateImageValidation(image entity.UpdateImage) (string, error) {
	imageUrl := ""
	if err := validator.ValidateImage(image); err != nil {
		return imageUrl, err
	}
	return imageUrl, nil
}

func FindCompanyById(id string) (*finduserbyid.CompanyResponse, error) {
	conn, err := grpc.Dial(":50050", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := finduserbyid.NewSharedServicesClient(conn)

	req := &finduserbyid.CompanyRequest{Id: id}
	ctx := context.TODO()
	if res, err := client.CompanyService(ctx, req); err == nil {
		return res, nil
	} else {
		return res, err
	}
}
