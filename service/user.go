package service

import (
	"encoding/json"
	"go-crud/dto"
	"go-crud/repository"
	"net/http"
)

func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	user, err := repository.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func GetAll(response http.ResponseWriter, request *http.Request) {
	users := repository.GetAllUser()
	responseWithJSON(response, http.StatusOK, users)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user dto.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := repository.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create user")
			return
		}
		responseWithJSON(response, http.StatusOK, user)
	}
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user dto.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := repository.UpdateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	result := repository.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete user successfully")
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
