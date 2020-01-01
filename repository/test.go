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
}

func responsWithError(w http.ResponseWriter,code int, msg string ) {
	responsWithJson(w,code,map[string]string{"error":msg)
}
