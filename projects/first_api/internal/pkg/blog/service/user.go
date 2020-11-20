package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhieu/golang-training/first_api/internal/pkg/model"
	"github.com/hhhieu/golang-training/first_api/pkg/database"
)

// Resource defines the service resources
type Resource struct {
	DBConnection database.Connector
}

// UserCreatingService implement the creating service
type UserCreatingService Resource

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

// UserGettingService implement the getting service
type UserGettingService Resource

// Serve creates user
func (S *UserGettingService) Serve(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	if err := S.DBConnection.First(&user, id).Error(); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Can not get user with id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

// UserUpdatingService implement the updating service
type UserUpdatingService Resource

// Serve updates user
func (S *UserUpdatingService) Serve(c *gin.Context) {
	// Bind the requested data
	id := c.Param("id")
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Can not update user with id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	// Find user
	var oldUser model.User
	if err := S.DBConnection.First(&oldUser, id).Error(); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Invalid user id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	// Update user
	newUser.ID = oldUser.ID
	if err := S.DBConnection.Save(&newUser).Error(); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Can not update user with id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	c.JSON(200, newUser)
}

// UserPartialUpdatingService implement the updating service
type UserPartialUpdatingService Resource

// Serve updates user
func (S *UserPartialUpdatingService) Serve(c *gin.Context) {
	// Bind the requested data
	id := c.Param("id")
	var partialNewUser model.User
	if err := c.ShouldBindJSON(&partialNewUser); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Can not partial update user with id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	// Find user
	var oldUser model.User
	if err := S.DBConnection.First(&oldUser, id).Error(); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Invalid user id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	// Update partial user
	if err := S.DBConnection.Model(&oldUser).Updates(&partialNewUser).Error(); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Can not partial update user with id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	c.JSON(200, oldUser)
}

// UserDeletingService implement the deleting service
type UserDeletingService Resource

// Serve creates user
func (S *UserDeletingService) Serve(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	if err := S.DBConnection.Delete(&user, id).Error(); err != nil {
		c.JSON(400, gin.H{
			"code":    1,
			"message": fmt.Sprintf("Can not delete user with id %v", id),
			"detail":  err.Error(),
		})
		return
	}
	c.JSON(204, nil)
}
