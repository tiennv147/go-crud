package repository

import (
	"github.com/jinzhu/gorm"

	"Service/Go-crud/go-crud/dto"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestRepository struct {
	DB *gorm.DB
}

func ProvideTestRepostiory(DB *gorm.DB) TestRepository {
	return TestRepository{DB: DB}
}

func (p *TestRepository) FindAll() []dto.Test {
	var tests []dto.Test
	p.DB.Find(&tests)

	return tests
}

func (p *TestRepository) FindByID(id int) dto.Test {
	var test dto.Test
	p.DB.First(&test, id)

	return test
}

func (p *TestRepository) Save(test dto.Test) dto.Test {
	p.DB.Save(&test)

	return test
}

func (p *TestRepository) Delete(test dto.Test) {
	p.DB.Delete(&test)
}

// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }

// func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
