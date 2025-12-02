// internal/handlers/user.go
package handlers

import (
	"errors"
	"net/http"

	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/database"
	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/models"
	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPass, _ := utils.HashPassword(input.Password)

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPass,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
