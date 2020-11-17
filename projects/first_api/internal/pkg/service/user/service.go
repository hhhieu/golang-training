package user

import "github.com/gin-gonic/gin"

// Creating implement the creating service
type Creating struct{}

// Serve creates user
func (C Creating) Serve(c *gin.Context) {
	c.String(200, "Hello world")
}
