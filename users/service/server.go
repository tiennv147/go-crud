package service

import (
	"context"
	"github.com/phamtanlong/go-crud/users/pb"
)

// AuthenticationServer can be embedded to have forward compatible implementations.
type AuthenticationServer struct {
	UserService UserService
}

func (au *AuthenticationServer) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	id, err := au.UserService.Verify(ctx, req.GetToken())
	res := &pb.VerifyResponse{
		Id: uint32(id),
	}
	if err != nil {
		res.Error = err.Error()
	}

	return res, nil
}
