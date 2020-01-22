package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/DataInfosec/faceappapi/locationservice/entity"
	"github.com/DataInfosec/faceappapi/locationservice/utils/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/DataInfosec/faceappapi/locationservice/utils/connection"
	"github.com/gin-gonic/gin"
	"github.com/DataInfosec/faceappapi/proto/finduserbyid"
	"google.golang.org/grpc"
)

var collection *mongo.Collection = connection.Connection()

type LocationServiceInterface interface {
	Create(ctx *gin.Context, body entity.Location) (entity.LocationDB, error)
	Find() ([]entity.LocationDB, error)
	FindById(ctx *gin.Context) (entity.LocationDB, error)
	FindByUser(ctx *gin.Context) ([]entity.LocationDB, error)
}

type Service struct{}

func LocationService() LocationServiceInterface {
	return &Service{}
}

func (s *Service) Create(ctx *gin.Context, location entity.Location) (entity.LocationDB, error) {
	var payload entity.LocationDB

	body, err_validator := CreateValidation(ctx, location)
	if err_validator != nil {
		return payload, err_validator
	}
	
	//authenticate the location with email and password
	res_u, err_u := collection.InsertOne(context.TODO(), body)
	if err_u != nil {
		return payload, err_u
	}
	 var id string
	 if oid, ok := res_u.InsertedID.(primitive.ObjectID); ok {
		 id = oid.Hex()
	}
	payload = entity.LocationDB{
	ID: id,
	Latitude: body.Latitude,
	Longitude: body.Longitude,
	User: body.User,
	}
    
	return payload, nil
}

func CreateValidation(ctx *gin.Context, location entity.Location)(entity.Location, error) {
	e := ctx.ShouldBindJSON(&location)
	if e != nil {
		return location, e
	}
	//validate if all the form are field filled with the correct input
	err := validator.Validate(location)
	if err != nil {
		return location, err
	}
	return location, nil
}

func FindLocationByID(id primitive.ObjectID) (entity.LocationDB, error) {
	var location entity.LocationDB
	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&location)
	if err != nil {
		return location, err
	}
	return location, nil
}

func (s *Service) Find() ([]entity.LocationDB, error) {
	var locations []entity.LocationDB
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil { return  locations, err}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var location entity.LocationDB
		err := cur.Decode(&location)
		if err != nil { return  locations, err }
		locations = append(locations,  location)

		// To get the raw bson bytes use cursor.Current
		// raw := cur.Current
		// fmt.Println("current raw :: ", raw)
	}
	if err := cur.Err(); err != nil {
		return  locations, err
	}
	return locations, nil
}

func (s *Service) FindById(ctx *gin.Context) (entity.LocationDB, error)  {
	id := ctx.Param("id")
	var location entity.LocationDB
	locationId, errInvalidId := primitive.ObjectIDFromHex(id)
	if errInvalidId != nil {
		return location, errors.New("Invalid _id supplied")
	}
	location, errIdNotFound := FindLocationByID(locationId)
	if errIdNotFound != nil {
		return location, errors.New("Location does not exist")
	}
    return location, nil
}

func (s *Service) FindByUser(ctx *gin.Context) ([]entity.LocationDB, error)  {
	id := ctx.Param("id")
	var locations []entity.LocationDB
	userId, errInvalidId := primitive.ObjectIDFromHex(id)

	userResponse, errUser := FindUserById(id) 
	if errUser != nil {
		return locations, errors.New("Location's owner does not exist")
	}else {
		fmt.Println("user Response :: ", userResponse)
	}
	if errInvalidId != nil {
		return locations, errors.New("Invalid _id supplied")
	}
	cur, err := collection.Find(context.Background(), bson.D{{"user", userId}})
	if err != nil { return  locations, err}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var location entity.LocationDB
		err := cur.Decode(&location)
		if err != nil { return  locations, err }
		locations = append(locations,  location)
	}
	if err := cur.Err(); err != nil {
		return  locations, err
	}
	return locations, nil
}

func FindUserById(id string) (*finduserbyid.UserResponse, error){
	conn, err := grpc.Dial(":50050", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client := finduserbyid.NewUserServiceClient(conn)

		req := &finduserbyid.UserRequest{Id: id}
		ctx := context.TODO();
		if res, err := client.UserService(ctx, req); err == nil {
				return res, nil
		} else {
			return res, err
		}
}