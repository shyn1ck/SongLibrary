package handlers

import (
	"SongLibrary/configs"
	"SongLibrary/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	r.GET("/ping", PingPong)

	if err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun)); err != nil {
		logger.Error.Fatalf("Error starting server: %v", err)
	}
	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
