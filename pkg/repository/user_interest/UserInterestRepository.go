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
	users := []model.UserInterest{}
	err := u.db.Find(&users).Error
	return users, err
}

// FindByID
func (u *Repository) FindByID(id uint64) (*model.UserInterest, error) {
	user := new(model.UserInterest)
	err := u.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

// FindByUserID
func (u *Repository) FindByUserID(userId uint64) (*model.UserInterest, error) {
	user := new(model.UserInterest)
	err := u.db.Where(`user_id = ?`, userId).First(&user).Error
	return user, err
}

// FindByInterestID
func (u *Repository) FindByInterestID(interestId uint) (*model.UserInterest, error) {
	user := new(model.UserInterest)
	err := u.db.Where(`interest_id = ?`, interestId).First(&user).Error
	return user, err
}

// Save
func (u *Repository) Save(user *model.UserInterest) (*model.UserInterest, error) {
	err := u.db.Save(&user).Error
	return user, err
}

// Delete
func (u *Repository) Delete(id uint64) error {
	err := u.db.Delete(&model.UserInterest{}).Error
	return err
}

// Migrate
func (u *Repository) Migrate() error {
	return u.db.AutoMigrate(&model.UserInterest{}).Error
}
