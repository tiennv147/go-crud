package repository

import (
	"Service/Go-crud/go-crud/dto"
	"errors"
)

var (
	listUser = make([]*dto.User, 0)
)

// function create user
func CreateUser(user *dto.User) bool {
	if user.Id != "" && user.Name != "" && user.Password != "" {
		if userF, _ := FindUser(user.Id); userF == nil {
			listUser = append(listUser, user)
			return true
		}
	}
	return false
}

func UpdateUser(eUser *dto.User) bool {
	for index, user := range listUser {
		if user.Id == eUser.Id {
			listUser[index] = eUser
			return true
		}
	}
	return false
}

func FindUser(id string) (*dto.User, error) {
	for _, user := range listUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("User does not exist")
}

func DeleteUser(id string) bool {
	for index, user := range listUser {
		if user.Id == id {
			copy(listUser[index:], listUser[index+1:])
			listUser[len(listUser)-1] = &dto.User{}
			listUser = listUser[:len(listUser)-1]
			return true
		}
	}
	return false
}

func GetAllUser() []*dto.User {
	return listUser
}
