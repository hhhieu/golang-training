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

// Connector defines interface of connection
type Connector interface {
	Open() error
	AutoMigrate(s ...interface{}) error
	Create(s interface{}) Connector
	First(s interface{}, conds ...interface{}) Connector
	Find(s interface{}, conds ...interface{}) Connector
	Save(value interface{}) Connector
	Error() error
}

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

// Create creates a row in table
func (P *Connection) Create(s interface{}) Connector {
	P.DB = P.DB.Create(s)
	return P
}

// Find finds a row in table
func (P *Connection) Find(s interface{}, conds ...interface{}) Connector {
	db := P.DB.Find(s, conds)
	return &Connection{
		DB:   db,
		Conf: P.Conf}
}

// First finds first row which matchs condition
func (P *Connection) First(s interface{}, conds ...interface{}) Connector {
	db := P.DB.First(s, conds)
	return &Connection{
		DB:   db,
		Conf: P.Conf}
}

// Save updates all fields of exist row
func (P *Connection) Save(value interface{}) Connector {
	db := P.DB.Save(value)
	return &Connection{
		DB:   db,
		Conf: P.Conf}
}

// Error gets current error
func (P *Connection) Error() error {
	return P.DB.Error
}

func generateMySQLDSN(c Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.URL, c.DBName)
}
