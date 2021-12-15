package service

import (
	"golang_web_api_application_sample/pkg/model"
	userInterest "golang_web_api_application_sample/pkg/repository/user_interest"
)

// UserInterestService
type UserInterestService struct {
	UserInterestRespository *userInterest.Repository
}

// UserInterestService Constructor
func NewUserInterestService(u *userInterest.Repository) UserInterestService {
	return UserInterestService{UserInterestRespository: u}
}

// Get All Service
func (u *UserInterestService) All() ([]model.UserInterest, error) {
	return u.UserInterestRespository.All()
}

// FindByID
func (u *UserInterestService) FindByID(id uint64) (*model.UserInterest, error) {
	return u.UserInterestRespository.FindByID(id)
}

// FindByUserID
func (u *UserInterestService) FindByUserID(userId uint64) ([]model.UserInterest, error) {
	return u.UserInterestRespository.FindByUserID(userId)
}

// FindByInterestID
func (u *UserInterestService) FindByInterestID(interestId uint) (*model.UserInterest, error) {
	return u.UserInterestRespository.FindByInterestID(interestId)
}

// Save
func (u *UserInterestService) Save(userInterest *model.UserInterest) (*model.UserInterest, error) {
	return u.UserInterestRespository.Save(userInterest)
}

// Delete
func (u *UserInterestService) Delete(id uint64) error {
	return u.UserInterestRespository.Delete(id)
}

// Migration
func (u *UserInterestService) Migrate() error {
	return u.UserInterestRespository.Migrate()
}
