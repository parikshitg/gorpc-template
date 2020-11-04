package main

import (
	"context"
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
	Name  string
	Phone string
	Email string
}

func (*user) Registration(ctx context.Context, request *protopb.UserRegistrationRequest) (*protopb.UserRegistrationResponse, error) {
	name := request.Name
	response := &protopb.UserRegistrationResponse{
		Message: "Name :" + name + "phone : " + request.Phone + "email : " + request.Email,
	}
	Create(request.Name, request.Phone, request.Email)
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

func Create(name, phone, email string) {
	db.Create(&user{Name: name, Phone: phone, Email: email})
}
