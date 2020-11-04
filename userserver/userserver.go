package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/parikshitg/gorpc-template/protopb"
	"google.golang.org/grpc"
)

type user struct {
	Name  string
	Phone string
	Email string
}

func (*user) Registration(ctx context.Context, request *protopb.UserRegistrationRequest) (*protopb.UserRegistrationResponse, error) {
	name := request.Name
	response := &protopb.UserRegistrationResponse{
		Message: "Name :" + name + "phone : " + request.Phone + "email : " + request.Email,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50060"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	protopb.RegisterUserServiceServer(s, &user{})

	s.Serve(lis)
}
