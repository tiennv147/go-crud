package repository

import (
	"Service/Go-crud/go-crud/dto"
	"database/sql"
	"encoding/json"
	"net/http"
)

type TestDto struct {
	Db *sql.DB
}
type TestRepository interface {
	Fetch(cursor string, num int64) ([]*dto.Test, error)
	GetByID(id int64) (*dto.Test, error)
	GetByTitle(title string) (*dto.Test, error)
	Update(article *dto.Test) (*dto.Test, error)
	Store(a *dto.Test) (int64, error)
	Delete(id int64) (bool, error)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
