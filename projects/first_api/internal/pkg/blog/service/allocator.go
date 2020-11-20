package service

import (
	"fmt"

	"github.com/hhhieu/golang-training/first_api/pkg/database"

	"github.com/gin-gonic/gin"
)

// List of supported service
const (
	UserCreating        string = "create_user"
	UserListing         string = "list_user"
	UserGetting         string = "get_user"
	UserUpdating        string = "update_user"
	UserPartialUpdating string = "partial_update_user"
	UserDeleting        string = "delete_user"
)

// Service define the service interface
type Service interface {
	Serve(c *gin.Context)
}

// IAllocator define the allocator interface
type IAllocator interface {
	Allocate(string) (Service, error)
}

// Allocator implements the service allocator
type Allocator struct {
	DBConnection database.Connector
}

// Allocate create a service assosiate with the specified name
func (A *Allocator) Allocate(name string) (Service, error) {
	switch name {
	case UserCreating:
		return &UserCreatingService{DBConnection: A.DBConnection}, nil
	case UserGetting:
		return &UserGettingService{DBConnection: A.DBConnection}, nil
	case UserUpdating:
		return &UserUpdatingService{DBConnection: A.DBConnection}, nil
	case UserPartialUpdating:
		return &UserPartialUpdatingService{DBConnection: A.DBConnection}, nil
	case UserDeleting:
		return &UserDeletingService{DBConnection: A.DBConnection}, nil
	default:
		return nil, fmt.Errorf("Unsupport service name: %v", name)
	}
}
