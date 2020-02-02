package user

import (
	"github.com/jbltx/rlauncher/cfg"
	repoPkg "github.com/jbltx/rlauncher/user/repository"

	"fmt"
)

// User ...
type User struct {
	cfg.BaseModel
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
}

type repository interface {
	GetByID(uuid string) (*User, error)
	Delete(user *User) error
	Create(user *User) (*User, error)
	Update(user *User) error
}

// Service ...
type Service struct {
	cfg        *cfg.Config
	repository *repository
}

// NewService ...
func NewService(cfg *cfg.Config) *Service {

	// init repository
	var repo repository
	switch cfg.Database.Type {
	case "mysql":
	case "postgres":
	case "sqlite3":
		repo = repoPkg.NewGormRepository(cfg)
	default:
		panic(fmt.Sprintf("Invalid database type: '%s'", cfg.Database.Type))
	}

	// init deliveries
	// [todo]

	return &Service{
		cfg:        cfg,
		repository: &repo,
	}
}
