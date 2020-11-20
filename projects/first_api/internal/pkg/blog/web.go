package blog

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhhieu/golang-training/first_api/internal/pkg/blog/service"
	"github.com/hhhieu/golang-training/first_api/pkg/database"
)

// Web contains the resources of web
type Web struct {
	Conf             WebConfig
	DBConnection     database.Connector
	ServiceAllocator service.IAllocator
}

const (
	routingGetMethod    = "GET"
	routingPostMethod   = "POST"
	routingPutMethod    = "PUT"
	routingPatchMethod  = "PATCH"
	routingDeleteMethod = "DELETE"
)

type routing struct {
	Method  string
	Path    string
	Service string
}

// NewWeb news the web services
func NewWeb(c WebConfig, dbConnection database.Connector) (*Web, error) {
	w := Web{
		Conf: c,
		ServiceAllocator: &service.Allocator{
			DBConnection: dbConnection,
		},
		DBConnection: dbConnection,
	}
	return &w, nil
}

// Serve start the web services
func (W *Web) Serve() error {
	// Define routings
	var routingList = []routing{
		routing{
			Method:  routingPostMethod,
			Path:    "/blog/user/create",
			Service: service.UserCreating,
		},
		routing{
			Method:  routingGetMethod,
			Path:    "/blog/user/:id",
			Service: service.UserGetting,
		},
		routing{
			Method:  routingPutMethod,
			Path:    "/blog/user/:id",
			Service: service.UserUpdating,
		},
		routing{
			Method:  routingPatchMethod,
			Path:    "/blog/user/:id",
			Service: service.UserPartialUpdating,
		},
		routing{
			Method:  routingDeleteMethod,
			Path:    "/blog/user/:id",
			Service: service.UserDeleting,
		}}
	// Connection to database
	W.DBConnection.Open()
	// Register the services
	router := gin.Default()
	for _, r := range routingList {
		s, err := W.ServiceAllocator.Allocate(r.Service)
		if err != nil {
			return err
		}
		switch r.Method {
		case routingGetMethod:
			router.GET(r.Path, s.Serve)
		case routingPostMethod:
			router.POST(r.Path, s.Serve)
		case routingPutMethod:
			router.PUT(r.Path, s.Serve)
		case routingPatchMethod:
			router.PATCH(r.Path, s.Serve)
		case routingDeleteMethod:
			router.DELETE(r.Path, s.Serve)
		default:
			panic(fmt.Errorf("Unsupport method %v", r.Method))
		}
	}
	// Start the services
	return router.Run(W.Conf.Address)
}
