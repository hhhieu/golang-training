package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhhieu/golang-training/first_api/internal/pkg/service/user"
)

// List of supported service
const (
	UserCreating string = "create_user"
	UserListing  string = "list_user"
	UserGetting  string = "get_user"
	UserUpdating string = "update_user"
	UserDeleting string = "delete_user"
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
type Allocator struct{}

// Allocate create a service assosiate with the specified name
func (A *Allocator) Allocate(name string) (Service, error) {
	switch name {
	case UserCreating:
		return &user.Creating{}, nil
	default:
		return nil, fmt.Errorf("Unsupport service name: %v", name)
	}
}
