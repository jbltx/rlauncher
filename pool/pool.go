package pool

import (
	"fmt"

	"github.com/jbltx/rlauncher/cfg"
	"github.com/jbltx/rlauncher/pool/model"
	"github.com/jbltx/rlauncher/pool/repository"
)

type poolRepository interface {
	GetByID(uuid string) (*model.Pool, error)
	Create(pool *model.Pool) (*model.Pool, error)
	Update(pool *model.Pool) error
	Delete(pool *model.Pool) error
}

// Service ...
type Service struct {
	appCfg     *cfg.Config
	repository poolRepository
}

// NewService ...
func NewService(appConfig *cfg.Config) *Service {

	// init repository
	var repo poolRepository
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
