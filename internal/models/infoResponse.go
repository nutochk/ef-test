package models

type AgeResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GenderResponse struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type Country struct {
	CountryId   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalityResponse struct {
	Name      string    `json:"name"`
	Countries []Country `json:"country"`
}
