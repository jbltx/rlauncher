package model

import "github.com/jbltx/rlauncher/cfg"

// User ...
type User struct {
	cfg.BaseModel
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
}
