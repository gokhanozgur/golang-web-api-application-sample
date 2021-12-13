package user

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
func (u *Repository) All() ([]model.User, error) {
	users := []model.User{}
	err := u.db.Find(&users).Error
	return users, err
}

// FindByID
func (u *Repository) FindByID(id uint64) (*model.User, error) {
	user := new(model.User)
	err := u.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

// Save
func (u *Repository) Save(user *model.User) (*model.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

// Delete
func (u *Repository) Delete(id uint64) error {
	err := u.db.Where(`id = ?`, id).Delete(&model.User{}).Error
	return err
}

// Migrate
func (u *Repository) Migrate() error {
	return u.db.AutoMigrate(&model.User{}).Error
}
