package main

import (
	"context"
	"fmt"
	"net"

	"github.com/DataInfosec/faceappapi/middlewareservice/utils/jwt"
	"github.com/DataInfosec/faceappapi/proto/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) JwtService(ctx context.Context, req *middleware.JwtRequest) (*middleware.JwtResponse, error) {
	token := req.GetToken()
	fmt.Println("get authorization:: ", token)
	user, err := jwt.DecodeJWT(token)
	if err != nil {
		fmt.Println("jwt error :: ", err)
		return nil, err
	}
	fmt.Println("jwt response :: ", user)
	return &middleware.JwtResponse{
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Authorized: user.Authorized,
		Exp:        user.Exp,
	}, nil
}

func main() {
	fmt.Println("starting middleware application")
	lis, _ := net.Listen("tcp", ":50052")
	srv := grpc.NewServer()
	middleware.RegisterJwtServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		panic(e)
	}
}
