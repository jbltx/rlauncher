package repository

import (
	"errors"

	"github.com/jbltx/rlauncher/cfg"
	pool "github.com/jbltx/rlauncher/pool/model"
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
func (repo *GormRepository) GetByID(uuid string) (*pool.Pool, error) {
	return nil, errors.New("Nothing here")
}

// Delete ...
func (repo *GormRepository) Delete(pool *pool.Pool) error {
	return errors.New("Nothing here")
}

// Create ...
func (repo *GormRepository) Create(pool *pool.Pool) (*pool.Pool, error) {
	return nil, errors.New("Nothing here")
}

// Update ...
func (repo *GormRepository) Update(pool *pool.Pool) error {
	return errors.New("Nothing here")
}
