package grpcserver

import (
	"google.golang.org/grpc"
)

func Grpcserver() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
	// return client
}
