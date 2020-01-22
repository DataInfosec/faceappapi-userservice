package service

import (
	"testing"

	"github.com/DataInfosec/faceappapi/authservice/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	mock.Mock
}

// user details of the person signing to the system
var userId, _ = primitive.ObjectIDFromHex("5e1f799738794b700e14cdc9")
var user = entity.User{Firstname: "Elvis", Lastname: "Okafor",
		Email: "okaforechezona1992@gmail.com", Type: "admin", ID: userId,
		Password: "$2y$12$P1MRu5psINXqpXkQfw2DK.95GSHIA/WVByTKT/nRWn0VrWKz6oBqG", Token: "dffjjjdj"}

// user login credentials of the person signing to the system
var loginBody = entity.Login{
		Email:    "okaforechezona1992@gmail.com",
		Password: "weekrjrjf"}

func (mock *Repository) Authenticate(loginBody entity.Login) (entity.User, error) {
	args := mock.Called()
	result := args.Get(1)
	return result.(entity.User), args.Error(1)
}

func TestAuthenticate(t *testing.T) {
	mockRepo := new(Repository)
	
	mockRepo.On("Authenticate").Return(user, nil)

	testService := Authentication()
	result, _ := testService.Authenticate(loginBody)

	assert.Equal(t, "Elvis", result.Firstname)
	assert.Equal(t, "Okafor", result.Lastname)
	assert.Equal(t, "okaforechezona1992@gmail.com", result.Email)
}

func TestEmptyPasswordField(t *testing.T) {
	testService := Authentication()
	loginBody := entity.Login{Email: "elvis@gmail.com", Password: ""}
	_, err := testService.Authenticate(loginBody)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Password is required")
}

func TestEmptyEmailField(t *testing.T) {
	testService := Authentication()
	loginBody := entity.Login{Email: "", Password: "233dkffk"}
	_, err := testService.Authenticate(loginBody)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Valid Email is required")
}

func TestVerifyPasswordWithWrongPassword(t *testing.T) {
	err := VerifyPassword(loginBody, user)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "crypto/bcrypt: hashedPassword is not the hash of the given password")
}

func TestVerifyPasswordWithCorrectPassword(t *testing.T) {
	body := entity.Login{
		Email:    "okaforechezona1992@gmail.com",
		Password: "elvisgo"}
	err := VerifyPassword(body, user)
	assert.Nil(t, err)
}

func TestGenerateJwtToken(t *testing.T) {
	data, errorToken := GenerateJwtToken(user)
	assert.NotNil(t, data);
	assert.Nil(t, errorToken)
}