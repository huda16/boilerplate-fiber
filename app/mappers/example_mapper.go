package mappers

import (
	"boilerplate-api/app/dto"
	"boilerplate-api/app/models"
)

func ToExampleDTO(example models.Example) dto.ExampleDTO {
    return dto.ExampleDTO{
        ID:    example.ID,
        Name:  example.Name,
        Email: example.Email,
    }
}

func ToExampleModel(dto dto.ExampleDTO) models.Example {
    return models.Example{
        ID:    dto.ID,
        Name:  dto.Name,
        Email: dto.Email,
    }
}

func ToExampleDTOs(examples []models.Example) []dto.ExampleDTO {
    dtos := make([]dto.ExampleDTO, len(examples))
    for i, example := range examples {
        dtos[i] = ToExampleDTO(example)
    }
    return dtos
}
