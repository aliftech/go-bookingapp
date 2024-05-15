package controllers

import (
	"net/http"

	"github.com/aliftech/go-bookingapp/models"
	"github.com/aliftech/go-bookingapp/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get data of req body
	var body struct {
		Name     string
		Email    string
		Password string
		Phone    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to read request body",
		})

		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to hash password",
		})

		return
	}

	// Insert user data into database
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash), Phone: body.Phone}

	result := utils.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to save data.",
		})

		return
	}

	// Return response
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   false,
		"message": "New user have been added.",
	})
}
