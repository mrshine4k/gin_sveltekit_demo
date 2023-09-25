package routes

import (
	"main/backend/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoute(router *gin.Engine) {
	//to export function, capitalize first letter
	router.POST("/album", controllers.CreateAlbum())
	router.GET("/album/:albumId", controllers.GetAnAlbum())
	router.PUT("/album/:albumId", controllers.UpdateAnAlbum())
	router.DELETE("/album/:albumId", controllers.DeleteAnAlbum())
	router.GET("/albums", controllers.GetAllAlbums())
}
