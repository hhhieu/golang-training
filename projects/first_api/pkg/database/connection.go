package database

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// List of supported database
const (
	MySQL string = "mysql"
)

// Connection contains the database connection
type Connection struct {
	DB   *gorm.DB
	Conf Config
}

// GetSupportedType gets list of supported database
func GetSupportedType() []string {
	return []string{MySQL}
}

// NewConnection new a connection
func NewConnection(c Config) (*Connection, error) {
	// Validate the configuration
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Creating Connection failed. The configuration is not valid:%v", err)
	}
	c.Type = strings.ToLower(c.Type)
	p := Connection{
		Conf: c,
	}
	return &p, nil
}

// Open create the database connection pool
func (P *Connection) Open() error {
	switch P.Conf.Type {
	case MySQL:
		dsn := generateMySQLDSN(P.Conf)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("Failed to connect to database %v", err)
		}
		P.DB = db
	default:
		return fmt.Errorf("Unsupport database type %v", P.Conf.Type)
	}
	return nil
}

// AutoMigrate creates table base on the struct definition
func (P *Connection) AutoMigrate(s ...interface{}) error {
	return P.DB.AutoMigrate(s...)
}

func generateMySQLDSN(c Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.URL, c.DBName)
}
