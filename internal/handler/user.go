package handler

import (
	"crypto-news-server/internal/data"
	"crypto-news-server/internal/models"
	"crypto-news-server/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": data.Users,
	})
}

func PostUsers(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	if err := service.ValidateUserPassword(&newUser); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidateUserEmail(&newUser, data.Users); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	newUser.ID = len(data.Users) + 1
	newUser.Created_at = time.Now()
	newUser.Updated_at = time.Now()
	data.Users = append(data.Users, newUser)

	data.SaveUser()

	c.JSON(http.StatusCreated, newUser)
}

func GetUsersById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for _, p := range data.Users {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func DeleteUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for i, p := range data.Users {
		if p.ID == id {
			data.Users = append(data.Users[:i], data.Users[i+1:]...)
			data.SaveUser()
			c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func UpdateUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var existingUser *models.User
	for i := range data.Users {
		if data.Users[i].ID == id {
			existingUser = &data.Users[i]
			break
		}
	}

	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if updatedUser.Email != "" && existingUser.Email != updatedUser.Email {
		if err := service.ValidateUserEmail(&updatedUser, data.Users); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
			return
		}
	} else {
		updatedUser.Email = existingUser.Email
	}

	if updatedUser.Password != "" {
		if err := service.ValidateUserPassword(&updatedUser); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
			return
		}
	} else {
		updatedUser.Password = existingUser.Password
	}

	if updatedUser.First_name == "" {
		updatedUser.First_name = existingUser.First_name
	}
	if updatedUser.Last_name == "" {
		updatedUser.Last_name = existingUser.Last_name
	}
	updatedUser.Created_at = existingUser.Created_at

	for i, u := range data.Users {
		if u.ID == id {
			data.Users[i] = updatedUser
			data.Users[i].ID = id
			data.Users[i].Updated_at = time.Now()
			data.SaveUser()
			c.JSON(http.StatusOK, gin.H{"message": "user updated"})
			return
		}
	}
}

