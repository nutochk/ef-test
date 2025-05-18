package service

import (
	"github.com/nutochk/ef-test/internal/models"
	"github.com/nutochk/ef-test/internal/repository"
	"github.com/nutochk/ef-test/pkg/logger"
)

type Service interface {
	Create(p *models.Person) (*models.PersonInfo, error)
	//Update(id int, i *models.Info) (*models.PersonInfo, error)
	//Delete(id int) (bool, error)
	//GetById(id int) (*models.PersonInfo, error)
}

type service struct {
	repo   repository.Repository
	logger logger.Logger
}

func New(repo repository.Repository, log logger.Logger) *service {
	return &service{repo: repo, logger: log}
}

func (s *service) Create(p *models.Person) (*models.PersonInfo, error) {
	s.logger.Debug("create method in service")
	return nil, nil
}
