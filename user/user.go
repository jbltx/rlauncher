package user

import (
	"github.com/jbltx/rlauncher/cfg"
	"github.com/jbltx/rlauncher/user/model"
	"github.com/jbltx/rlauncher/user/repository"

	"fmt"
)

type userRepository interface {
	GetByID(uuid string) (*model.User, error)
	Delete(user *model.User) error
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) error
}

// Service ...
type Service struct {
	appCfg     *cfg.Config
	repository *userRepository
}

// NewService ...
func NewService(appConfig *cfg.Config) *Service {

	// init repository
	var repo userRepository
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
		repository: &repo,
	}
}
