package userserver

import (
	"context"
	"errors"

	"github.com/gorpc-template/models"
	"github.com/gorpc-template/protopb"
)

// User - user structure
type User struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

// Registration - user Registration rpc method
func (User) Registration(ctx context.Context, request *protopb.UserRegistrationRequest) (*protopb.UserRegistrationResponse, error) {

	if models.ExistingUser(request.Email, request.Password) {
		return nil, errors.New("User Already Exists")
	}
	response := &protopb.UserRegistrationResponse{
		Message: "Successfully Registered user :" + request.Name,
	}
	models.CreateUser(request.Name, request.Phone, request.Email, request.Password)
	return response, nil
}

// Login - user Login rpc method
func (*User) Login(ctx context.Context, request *protopb.UserLoginRequest) (*protopb.UserLoginResponse, error) {
	response := &protopb.UserLoginResponse{
		Message: "Successfully Login user with email :" + request.Email,
	}
	if !models.ExistingUser(request.Email, request.Password) {
		return nil, errors.New("User does not exists")
	}
	return response, nil
}

// List - users List rpc method
func (*User) List(ctx context.Context, request *protopb.UserListRequest) (*protopb.UserListResponse, error) {

	if !models.ExistingUser(request.Email, request.Password) {
		return nil, errors.New("User does not exists")
	}
	users := models.ListUser(request.Email, request.Password)
	response := &protopb.UserListResponse{
		UsersList: make([]*protopb.User, len(users)),
	}

	for i := range users {
		response.UsersList[i] = &protopb.User{
			Name:     users[i].Name,
			Phone:    users[i].Phone,
			Email:    users[i].Email,
			Password: users[i].Password,
		}
	}

	return response, nil
}
