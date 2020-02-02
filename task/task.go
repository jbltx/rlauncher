package task

import (
	"time"

	agent "github.com/jbltx/rlauncher/agent"
)

// Status ...
type Status uint8

const (
	// Idle ...
	Idle Status = 0
	// Queued ...
	Queued Status = 1
	// Pending ...
	Pending Status = 2
	// Processing ...
	Processing Status = 3
	// Suspended ...
	Suspended Status = 4
	// Failed ...
	Failed Status = 5
	// Completed ...
	Completed Status = 6
)

// Task ...
type Task struct {
	TaskID      uint32       `json:"taskID"`
	Command     string       `json:"command"`
	Status      Status       `json:"status"`
	AssignedTo  *agent.Agent `json:"assignedTo"`
	StartedAt   time.Time    `json:"startedAt"`
	CompletedAt time.Time    `json:"completedAt"`
	Report      struct {
		Log        string    `json:"log"`
		Errors     uint32    `json:"errors"`
		ReportDate time.Time `json:"reportDate"`
	} `json:"report"`
	Progress    float32 `json:"progress"`
	SystemUsage struct {
		CPU        uint32 `json:"cpu_rt"`
		AverageCPU uint32 `json:"cpu_avg"`
		RAM        uint32 `json:"ram_rt"`
		AverageRAM uint32 `json:"ram_avg"`
		HDD        uint32 `json:"hdd_rt"`
		AverageHDD uint32 `json:"hdd_avg"`
	} `json:"systemUsage"`
	SystemInfo struct {
		CPU       string `json:"cpu"`
		Frequency uint32 `json:"freq"`
		Cores     uint32 `json:"cores"`
		RAM       uint32 `json:"ram"`
	} `json:"systemInfo"`
}
