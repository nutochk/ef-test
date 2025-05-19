package models

type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type PersonInfo struct {
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Patronymic        string    `json:"patronymic"`
	Age               int       `json:"age"`
	Gender            string    `json:"gender"`
	GenderProbability float64   `json:"gender_probability"`
	Nationality       []Country `json:"nationality"`
}
