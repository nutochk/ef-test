package models

// Person personal data
type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

// PersonInfo information about person
type PersonInfo struct {
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Patronymic        string    `json:"patronymic"`
	Age               int       `json:"age"`
	Gender            string    `json:"gender"`
	GenderProbability float64   `json:"gender_probability"`
	Nationality       []Country `json:"nationality"`
}
