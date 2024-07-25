package repositories

import (
	"boilerplate-api/app/models"

	"gorm.io/gorm"
)

type ExampleRepository struct {
    DB *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
    return &ExampleRepository{DB: db}
}

func (r *ExampleRepository) GetAll() ([]models.Example, error) {
    var examples []models.Example
    err := r.DB.Find(&examples).Error
    return examples, err
}
