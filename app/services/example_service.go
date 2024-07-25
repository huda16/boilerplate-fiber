package services

import (
	"boilerplate-api/app/models"
	"boilerplate-api/app/repositories"
	"boilerplate-api/database"
)

type ExampleService struct {
    Repository *repositories.ExampleRepository
}

func GetExampleService() *ExampleService {
    return &ExampleService{
        Repository: repositories.NewExampleRepository(database.DB),
    }
}

func (s *ExampleService) GetAllExamples() ([]models.Example, error) {
    return s.Repository.GetAll()
}
