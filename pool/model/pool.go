package model

import (
	"github.com/jbltx/rlauncher/cfg"
	"golang.org/x/crypto/ssh/agent"
)

// Type defines the type of a Pool instance
type Type uint8

const (
	// AgentPoolType ...
	AgentPoolType Type = 0
	// PoolSetType ...
	PoolSetType Type = 1
)

// Pool ...
type Pool struct {
	cfg.BaseModel
	Name   string         `json:"name"`
	Type   Type           `json:"type"`
	Agents []*agent.Agent `json:"agents"`
}

// RelationType ...
type RelationType uint8

const (
	// Union ...
	Union RelationType = 0
	// Intersection ...
	Intersection RelationType = 1
)

// PoolSet ...
type PoolSet struct {
	cfg.BaseModel
	A        *Pool        `json:"a"`
	B        *Pool        `json:"b"`
	Relation RelationType `json:"relation"`
}
