package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhieu/golang-training/first_api/internal/pkg/model"
	"github.com/hhhieu/golang-training/first_api/pkg/database"
)

// UserCreatingService implement the creating service
type UserCreatingService struct {
	DBConnection *database.Connection
}

// Serve creates user
func (S *UserCreatingService) Serve(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    ErrDatabase,
			"message": " Invalid format data",
			"detail":  err.Error(),
		})
		return
	}
	if err := S.DBConnection.Create(&user).Error(); err != nil {
		c.JSON(500, gin.H{
			"code":    1,
			"message": "Can not create user",
			"detail":  err.Error(),
		})
		return
	}
	c.JSON(200, user)
}
