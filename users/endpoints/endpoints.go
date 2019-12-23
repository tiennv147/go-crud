package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/phamtanlong/go-crud/users/service"
)

type RegisterRequest struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}
type RegisterResponse struct {
	Id    uint	`json:"id"`
	Error error	`json:"error"`
}

func MakeRegisterEndpoint(u service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		id, error := u.Register(ctx, req.Username, req.Password)
		return RegisterResponse{Id: id, Error: error}, nil
	}
}

type LoginRequest struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}
type LoginResponse struct {
	Token string	`json:"token"`
	Error error		`json:"error"`
}

func MakeLoginEndpoint(u service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, error := u.Login(ctx, req.Username, req.Password)
		return LoginResponse{Token: token, Error: error}, nil
	}
}

type VerifyRequest struct {
	Token string	`json:"token"`
}
type VerifyResponse struct {
	Id    uint	`json:"id"`
	Error error	`json:"error"`
}

func MakeVerifyEndpoint(u service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyRequest)
		id, error := u.Verify(ctx, req.Token)
		return VerifyResponse{Id: id, Error: error}, nil
	}
}

type Endpoints struct {
	Register endpoint.Endpoint
	Login    endpoint.Endpoint
	Verify   endpoint.Endpoint
}
