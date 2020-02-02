package model

import (
	"time"

	pool "github.com/jbltx/rlauncher/pool/model"
	task "github.com/jbltx/rlauncher/task/model"
	user "github.com/jbltx/rlauncher/user/model"
)

// Job ...
type Job struct {
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	AssignedPools   []*pool.Pool `json:"assignedPools"`
	ConcurrentTasks uint32       `json:"concurrentTasks"`
	Priority        uint32       `json:"priority"`
	Author          *user.User   `json:"author"`
	Tasks           []*task.Task `json:"tasks"`
	Status          task.Status  `json:"status"`
	SubmittedAt     time.Time    `json:"submittedAt"`
	CompletedAt     time.Time    `json:"completedAt"`
	StartedAt       time.Time    `json:"startedAt"`
}
