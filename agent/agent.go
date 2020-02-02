package agent

import (
	"time"

	user "github.com/jbltx/rlauncher/user"
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
	Name            string     `json:"name"`
	Status          Status     `json:"status"`
	MachineName     string     `json:"machineName"`
	MachineUser     *user.User `json:"machineUser"`
	StatusUpdatedAt time.Time  `json:"statusUpdatedAt"`
}
