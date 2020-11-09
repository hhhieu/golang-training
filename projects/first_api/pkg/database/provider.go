package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// List of supported database
const (
	MySQL string = "mysql"
)

// Provider contains the database connection
type Provider struct {
	DB     *gorm.DB
	dSN    string
	dBType string
}

// GetSupportedType gets list of supported database
func GetSupportedType() []string {
	return []string{MySQL}
}

// NewProvider new a provider
func NewProvider(c Config) (*Provider, error) {
	p := Provider{
		dBType: c.Type,
	}
	// Validate the configuration
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Creating provider failed. The configuration is not valid:%v", err)
	}
	// Generate the data source name (DSN)
	switch c.Type {
	case MySQL:
		p.dSN = generateMySQLDSN(c)
	default:
		return nil, fmt.Errorf("Unsupport database type %v", c.Type)
	}
	return &p, nil
}

// Open create the database connection pool
func (P *Provider) Open() error {
	switch P.dBType {
	case MySQL:
		db, err := gorm.Open(mysql.Open(P.dSN), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("Failed to connect to database %v", err)
		}
		P.DB = db
	default:
		return fmt.Errorf("Unsupport database type %v", P.dBType)
	}
	return nil
}

func generateMySQLDSN(c Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.URL, c.DBName)
}
