package command

import (
	"fmt"
	"log"
	"time"

	"github.com/hhhieu/golang-training/first_api/internal/pkg/blog"
	"github.com/hhhieu/golang-training/first_api/pkg/database"
)

// ServiceCommander interface for a commander
type ServiceCommander interface {
	Serve() error
	WaitForDatabase() error
}

// BlogCommander implements commands in blog
type BlogCommander struct {
	BlogContainer blog.AppContainer
}

// Serve start all blog services
func (C *BlogCommander) Serve() error {
	// Check object
	if err := check(C); err != nil {
		return err
	}
	return nil
}

// WaitForDatabase waits until database ready or timeout
func (C *BlogCommander) WaitForDatabase(startPeriod int, retryTime int, interval int) error {
	// Check object
	if err := check(C); err != nil {
		return err
	}
	// Create the database object
	var db database.Provider
	err := C.BlogContainer.Make(&db)
	if err != nil {
		return err
	}
	// Try to connect to database
	log.Printf("Try to connect to database")
	for err = db.Open(); err != nil && retryTime > 0; retryTime-- {
		time.Sleep(time.Duration(interval) * time.Second)
		log.Printf("Retry to connect to database")
		err = db.Open()
	}
	if err != nil {
		log.Printf("Fail to wait for database")
	} else {
		log.Printf("Success to wait for database")
	}
	return err
}

// NewCommander create a commander object
func NewCommander() (c *BlogCommander) {
	return &BlogCommander{
		BlogContainer: &blog.Container{},
	}
}

func check(c *BlogCommander) error {
	// Check object
	if c.BlogContainer == nil {
		return fmt.Errorf("Blog container is nil. Please use the provided function to create this object")
	}
	return nil
}
