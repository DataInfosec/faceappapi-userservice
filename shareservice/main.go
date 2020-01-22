package main

import (
	"context"
	"fmt"
	"github.com/DataInfosec/faceappapi/proto/finduserbyid"
	"github.com/DataInfosec/faceappapi/shareservice/entity"
	"github.com/DataInfosec/faceappapi/shareservice/utils/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var collection *mongo.Collection = connection.Connection()

type server struct {
}

func (s *server) UserService(ctx context.Context, req *finduserbyid.UserRequest) (*finduserbyid.UserResponse, error) {
	id := req.GetId()
	var user entity.User
	userId, errInvalidId := primitive.ObjectIDFromHex(id)
	fmt.Println("get user id:: ", userId)
	if errInvalidId != nil {
		return nil, errInvalidId
	}
	err := collection.FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &finduserbyid.UserResponse{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Type:      user.Type,
		Id:        user.ID.Hex(),
	}, nil
}

func main() {
	fmt.Println("starting userbyidapplication")
	lis, _ := net.Listen("tcp", ":50050")
	srv := grpc.NewServer()
	finduserbyid.RegisterUserServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		panic(e)
	}
}
