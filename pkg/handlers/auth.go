package handlers

import (
	"SongLibrary/logger"
	"SongLibrary/models"
	"SongLibrary/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("Client with IP: %s requested to sign in", ip)

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("Client with IP: %s failed to sign in: Error parsing request body: %v", ip, err)
		handleError(c, err)
		return
	}
	accessToken, err := services.SignIn(user.Username, user.Password)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("Client with IP: %s successfully signed in", ip)
	c.JSON(http.StatusOK, AccessTokenResponse{accessToken})
}

func SignUp(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("Client with IP: %s requested to create a new user", ip)
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("Client with IP: %s failed to create new user: Error parsing request body: %v", ip, err)
		handleError(c, err)
		return
	}

	id, err := services.CreateUser(user)
	if err != nil {
		handleError(c, err)
		return
	}
	response := NewDefaultResponse("User created successfully")
	logger.Info.Printf("Client with IP: %s successfully created a new user with ID: %d", ip, id)
	c.JSON(http.StatusCreated, gin.H{
		"message": response.Message,
		"user_id": id,
	})
}
