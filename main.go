package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gorpc-template/config"
	"github.com/gorpc-template/models"
	"github.com/gorpc-template/protopb"
	"github.com/gorpc-template/userserver"
	"github.com/sethvargo/go-envconfig"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var c config.AppConfig

// init - initlizes DB and env config
func init() {

	ctx := context.Background()

	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	var err error
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v port=%v", c.Dbusername, c.Dbpassword, c.Dbdatabase, c.Dbport)
	models.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gorm open error : ", err)
		return
	}
	fmt.Println("Connected to database")
	models.Db.AutoMigrate(&models.User{})
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.GrpcServerPort))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", c.GrpcServerPort)

	s := grpc.NewServer()

	protopb.RegisterUserServiceServer(s, &userserver.User{})

	s.Serve(lis)
}
