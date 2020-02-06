package model

import (
	"time"

	"github.com/jbltx/rlauncher/cfg"
	user "github.com/jbltx/rlauncher/user/model"
)

// Status ...
type Status uint8

const (
	// Idle ...
	Idle Status = 0
	// Processing ...
	Processing Status = 1
	// Offline ...
	Offline Status = 2
)

// Agent ...
type Agent struct {
	cfg.BaseModel
	Name            string     `json:"name" gorm:"type:varchar(100);unique_index"`
	Status          Status     `json:"status" gorm:""`
	MachineName     string     `json:"machineName" gorm:"type:varchar(100);unique"`
	MachineUser     *user.User `json:"machineUser"`
	StatusUpdatedAt time.Time  `json:"statusUpdatedAt"`
}
