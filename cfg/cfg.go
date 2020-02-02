package cfg

import "github.com/jinzhu/gorm"

// Config ...
type Config struct {
	Database struct {
		// Type can be either "mysql" or "postgres"
		Type string `json:"type"`
		Host string `json:"host"`
		Port uint16 `json:"port"`
		Name string `json:"name"`
	} `json:"database"`

	// Non serlialized
	GormDB *gorm.DB
}
