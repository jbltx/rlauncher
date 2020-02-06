package task

import (
	"github.com/jbltx/rlauncher/cfg"
	job "github.com/jbltx/rlauncher/job/model"
	"github.com/jbltx/rlauncher/task/model"
	"github.com/jbltx/rlauncher/task/repository"

	"fmt"
)

type taskRepository interface {
	GetByID(uuid string) (*model.Task, error)
	GetJob(task *model.Task) (*job.Job, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) error
	Delete(task *model.Task) error
}

// Service ...
type Service struct {
	appCfg     *cfg.Config
	repository taskRepository
}

// NewService ...
func NewService(appConfig *cfg.Config) *Service {

	// init repository
	var repo taskRepository
	switch appConfig.Database.Type {
	case "mysql":
	case "postgres":
	case "sqlite3":
		repo = repository.NewGormRepository(appConfig)
	default:
		panic(fmt.Sprintf("Invalid database type: '%s'", appConfig.Database.Type))
	}

	// init deliveries
	// [TODO]

	return &Service{
		appCfg:     appConfig,
		repository: repo,
	}
}
