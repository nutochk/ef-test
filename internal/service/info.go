package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nutochk/ef-test/internal/models"
)

func getAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return 0, ErrRequest(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, ErrResponse(err)
	}
	var result models.AgeResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return 0, ErrParsing(err)
	}
	return result.Age, nil
}

func getGender(name string) (string, float64, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, ErrRequest(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, ErrResponse(err)
	}
	var result models.GenderResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return "", 0, ErrParsing(err)
	}
	return result.Gender, result.Probability, nil
}

func getCountries(name string) ([]models.Country, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrRequest(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrResponse(err)
	}
	var result models.NationalityResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, ErrParsing(err)
	}
	return result.Countries, nil
}
