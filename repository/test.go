package repository

import (
	"database/sql"
	"dto"
	"net/http"
)

type TestDto struct {
	Db *sql.DB
}

func (testDto TestDto) findAll() (test []dto.Test, err error) {
	rows,err := testDto.Db.Query("select * from test")
	if err != nill {
		return nill,err
	} 
}

func respondWithError(w http.ResponseWriter,code int, msg string ) {
	respondWithJson(w,code,map[string]string{"error":msg)
}

func respondWithJson(w http.ResponseWriter, code int , payload interface{}){
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}