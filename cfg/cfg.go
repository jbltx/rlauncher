package cfg

import "github.com/jinzhu/gorm"

type AuthProvider string

const (
	GoogleAP AuthProvider = "google"
)

// Config ...
type Config struct {
	Database struct {
		// Type can be either "mysql" or "postgres"
		Type string `json:"type"`
		Host string `json:"host"`
		Port uint16 `json:"port"`
		Name string `json:"name"`
	} `json:"database"`

	Auth struct {
		Provider AuthProvider `json:"provider"`
	} `json:"auth"`

	// Non serlialized
	GormDB *gorm.DB
}
