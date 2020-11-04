package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/parikshitg/gorpc-template/config"
	"github.com/parikshitg/gorpc-template/protopb"
	"github.com/sethvargo/go-envconfig"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type user struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

func (*user) Registration(ctx context.Context, request *protopb.UserRegistrationRequest) (*protopb.UserRegistrationResponse, error) {

	if IfUserExists(request.Email, request.Password) {
		return nil, errors.New("User Already Exists")
	}
	response := &protopb.UserRegistrationResponse{
		Message: "Successfully Registered user :" + request.Name,
	}
	CreateUser(request.Name, request.Phone, request.Email, request.Password)
	return response, nil
}

func (*user) Login(ctx context.Context, request *protopb.UserLoginRequest) (*protopb.UserLoginResponse, error) {
	response := &protopb.UserLoginResponse{
		Message: "Successfully Login user with email :" + request.Email,
	}
	if !IfUserExists(request.Email, request.Password) {
		return nil, errors.New("User does not exists")
	}
	return response, nil
}

var db *gorm.DB

func main() {

	ctx := context.Background()

	var c config.AppConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	var err error
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v port=%v", c.Dbusername, c.Dbpassword, c.Dbdatabase, c.Dbport)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gorm open error : ", err)
		return
	}
	fmt.Println("Connected to database")

	// Create a table
	db.AutoMigrate(&user{})

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

func CreateUser(name, phone, email, password string) {
	db.Create(&user{Name: name, Phone: phone, Email: email, Password: password})
}

func IfUserExists(email, password string) bool {
	var u user
	db.Where("email = ?", email).First(&u)
	if email != u.Email && password != u.Password {
		return false
	}
	return true
}
