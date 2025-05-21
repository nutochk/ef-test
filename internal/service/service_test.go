package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nutochk/ef-test/internal/dto"
	"github.com/nutochk/ef-test/internal/models"
	"github.com/nutochk/ef-test/internal/repository"
	logger2 "github.com/nutochk/ef-test/pkg/logger"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	logger, _ := logger2.New()
	svc := New(mockRepo, *logger)

	person := &models.Person{Name: "John", Surname: "Doe", Patronymic: "Smith"}
	expectedPersonInfo := &dto.PersonInfo{Id: 1, Name: "John", Surname: "Doe", Patronymic: "Smith"}

	mockRepo.EXPECT().Create(gomock.Any()).Return(1, nil)

	result, err := svc.Create(person)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.Id != expectedPersonInfo.Id {
		t.Errorf("Expected Id %d, got %d", expectedPersonInfo.Id, result.Id)
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	logger, _ := logger2.New()
	svc := New(mockRepo, *logger)

	person := &models.Person{Name: "John", Surname: "Doe", Patronymic: "Smith"}
	updatedInfo := &models.PersonInfo{}

	mockRepo.EXPECT().Update(1, person).Return(updatedInfo, nil)

	result, err := svc.Update(1, person)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != updatedInfo {
		t.Errorf("Expected %v, got %v", updatedInfo, result)
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	logger, _ := logger2.New()
	svc := New(mockRepo, *logger)

	mockRepo.EXPECT().Delete(1).Return(true, nil)

	err := svc.Delete(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	logger, _ := logger2.New()
	svc := New(mockRepo, *logger)

	expectedPersonInfo := &models.PersonInfo{Name: "John", Surname: "Doe"}

	mockRepo.EXPECT().GetById(1).Return(expectedPersonInfo, nil)

	result, err := svc.GetById(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expectedPersonInfo {
		t.Errorf("Expected %v, got %v", expectedPersonInfo, result)
	}
}

func TestGetPeople(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	logger, _ := logger2.New()
	svc := New(mockRepo, *logger)

	filters := &dto.PersonFilter{}
	pagination := &dto.Pagination{Page: 1, PerPage: 10}
	people := &[]dto.PersonInfo{{Id: 1, Name: "John", Surname: "Doe"}}
	total := 1

	mockRepo.EXPECT().GetPeople(filters, pagination).Return(people, total, nil)

	result, err := svc.GetPeople(filters, pagination)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Errorf("Expected people , got %v", result)
	}
}
