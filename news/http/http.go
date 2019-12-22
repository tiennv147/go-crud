package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phamtanlong/go-crud/news/endpoints"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := mux.NewRouter()
	m.Methods("POST").Path("/create").Handler(httptransport.NewServer(endpoints.Create, DecodeCreateRequest, EncodeCreateResponse))
	m.Methods("POST").Path("/update").Handler(httptransport.NewServer(endpoints.Update, DecodeUpdateRequest, EncodeUpdateResponse))
	m.Methods("DELETE").Path("/delete").Handler(httptransport.NewServer(endpoints.Delete, DecodeDeleteRequest, EncodeDeleteResponse))
	m.Path("/read").Handler(httptransport.NewServer(endpoints.Read, DecodeReadRequest, EncodeReadResponse))
	return m
}
func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeCreateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeUpdateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func DecodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeDeleteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func DecodeReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ReadRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeReadResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
