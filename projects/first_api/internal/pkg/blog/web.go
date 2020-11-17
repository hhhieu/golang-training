package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/hhhieu/golang-training/first_api/pkg/database"
)

// Web contains the resources of web
type Web struct {
	Conf         WebConfig
	DBConnection *database.Connection
}

// NewWeb news the web services
func NewWeb(c WebConfig, dbConnection *database.Connection) (*Web, error) {
	w := Web{
		Conf:         c,
		DBConnection: dbConnection,
	}
	return &w, nil
}

// Serve start the web services
func (W *Web) Serve() error {
	router := gin.Default()
	return router.Run(W.Conf.Address)
}
