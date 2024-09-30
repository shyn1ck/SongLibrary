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
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)
	//
	//songGroup := r.Group("/songs")
	//{
	//	songGroup.GET("/", GetAllSongs)
	//	songGroup.GET("/:id", GetSongByID)
	//	songGroup.POST("/", AddSong)
	//	songGroup.PUT("/:id", UpdateSong)
	//	songGroup.DELETE("/:id", DeleteSong)
	//}
	//
	//artistGroup := r.Group("/artists")
	//{
	//	artistGroup.GET("/", GetAllArtists)
	//	artistGroup.GET("/:id", GetArtistByID)
	//	artistGroup.POST("/", AddArtist)
	//	artistGroup.PUT("/:id", UpdateArtist)
	//	artistGroup.DELETE("/:id", DeleteArtist)
	//}
	//
	//albumGroup := r.Group("/albums")
	//{
	//	albumGroup.GET("/", GetAllAlbums)
	//	albumGroup.GET("/:id", GetAlbumByID)
	//	albumGroup.POST("/", AddAlbum)
	//	albumGroup.PUT("/:id", UpdateAlbum)
	//	albumGroup.DELETE("/:id", DeleteAlbum)
	//}

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
