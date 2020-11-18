package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/hhhieu/golang-training/first_api/internal/pkg/service"
	"github.com/hhhieu/golang-training/first_api/pkg/database"
)

// Web contains the resources of web
type Web struct {
	Conf             WebConfig
	DBConnection     *database.Connection
	ServiceAllocator service.IAllocator
}

// NewWeb news the web services
func NewWeb(c WebConfig, dbConnection *database.Connection) (*Web, error) {
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
	W.DBConnection.Open()
	router := gin.Default()
	s, _ := W.ServiceAllocator.Allocate(service.UserCreating)
	router.POST("/blog/user/create", s.Serve)
	return router.Run(W.Conf.Address)
}
