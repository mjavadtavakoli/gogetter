package controllers

import (
	"gogetters/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	values := []models.User{}
	for _, v := range models.Users {
		values = append(values, v)
	}
	c.JSON(http.StatusOK, values)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, exists := models.Users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	user.ID = uuid.New().String()
	models.Users[user.ID] = user
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, exists := models.Users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	delete(models.Users, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
