package service

import (
	"testing"

	"github.com/DataInfosec/faceappapi-userservice/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	mock.Mock
}

// user details of the person signing to the system
var idStr = "5e1f799738794b700e14cdc9"
var userId, _ = primitive.ObjectIDFromHex(idStr)
var user = entity.User{Type: "admin",
	Firstname:       "elvis",
	Lastname:        "eche",
	Email:           "okaforechezona@gmail.com",
	Password:        "password",
	Confirmpassword: "password",
}

func TestCreate(t *testing.T) {
	var body entity.User
	var result entity.UserDB
	var err error

	testService := UserService()
	result, err = testService.Create(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	var body entity.UpdateUser
	var result entity.UpdateUser
	var err error

	testService := UserService()
	result, err = testService.Update(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
}

func TestInvalidType(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "",
		Firstname:       "elvis",
		Lastname:        "eche",
		Email:           "okaforechezona@gmail.com",
		Password:        "password",
		Confirmpassword: "password",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Type is required")
}

func TestInvalidEmail(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "admin",
		Firstname:       "elvis",
		Lastname:        "eche",
		Email:           "",
		Password:        "password",
		Confirmpassword: "password",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Valid Email is required")
}

func TestInvalidFirstname(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "admin",
		Firstname:       "",
		Lastname:        "eche",
		Email:           "okaforechezona@gmail.com",
		Password:        "password",
		Confirmpassword: "password",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "First Name is required")
}

func TestInvalidLastname(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "admin",
		Firstname:       "Elvis",
		Lastname:        "",
		Email:           "okaforechezona@gmail.com",
		Password:        "password",
		Confirmpassword: "password",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Last Name is required")
}

func TestInvalidPassword(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "admin",
		Firstname:       "Elvis",
		Lastname:        "eche",
		Email:           "okaforechezona@gmail.com",
		Password:        "",
		Confirmpassword: "password",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Password is required")
}

func TestInvalidConfirmpassword(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "admin",
		Firstname:       "Elvis",
		Lastname:        "eche",
		Email:           "okaforechezona@gmail.com",
		Password:        "password",
		Confirmpassword: "",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Confirm password is required")
}

func TestValidUserInput(t *testing.T) {
	var result entity.User
	var err error
	var body = entity.User{
		Type:            "admin",
		Firstname:       "Elvis",
		Lastname:        "eche",
		Email:           "okaforechezona@gmail.com",
		Password:        "password",
		Confirmpassword: "password",
	}
	result, err = CreateValidation(body)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFindUserByEmail(t *testing.T) {
	var result entity.UserDB
	var err error

	result, err = FindUserByEmail("elvis@gmail.com")

	assert.NotNil(t, result)
	assert.NotNil(t, err)
}

func TestFindUserByID(t *testing.T) {
	var result entity.UserDB
	var err error

	result, err = FindUserByID(userId)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
}

func TestFind(t *testing.T) {
	var result []entity.UserDB
	var err error

	testService := UserService()
	result, err = testService.Find()

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestFindById(t *testing.T) {
	var result entity.UserDB
	var err error

	testService := UserService()
	result, err = testService.FindById(idStr)

	assert.NotNil(t, result)
	assert.NotNil(t, err)
}
