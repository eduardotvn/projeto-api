package controllers

import (
	"github.com/eduardotvn/projeto-api/src/models"
	"github.com/eduardotvn/projeto-api/src/postgres"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct {
		Name     string
		Password string
	}

	c.Bind(&body)

	user := models.User{Name: body.Name, Password: body.Password}

	result := postgres.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {

	var users []models.User

	postgres.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	postgres.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}
