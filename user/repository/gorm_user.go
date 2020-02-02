package repository

import (
	"errors"

	"github.com/jbltx/rlauncher/cfg"
	"github.com/jbltx/rlauncher/user/model"
	"github.com/jinzhu/gorm"
)

// GormRepository ...
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository ...
func NewGormRepository(cfg *cfg.Config) *GormRepository {
	return &GormRepository{
		db: cfg.GormDB,
	}
}

// GetByID ...
func (repo *GormRepository) GetByID(uuid string) (*model.User, error) {
	return nil, errors.New("Nothing here")
}

// Delete ...
func (repo *GormRepository) Delete(user *model.User) error {
	return errors.New("Nothing here")
}

// Create ...
func (repo *GormRepository) Create(user *model.User) (*model.User, error) {
	return nil, errors.New("Nothing here")
}

// Update ...
func (repo *GormRepository) Update(user *model.User) error {
	return errors.New("Nothing here")
}
