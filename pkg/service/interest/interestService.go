package service

import (
	"golang_web_api_application_sample/pkg/model"
	"golang_web_api_application_sample/pkg/repository/interest"
)

// InterestService
type InterestService struct {
	InterestRespository *interest.Repository
}

// InterestService Constructor
func NewInterestService(i *interest.Repository) InterestService {
	return InterestService{InterestRespository: i}
}

// Get All Service
func (i *InterestService) All() ([]model.Interest, error) {
	return i.InterestRespository.All()
}

// FindByID
func (i *InterestService) FindByID(id uint) (*model.Interest, error) {
	return i.InterestRespository.FindByID(id)
}

// Save
func (i *InterestService) Save(interest *model.Interest) (*model.Interest, error) {
	return i.InterestRespository.Save(interest)
}

// Delete
func (i *InterestService) Delete(id uint) error {
	return i.InterestRespository.Delete(id)
}

// Migration
func (i *InterestService) Migrate() error {
	return i.InterestRespository.Migrate()
}
