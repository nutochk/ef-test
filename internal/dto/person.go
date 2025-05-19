package dto

import "github.com/nutochk/ef-test/internal/models"

type PersonInfo struct {
	Id                int              `json:"id"`
	Name              string           `json:"name"`
	Surname           string           `json:"surname"`
	Patronymic        string           `json:"patronymic"`
	Age               int              `json:"age"`
	Gender            string           `json:"gender"`
	GenderProbability float64          `json:"gender_probability"`
	Nationality       []models.Country `json:"nationality"`
}
