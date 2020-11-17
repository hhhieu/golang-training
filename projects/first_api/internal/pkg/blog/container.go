package blog

import (
	"fmt"
	"reflect"

	"github.com/hhhieu/golang-training/first_api/pkg/database"
	"github.com/hhhieu/golang-training/first_api/pkg/system"
)

// IContainer provides function to make an object
type IContainer interface {
	Make(t interface{}) error
}

// Container provides function to make object in blog app
type Container struct{}

const configFile = "/home/configs/blog_service.yml"

// Make object with specified type
func (C *Container) Make(t interface{}) error {
	switch v := t.(type) {
	case *Config:
		return makeConfig(v)
	case *database.Connection:
		return makeDatabase(v)
	default:
		return fmt.Errorf("Unsupport type %v", reflect.TypeOf(v))
	}
}

func makeConfig(c *Config) error {
	conf, err := LoadConfig(system.IOUtil{}, configFile)
	if err == nil {
		*c = *conf
	}
	return err
}

func makeDatabase(d *database.Connection) error {
	// Make configuration
	var c Config
	err := makeConfig(&c)
	if err != nil {
		return err
	}
	// Make database object
	database, err := database.NewConnection(c.Database)
	if err == nil {
		*d = *database
	}
	return err
}