package service

import (
	"golang_web_api_application_sample/pkg/model"
	"golang_web_api_application_sample/pkg/repository/user"
)

// UserService
type UserService struct {
	UserRespository *user.Repository
}

// UserService Constructor
func NewUserService(u *user.Repository) UserService {
	return UserService{UserRespository: u}
}

// Get All Service
func (u *UserService) All() ([]model.User, error) {
	return u.UserRespository.All()
}

// FindByID
func (u *UserService) FindByID(id uint64) (*model.User, error) {
	return u.UserRespository.FindByID(id)
}

// Save
func (u *UserService) Save(user *model.User) (*model.User, error) {
	return u.UserRespository.Save(user)
}

// Delete
func (u *UserService) Delete(id uint64) error {
	return u.UserRespository.Delete(id)
}

// Migration
func (u *UserService) Migrate() error {
	return u.UserRespository.Migrate()
}
