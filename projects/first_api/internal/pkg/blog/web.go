package blog

import "github.com/hhhieu/golang-training/first_api/pkg/database"

// Web contains the resources of web
type Web struct {
	DBConnection *database.Connection
}

// NewWeb news the web services
func NewWeb(c *database.Connection) (*Web, error) {
	w := Web{
		DBConnection: c,
	}
	return &w, nil
}

// Serve start the web services
func (W *Web) Serve() error {
	return nil
}
