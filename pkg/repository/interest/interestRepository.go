package interest

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
func (u *Repository) All() ([]model.Interest, error) {
	interests := []model.Interest{}
	err := u.db.Find(&interests).Error
	return interests, err
}

// FindByID
func (u *Repository) FindByID(id uint64) (*model.Interest, error) {
	interest := new(model.Interest)
	err := u.db.Where(`id = ?`, id).First(&interest).Error
	return interest, err
}

// Save
func (u *Repository) Save(interest *model.Interest) (*model.Interest, error) {
	err := u.db.Save(&interest).Error
	return interest, err
}

// Delete
func (u *Repository) Delete(id uint64) error {
	err := u.db.Delete(&model.Interest{}).Error
	return err
}

// Migrate
func (u *Repository) Migrate() error {
	return u.db.AutoMigrate(&model.Interest{}).Error
}
