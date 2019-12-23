package main

import (
	"context"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	Register(ctx context.Context, username string, password string) (id uint, error error)
	Login(ctx context.Context, username string, password string) (token string, error error)
	Verify(ctx context.Context, token string) (id uint, error error)
}

type User struct {
	gorm.Model
	Username string	`json:"username"`
	Password string	`json:"password"`
}
