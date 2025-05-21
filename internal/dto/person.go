package dto

import "github.com/nutochk/ef-test/internal/models"

// PersonInfo information about person with id
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

type PersonFilter struct {
	Name    string `form:"name"`
	Surname string `form:"surname"`
	AgeMin  int    `form:"age_min"`
	AgeMax  int    `form:"age_max"`
	Gender  string `form:"gender"`
}

type Pagination struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

// PaginatedResponse information about people and pagination
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination struct {
		Total       int `json:"total"`
		CurrentPage int `json:"current_page"`
		PerPage     int `json:"per_page"`
	} `json:"pagination"`
}
