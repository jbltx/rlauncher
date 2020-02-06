package repository

import (
	"errors"

	"github.com/jbltx/rlauncher/cfg"
	job "github.com/jbltx/rlauncher/job/model"
	task "github.com/jbltx/rlauncher/task/model"
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
func (repo *GormRepository) GetByID(uuid string) (*task.Task, error) {
	return nil, errors.New("Nothing here")
}

// GetJob ...
func (repo *GormRepository) GetJob(task *task.Task) (*job.Job, error) {
	return nil, errors.New("Nothing here")
}

// Delete ...
func (repo *GormRepository) Delete(task *task.Task) error {
	return errors.New("Nothing here")
}

// Create ...
func (repo *GormRepository) Create(task *task.Task) (*task.Task, error) {
	return nil, errors.New("Nothing here")
}

// Update ...
func (repo *GormRepository) Update(task *task.Task) error {
	return errors.New("Nothing here")
}
