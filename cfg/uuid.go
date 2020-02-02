package cfg

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BaseModel ...
type BaseModel struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
