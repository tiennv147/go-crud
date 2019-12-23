package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phamtanlong/go-crud/users/endpoints"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := mux.NewRouter()
	m.Methods("POST").Path("/register").Handler(httptransport.NewServer(endpoints.Register, DecodeRegisterRequest, EncodeRegisterResponse))
	m.Methods("POST").Path("/login").Handler(httptransport.NewServer(endpoints.Login, DecodeLoginRequest, EncodeLoginResponse))
	m.Methods("POST").Path("/verify").Handler(httptransport.NewServer(endpoints.Verify, DecodeVerifyRequest, EncodeVerifyResponse))
	return m
}
func DecodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeRegisterResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func DecodeVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	req.Token = r.Header.Get("Authorization")
	return req, err
}
func EncodeVerifyResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
