package service

import (
	"github.com/nutochk/ef-test/internal/dto"
	"github.com/nutochk/ef-test/internal/models"
	"github.com/nutochk/ef-test/internal/repository"
	"github.com/nutochk/ef-test/pkg/logger"
	"go.uber.org/zap"
)

type Service interface {
	Create(p *models.Person) (*dto.PersonInfo, error)
	Update(id int, i *models.Person) (*models.PersonInfo, error)
	Delete(id int) error
	GetById(id int) (*models.PersonInfo, error)
	GetPeople(filters *dto.PersonFilter, pagination *dto.Pagination) (*dto.PaginatedResponse, error)
}

type service struct {
	repo   repository.Repository
	logger logger.Logger
}

func New(repo repository.Repository, log logger.Logger) *service {
	return &service{repo: repo, logger: log}
}

func (s *service) Create(p *models.Person) (*dto.PersonInfo, error) {
	s.logger.Debug("create method in service")
	age, err := getAge(p.Name)
	if err != nil {
		s.logger.Error("failed to get age in create method", zap.Error(err))
		return nil, err
	}
	gender, prob, err := getGender(p.Name)
	if err != nil {
		s.logger.Error("failed to get gender in create method", zap.Error(err))
		return nil, err
	}
	countries, err := getCountries(p.Name)
	if err != nil {
		s.logger.Error("failed to get countries in create method", zap.Error(err))
	}
	var pi models.PersonInfo
	pi.Name = p.Name
	pi.Surname = p.Surname
	pi.Patronymic = p.Patronymic
	pi.Age = age
	pi.Gender = gender
	pi.GenderProbability = prob
	pi.Nationality = countries
	id, err := s.repo.Create(&pi)
	if err != nil {
		s.logger.Error("failed to create in repository", zap.Error(err))
		return nil, err
	}
	person := dto.PersonInfo{
		Id:                id,
		Name:              p.Name,
		Surname:           p.Surname,
		Patronymic:        p.Patronymic,
		Age:               age,
		Gender:            gender,
		GenderProbability: prob,
		Nationality:       countries,
	}
	return &person, nil
}

func (s *service) Update(id int, p *models.Person) (*models.PersonInfo, error) {
	s.logger.Debug("update method in service")
	pi, err := s.repo.Update(id, p)
	if err != nil {
		s.logger.Error("failed to update in repository", zap.Error(err))
		return nil, err
	}
	return pi, nil
}

func (s *service) Delete(id int) error {
	s.logger.Debug("delete method in service")
	_, err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete in repository", zap.Error(err))
		return err
	}
	return nil
}

func (s *service) GetById(id int) (*models.PersonInfo, error) {
	s.logger.Debug("get by id method in service")
	pi, err := s.repo.GetById(id)
	if err != nil {
		s.logger.Error("failed to get by id in repository", zap.Error(err))
		return nil, err
	}
	return pi, nil
}

func (s *service) GetPeople(filters *dto.PersonFilter, pagination *dto.Pagination) (*dto.PaginatedResponse, error) {
	s.logger.Debug("get people method in service")
	people, total, err := s.repo.GetPeople(filters, pagination)
	if err != nil {
		s.logger.Error("failed to get people in repository", zap.Error(err))
		return nil, err
	}
	response := dto.PaginatedResponse{
		Data: people,
		Pagination: struct {
			Total       int `json:"total"`
			CurrentPage int `json:"current_page"`
			PerPage     int `json:"per_page"`
		}{Total: total, CurrentPage: pagination.Page, PerPage: pagination.PerPage},
	}
	return &response, nil
}
