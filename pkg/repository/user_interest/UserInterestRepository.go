package repository

import (
	"golang_web_api_application_sample/pkg/model"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

// New repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Get All
func (u *Repository) All() ([]model.UserInterest, error) {
	userInterests := []model.UserInterest{}
	err := u.db.Find(&userInterests).Error
	return userInterests, err
}

// FindByID
func (u *Repository) FindByID(id uint64) (*model.UserInterest, error) {
	interest := new(model.UserInterest)
	err := u.db.Where(`id = ?`, id).First(&interest).Error
	return interest, err
}

// FindByUserID
func (u *Repository) FindByUserID(userId uint64) (*model.UserInterest, error) {
	userInterests := new(model.UserInterest)
	err := u.db.Where(`user_id = ?`, userId).First(&userInterests).Error
	return userInterests, err
}

// FindByInterestID
func (u *Repository) FindByInterestID(interestId uint) (*model.UserInterest, error) {
	userInterests := new(model.UserInterest)
	err := u.db.Where(`interest_id = ?`, interestId).First(&userInterests).Error
	return userInterests, err
}

// Save
func (u *Repository) Save(user *model.UserInterest) (*model.UserInterest, error) {
	err := u.db.Save(&user).Error
	return user, err
}

// Delete
func (u *Repository) Delete(id uint64) error {
	err := u.db.Where(`id = ?`, id).Delete(&model.UserInterest{}).Error
	return err
}

// Migrate
func (u *Repository) Migrate() error {
	return u.db.AutoMigrate(&model.UserInterest{}).Error
}
