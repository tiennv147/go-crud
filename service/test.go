package service

import (
	"Service/Go-crud/go-crud/dto"
	"Service/Go-crud/go-crud/repository"
)

type TestService struct {
	TestRepository repository.TestRepository
}

// func ProvideTestService(t TestRepository) TestService {
// 	return TestService{Testrepository: t}

// }

func (p *TestService) FindAll() []dto.Test {
	return p.TestRepository.FindAll()
}

func (p *TestService) FindByID(id int) dto.Test {
	return p.TestRepository.FindByID(id)
}

func (p *TestService) Save(test dto.Test) dto.Test {
	p.TestRepository.Save(test)
	return test
}

func (p *TestService) Delete(test dto.Test) {
	p.TestRepository.Delete(test)
}
